package remainder

import (
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/registry"
	"github.com/mitchellh/hashstructure"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// 取余模式负载均衡实现
type RemainderLoadBalance struct {
	serviceName string
	reg         registry.Registry
	ready       int32
	reloadFunc  func() error
	Nodes       []*registry.Node

	sync.RWMutex
}

func (b *RemainderLoadBalance) SetServiceName(name string) {
	b.serviceName = name
}

func (b *RemainderLoadBalance) SetRegistry(reg registry.Registry) {
	b.reg = reg
}

func (b *RemainderLoadBalance) SetReloadFunc(f func() error) {
	b.reloadFunc = f
}

func (b *RemainderLoadBalance) Ready() bool {
	return atomic.LoadInt32(&b.ready) == 1
}

func (b *RemainderLoadBalance) GetService(key string) *registry.Node {
	id, err := hashstructure.Hash(key, nil)
	if err != nil {
		println(err.Error())
		return nil
	}

	b.RLock()
	defer b.RUnlock()

	return b.Nodes[int(id%uint64(len(b.Nodes)))]
}

func (b *RemainderLoadBalance) Start(TTL time.Duration) error {
	if b.serviceName == "" {
		panic("serviceName empty")
	}
	if b.reg == nil {
		panic("registry is nil")
	}

	if err := b.reload(); err != nil {
		return err
	}

	go b.watch()
	if TTL > 0 {
		go b.keepAlive(TTL)
	}

	return nil
}

func (b *RemainderLoadBalance) reload() error {
	atomic.StoreInt32(&b.ready, 0)
	ss, err := b.reg.GetService(b.serviceName)
	if err != nil {
		return err
	}

	b.Lock()
	b.Nodes = make([]*registry.Node, 0)
	for _, s := range ss {
		b.Nodes = append(b.Nodes, s.Nodes...)
	}
	sort.Slice(b.Nodes, func(i, j int) bool {
		return b.Nodes[i].Id < b.Nodes[j].Id
	})
	b.Unlock()

	if b.reloadFunc != nil {
		if err = b.reloadFunc(); err != nil {
			return err
		}
	}

	atomic.StoreInt32(&b.ready, 1)
	return nil
}

func (b *RemainderLoadBalance) watch() {
	var (
		watch     registry.Watcher
		err       error
		noWatcher = true
	)

	for {
		if noWatcher {
			watch, err = b.reg.Watch(registry.WatchService(b.serviceName))
			if err == nil {
				noWatcher = false
			} else {
				common.ErrorLog("new watcher err", nil, err.Error())
				continue
			}
		}

		_, err := watch.Next()
		if err != nil {
			common.ErrorLog("load balance watch err", nil, err.Error())
			continue
		}

		if err = b.reload(); err != nil {
			common.ErrorLog("load balance relaod err", nil, err.Error())
		}
	}
}

func (b *RemainderLoadBalance) keepAlive(TTL time.Duration) {
	TTL = TTL / 3
	ticker := time.NewTicker(TTL)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			b.reg.KeepAliveOnce()
		}
	}
}

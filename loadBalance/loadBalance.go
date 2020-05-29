package loadBalance

import (
	"time"

	"github.com/becent/golang-common/registry"
)

type LoadBalance interface {
	Ready() bool
	SetServiceName(name string)
	SetRegistry(reg registry.Registry)
	SetEndPoints(nodes []*registry.Node)
	SetReloadFunc(func() error)
	GetNode(key string) *registry.Node
	GetNodes() []*registry.Node
	Start(TTL time.Duration) error
	Close()
}

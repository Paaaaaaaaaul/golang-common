package base_server_sdk

import (
	"context"
	"errors"
	"github.com/SongLiangChen/grpc_pool"
	"github.com/becent/commom"
	"github.com/becent/commom/base_server_sdk/end"
	json "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Hosts struct {
	UserServerHost       string
	AccountServerHost    string
	StatisticServerHost  string
	ThirdPartyServerHost string
	// TODO add server host here
}

type BaseServerSdkClient struct {
	gRpcMapPool *grpc_pool.MapPool
	httpClient  *http.Client

	appId        string
	appSecretKey string

	requestTimeout time.Duration

	cp sync.Pool

	gRpcOnly bool

	Hosts Hosts
}

var (
	ErrHostEmpty   = errors.New("host empty, please config it when do InitBaseServerSdk")
	ErrServiceBusy = errors.New("service busy")
)

type Response struct {
	Success bool        `json:"success"`
	PayLoad interface{} `json:"payload"`
	Err     Error       `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *BaseServerSdkClient) DoRequest(host string, controller, action string, params map[string]string) ([]byte, error) {
	if host == "" {
		return nil, ErrHostEmpty
	}

	var (
		data []byte
		err  error
	)

	if c.gRpcOnly {
		data, err = c.doGRpcRequest(host, controller, action, params)
	} else {
		data, err = c.doHttpRequest(host, controller, action, params)
	}
	if err != nil {
		return nil, ErrServiceBusy
	}

	resp := &Response{}
	if err = json.Unmarshal(data, resp); err != nil {
		return nil, ErrServiceBusy
	}

	if !resp.Success {
		return nil, errors.New(resp.Err.Message)
	}

	return json.Marshal(resp.PayLoad)
}

func (c *BaseServerSdkClient) doHttpRequest(host string, controller, action string, params map[string]string) ([]byte, error) {
	// Assembly body
	var v url.Values
	for key, val := range params {
		v.Set(key, val)
	}
	body := strings.NewReader(v.Encode())

	// Make new request
	request, err := http.NewRequest("POST", strings.Join([]string{host, controller, action}, "/"), body)
	if err != nil {
		return nil, err
	}

	// Fill header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	request.Header.Set("Signature", c.makeSignature())
	request.Header.Set("RequestId", c.requestId())

	// Do http request
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *BaseServerSdkClient) doGRpcRequest(host string, controller, action string, params map[string]string) ([]byte, error) {
	// Make new request
	request := &grpc_end.Request{
		Controller: controller,
		Action:     action,
		Params:     params,
		Header: map[string]string{
			"requestId": c.requestId(),
			"signature": c.makeSignature(),
		},
	}

	// Get a gRpc conn from pool
	pool := c.gRpcMapPool.GetPool(host)
	conn, err := pool.Get()
	if err != nil {
		return nil, err
	}

	// Set the request timeout
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()

	// Do gRpc request
	resp, err := grpc_end.NewEndClient(conn.GetConn()).DoRequest(ctx, request)
	if err != nil {
		pool.DelErrorClient(conn)
		return nil, err
	}
	pool.Put(conn)

	return resp.Data, nil
}

func (c *BaseServerSdkClient) requestId() string {
	return common.Md5Encode(time.Now().String())
}

func (c *BaseServerSdkClient) makeSignature() string {
	now := strconv.FormatInt(time.Now().Unix(), 10)

	b := c.cp.Get().(*strings.Builder)
	defer c.cp.Put(b)

	b.WriteString(c.appId)
	b.WriteString("-")
	b.WriteString(now)
	b.WriteString("-")
	b.WriteString(c.appId)

	data := b.String()

	b.Reset()
	b.WriteString(now)
	b.WriteString(":")
	b.WriteString(c.appId)
	b.WriteString(":")
	b.WriteString(common.Hmac(c.appSecretKey, data))

	return b.String()
}

func DialFunc(addr string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure())
}

type Config struct {
	AppId           string
	AppSecretKey    string
	RequestTimeout  time.Duration
	IdleConnTimeout time.Duration
	Hosts           Hosts
	GRpcOnly        bool
}

var Instance *BaseServerSdkClient

func InitBaseServerSdk(c *Config) {
	Instance = &BaseServerSdkClient{
		gRpcMapPool: grpc_pool.NewMapPool(DialFunc, 0, c.IdleConnTimeout),

		httpClient: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   c.RequestTimeout,
					KeepAlive: 0,
					DualStack: true,
				}).DialContext,
				IdleConnTimeout:       c.IdleConnTimeout,
				ResponseHeaderTimeout: c.RequestTimeout,
			},
		},

		appId:        c.AppId,
		appSecretKey: c.AppSecretKey,

		gRpcOnly: c.GRpcOnly,

		Hosts: c.Hosts,
	}

	Instance.cp.New = func() interface{} {
		return &strings.Builder{}
	}
}

func ReleaseBaseServerSdk() {
	Instance.gRpcMapPool.ReleaseAllPool()
}

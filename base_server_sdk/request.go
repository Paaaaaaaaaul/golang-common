package base_server_sdk

import (
	"context"
	"encoding/json"
	"github.com/becent/commom"
	"github.com/becent/commom/base_server_sdk/end"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	ErrServiceBusy = &Error{
		Code:    900000,
		Message: "service busy",
	}
	ErrHostEmpty = &Error{
		Code:    900001,
		Message: "host empty, please config it when do InitBaseServerSdk",
	}
	ErrInvalidParams = &Error{
		Code:    900002,
		Message: "params invalid or missed",
	}
)

type Response struct {
	Success bool        `json:"success"`
	PayLoad interface{} `json:"payload"`
	Err     *Error      `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) String() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}

func (c *BaseServerSdkClient) DoRequest(host string, controller, action string, params map[string]string) ([]byte, *Error) {
	if host == "" {
		return nil, ErrHostEmpty
	}

	params["orgId"] = strconv.Itoa(c.OrgId)

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
		common.ErrorLog("baseServerSdk_DoRequest", map[string]interface{}{
			"host":       host,
			"controller": controller,
			"action":     action,
			"params":     params,
		}, err.Error())
		return nil, ErrServiceBusy
	}

	resp := &Response{}
	if err = json.Unmarshal(data, resp); err != nil {
		common.ErrorLog("baseServerSdk_DoRequest", map[string]interface{}{
			"host":       host,
			"controller": controller,
			"action":     action,
			"params":     params,
		}, err.Error())
		return nil, ErrServiceBusy
	}

	if !resp.Success {
		return nil, resp.Err
	}

	if data, err = json.Marshal(resp.PayLoad); err != nil {
		common.ErrorLog("baseServerSdk_DoRequest", map[string]interface{}{
			"host":       host,
			"controller": controller,
			"action":     action,
			"params":     params,
		}, err.Error())
		return nil, ErrServiceBusy
	}

	return data, nil
}

func (c *BaseServerSdkClient) doHttpRequest(host string, controller, action string, params map[string]string) ([]byte, error) {
	// Assembly body
	v := make(url.Values)
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
	_ = pool.Put(conn)

	return resp.Data, nil
}

func (c *BaseServerSdkClient) requestId() string {
	return common.Md5Encode(time.Now().String())
}

func (c *BaseServerSdkClient) makeSignature() string {
	now := strconv.FormatInt(time.Now().Unix(), 10)

	b := c.cp.Get().(*strings.Builder)
	defer c.cp.Put(b)
	b.Reset()

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

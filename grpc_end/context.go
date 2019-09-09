package grpc_end

import (
	"context"
	json "github.com/json-iterator/go"
	"math"
	"strconv"
)

const abortIndex int8 = math.MaxInt8 / 2

type GRpcContext struct {
	handlers HandlersChain // Execution chain
	index    int8          // index of current execution handler

	ctx    context.Context
	engine *GRpcEngine

	req  *Request
	resp *Response

	// Keys is a key/value pair exclusively for the context of each request.
	Keys map[string]interface{}

	// appName, for log
	appName string
}

// ---------------------------------------------------------------------------------------------------------------------

func (c *GRpcContext) reset() {
	c.handlers = nil
	c.index = -1
	c.ctx = nil
	c.engine = nil
	c.req = nil
	c.resp = &Response{}
	c.Keys = nil
}

func (c *GRpcContext) GetContext() context.Context {
	return c.ctx
}

// GetRequest returns GRpc Request
func (c *GRpcContext) GetRequest() *Request {
	return c.req
}

// GetResponse returns GRpc Response of this request
func (c *GRpcContext) GetResponse() *Response {
	return c.resp
}

// GetAppName returns the appName
func (c *GRpcContext) GetAppName() string {
	return c.appName
}

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes c.Keys if it was not used previously.
func (c *GRpcContext) Set(key string, val interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = val
}

// Get returns the value for the given key.
// If the value does not exists it returns nil
func (c *GRpcContext) Get(key string) interface{} {
	if c.Keys == nil {
		return nil
	}
	return c.Keys[key]
}

// GetStringMap returns string for the given key
// If the value does not exists it returns empty string
func (c *GRpcContext) GetString(key string) string {
	if c.Keys == nil {
		return ""
	}
	if val, ok := c.Keys[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}

	return ""
}

// GetStringMap returns the map[string]string for the given key
// If the value does not exists it returns nil
func (c *GRpcContext) GetStringMap(key string) map[string]string {
	if c.Keys == nil {
		return nil
	}
	if val, ok := c.Keys[key]; ok {
		if m, ok := val.(map[string]string); ok {
			return m
		}
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// ParamString returns string val from request's Params for the given key
// and returns empty string if val not exists
func (c *GRpcContext) ParamString(key string) string {
	return c.req.Params[key]
}

// ParamString returns string val from request's Params for the given key
// and returns defVal if val not exists.
func (c *GRpcContext) ParamStringDefault(key string, defVal string) string {
	if val, ok := c.req.Params[key]; ok && val != "" {
		return val
	}
	return defVal
}

// ParamInt returns int val from request's Params for the given key
// and returns zero if val not exists.
func (c *GRpcContext) ParamInt(key string) int {
	if val := c.ParamString(key); val != "" {
		if n, err := strconv.Atoi(val); err == nil {
			return n
		}
	}
	return 0
}

// ParamInt returns int val from request's Params for the given key
// and returns defVal if val not exists.
func (c *GRpcContext) ParamIntDefault(key string, defVal int) int {
	if val := c.ParamString(key); val != "" {
		if n, err := strconv.Atoi(val); err == nil {
			return n
		}
	}
	return defVal
}

// ParamInt returns int64 val from request's Params for the given key
// and returns zero if val not exists.
func (c *GRpcContext) ParamInt64(key string) int64 {
	if val := c.ParamString(key); val != "" {
		if n, err := strconv.ParseInt(val, 10, 64); err == nil {
			return n
		}
	}
	return 0
}

// ParamInt returns int64 val from request's Params for the given key
// and returns defVal if val not exists.
func (c *GRpcContext) ParamInt64Default(key string, defVal int64) int64 {
	if val := c.ParamString(key); val != "" {
		if n, err := strconv.ParseInt(val, 10, 64); err == nil {
			return n
		}
	}
	return defVal
}

// ---------------------------------------------------------------------------------------------------------------------

// HeaderString returns string val from request's Header for the given key
// and returns empty string if val not exists
//
// The GateWay will fill in some val to request's Header, like:
// --------------------------------------------------------------------
// | key     | desc                                                   |
// --------------------------------------------------------------------
// | ip      | client's ip address                                    |
// | lang    | the Language client use, ie: 'zh', 'en', 'ko'...       |
// | device  | the Device client use, ie: 'iphone 7 Plus', 'chrome'   |
// | dt      | 'N' mean web, 'a' mean android, 'i' mean ios           |
// | userId  | the client's unique id                                 |
// | orgId   | the client's organization id                           |
// | host    | the host of this request belong                        |
// | account | phone num or email address                             |
// --------------------------------------------------------------------

var (
	HeaderKeyIp      = "ip"
	HeaderKeyLang    = "lang"
	HeaderKeyDevice  = "device"
	HeaderKeyDt      = "dt"
	HeaderKeyUserId  = "userId"
	HeaderKeyOrgId   = "orgId"
	HeaderKeyHost    = "host"
	HeaderKeyAccount = "account"
)

func (c *GRpcContext) HeaderString(key string) string {
	return c.req.Header[key]
}

// HeaderString returns string val from request's Header for the given key
// and returns defVal if val not exists
func (c *GRpcContext) HeaderStringDefault(key string, defVal string) string {
	if _, ok := c.req.Header[key]; !ok {
		return defVal
	}

	return c.HeaderString(key)
}

// HeaderString returns int val from request's Header for the given key
// and returns zero if val not exists
func (c *GRpcContext) HeaderInt(key string) int {
	if val := c.HeaderString(key); val != "" {
		if n, err := strconv.Atoi(val); err == nil {
			return n
		}
	}
	return 0
}

// HeaderString returns int64 val from request's Header for the given key
// and returns zero if val not exists
func (c *GRpcContext) HeaderInt64(key string) int64 {
	if val := c.HeaderString(key); val != "" {
		if n, err := strconv.ParseInt(val, 10, 64); err == nil {
			return n
		}
	}
	return 0
}

// ---------------------------------------------------------------------------------------------------------------------

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
func (c *GRpcContext) Next() {
	c.index++
	for s := int8(len(c.handlers)); c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *GRpcContext) Abort() {
	c.index = abortIndex
}

// IsAborted returns true if the current context was aborted.
func (c *GRpcContext) IsAbort() bool {
	return c.index >= abortIndex
}

// ---------------------------------------------------------------------------------------------------------------------

type SResponse struct {
	Success bool        `json:"success"`
	PayLoad interface{} `json:"payload"`
}

type EResponse struct {
	Success bool  `json:"success"`
	Err     Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *GRpcContext) SuccessResponse(v interface{}) {
	jsonStr, _ := json.Marshal(&SResponse{
		Success: true,
		PayLoad: v,
	})
	c.resp.Data = jsonStr
}

func (c *GRpcContext) ErrorResponse(code int, message string) {
	jsonStr, _ := json.Marshal(&EResponse{
		Success: false,
		Err: Error{
			Code:    code,
			Message: message,
		},
	})
	c.resp.Data = jsonStr
}

// ---------------------------------------------------------------------------------------------------------------------
//                                      Gateway related operations
// ---------------------------------------------------------------------------------------------------------------------

func (c *GRpcContext) MarkGateWaySuccess(v bool) {
	c.resp.Success = v
}

func (c *GRpcContext) MarkGateWayUserId(userId string) {
	c.resp.UserId = userId
}

func (c *GRpcContext) MarkGateWayOrgId(orgId string) {
	c.resp.OrgId = orgId
}

func (c *GRpcContext) MarkGateWayUserLevel(userLevel int32) {
	c.resp.UserLevel = userLevel
}

func (c *GRpcContext) MarkGateWayResponse(data []byte) {
	c.resp.Data = data
}

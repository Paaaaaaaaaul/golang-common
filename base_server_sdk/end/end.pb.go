// Code generated by protoc-gen-go. DO NOT EDIT.
// source: end.proto

package grpc_end

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	// 请求的controller
	Controller string `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	// 请求的action
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	// 请求附带的所有参数
	Params map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 请求附带的文件, key为文件字段名，val为文件的原始数据
	Files map[string][]byte `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 网关补全的字段
	Header               map[string]string `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5c7823ce1477e4, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *Request) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Request) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Request) GetFiles() map[string][]byte {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Request) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

type Response struct {
	// 供网关确认业务是否成功
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// 供网关执行业务
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	// 供网关执行业务
	OrgId string `protobuf:"bytes,3,opt,name=orgId,proto3" json:"orgId,omitempty"`
	// 用户的访问权限
	UserLevel int32 `protobuf:"varint,4,opt,name=userLevel,proto3" json:"userLevel,omitempty"`
	// 实际返回给客户端的数据，网关将原样交给客户端
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5c7823ce1477e4, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Response) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Response) GetUserLevel() int32 {
	if m != nil {
		return m.UserLevel
	}
	return 0
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "grpc_end.Request")
	proto.RegisterMapType((map[string][]byte)(nil), "grpc_end.Request.FilesEntry")
	proto.RegisterMapType((map[string]string)(nil), "grpc_end.Request.HeaderEntry")
	proto.RegisterMapType((map[string]string)(nil), "grpc_end.Request.ParamsEntry")
	proto.RegisterType((*Response)(nil), "grpc_end.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndClient is the client API for End service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndClient interface {
	// 处理网关请求
	DoRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type endClient struct {
	cc *grpc.ClientConn
}

func NewEndClient(cc *grpc.ClientConn) EndClient {
	return &endClient{cc}
}

func (c *endClient) DoRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc_end.End/DoRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndServer is the server API for End service.
type EndServer interface {
	// 处理网关请求
	DoRequest(context.Context, *Request) (*Response, error)
}

func RegisterEndServer(s *grpc.Server, srv EndServer) {
	s.RegisterService(&_End_serviceDesc, srv)
}

func _End_DoRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndServer).DoRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_end.End/DoRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndServer).DoRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _End_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_end.End",
	HandlerType: (*EndServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoRequest",
			Handler:    _End_DoRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "end.proto",
}

func init() { proto.RegisterFile("end.proto", fileDescriptor_dc5c7823ce1477e4) }

var fileDescriptor_dc5c7823ce1477e4 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0xd2, 0x34, 0x6d, 0x33, 0xfd, 0x0e, 0x3a, 0x88, 0x2c, 0xa5, 0x4a, 0xe9, 0xa9, 0xa7,
	0x1c, 0xa2, 0x42, 0xab, 0x57, 0x2b, 0x16, 0x3c, 0x48, 0xfe, 0x80, 0xac, 0xd9, 0xb1, 0x16, 0xe3,
	0x6e, 0xdc, 0x4d, 0x0a, 0xbd, 0xfb, 0x0f, 0xfc, 0xc3, 0x92, 0xdd, 0x84, 0x06, 0x8a, 0x88, 0xb7,
	0x7d, 0xf3, 0xe6, 0xcd, 0x7b, 0xb3, 0xbb, 0x10, 0x92, 0x14, 0x51, 0xae, 0x55, 0xa1, 0x70, 0xb0,
	0xd6, 0x79, 0xfa, 0x44, 0x52, 0x4c, 0xbf, 0x7c, 0xe8, 0x27, 0xf4, 0x51, 0x92, 0x29, 0xf0, 0x1c,
	0x20, 0x55, 0xb2, 0xd0, 0x2a, 0xcb, 0x48, 0x33, 0x6f, 0xe2, 0xcd, 0xc2, 0xa4, 0x55, 0xc1, 0x53,
	0xe8, 0xf1, 0xb4, 0xd8, 0x28, 0xc9, 0x3a, 0x96, 0xab, 0x11, 0x5e, 0x41, 0x2f, 0xe7, 0x9a, 0xbf,
	0x1b, 0xe6, 0x4f, 0xfc, 0xd9, 0x30, 0x3e, 0x8b, 0x9a, 0xf1, 0x51, 0x3d, 0x3a, 0x7a, 0xb4, 0xfc,
	0x52, 0x16, 0x7a, 0x97, 0xd4, 0xcd, 0x18, 0x43, 0xf0, 0xb2, 0xc9, 0xc8, 0xb0, 0xae, 0x55, 0x8d,
	0x0f, 0x55, 0x77, 0x15, 0xed, 0x44, 0xae, 0xb5, 0xb2, 0x7a, 0x25, 0x2e, 0x48, 0xb3, 0xe0, 0x27,
	0xab, 0x7b, 0xcb, 0xd7, 0x56, 0xae, 0x79, 0xb4, 0x80, 0x61, 0x2b, 0x01, 0x1e, 0x81, 0xff, 0x46,
	0xbb, 0x7a, 0xc3, 0xea, 0x88, 0x27, 0x10, 0x6c, 0x79, 0x56, 0x52, 0xbd, 0x99, 0x03, 0xd7, 0x9d,
	0xb9, 0x37, 0x9a, 0x03, 0xec, 0x63, 0xfc, 0xa6, 0xfc, 0xdf, 0x56, 0x2e, 0x60, 0xd8, 0xca, 0xf2,
	0x17, 0xd3, 0xe9, 0xa7, 0x07, 0x83, 0x84, 0x4c, 0xae, 0xa4, 0x21, 0x64, 0xd0, 0x37, 0x65, 0x9a,
	0x92, 0x31, 0x56, 0x3c, 0x48, 0x1a, 0x58, 0x3d, 0x48, 0x69, 0x48, 0xaf, 0x44, 0xf3, 0x20, 0x0e,
	0x55, 0x83, 0x95, 0x5e, 0xaf, 0x04, 0xf3, 0xdd, 0x60, 0x0b, 0x70, 0x0c, 0x61, 0xc5, 0x3f, 0xd0,
	0x96, 0x32, 0xd6, 0x9d, 0x78, 0xb3, 0x20, 0xd9, 0x17, 0x10, 0xa1, 0x2b, 0x78, 0xc1, 0x59, 0x60,
	0xd7, 0xb0, 0xe7, 0xf8, 0x06, 0xfc, 0xa5, 0x14, 0x78, 0x09, 0xe1, 0xad, 0x6a, 0x3e, 0xc9, 0xf1,
	0xc1, 0x8d, 0x8f, 0xb0, 0x5d, 0x72, 0xa1, 0xa7, 0xff, 0x9e, 0x7b, 0xf6, 0xab, 0x5d, 0x7c, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0xcf, 0xca, 0xa0, 0x77, 0x02, 0x00, 0x00,
}
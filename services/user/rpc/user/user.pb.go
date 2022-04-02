// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WeChatLoginRequest struct {
	Appid                string   `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatLoginRequest) Reset()         { *m = WeChatLoginRequest{} }
func (m *WeChatLoginRequest) String() string { return proto.CompactTextString(m) }
func (*WeChatLoginRequest) ProtoMessage()    {}
func (*WeChatLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *WeChatLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatLoginRequest.Unmarshal(m, b)
}
func (m *WeChatLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatLoginRequest.Marshal(b, m, deterministic)
}
func (m *WeChatLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatLoginRequest.Merge(m, src)
}
func (m *WeChatLoginRequest) XXX_Size() int {
	return xxx_messageInfo_WeChatLoginRequest.Size(m)
}
func (m *WeChatLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatLoginRequest proto.InternalMessageInfo

func (m *WeChatLoginRequest) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *WeChatLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type WeChatLoginResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeChatLoginResponse) Reset()         { *m = WeChatLoginResponse{} }
func (m *WeChatLoginResponse) String() string { return proto.CompactTextString(m) }
func (*WeChatLoginResponse) ProtoMessage()    {}
func (*WeChatLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *WeChatLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeChatLoginResponse.Unmarshal(m, b)
}
func (m *WeChatLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeChatLoginResponse.Marshal(b, m, deterministic)
}
func (m *WeChatLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeChatLoginResponse.Merge(m, src)
}
func (m *WeChatLoginResponse) XXX_Size() int {
	return xxx_messageInfo_WeChatLoginResponse.Size(m)
}
func (m *WeChatLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WeChatLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WeChatLoginResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*WeChatLoginRequest)(nil), "user.WeChatLoginRequest")
	proto.RegisterType((*WeChatLoginResponse)(nil), "user.WeChatLoginResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xec, 0xb8, 0x84, 0xc2, 0x53,
	0x9d, 0x33, 0x12, 0x4b, 0x7c, 0xf2, 0xd3, 0x33, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x44, 0xb8, 0x58, 0x13, 0x0b, 0x0a, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0x09, 0x26, 0xb0, 0x20, 0x98, 0xad,
	0x24, 0xca, 0x25, 0x8c, 0xa2, 0xbf, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xc8, 0x8b, 0x8b, 0x25,
	0xb4, 0x38, 0xb5, 0x48, 0xc8, 0x89, 0x8b, 0x1b, 0x49, 0x5a, 0x48, 0x42, 0x0f, 0xec, 0x00, 0x4c,
	0x1b, 0xa5, 0x24, 0xb1, 0xc8, 0x40, 0xcc, 0x4a, 0x62, 0x03, 0xbb, 0xd7, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x62, 0xf1, 0x3c, 0x35, 0xbd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	WeChatLogin(ctx context.Context, in *WeChatLoginRequest, opts ...grpc.CallOption) (*WeChatLoginResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) WeChatLogin(ctx context.Context, in *WeChatLoginRequest, opts ...grpc.CallOption) (*WeChatLoginResponse, error) {
	out := new(WeChatLoginResponse)
	err := c.cc.Invoke(ctx, "/user.User/WeChatLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	WeChatLogin(context.Context, *WeChatLoginRequest) (*WeChatLoginResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) WeChatLogin(ctx context.Context, req *WeChatLoginRequest) (*WeChatLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeChatLogin not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_WeChatLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeChatLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).WeChatLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/WeChatLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).WeChatLogin(ctx, req.(*WeChatLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WeChatLogin",
			Handler:    _User_WeChatLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallpaper.proto

package wallpaper

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

type WallPaperInfo struct {
	Wid                  string   `protobuf:"bytes,1,opt,name=Wid,proto3" json:"Wid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageURL             string   `protobuf:"bytes,3,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Author               string   `protobuf:"bytes,4,opt,name=Author,proto3" json:"Author,omitempty"`
	Desc                 string   `protobuf:"bytes,5,opt,name=Desc,proto3" json:"Desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WallPaperInfo) Reset()         { *m = WallPaperInfo{} }
func (m *WallPaperInfo) String() string { return proto.CompactTextString(m) }
func (*WallPaperInfo) ProtoMessage()    {}
func (*WallPaperInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdca87ab84ee5c8e, []int{0}
}

func (m *WallPaperInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WallPaperInfo.Unmarshal(m, b)
}
func (m *WallPaperInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WallPaperInfo.Marshal(b, m, deterministic)
}
func (m *WallPaperInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WallPaperInfo.Merge(m, src)
}
func (m *WallPaperInfo) XXX_Size() int {
	return xxx_messageInfo_WallPaperInfo.Size(m)
}
func (m *WallPaperInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WallPaperInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WallPaperInfo proto.InternalMessageInfo

func (m *WallPaperInfo) GetWid() string {
	if m != nil {
		return m.Wid
	}
	return ""
}

func (m *WallPaperInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WallPaperInfo) GetImageURL() string {
	if m != nil {
		return m.ImageURL
	}
	return ""
}

func (m *WallPaperInfo) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *WallPaperInfo) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ImportRequest struct {
	List                 []*WallPaperInfo `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ImportRequest) Reset()         { *m = ImportRequest{} }
func (m *ImportRequest) String() string { return proto.CompactTextString(m) }
func (*ImportRequest) ProtoMessage()    {}
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdca87ab84ee5c8e, []int{1}
}

func (m *ImportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportRequest.Unmarshal(m, b)
}
func (m *ImportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportRequest.Marshal(b, m, deterministic)
}
func (m *ImportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportRequest.Merge(m, src)
}
func (m *ImportRequest) XXX_Size() int {
	return xxx_messageInfo_ImportRequest.Size(m)
}
func (m *ImportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportRequest proto.InternalMessageInfo

func (m *ImportRequest) GetList() []*WallPaperInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ImportResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportResponse) Reset()         { *m = ImportResponse{} }
func (m *ImportResponse) String() string { return proto.CompactTextString(m) }
func (*ImportResponse) ProtoMessage()    {}
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdca87ab84ee5c8e, []int{2}
}

func (m *ImportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportResponse.Unmarshal(m, b)
}
func (m *ImportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportResponse.Marshal(b, m, deterministic)
}
func (m *ImportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportResponse.Merge(m, src)
}
func (m *ImportResponse) XXX_Size() int {
	return xxx_messageInfo_ImportResponse.Size(m)
}
func (m *ImportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportResponse proto.InternalMessageInfo

type RemoveRequest struct {
	List                 []string `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdca87ab84ee5c8e, []int{3}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type RemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdca87ab84ee5c8e, []int{4}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

type DetailRequest struct {
	Wid                  string   `protobuf:"bytes,1,opt,name=Wid,proto3" json:"Wid,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=Start,proto3" json:"Start,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdca87ab84ee5c8e, []int{5}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

func (m *DetailRequest) GetWid() string {
	if m != nil {
		return m.Wid
	}
	return ""
}

func (m *DetailRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DetailRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DetailResponse struct {
	List                 []*WallPaperInfo `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Total                int64            `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DetailResponse) Reset()         { *m = DetailResponse{} }
func (m *DetailResponse) String() string { return proto.CompactTextString(m) }
func (*DetailResponse) ProtoMessage()    {}
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdca87ab84ee5c8e, []int{6}
}

func (m *DetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResponse.Unmarshal(m, b)
}
func (m *DetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResponse.Marshal(b, m, deterministic)
}
func (m *DetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResponse.Merge(m, src)
}
func (m *DetailResponse) XXX_Size() int {
	return xxx_messageInfo_DetailResponse.Size(m)
}
func (m *DetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResponse proto.InternalMessageInfo

func (m *DetailResponse) GetList() []*WallPaperInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *DetailResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*WallPaperInfo)(nil), "wallpaper.WallPaperInfo")
	proto.RegisterType((*ImportRequest)(nil), "wallpaper.ImportRequest")
	proto.RegisterType((*ImportResponse)(nil), "wallpaper.ImportResponse")
	proto.RegisterType((*RemoveRequest)(nil), "wallpaper.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "wallpaper.RemoveResponse")
	proto.RegisterType((*DetailRequest)(nil), "wallpaper.DetailRequest")
	proto.RegisterType((*DetailResponse)(nil), "wallpaper.DetailResponse")
}

func init() { proto.RegisterFile("wallpaper.proto", fileDescriptor_bdca87ab84ee5c8e) }

var fileDescriptor_bdca87ab84ee5c8e = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x95, 0x7f, 0xd2, 0xe8, 0x9f, 0x43, 0x29, 0x95, 0x85, 0x90, 0xe9, 0x54, 0x85, 0xa5,
	0x03, 0xea, 0x50, 0xe6, 0x0c, 0x48, 0x5d, 0x22, 0x05, 0x84, 0x4c, 0x51, 0x66, 0x03, 0x06, 0x22,
	0x39, 0x75, 0x48, 0x5c, 0x98, 0xf8, 0x7e, 0x7c, 0x2c, 0x64, 0x3b, 0x6e, 0x7b, 0x82, 0x85, 0xed,
	0xde, 0xd9, 0xbf, 0xe7, 0xbb, 0x27, 0xc3, 0xf1, 0x07, 0x97, 0xb2, 0xe5, 0xad, 0xe8, 0x16, 0x6d,
	0xa7, 0xb4, 0x22, 0xc9, 0xae, 0x91, 0x7d, 0x42, 0x5a, 0x71, 0x29, 0x6f, 0x8d, 0x28, 0x36, 0xcf,
	0x8a, 0x4c, 0x20, 0xac, 0xea, 0x27, 0x1a, 0xcc, 0x82, 0x79, 0xc2, 0x4c, 0x49, 0x08, 0x44, 0x37,
	0xbc, 0x11, 0xf4, 0x9f, 0x6d, 0xd9, 0x9a, 0x4c, 0xe1, 0x7f, 0xd1, 0xf0, 0x17, 0x71, 0xcf, 0x4a,
	0x1a, 0xda, 0xfe, 0x4e, 0x93, 0x53, 0x88, 0xaf, 0xb6, 0xfa, 0x55, 0x75, 0x34, 0xb2, 0x27, 0x83,
	0x32, 0x3e, 0x2b, 0xd1, 0x3f, 0xd2, 0x91, 0xf3, 0x31, 0x75, 0x96, 0x43, 0x5a, 0x34, 0xad, 0xea,
	0x34, 0x13, 0x6f, 0x5b, 0xd1, 0x6b, 0x72, 0x01, 0x51, 0x59, 0xf7, 0x9a, 0x06, 0xb3, 0x70, 0x7e,
	0xb4, 0xa4, 0x8b, 0xfd, 0xe8, 0x68, 0x4c, 0x66, 0x6f, 0x65, 0x13, 0x18, 0x7b, 0xbc, 0x6f, 0xd5,
	0xa6, 0x17, 0xd9, 0x39, 0xa4, 0x4c, 0x34, 0xea, 0x5d, 0x78, 0x43, 0x72, 0x60, 0x98, 0xec, 0x31,
	0x7f, 0x69, 0xc0, 0xae, 0x21, 0x5d, 0x09, 0xcd, 0x6b, 0xe9, 0xb1, 0x9f, 0x31, 0x9c, 0xc0, 0xe8,
	0x4e, 0xf3, 0x4e, 0xdb, 0x1c, 0x42, 0xe6, 0x84, 0xe9, 0x96, 0x75, 0x53, 0x6b, 0x9b, 0x42, 0xc8,
	0x9c, 0xc8, 0xd6, 0x30, 0xf6, 0x76, 0xee, 0x81, 0xbf, 0xed, 0x65, 0x5c, 0xd7, 0x4a, 0x73, 0xe9,
	0xdf, 0xb2, 0x62, 0xf9, 0x15, 0x40, 0x52, 0x79, 0x8e, 0xe4, 0x10, 0xbb, 0xdd, 0xc9, 0xa1, 0x1b,
	0x4a, 0x73, 0x7a, 0xf6, 0xcb, 0xc9, 0x30, 0x50, 0x0e, 0xb1, 0xcb, 0x00, 0xe1, 0x28, 0x3b, 0x84,
	0xe3, 0xc0, 0x0c, 0xee, 0x36, 0x44, 0x38, 0xca, 0x10, 0xe1, 0x38, 0x8e, 0x87, 0xd8, 0x7e, 0xc4,
	0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0xad, 0x1d, 0x7d, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WallpaperClient is the client API for Wallpaper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WallpaperClient interface {
	Import(ctx context.Context, in *ImportRequest, opts ...grpc.CallOption) (*ImportResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
}

type wallpaperClient struct {
	cc *grpc.ClientConn
}

func NewWallpaperClient(cc *grpc.ClientConn) WallpaperClient {
	return &wallpaperClient{cc}
}

func (c *wallpaperClient) Import(ctx context.Context, in *ImportRequest, opts ...grpc.CallOption) (*ImportResponse, error) {
	out := new(ImportResponse)
	err := c.cc.Invoke(ctx, "/wallpaper.Wallpaper/Import", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wallpaperClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/wallpaper.Wallpaper/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wallpaperClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, "/wallpaper.Wallpaper/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WallpaperServer is the server API for Wallpaper service.
type WallpaperServer interface {
	Import(context.Context, *ImportRequest) (*ImportResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
}

// UnimplementedWallpaperServer can be embedded to have forward compatible implementations.
type UnimplementedWallpaperServer struct {
}

func (*UnimplementedWallpaperServer) Import(ctx context.Context, req *ImportRequest) (*ImportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Import not implemented")
}
func (*UnimplementedWallpaperServer) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedWallpaperServer) Detail(ctx context.Context, req *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}

func RegisterWallpaperServer(s *grpc.Server, srv WallpaperServer) {
	s.RegisterService(&_Wallpaper_serviceDesc, srv)
}

func _Wallpaper_Import_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WallpaperServer).Import(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallpaper.Wallpaper/Import",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WallpaperServer).Import(ctx, req.(*ImportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallpaper_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WallpaperServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallpaper.Wallpaper/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WallpaperServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallpaper_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WallpaperServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallpaper.Wallpaper/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WallpaperServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wallpaper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wallpaper.Wallpaper",
	HandlerType: (*WallpaperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Import",
			Handler:    _Wallpaper_Import_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Wallpaper_Remove_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _Wallpaper_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallpaper.proto",
}

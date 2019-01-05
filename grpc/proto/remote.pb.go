// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/remote.proto

// Web exposes a backend server over gRPC.

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type PlayerID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerID) Reset()         { *m = PlayerID{} }
func (m *PlayerID) String() string { return proto.CompactTextString(m) }
func (*PlayerID) ProtoMessage()    {}
func (*PlayerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{0}
}

func (m *PlayerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerID.Unmarshal(m, b)
}
func (m *PlayerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerID.Marshal(b, m, deterministic)
}
func (m *PlayerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerID.Merge(m, src)
}
func (m *PlayerID) XXX_Size() int {
	return xxx_messageInfo_PlayerID.Size(m)
}
func (m *PlayerID) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerID.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerID proto.InternalMessageInfo

func (m *PlayerID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type WorldResponse struct {
	Tiles                []byte   `protobuf:"bytes,1,opt,name=tiles,proto3" json:"tiles,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorldResponse) Reset()         { *m = WorldResponse{} }
func (m *WorldResponse) String() string { return proto.CompactTextString(m) }
func (*WorldResponse) ProtoMessage()    {}
func (*WorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{2}
}

func (m *WorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorldResponse.Unmarshal(m, b)
}
func (m *WorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorldResponse.Marshal(b, m, deterministic)
}
func (m *WorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorldResponse.Merge(m, src)
}
func (m *WorldResponse) XXX_Size() int {
	return xxx_messageInfo_WorldResponse.Size(m)
}
func (m *WorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorldResponse proto.InternalMessageInfo

func (m *WorldResponse) GetTiles() []byte {
	if m != nil {
		return m.Tiles
	}
	return nil
}

func (m *WorldResponse) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *WorldResponse) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type EntityResponse struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityResponse) Reset()         { *m = EntityResponse{} }
func (m *EntityResponse) String() string { return proto.CompactTextString(m) }
func (*EntityResponse) ProtoMessage()    {}
func (*EntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{3}
}

func (m *EntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityResponse.Unmarshal(m, b)
}
func (m *EntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityResponse.Marshal(b, m, deterministic)
}
func (m *EntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityResponse.Merge(m, src)
}
func (m *EntityResponse) XXX_Size() int {
	return xxx_messageInfo_EntityResponse.Size(m)
}
func (m *EntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntityResponse proto.InternalMessageInfo

func (m *EntityResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayerID)(nil), "web.PlayerID")
	proto.RegisterType((*Empty)(nil), "web.Empty")
	proto.RegisterType((*WorldResponse)(nil), "web.WorldResponse")
	proto.RegisterType((*EntityResponse)(nil), "web.EntityResponse")
}

func init() { proto.RegisterFile("proto/remote.proto", fileDescriptor_b59cb7d8fa836ac0) }

var fileDescriptor_b59cb7d8fa836ac0 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x6a, 0x83, 0x40,
	0x10, 0x86, 0x63, 0x42, 0x92, 0x66, 0x48, 0x72, 0xd8, 0x96, 0x22, 0x39, 0x85, 0x3d, 0x49, 0x0f,
	0x2a, 0xc9, 0xb1, 0xb7, 0x60, 0x0e, 0x5e, 0x4a, 0x31, 0x87, 0x42, 0x6f, 0xab, 0x0e, 0x2a, 0x51,
	0x77, 0xbb, 0x4e, 0x2a, 0xbe, 0x4a, 0x9f, 0xb6, 0x74, 0x35, 0x05, 0x6f, 0xff, 0x37, 0xcc, 0x0e,
	0xdf, 0xbf, 0xc0, 0x94, 0x96, 0x24, 0x3d, 0x8d, 0x95, 0x24, 0x74, 0x0d, 0xb0, 0x59, 0x8b, 0x31,
	0xdf, 0xc1, 0xc3, 0x7b, 0x29, 0x3a, 0xd4, 0x61, 0xc0, 0xb6, 0x30, 0x0d, 0x03, 0xdb, 0xda, 0x5b,
	0xce, 0x2a, 0x9a, 0x86, 0x01, 0x5f, 0xc2, 0xfc, 0x5c, 0x29, 0xea, 0xf8, 0x05, 0x36, 0x1f, 0x52,
	0x97, 0x69, 0x84, 0x8d, 0x92, 0x75, 0x83, 0xec, 0x09, 0xe6, 0x54, 0x94, 0xd8, 0x98, 0xe5, 0x75,
	0xd4, 0xc3, 0xdf, 0xb4, 0x2d, 0x52, 0xca, 0xed, 0xe9, 0xde, 0x72, 0xe6, 0x51, 0x0f, 0xec, 0x19,
	0x16, 0x39, 0x16, 0x59, 0x4e, 0xf6, 0xcc, 0x8c, 0x07, 0xe2, 0x2f, 0xb0, 0x3d, 0xd7, 0x54, 0x50,
	0xf7, 0x7f, 0xd5, 0x86, 0xa5, 0x12, 0x5d, 0x29, 0x45, 0x3a, 0xdc, 0xbd, 0xe3, 0xe1, 0xc7, 0x82,
	0xe5, 0x49, 0x24, 0x57, 0xac, 0x53, 0xe6, 0xc0, 0xea, 0x0d, 0xdb, 0x5e, 0x9a, 0x81, 0xdb, 0x62,
	0xec, 0x1a, 0xcb, 0xdd, 0xc6, 0xe4, 0x7b, 0x1b, 0x3e, 0x61, 0x3e, 0xac, 0x07, 0xed, 0xaf, 0x1b,
	0x36, 0x34, 0x5a, 0x66, 0x26, 0x8f, 0x5a, 0xf1, 0x09, 0x3b, 0xc2, 0xba, 0x77, 0xba, 0x90, 0x46,
	0x51, 0x8d, 0x5e, 0x3c, 0xf6, 0x79, 0xa4, 0xcc, 0x27, 0xbe, 0x75, 0x3a, 0x7c, 0xfa, 0x59, 0x41,
	0xf9, 0x2d, 0x76, 0x13, 0x59, 0x79, 0xd7, 0x0e, 0x91, 0xc8, 0xcb, 0x44, 0x85, 0x0d, 0xea, 0x6f,
	0xd4, 0x5e, 0xa6, 0x55, 0x32, 0x44, 0xf3, 0xf1, 0xaf, 0x2a, 0x8e, 0x17, 0x26, 0x1d, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x67, 0x4a, 0x62, 0x96, 0x98, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendClient interface {
	NewPlayer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerID, error)
	// rpc PerformAction(ActionRequest) returns (ActionRequest) {}
	WorldRequest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WorldResponse, error)
	EntityStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Backend_EntityStreamClient, error)
}

type backendClient struct {
	cc *grpc.ClientConn
}

func NewBackendClient(cc *grpc.ClientConn) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) NewPlayer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerID, error) {
	out := new(PlayerID)
	err := c.cc.Invoke(ctx, "/web.Backend/NewPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) WorldRequest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WorldResponse, error) {
	out := new(WorldResponse)
	err := c.cc.Invoke(ctx, "/web.Backend/WorldRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) EntityStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Backend_EntityStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Backend_serviceDesc.Streams[0], "/web.Backend/EntityStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &backendEntityStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backend_EntityStreamClient interface {
	Recv() (*EntityResponse, error)
	grpc.ClientStream
}

type backendEntityStreamClient struct {
	grpc.ClientStream
}

func (x *backendEntityStreamClient) Recv() (*EntityResponse, error) {
	m := new(EntityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BackendServer is the server API for Backend service.
type BackendServer interface {
	NewPlayer(context.Context, *Empty) (*PlayerID, error)
	// rpc PerformAction(ActionRequest) returns (ActionRequest) {}
	WorldRequest(context.Context, *Empty) (*WorldResponse, error)
	EntityStream(*Empty, Backend_EntityStreamServer) error
}

func RegisterBackendServer(s *grpc.Server, srv BackendServer) {
	s.RegisterService(&_Backend_serviceDesc, srv)
}

func _Backend_NewPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).NewPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Backend/NewPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).NewPlayer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_WorldRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).WorldRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Backend/WorldRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).WorldRequest(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_EntityStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServer).EntityStream(m, &backendEntityStreamServer{stream})
}

type Backend_EntityStreamServer interface {
	Send(*EntityResponse) error
	grpc.ServerStream
}

type backendEntityStreamServer struct {
	grpc.ServerStream
}

func (x *backendEntityStreamServer) Send(m *EntityResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Backend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "web.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPlayer",
			Handler:    _Backend_NewPlayer_Handler,
		},
		{
			MethodName: "WorldRequest",
			Handler:    _Backend_WorldRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EntityStream",
			Handler:       _Backend_EntityStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/remote.proto",
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/remote.proto

// Web exposes a backend server over gRPC.

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/johanbrandhorst/protobuf/proto"
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

type ActionRequest_Action int32

const (
	ActionRequest_Move ActionRequest_Action = 0
)

var ActionRequest_Action_name = map[int32]string{
	0: "Move",
}

var ActionRequest_Action_value = map[string]int32{
	"Move": 0,
}

func (x ActionRequest_Action) String() string {
	return proto.EnumName(ActionRequest_Action_name, int32(x))
}

func (ActionRequest_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{2, 0}
}

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

type Entity struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	X                    int32    `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32    `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
	Theta                int32    `protobuf:"varint,4,opt,name=Theta,proto3" json:"Theta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{1}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Entity) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Entity) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Entity) GetTheta() int32 {
	if m != nil {
		return m.Theta
	}
	return 0
}

type ActionRequest struct {
	Entity               *Entity              `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Action               ActionRequest_Action `protobuf:"varint,2,opt,name=action,proto3,enum=web.ActionRequest_Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{2}
}

func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ActionRequest) GetAction() ActionRequest_Action {
	if m != nil {
		return m.Action
	}
	return ActionRequest_Move
}

type ActionResponse struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionResponse) Reset()         { *m = ActionResponse{} }
func (m *ActionResponse) String() string { return proto.CompactTextString(m) }
func (*ActionResponse) ProtoMessage()    {}
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{3}
}

func (m *ActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionResponse.Unmarshal(m, b)
}
func (m *ActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionResponse.Marshal(b, m, deterministic)
}
func (m *ActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionResponse.Merge(m, src)
}
func (m *ActionResponse) XXX_Size() int {
	return xxx_messageInfo_ActionResponse.Size(m)
}
func (m *ActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActionResponse proto.InternalMessageInfo

func (m *ActionResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
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
	return fileDescriptor_b59cb7d8fa836ac0, []int{4}
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
	return fileDescriptor_b59cb7d8fa836ac0, []int{5}
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
	return fileDescriptor_b59cb7d8fa836ac0, []int{6}
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
	proto.RegisterEnum("web.ActionRequest_Action", ActionRequest_Action_name, ActionRequest_Action_value)
	proto.RegisterType((*PlayerID)(nil), "web.PlayerID")
	proto.RegisterType((*Entity)(nil), "web.Entity")
	proto.RegisterType((*ActionRequest)(nil), "web.ActionRequest")
	proto.RegisterType((*ActionResponse)(nil), "web.ActionResponse")
	proto.RegisterType((*Empty)(nil), "web.Empty")
	proto.RegisterType((*WorldResponse)(nil), "web.WorldResponse")
	proto.RegisterType((*EntityResponse)(nil), "web.EntityResponse")
}

func init() { proto.RegisterFile("proto/remote.proto", fileDescriptor_b59cb7d8fa836ac0) }

var fileDescriptor_b59cb7d8fa836ac0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xd3, 0xc6, 0x69, 0xa7, 0x49, 0x84, 0xb6, 0x15, 0x32, 0x39, 0x55, 0x8b, 0x84, 0x22,
	0x0e, 0x71, 0x49, 0x85, 0x84, 0x38, 0x41, 0x95, 0x1e, 0x82, 0x54, 0x54, 0xb9, 0x48, 0xb4, 0xdc,
	0xd6, 0xf1, 0x34, 0xeb, 0x36, 0xf6, 0x9a, 0xf5, 0xa4, 0x96, 0xc5, 0x7f, 0xf2, 0x31, 0x9c, 0x50,
	0x76, 0xd7, 0x51, 0x02, 0x07, 0x38, 0xd9, 0x6f, 0x76, 0xde, 0x7b, 0xb3, 0x6f, 0x16, 0x58, 0xa1,
	0x15, 0xa9, 0x50, 0x63, 0xa6, 0x08, 0xc7, 0x06, 0xb0, 0xbd, 0x0a, 0xe3, 0xe1, 0xbb, 0x45, 0x4a,
	0x72, 0x15, 0x8f, 0xe7, 0x2a, 0x0b, 0x1f, 0x94, 0x14, 0x79, 0xac, 0x45, 0x9e, 0x48, 0xa5, 0x4b,
	0x0a, 0x4d, 0x5b, 0xbc, 0xba, 0xb7, 0x3f, 0xe1, 0x42, 0x15, 0x12, 0xf5, 0x43, 0x69, 0xe9, 0x7c,
	0x08, 0x07, 0xd7, 0x4b, 0x51, 0xa3, 0x9e, 0x4d, 0xd9, 0x00, 0xda, 0xb3, 0x69, 0xe0, 0x9d, 0x7a,
	0xa3, 0xc3, 0xa8, 0x3d, 0x9b, 0xf2, 0x4f, 0xe0, 0x5f, 0xe6, 0x94, 0x52, 0xfd, 0xe7, 0x09, 0xeb,
	0x81, 0x77, 0x1b, 0xb4, 0x4f, 0xbd, 0x51, 0x27, 0xf2, 0x6e, 0xd7, 0xe8, 0x2e, 0xd8, 0xb3, 0xe8,
	0x8e, 0x9d, 0x40, 0xe7, 0x8b, 0x44, 0x12, 0xc1, 0xbe, 0xa9, 0x58, 0xc0, 0x7f, 0x40, 0xff, 0xe3,
	0x9c, 0x52, 0x95, 0x47, 0xf8, 0x7d, 0x85, 0x25, 0xb1, 0x97, 0xe0, 0xa3, 0x11, 0x37, 0xb2, 0x47,
	0x93, 0xa3, 0x71, 0x85, 0xf1, 0xd8, 0xfa, 0x45, 0xee, 0x88, 0xbd, 0x01, 0x5f, 0x18, 0x96, 0x31,
	0x1b, 0x4c, 0x5e, 0x98, 0xa6, 0x1d, 0xa1, 0x06, 0xb9, 0x46, 0xce, 0xc0, 0xb7, 0x15, 0x76, 0x00,
	0xfb, 0x57, 0xea, 0x09, 0x9f, 0xb5, 0xf8, 0x5b, 0x18, 0x34, 0x9c, 0xb2, 0x50, 0x79, 0x89, 0xff,
	0xe5, 0xce, 0xbb, 0xd0, 0xb9, 0xcc, 0x0a, 0xaa, 0xf9, 0x0d, 0xf4, 0xbf, 0x2a, 0xbd, 0x4c, 0x36,
	0xf4, 0x13, 0xe8, 0x50, 0xba, 0xc4, 0xd2, 0xb0, 0x7b, 0x91, 0x05, 0xeb, 0x6a, 0x95, 0x26, 0x24,
	0x5d, 0x32, 0x16, 0xb0, 0xe7, 0xe0, 0x4b, 0x4c, 0x17, 0x92, 0x5c, 0x44, 0x0e, 0xf1, 0xd7, 0x30,
	0x70, 0x7e, 0x8d, 0x6a, 0x00, 0xdd, 0x42, 0xd4, 0x4b, 0x25, 0x12, 0xa7, 0xdb, 0xc0, 0xc9, 0x4f,
	0x0f, 0xba, 0x17, 0x62, 0xfe, 0x88, 0x79, 0xc2, 0x46, 0x70, 0xf8, 0x19, 0x2b, 0xbb, 0x34, 0x06,
	0x76, 0xee, 0xf5, 0x94, 0xc3, 0xbe, 0xf9, 0x6f, 0xb6, 0xc9, 0x5b, 0xec, 0x3d, 0xf4, 0xaf, 0x51,
	0xdf, 0x2b, 0x9d, 0xb9, 0x44, 0xd8, 0xdf, 0xf1, 0x0d, 0x8f, 0x77, 0x6a, 0x76, 0x12, 0xde, 0x62,
	0x67, 0xd0, 0x73, 0x57, 0xb6, 0xeb, 0xda, 0x36, 0xb2, 0x32, 0x3b, 0x89, 0xf0, 0x16, 0x3b, 0x87,
	0x9e, 0xbd, 0xcf, 0x0d, 0x69, 0x14, 0xd9, 0x0e, 0xe3, 0x78, 0x3b, 0xde, 0x0d, 0xe5, 0xcc, 0xbb,
	0xb8, 0xfa, 0xf5, 0xe1, 0xd5, 0xd6, 0xdb, 0x7d, 0xac, 0x11, 0x89, 0xc2, 0x85, 0xc8, 0xb0, 0x44,
	0xfd, 0x84, 0xda, 0x3d, 0xda, 0x0a, 0xe3, 0x6f, 0xfc, 0xdf, 0x7d, 0xb1, 0x6f, 0x3e, 0xe7, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x5f, 0x1f, 0x7f, 0x29, 0x03, 0x00, 0x00,
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
	PerformAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
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

func (c *backendClient) PerformAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/web.Backend/PerformAction", in, out, opts...)
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
	PerformAction(context.Context, *ActionRequest) (*ActionResponse, error)
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

func _Backend_PerformAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).PerformAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Backend/PerformAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).PerformAction(ctx, req.(*ActionRequest))
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
			MethodName: "PerformAction",
			Handler:    _Backend_PerformAction_Handler,
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

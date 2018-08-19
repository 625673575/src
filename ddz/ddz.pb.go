// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ddz.proto

package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type IDENTITY int32

const (
	IDENTITY_LORD    IDENTITY = 0
	IDENTITY_FARMER1 IDENTITY = 1
	IDENTITY_FARMER2 IDENTITY = 2
)

var IDENTITY_name = map[int32]string{
	0: "LORD",
	1: "FARMER1",
	2: "FARMER2",
}
var IDENTITY_value = map[string]int32{
	"LORD":    0,
	"FARMER1": 1,
	"FARMER2": 2,
}

func (x IDENTITY) String() string {
	return proto.EnumName(IDENTITY_name, int32(x))
}
func (IDENTITY) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ddz_1d78e3b4767c5460, []int{0}
}

type DEALCARD_TYPE int32

const (
	DEALCARD_TYPE_OPTIMIZED DEALCARD_TYPE = 0
	DEALCARD_TYPE_SMOOTH    DEALCARD_TYPE = 1
)

var DEALCARD_TYPE_name = map[int32]string{
	0: "OPTIMIZED",
	1: "SMOOTH",
}
var DEALCARD_TYPE_value = map[string]int32{
	"OPTIMIZED": 0,
	"SMOOTH":    1,
}

func (x DEALCARD_TYPE) String() string {
	return proto.EnumName(DEALCARD_TYPE_name, int32(x))
}
func (DEALCARD_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ddz_1d78e3b4767c5460, []int{1}
}

type BoolReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolReply) Reset()         { *m = BoolReply{} }
func (m *BoolReply) String() string { return proto.CompactTextString(m) }
func (*BoolReply) ProtoMessage()    {}
func (*BoolReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddz_1d78e3b4767c5460, []int{0}
}
func (m *BoolReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolReply.Unmarshal(m, b)
}
func (m *BoolReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolReply.Marshal(b, m, deterministic)
}
func (dst *BoolReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolReply.Merge(dst, src)
}
func (m *BoolReply) XXX_Size() int {
	return xxx_messageInfo_BoolReply.Size(m)
}
func (m *BoolReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolReply.DiscardUnknown(m)
}

var xxx_messageInfo_BoolReply proto.InternalMessageInfo

func (m *BoolReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type TrustShipRequest struct {
	PlayerIdentity       int32    `protobuf:"varint,1,opt,name=player_identity,json=playerIdentity,proto3" json:"player_identity,omitempty"`
	PlayerHandcard       []byte   `protobuf:"bytes,2,opt,name=player_handcard,json=playerHandcard,proto3" json:"player_handcard,omitempty"`
	LastIdentity         int32    `protobuf:"varint,3,opt,name=last_identity,json=lastIdentity,proto3" json:"last_identity,omitempty"`
	LastPlaycard         []byte   `protobuf:"bytes,4,opt,name=last_playcard,json=lastPlaycard,proto3" json:"last_playcard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrustShipRequest) Reset()         { *m = TrustShipRequest{} }
func (m *TrustShipRequest) String() string { return proto.CompactTextString(m) }
func (*TrustShipRequest) ProtoMessage()    {}
func (*TrustShipRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddz_1d78e3b4767c5460, []int{1}
}
func (m *TrustShipRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrustShipRequest.Unmarshal(m, b)
}
func (m *TrustShipRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrustShipRequest.Marshal(b, m, deterministic)
}
func (dst *TrustShipRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrustShipRequest.Merge(dst, src)
}
func (m *TrustShipRequest) XXX_Size() int {
	return xxx_messageInfo_TrustShipRequest.Size(m)
}
func (m *TrustShipRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrustShipRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrustShipRequest proto.InternalMessageInfo

func (m *TrustShipRequest) GetPlayerIdentity() int32 {
	if m != nil {
		return m.PlayerIdentity
	}
	return 0
}

func (m *TrustShipRequest) GetPlayerHandcard() []byte {
	if m != nil {
		return m.PlayerHandcard
	}
	return nil
}

func (m *TrustShipRequest) GetLastIdentity() int32 {
	if m != nil {
		return m.LastIdentity
	}
	return 0
}

func (m *TrustShipRequest) GetLastPlaycard() []byte {
	if m != nil {
		return m.LastPlaycard
	}
	return nil
}

type RobotRequest struct {
	Playeridentity       int32    `protobuf:"varint,1,opt,name=playeridentity,proto3" json:"playeridentity,omitempty"`
	LordHandcard         []byte   `protobuf:"bytes,2,opt,name=lord_handcard,json=lordHandcard,proto3" json:"lord_handcard,omitempty"`
	Farmer1Handcard      []byte   `protobuf:"bytes,3,opt,name=farmer1_handcard,json=farmer1Handcard,proto3" json:"farmer1_handcard,omitempty"`
	Farmer2Handcard      []byte   `protobuf:"bytes,4,opt,name=farmer2_handcard,json=farmer2Handcard,proto3" json:"farmer2_handcard,omitempty"`
	LastIdentity         int32    `protobuf:"varint,5,opt,name=last_identity,json=lastIdentity,proto3" json:"last_identity,omitempty"`
	LastPlaycard         []byte   `protobuf:"bytes,6,opt,name=last_playcard,json=lastPlaycard,proto3" json:"last_playcard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RobotRequest) Reset()         { *m = RobotRequest{} }
func (m *RobotRequest) String() string { return proto.CompactTextString(m) }
func (*RobotRequest) ProtoMessage()    {}
func (*RobotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddz_1d78e3b4767c5460, []int{2}
}
func (m *RobotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotRequest.Unmarshal(m, b)
}
func (m *RobotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotRequest.Marshal(b, m, deterministic)
}
func (dst *RobotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotRequest.Merge(dst, src)
}
func (m *RobotRequest) XXX_Size() int {
	return xxx_messageInfo_RobotRequest.Size(m)
}
func (m *RobotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RobotRequest proto.InternalMessageInfo

func (m *RobotRequest) GetPlayeridentity() int32 {
	if m != nil {
		return m.Playeridentity
	}
	return 0
}

func (m *RobotRequest) GetLordHandcard() []byte {
	if m != nil {
		return m.LordHandcard
	}
	return nil
}

func (m *RobotRequest) GetFarmer1Handcard() []byte {
	if m != nil {
		return m.Farmer1Handcard
	}
	return nil
}

func (m *RobotRequest) GetFarmer2Handcard() []byte {
	if m != nil {
		return m.Farmer2Handcard
	}
	return nil
}

func (m *RobotRequest) GetLastIdentity() int32 {
	if m != nil {
		return m.LastIdentity
	}
	return 0
}

func (m *RobotRequest) GetLastPlaycard() []byte {
	if m != nil {
		return m.LastPlaycard
	}
	return nil
}

type PlayReply struct {
	Handcard             []byte   `protobuf:"bytes,1,opt,name=handcard,proto3" json:"handcard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayReply) Reset()         { *m = PlayReply{} }
func (m *PlayReply) String() string { return proto.CompactTextString(m) }
func (*PlayReply) ProtoMessage()    {}
func (*PlayReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddz_1d78e3b4767c5460, []int{3}
}
func (m *PlayReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayReply.Unmarshal(m, b)
}
func (m *PlayReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayReply.Marshal(b, m, deterministic)
}
func (dst *PlayReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayReply.Merge(dst, src)
}
func (m *PlayReply) XXX_Size() int {
	return xxx_messageInfo_PlayReply.Size(m)
}
func (m *PlayReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayReply.DiscardUnknown(m)
}

var xxx_messageInfo_PlayReply proto.InternalMessageInfo

func (m *PlayReply) GetHandcard() []byte {
	if m != nil {
		return m.Handcard
	}
	return nil
}

type DealCardRequest struct {
	Type                 DEALCARD_TYPE `protobuf:"varint,1,opt,name=type,proto3,enum=ddz.DEALCARD_TYPE" json:"type,omitempty"`
	Params               []byte        `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DealCardRequest) Reset()         { *m = DealCardRequest{} }
func (m *DealCardRequest) String() string { return proto.CompactTextString(m) }
func (*DealCardRequest) ProtoMessage()    {}
func (*DealCardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddz_1d78e3b4767c5460, []int{4}
}
func (m *DealCardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealCardRequest.Unmarshal(m, b)
}
func (m *DealCardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealCardRequest.Marshal(b, m, deterministic)
}
func (dst *DealCardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealCardRequest.Merge(dst, src)
}
func (m *DealCardRequest) XXX_Size() int {
	return xxx_messageInfo_DealCardRequest.Size(m)
}
func (m *DealCardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DealCardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DealCardRequest proto.InternalMessageInfo

func (m *DealCardRequest) GetType() DEALCARD_TYPE {
	if m != nil {
		return m.Type
	}
	return DEALCARD_TYPE_OPTIMIZED
}

func (m *DealCardRequest) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

type DealCardReply struct {
	Player0              []byte   `protobuf:"bytes,1,opt,name=player0,proto3" json:"player0,omitempty"`
	Player1              []byte   `protobuf:"bytes,2,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2              []byte   `protobuf:"bytes,3,opt,name=player2,proto3" json:"player2,omitempty"`
	Extra                []byte   `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DealCardReply) Reset()         { *m = DealCardReply{} }
func (m *DealCardReply) String() string { return proto.CompactTextString(m) }
func (*DealCardReply) ProtoMessage()    {}
func (*DealCardReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddz_1d78e3b4767c5460, []int{5}
}
func (m *DealCardReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealCardReply.Unmarshal(m, b)
}
func (m *DealCardReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealCardReply.Marshal(b, m, deterministic)
}
func (dst *DealCardReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealCardReply.Merge(dst, src)
}
func (m *DealCardReply) XXX_Size() int {
	return xxx_messageInfo_DealCardReply.Size(m)
}
func (m *DealCardReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DealCardReply.DiscardUnknown(m)
}

var xxx_messageInfo_DealCardReply proto.InternalMessageInfo

func (m *DealCardReply) GetPlayer0() []byte {
	if m != nil {
		return m.Player0
	}
	return nil
}

func (m *DealCardReply) GetPlayer1() []byte {
	if m != nil {
		return m.Player1
	}
	return nil
}

func (m *DealCardReply) GetPlayer2() []byte {
	if m != nil {
		return m.Player2
	}
	return nil
}

func (m *DealCardReply) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

func init() {
	proto.RegisterType((*BoolReply)(nil), "ddz.BoolReply")
	proto.RegisterType((*TrustShipRequest)(nil), "ddz.TrustShipRequest")
	proto.RegisterType((*RobotRequest)(nil), "ddz.RobotRequest")
	proto.RegisterType((*PlayReply)(nil), "ddz.PlayReply")
	proto.RegisterType((*DealCardRequest)(nil), "ddz.DealCardRequest")
	proto.RegisterType((*DealCardReply)(nil), "ddz.DealCardReply")
	proto.RegisterEnum("ddz.IDENTITY", IDENTITY_name, IDENTITY_value)
	proto.RegisterEnum("ddz.DEALCARD_TYPE", DEALCARD_TYPE_name, DEALCARD_TYPE_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DealCardServiceClient is the client API for DealCardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DealCardServiceClient interface {
	GetCard(ctx context.Context, in *DealCardRequest, opts ...grpc.CallOption) (*DealCardReply, error)
}

type dealCardServiceClient struct {
	cc *grpc.ClientConn
}

func NewDealCardServiceClient(cc *grpc.ClientConn) DealCardServiceClient {
	return &dealCardServiceClient{cc}
}

func (c *dealCardServiceClient) GetCard(ctx context.Context, in *DealCardRequest, opts ...grpc.CallOption) (*DealCardReply, error) {
	out := new(DealCardReply)
	err := c.cc.Invoke(ctx, "/ddz.DealCardService/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DealCardServiceServer is the server API for DealCardService service.
type DealCardServiceServer interface {
	GetCard(context.Context, *DealCardRequest) (*DealCardReply, error)
}

func RegisterDealCardServiceServer(s *grpc.Server, srv DealCardServiceServer) {
	s.RegisterService(&_DealCardService_serviceDesc, srv)
}

func _DealCardService_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DealCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealCardServiceServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddz.DealCardService/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealCardServiceServer).GetCard(ctx, req.(*DealCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DealCardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddz.DealCardService",
	HandlerType: (*DealCardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCard",
			Handler:    _DealCardService_GetCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ddz.proto",
}

// TrustshipServiceClient is the client API for TrustshipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrustshipServiceClient interface {
	Ship(ctx context.Context, in *TrustShipRequest, opts ...grpc.CallOption) (*PlayReply, error)
}

type trustshipServiceClient struct {
	cc *grpc.ClientConn
}

func NewTrustshipServiceClient(cc *grpc.ClientConn) TrustshipServiceClient {
	return &trustshipServiceClient{cc}
}

func (c *trustshipServiceClient) Ship(ctx context.Context, in *TrustShipRequest, opts ...grpc.CallOption) (*PlayReply, error) {
	out := new(PlayReply)
	err := c.cc.Invoke(ctx, "/ddz.TrustshipService/Ship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrustshipServiceServer is the server API for TrustshipService service.
type TrustshipServiceServer interface {
	Ship(context.Context, *TrustShipRequest) (*PlayReply, error)
}

func RegisterTrustshipServiceServer(s *grpc.Server, srv TrustshipServiceServer) {
	s.RegisterService(&_TrustshipService_serviceDesc, srv)
}

func _TrustshipService_Ship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustShipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustshipServiceServer).Ship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddz.TrustshipService/Ship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustshipServiceServer).Ship(ctx, req.(*TrustShipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrustshipService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddz.TrustshipService",
	HandlerType: (*TrustshipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ship",
			Handler:    _TrustshipService_Ship_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ddz.proto",
}

// RobotServiceClient is the client API for RobotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RobotServiceClient interface {
	Play(ctx context.Context, in *RobotRequest, opts ...grpc.CallOption) (*PlayReply, error)
}

type robotServiceClient struct {
	cc *grpc.ClientConn
}

func NewRobotServiceClient(cc *grpc.ClientConn) RobotServiceClient {
	return &robotServiceClient{cc}
}

func (c *robotServiceClient) Play(ctx context.Context, in *RobotRequest, opts ...grpc.CallOption) (*PlayReply, error) {
	out := new(PlayReply)
	err := c.cc.Invoke(ctx, "/ddz.RobotService/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotServiceServer is the server API for RobotService service.
type RobotServiceServer interface {
	Play(context.Context, *RobotRequest) (*PlayReply, error)
}

func RegisterRobotServiceServer(s *grpc.Server, srv RobotServiceServer) {
	s.RegisterService(&_RobotService_serviceDesc, srv)
}

func _RobotService_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddz.RobotService/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).Play(ctx, req.(*RobotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RobotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddz.RobotService",
	HandlerType: (*RobotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _RobotService_Play_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ddz.proto",
}

func init() { proto.RegisterFile("ddz.proto", fileDescriptor_ddz_1d78e3b4767c5460) }

var fileDescriptor_ddz_1d78e3b4767c5460 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x6d, 0xd6, 0xf4, 0xdf, 0xfd, 0xf5, 0x4f, 0x7e, 0xd6, 0x40, 0x55, 0x9f, 0xa6, 0x4c, 0x62,
	0xa3, 0x48, 0x85, 0x06, 0xf1, 0xc4, 0x53, 0x69, 0x02, 0x8d, 0xb4, 0xd2, 0xe2, 0xe6, 0x65, 0xbc,
	0x4c, 0xde, 0x62, 0xd4, 0x48, 0xd9, 0x12, 0x1c, 0x17, 0x91, 0x7d, 0x28, 0xbe, 0x20, 0x2f, 0xc8,
	0x8e, 0x13, 0x77, 0xdd, 0x90, 0x78, 0xcb, 0x39, 0x3e, 0xf7, 0xd8, 0xf7, 0xe4, 0x5e, 0xe8, 0x84,
	0xe1, 0xfd, 0x24, 0x65, 0x09, 0x4f, 0x50, 0x3d, 0x0c, 0xef, 0xed, 0x53, 0xe8, 0x7c, 0x48, 0x92,
	0x18, 0xd3, 0x34, 0xce, 0xd1, 0x73, 0x68, 0x32, 0x9a, 0xed, 0x62, 0x3e, 0x34, 0x4e, 0x8c, 0xf3,
	0x36, 0x56, 0xc8, 0xfe, 0x65, 0x80, 0x15, 0xb0, 0x5d, 0xc6, 0x37, 0xdb, 0x28, 0xc5, 0xf4, 0xfb,
	0x8e, 0x66, 0x1c, 0x9d, 0xc1, 0x20, 0x8d, 0x49, 0x4e, 0xd9, 0x55, 0x14, 0xd2, 0x3b, 0x1e, 0xf1,
	0x5c, 0x56, 0x35, 0x70, 0xbf, 0xa0, 0x7d, 0xc5, 0xee, 0x09, 0xb7, 0xe4, 0x2e, 0xbc, 0x21, 0x2c,
	0x1c, 0x1e, 0x9d, 0x18, 0xe7, 0xdd, 0x52, 0xb8, 0x50, 0x2c, 0x3a, 0x85, 0x5e, 0x4c, 0x32, 0xae,
	0xfd, 0xea, 0xd2, 0xaf, 0x2b, 0xc8, 0xca, 0xad, 0x14, 0x89, 0x5a, 0xe9, 0x65, 0x4a, 0x2f, 0x29,
	0x5a, 0x2b, 0xce, 0xfe, 0x6d, 0x40, 0x17, 0x27, 0xd7, 0x09, 0x2f, 0x1f, 0xfb, 0x02, 0xd4, 0x65,
	0x4f, 0xbf, 0x35, 0xda, 0x77, 0x4f, 0x58, 0x78, 0xf8, 0xd2, 0xae, 0x20, 0xab, 0x77, 0xbe, 0x04,
	0xeb, 0x1b, 0x61, 0xb7, 0x94, 0x4d, 0xb5, 0xae, 0x2e, 0x75, 0x03, 0xc5, 0x3f, 0x96, 0x3a, 0x5a,
	0x6a, 0xee, 0x4b, 0x9d, 0xbf, 0x77, 0xdf, 0xf8, 0x97, 0xee, 0x9b, 0x4f, 0x74, 0x7f, 0x06, 0x1d,
	0xf1, 0x5d, 0xfc, 0xd3, 0x11, 0xb4, 0xab, 0x9b, 0x0d, 0x29, 0xae, 0xb0, 0xfd, 0x05, 0x06, 0x2e,
	0x25, 0xf1, 0x9c, 0xb0, 0x50, 0x07, 0x65, 0xf2, 0x3c, 0xa5, 0x52, 0xda, 0x77, 0xd0, 0x44, 0x8c,
	0x8b, 0xeb, 0xcd, 0x2e, 0xe6, 0x33, 0xec, 0x5e, 0x05, 0x97, 0x6b, 0x0f, 0xcb, 0x73, 0x31, 0x2a,
	0x29, 0x61, 0xe4, 0x36, 0x53, 0x09, 0x29, 0x64, 0x67, 0xd0, 0xd3, 0x96, 0xe2, 0xfe, 0x21, 0xb4,
	0x8a, 0x8c, 0xdf, 0xa8, 0xeb, 0x4b, 0xa8, 0x4f, 0xa6, 0xca, 0xa3, 0x84, 0xfa, 0xc4, 0x51, 0xb9,
	0x96, 0x10, 0x1d, 0x43, 0x83, 0xfe, 0xe4, 0x8c, 0xa8, 0x10, 0x0b, 0x30, 0x9e, 0x40, 0xdb, 0x77,
	0xbd, 0xcf, 0x81, 0x1f, 0x5c, 0xa2, 0x36, 0x98, 0x17, 0x2b, 0xec, 0x5a, 0x35, 0xf4, 0x1f, 0xb4,
	0x3e, 0xce, 0xf0, 0xd2, 0xc3, 0x53, 0xcb, 0xd0, 0xc0, 0xb1, 0x8e, 0xc6, 0x63, 0xe8, 0x3d, 0xe8,
	0x09, 0xf5, 0xa0, 0xb3, 0x5a, 0x07, 0xfe, 0xd2, 0xff, 0xea, 0x89, 0x4a, 0x80, 0xe6, 0x66, 0xb9,
	0x5a, 0x05, 0x0b, 0xcb, 0x70, 0x16, 0x3a, 0xa3, 0x0d, 0x65, 0x3f, 0xa2, 0x1b, 0x8a, 0xde, 0x41,
	0xeb, 0x13, 0xe5, 0x82, 0x41, 0xc7, 0x45, 0x40, 0x0f, 0x43, 0x1c, 0xa1, 0x03, 0x36, 0x8d, 0x73,
	0xbb, 0xe6, 0xcc, 0xd5, 0x12, 0x65, 0xdb, 0x28, 0x2d, 0xad, 0x5e, 0x83, 0x29, 0x76, 0x0a, 0x3d,
	0x93, 0x15, 0x87, 0x3b, 0x36, 0xea, 0x4b, 0xba, 0xfa, 0x99, 0x76, 0xcd, 0x79, 0xaf, 0x06, 0xbb,
	0x34, 0x78, 0x05, 0xa6, 0x38, 0x46, 0xff, 0x4b, 0xe5, 0xfe, 0xcc, 0x3f, 0x2e, 0xbe, 0x6e, 0xca,
	0xc5, 0x7f, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x87, 0xdf, 0x6e, 0x52, 0x05, 0x04, 0x00, 0x00,
}
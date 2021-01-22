// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: topup/v1beta/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the topup module's genesis state.
type MsgTopup struct {
	FromAddress string                                 `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	User        string                                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Fee         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"varint,3,opt,name=fee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Int" json:"fee,omitempty"`
	TxHash      string                                 `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" yaml:"tx_hash"`
	LogIndex    uint64                                 `protobuf:"varint,5,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty" yaml:"log_index"`
	BlockNumber uint64                                 `protobuf:"varint,6,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty" yaml:"block_number"`
}

func (m *MsgTopup) Reset()         { *m = MsgTopup{} }
func (m *MsgTopup) String() string { return proto.CompactTextString(m) }
func (*MsgTopup) ProtoMessage()    {}
func (*MsgTopup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d59e70ac70f1c423, []int{0}
}
func (m *MsgTopup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTopup.Unmarshal(m, b)
}
func (m *MsgTopup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTopup.Marshal(b, m, deterministic)
}
func (m *MsgTopup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTopup.Merge(m, src)
}
func (m *MsgTopup) XXX_Size() int {
	return xxx_messageInfo_MsgTopup.Size(m)
}
func (m *MsgTopup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTopup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTopup proto.InternalMessageInfo

type MsgTopupResponse struct {
}

func (m *MsgTopupResponse) Reset()         { *m = MsgTopupResponse{} }
func (m *MsgTopupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTopupResponse) ProtoMessage()    {}
func (*MsgTopupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d59e70ac70f1c423, []int{1}
}
func (m *MsgTopupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTopupResponse.Unmarshal(m, b)
}
func (m *MsgTopupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTopupResponse.Marshal(b, m, deterministic)
}
func (m *MsgTopupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTopupResponse.Merge(m, src)
}
func (m *MsgTopupResponse) XXX_Size() int {
	return xxx_messageInfo_MsgTopupResponse.Size(m)
}
func (m *MsgTopupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTopupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTopupResponse proto.InternalMessageInfo

type MsgWithdrawFee struct {
	UserAddress string                                 `protobuf:"bytes,1,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty" yaml:"user_address"`
	Amount      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"varint,2,opt,name=amount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount,omitempty"`
}

func (m *MsgWithdrawFee) Reset()         { *m = MsgWithdrawFee{} }
func (m *MsgWithdrawFee) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawFee) ProtoMessage()    {}
func (*MsgWithdrawFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_d59e70ac70f1c423, []int{2}
}
func (m *MsgWithdrawFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithdrawFee.Unmarshal(m, b)
}
func (m *MsgWithdrawFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithdrawFee.Marshal(b, m, deterministic)
}
func (m *MsgWithdrawFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawFee.Merge(m, src)
}
func (m *MsgWithdrawFee) XXX_Size() int {
	return xxx_messageInfo_MsgWithdrawFee.Size(m)
}
func (m *MsgWithdrawFee) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawFee.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawFee proto.InternalMessageInfo

type MsgWithdrawFeeResponse struct {
}

func (m *MsgWithdrawFeeResponse) Reset()         { *m = MsgWithdrawFeeResponse{} }
func (m *MsgWithdrawFeeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawFeeResponse) ProtoMessage()    {}
func (*MsgWithdrawFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d59e70ac70f1c423, []int{3}
}
func (m *MsgWithdrawFeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWithdrawFeeResponse.Unmarshal(m, b)
}
func (m *MsgWithdrawFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWithdrawFeeResponse.Marshal(b, m, deterministic)
}
func (m *MsgWithdrawFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawFeeResponse.Merge(m, src)
}
func (m *MsgWithdrawFeeResponse) XXX_Size() int {
	return xxx_messageInfo_MsgWithdrawFeeResponse.Size(m)
}
func (m *MsgWithdrawFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawFeeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTopup)(nil), "heimdall.topup.v1beta1.MsgTopup")
	proto.RegisterType((*MsgTopupResponse)(nil), "heimdall.topup.v1beta1.MsgTopupResponse")
	proto.RegisterType((*MsgWithdrawFee)(nil), "heimdall.topup.v1beta1.MsgWithdrawFee")
	proto.RegisterType((*MsgWithdrawFeeResponse)(nil), "heimdall.topup.v1beta1.MsgWithdrawFeeResponse")
}

func init() { proto.RegisterFile("topup/v1beta/msg.proto", fileDescriptor_d59e70ac70f1c423) }

var fileDescriptor_d59e70ac70f1c423 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3f, 0x8f, 0xd3, 0x30,
	0x1c, 0x4d, 0xae, 0xbd, 0x72, 0xe7, 0xa2, 0xd3, 0xc9, 0x9c, 0x4a, 0xd4, 0x21, 0xa9, 0x3c, 0x9c,
	0x2a, 0x10, 0x8e, 0x0a, 0x5b, 0xc5, 0x42, 0x07, 0x74, 0x37, 0x94, 0x21, 0x20, 0x21, 0xb1, 0x54,
	0x4e, 0xe3, 0x73, 0xa2, 0xc6, 0x71, 0x14, 0x3b, 0x5c, 0xef, 0x1b, 0x30, 0x32, 0x31, 0xf3, 0x39,
	0xd8, 0xd8, 0x18, 0x3b, 0x32, 0x55, 0xa8, 0xfd, 0x06, 0x1d, 0x99, 0x90, 0x9d, 0x06, 0x05, 0x89,
	0x7f, 0x37, 0xc5, 0x4f, 0xef, 0xbd, 0xfc, 0xde, 0xef, 0xc9, 0x06, 0x3d, 0x25, 0xf2, 0x32, 0xf7,
	0xdf, 0x8e, 0x42, 0xaa, 0x88, 0xcf, 0x25, 0xc3, 0x79, 0x21, 0x94, 0x80, 0xbd, 0x98, 0x26, 0x3c,
	0x22, 0x69, 0x8a, 0x8d, 0x00, 0x57, 0x82, 0x51, 0xff, 0x8c, 0x09, 0x26, 0x8c, 0xc4, 0xd7, 0xa7,
	0x4a, 0x8d, 0x3e, 0x1d, 0x80, 0xa3, 0xa9, 0x64, 0xaf, 0xb4, 0x14, 0x8e, 0xc1, 0xdd, 0xab, 0x42,
	0xf0, 0x19, 0x89, 0xa2, 0x82, 0x4a, 0xe9, 0xd8, 0x03, 0x7b, 0x78, 0x3c, 0xb9, 0xbf, 0x5b, 0x7b,
	0xf7, 0x6e, 0x08, 0x4f, 0xc7, 0xa8, 0xc9, 0xa2, 0xa0, 0xab, 0xe1, 0xb3, 0x0a, 0x41, 0x08, 0xda,
	0xa5, 0xa4, 0x85, 0x73, 0xa0, 0x3d, 0x81, 0x39, 0xc3, 0xa7, 0xa0, 0x75, 0x45, 0xa9, 0xd3, 0x1a,
	0xd8, 0xc3, 0xf6, 0xe4, 0xc1, 0xf7, 0xb5, 0x77, 0xce, 0x12, 0x15, 0x97, 0x21, 0x9e, 0x0b, 0xee,
	0xcf, 0x85, 0xe4, 0x42, 0xee, 0x3f, 0x8f, 0x64, 0xb4, 0xf0, 0xd5, 0x4d, 0x4e, 0x25, 0xbe, 0xcc,
	0x54, 0xa0, 0x6d, 0xf0, 0x21, 0xb8, 0xa3, 0x96, 0xb3, 0x98, 0xc8, 0xd8, 0x69, 0x9b, 0x20, 0x70,
	0xb7, 0xf6, 0x4e, 0xaa, 0x20, 0x7b, 0x02, 0x05, 0x1d, 0xb5, 0xbc, 0x20, 0x32, 0x86, 0x23, 0x70,
	0x9c, 0x0a, 0x36, 0x4b, 0xb2, 0x88, 0x2e, 0x9d, 0x43, 0x33, 0xf0, 0x6c, 0xb7, 0xf6, 0x4e, 0x2b,
	0xf9, 0x4f, 0x0a, 0x05, 0x47, 0xa9, 0x60, 0x97, 0xfa, 0xa8, 0xb7, 0x0d, 0x53, 0x31, 0x5f, 0xcc,
	0xb2, 0x92, 0x87, 0xb4, 0x70, 0x3a, 0xc6, 0xd5, 0xd8, 0xb6, 0xc9, 0xa2, 0xa0, 0x6b, 0xe0, 0x0b,
	0x83, 0xc6, 0xed, 0x77, 0x1f, 0x3d, 0x0b, 0x41, 0x70, 0x5a, 0x77, 0x17, 0x50, 0x99, 0x8b, 0x4c,
	0x52, 0xf4, 0xc1, 0x06, 0x27, 0x53, 0xc9, 0x5e, 0x27, 0x2a, 0x8e, 0x0a, 0x72, 0xfd, 0x9c, 0x52,
	0x3d, 0x48, 0xd7, 0xf1, 0xe7, 0x5a, 0x9b, 0x2c, 0x0a, 0xba, 0x1a, 0xd6, 0xb5, 0x4e, 0x40, 0x87,
	0x70, 0x51, 0x66, 0xca, 0x14, 0x7b, 0xbb, 0x16, 0xf7, 0xce, 0x7d, 0x58, 0x07, 0xf4, 0x7e, 0xcd,
	0x55, 0x47, 0x7e, 0xfc, 0xd9, 0x06, 0xad, 0xa9, 0x64, 0xf0, 0x25, 0x38, 0xac, 0xee, 0xc1, 0x00,
	0xff, 0xfe, 0x0e, 0xe1, 0x7a, 0xdb, 0xfe, 0xf0, 0x5f, 0x8a, 0xfa, 0xe7, 0x90, 0x82, 0x6e, 0xb3,
	0x8b, 0xf3, 0xbf, 0x18, 0x1b, 0xba, 0x3e, 0xfe, 0x3f, 0x5d, 0x3d, 0x66, 0x72, 0xf1, 0x65, 0xe3,
	0x5a, 0xab, 0x8d, 0x6b, 0x7d, 0xdb, 0xb8, 0xd6, 0xfb, 0xad, 0x6b, 0xad, 0xb6, 0xae, 0xf5, 0x75,
	0xeb, 0x5a, 0x6f, 0x70, 0xa3, 0x2d, 0x4e, 0x54, 0x32, 0xcf, 0xa8, 0xba, 0x16, 0xc5, 0xc2, 0xaf,
	0x07, 0xf8, 0x4b, 0xbf, 0x7a, 0x4a, 0xa6, 0xb9, 0xb0, 0x63, 0x1e, 0xc6, 0x93, 0x1f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf6, 0x8a, 0x7a, 0x01, 0x60, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Topup defines a method to topup validator account with additional tokens.
	Topup(ctx context.Context, in *MsgTopup, opts ...grpc.CallOption) (*MsgTopupResponse, error)
	// WithdrawFee defines a method for validators to withdraw tokens from
	// heimdall.
	WithdrawFee(ctx context.Context, in *MsgWithdrawFee, opts ...grpc.CallOption) (*MsgWithdrawFeeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Topup(ctx context.Context, in *MsgTopup, opts ...grpc.CallOption) (*MsgTopupResponse, error) {
	out := new(MsgTopupResponse)
	err := c.cc.Invoke(ctx, "/heimdall.topup.v1beta1.Msg/Topup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawFee(ctx context.Context, in *MsgWithdrawFee, opts ...grpc.CallOption) (*MsgWithdrawFeeResponse, error) {
	out := new(MsgWithdrawFeeResponse)
	err := c.cc.Invoke(ctx, "/heimdall.topup.v1beta1.Msg/WithdrawFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Topup defines a method to topup validator account with additional tokens.
	Topup(context.Context, *MsgTopup) (*MsgTopupResponse, error)
	// WithdrawFee defines a method for validators to withdraw tokens from
	// heimdall.
	WithdrawFee(context.Context, *MsgWithdrawFee) (*MsgWithdrawFeeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Topup(ctx context.Context, req *MsgTopup) (*MsgTopupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Topup not implemented")
}
func (*UnimplementedMsgServer) WithdrawFee(ctx context.Context, req *MsgWithdrawFee) (*MsgWithdrawFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawFee not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Topup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTopup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Topup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.topup.v1beta1.Msg/Topup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Topup(ctx, req.(*MsgTopup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.topup.v1beta1.Msg/WithdrawFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawFee(ctx, req.(*MsgWithdrawFee))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall.topup.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Topup",
			Handler:    _Msg_Topup_Handler,
		},
		{
			MethodName: "WithdrawFee",
			Handler:    _Msg_WithdrawFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topup/v1beta/msg.proto",
}
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type MsgSubmitOrder struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	FunctionId string `protobuf:"bytes,2,opt,name=functionId,proto3" json:"functionId,omitempty"`
	Fuel       string `protobuf:"bytes,3,opt,name=fuel,proto3" json:"fuel,omitempty"`
}

func (m *MsgSubmitOrder) Reset()         { *m = MsgSubmitOrder{} }
func (m *MsgSubmitOrder) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitOrder) ProtoMessage()    {}
func (*MsgSubmitOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_2966ca2342567dca, []int{0}
}
func (m *MsgSubmitOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitOrder.Merge(m, src)
}
func (m *MsgSubmitOrder) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitOrder proto.InternalMessageInfo

func (m *MsgSubmitOrder) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitOrder) GetFunctionId() string {
	if m != nil {
		return m.FunctionId
	}
	return ""
}

func (m *MsgSubmitOrder) GetFuel() string {
	if m != nil {
		return m.Fuel
	}
	return ""
}

type MsgSubmitOrderResponse struct {
}

func (m *MsgSubmitOrderResponse) Reset()         { *m = MsgSubmitOrderResponse{} }
func (m *MsgSubmitOrderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitOrderResponse) ProtoMessage()    {}
func (*MsgSubmitOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2966ca2342567dca, []int{1}
}
func (m *MsgSubmitOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitOrderResponse.Merge(m, src)
}
func (m *MsgSubmitOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitOrderResponse proto.InternalMessageInfo

type MsgSubmitCompletedOrder struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	OrderIndex string `protobuf:"bytes,2,opt,name=orderIndex,proto3" json:"orderIndex,omitempty"`
	FuelUsed   string `protobuf:"bytes,3,opt,name=fuelUsed,proto3" json:"fuelUsed,omitempty"`
}

func (m *MsgSubmitCompletedOrder) Reset()         { *m = MsgSubmitCompletedOrder{} }
func (m *MsgSubmitCompletedOrder) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitCompletedOrder) ProtoMessage()    {}
func (*MsgSubmitCompletedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_2966ca2342567dca, []int{2}
}
func (m *MsgSubmitCompletedOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitCompletedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitCompletedOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitCompletedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitCompletedOrder.Merge(m, src)
}
func (m *MsgSubmitCompletedOrder) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitCompletedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitCompletedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitCompletedOrder proto.InternalMessageInfo

func (m *MsgSubmitCompletedOrder) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitCompletedOrder) GetOrderIndex() string {
	if m != nil {
		return m.OrderIndex
	}
	return ""
}

func (m *MsgSubmitCompletedOrder) GetFuelUsed() string {
	if m != nil {
		return m.FuelUsed
	}
	return ""
}

type MsgSubmitCompletedOrderResponse struct {
}

func (m *MsgSubmitCompletedOrderResponse) Reset()         { *m = MsgSubmitCompletedOrderResponse{} }
func (m *MsgSubmitCompletedOrderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitCompletedOrderResponse) ProtoMessage()    {}
func (*MsgSubmitCompletedOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2966ca2342567dca, []int{3}
}
func (m *MsgSubmitCompletedOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitCompletedOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitCompletedOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitCompletedOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitCompletedOrderResponse.Merge(m, src)
}
func (m *MsgSubmitCompletedOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitCompletedOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitCompletedOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitCompletedOrderResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitOrder)(nil), "txlabs.blocklesschain.market.MsgSubmitOrder")
	proto.RegisterType((*MsgSubmitOrderResponse)(nil), "txlabs.blocklesschain.market.MsgSubmitOrderResponse")
	proto.RegisterType((*MsgSubmitCompletedOrder)(nil), "txlabs.blocklesschain.market.MsgSubmitCompletedOrder")
	proto.RegisterType((*MsgSubmitCompletedOrderResponse)(nil), "txlabs.blocklesschain.market.MsgSubmitCompletedOrderResponse")
}

func init() { proto.RegisterFile("market/tx.proto", fileDescriptor_2966ca2342567dca) }

var fileDescriptor_2966ca2342567dca = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0xd1, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0xa9, 0xc8,
	0x49, 0x4c, 0x2a, 0xd6, 0x4b, 0xca, 0xc9, 0x4f, 0xce, 0xce, 0x49, 0x2d, 0x2e, 0x4e, 0xce, 0x48,
	0xcc, 0xcc, 0xd3, 0x83, 0x28, 0x53, 0x8a, 0xe3, 0xe2, 0xf3, 0x2d, 0x4e, 0x0f, 0x2e, 0x4d, 0xca,
	0xcd, 0x2c, 0xf1, 0x2f, 0x4a, 0x49, 0x2d, 0x12, 0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c,
	0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xe4, 0xb8, 0xb8, 0xd2,
	0x4a, 0xf3, 0x92, 0x4b, 0x32, 0xf3, 0xf3, 0x3c, 0x53, 0x24, 0x98, 0xc0, 0x92, 0x48, 0x22, 0x42,
	0x42, 0x5c, 0x2c, 0x69, 0xa5, 0xa9, 0x39, 0x12, 0xcc, 0x60, 0x19, 0x30, 0x5b, 0x49, 0x82, 0x4b,
	0x0c, 0xd5, 0xfc, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa5, 0x7c, 0x2e, 0x71, 0xb8,
	0x8c, 0x73, 0x7e, 0x6e, 0x41, 0x4e, 0x6a, 0x49, 0x6a, 0x0a, 0x11, 0x4e, 0xc8, 0x07, 0x29, 0xf1,
	0xcc, 0x4b, 0x49, 0xad, 0x80, 0x39, 0x01, 0x21, 0x22, 0x24, 0xc5, 0xc5, 0x01, 0xb2, 0x36, 0xb4,
	0x38, 0x35, 0x05, 0xea, 0x0c, 0x38, 0x5f, 0x49, 0x91, 0x4b, 0x1e, 0x87, 0x85, 0x30, 0x37, 0x19,
	0xb5, 0x33, 0x71, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x15, 0x72, 0x71, 0x23, 0x07, 0x89, 0x8e, 0x1e,
	0xbe, 0x30, 0xd4, 0x43, 0xf5, 0xa0, 0x94, 0x09, 0x29, 0xaa, 0x61, 0x56, 0x0b, 0xf5, 0x30, 0x72,
	0x89, 0x60, 0x0d, 0x0c, 0x53, 0x22, 0x8d, 0x43, 0xd5, 0x26, 0x65, 0x4b, 0x96, 0x36, 0x98, 0x73,
	0x9c, 0xbc, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x20, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x62, 0x85, 0x3e, 0xdc, 0x0a, 0x5d, 0xb0,
	0x1d, 0xfa, 0x15, 0xfa, 0xb0, 0x44, 0x58, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x4e, 0x88, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xb7, 0x96, 0x80, 0x9b, 0x02, 0x00, 0x00,
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
	SubmitOrder(ctx context.Context, in *MsgSubmitOrder, opts ...grpc.CallOption) (*MsgSubmitOrderResponse, error)
	SubmitCompletedOrder(ctx context.Context, in *MsgSubmitCompletedOrder, opts ...grpc.CallOption) (*MsgSubmitCompletedOrderResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitOrder(ctx context.Context, in *MsgSubmitOrder, opts ...grpc.CallOption) (*MsgSubmitOrderResponse, error) {
	out := new(MsgSubmitOrderResponse)
	err := c.cc.Invoke(ctx, "/txlabs.blocklesschain.market.Msg/SubmitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitCompletedOrder(ctx context.Context, in *MsgSubmitCompletedOrder, opts ...grpc.CallOption) (*MsgSubmitCompletedOrderResponse, error) {
	out := new(MsgSubmitCompletedOrderResponse)
	err := c.cc.Invoke(ctx, "/txlabs.blocklesschain.market.Msg/SubmitCompletedOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitOrder(context.Context, *MsgSubmitOrder) (*MsgSubmitOrderResponse, error)
	SubmitCompletedOrder(context.Context, *MsgSubmitCompletedOrder) (*MsgSubmitCompletedOrderResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitOrder(ctx context.Context, req *MsgSubmitOrder) (*MsgSubmitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (*UnimplementedMsgServer) SubmitCompletedOrder(ctx context.Context, req *MsgSubmitCompletedOrder) (*MsgSubmitCompletedOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCompletedOrder not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/txlabs.blocklesschain.market.Msg/SubmitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitOrder(ctx, req.(*MsgSubmitOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitCompletedOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitCompletedOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitCompletedOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/txlabs.blocklesschain.market.Msg/SubmitCompletedOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitCompletedOrder(ctx, req.(*MsgSubmitCompletedOrder))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "txlabs.blocklesschain.market.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrder",
			Handler:    _Msg_SubmitOrder_Handler,
		},
		{
			MethodName: "SubmitCompletedOrder",
			Handler:    _Msg_SubmitCompletedOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market/tx.proto",
}

func (m *MsgSubmitOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fuel) > 0 {
		i -= len(m.Fuel)
		copy(dAtA[i:], m.Fuel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Fuel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FunctionId) > 0 {
		i -= len(m.FunctionId)
		copy(dAtA[i:], m.FunctionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FunctionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitOrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitCompletedOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitCompletedOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitCompletedOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FuelUsed) > 0 {
		i -= len(m.FuelUsed)
		copy(dAtA[i:], m.FuelUsed)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FuelUsed)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OrderIndex) > 0 {
		i -= len(m.OrderIndex)
		copy(dAtA[i:], m.OrderIndex)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OrderIndex)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitCompletedOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitCompletedOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitCompletedOrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.FunctionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Fuel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitOrderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitCompletedOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OrderIndex)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.FuelUsed)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitCompletedOrderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunctionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunctionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fuel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fuel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitOrderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitCompletedOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitCompletedOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitCompletedOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuelUsed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FuelUsed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitCompletedOrderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitCompletedOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitCompletedOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

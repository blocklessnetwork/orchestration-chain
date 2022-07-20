// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/order.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Order struct {
	Index      string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	FunctionId string `protobuf:"bytes,2,opt,name=functionId,proto3" json:"functionId,omitempty"`
	Fuel       string `protobuf:"bytes,3,opt,name=fuel,proto3" json:"fuel,omitempty"`
	Customer   string `protobuf:"bytes,4,opt,name=customer,proto3" json:"customer,omitempty"`
	Height     string `protobuf:"bytes,5,opt,name=height,proto3" json:"height,omitempty"`
	Date       string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	State      string `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6375df0c4a1904, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Order) GetFunctionId() string {
	if m != nil {
		return m.FunctionId
	}
	return ""
}

func (m *Order) GetFuel() string {
	if m != nil {
		return m.Fuel
	}
	return ""
}

func (m *Order) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *Order) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

func (m *Order) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Order) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type OrderFilter struct {
	Customer string `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	State    string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *OrderFilter) Reset()         { *m = OrderFilter{} }
func (m *OrderFilter) String() string { return proto.CompactTextString(m) }
func (*OrderFilter) ProtoMessage()    {}
func (*OrderFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6375df0c4a1904, []int{1}
}
func (m *OrderFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFilter.Merge(m, src)
}
func (m *OrderFilter) XXX_Size() int {
	return m.Size()
}
func (m *OrderFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFilter.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFilter proto.InternalMessageInfo

func (m *OrderFilter) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *OrderFilter) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "txlabs.blocklesschain.market.Order")
	proto.RegisterType((*OrderFilter)(nil), "txlabs.blocklesschain.market.OrderFilter")
}

func init() { proto.RegisterFile("market/order.proto", fileDescriptor_8c6375df0c4a1904) }

var fileDescriptor_8c6375df0c4a1904 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x9b, 0x75, 0x5b, 0x35, 0xde, 0x82, 0x48, 0x10, 0x09, 0xb2, 0x27, 0x2f, 0xb6, 0x82,
	0x0f, 0x20, 0x78, 0x10, 0xf4, 0x22, 0x78, 0xf4, 0xd6, 0x3f, 0xb3, 0xdb, 0xb0, 0x69, 0xb3, 0x24,
	0x53, 0xa8, 0x6f, 0xe1, 0x9b, 0xf8, 0x1a, 0x1e, 0xf7, 0xe8, 0x51, 0xda, 0x17, 0x91, 0x24, 0x4b,
	0x59, 0xf7, 0x36, 0xdf, 0x7c, 0xc3, 0x37, 0x1f, 0x3f, 0xca, 0x9a, 0xdc, 0xac, 0x01, 0x33, 0x6d,
	0x2a, 0x30, 0xe9, 0xc6, 0x68, 0xd4, 0xec, 0x0a, 0x7b, 0x95, 0x17, 0x36, 0x2d, 0x94, 0x2e, 0xd7,
	0x0a, 0xac, 0x2d, 0xeb, 0x5c, 0xb6, 0x69, 0xb8, 0x5c, 0x7c, 0x11, 0x1a, 0xbf, 0xba, 0x6b, 0x76,
	0x4e, 0x63, 0xd9, 0x56, 0xd0, 0x73, 0x72, 0x4d, 0x6e, 0x4e, 0xdf, 0x82, 0x60, 0x82, 0xd2, 0x65,
	0xd7, 0x96, 0x28, 0x75, 0xfb, 0x5c, 0xf1, 0x99, 0xb7, 0xf6, 0x36, 0x8c, 0xd1, 0xf9, 0xb2, 0x03,
	0xc5, 0x8f, 0xbc, 0xe3, 0x67, 0x76, 0x49, 0x4f, 0xca, 0xce, 0xa2, 0x6e, 0xc0, 0xf0, 0xb9, 0xdf,
	0x4f, 0x9a, 0x5d, 0xd0, 0xa4, 0x06, 0xb9, 0xaa, 0x91, 0xc7, 0xde, 0xd9, 0x29, 0x97, 0x53, 0xe5,
	0x08, 0x3c, 0x09, 0x39, 0x6e, 0x76, 0x8d, 0x2c, 0xba, 0xe5, 0x71, 0x68, 0xe4, 0xc5, 0xe2, 0x81,
	0x9e, 0xf9, 0xc2, 0x4f, 0x52, 0x21, 0x98, 0x7f, 0xcf, 0xc8, 0xc1, 0xb3, 0x29, 0x60, 0xb6, 0x17,
	0xf0, 0xf8, 0xf2, 0x3d, 0x08, 0xb2, 0x1d, 0x04, 0xf9, 0x1d, 0x04, 0xf9, 0x1c, 0x45, 0xb4, 0x1d,
	0x45, 0xf4, 0x33, 0x8a, 0xe8, 0xfd, 0x6e, 0x25, 0xb1, 0xee, 0x8a, 0xb4, 0xd4, 0x4d, 0x16, 0xa8,
	0x65, 0x13, 0xb5, 0x5b, 0x8f, 0x2d, 0xeb, 0xb3, 0x1d, 0x62, 0xfc, 0xd8, 0x80, 0x2d, 0x12, 0xcf,
	0xf8, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x55, 0xef, 0x1d, 0x3f, 0x79, 0x01, 0x00, 0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Height) > 0 {
		i -= len(m.Height)
		copy(dAtA[i:], m.Height)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Height)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Customer) > 0 {
		i -= len(m.Customer)
		copy(dAtA[i:], m.Customer)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Customer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Fuel) > 0 {
		i -= len(m.Fuel)
		copy(dAtA[i:], m.Fuel)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Fuel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FunctionId) > 0 {
		i -= len(m.FunctionId)
		copy(dAtA[i:], m.FunctionId)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.FunctionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OrderFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Customer) > 0 {
		i -= len(m.Customer)
		copy(dAtA[i:], m.Customer)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Customer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.FunctionId)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Fuel)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Customer)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Height)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func (m *OrderFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Customer)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunctionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
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
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fuel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Customer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *OrderFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: OrderFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Customer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Customer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the market module's genesis state.
type GenesisState struct {
	Params               Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	OrderList            []Order            `protobuf:"bytes,2,rep,name=orderList,proto3" json:"orderList"`
	CompletedOrderList   []CompletedOrder   `protobuf:"bytes,3,rep,name=completedOrderList,proto3" json:"completedOrderList"`
	ClaimedOrderList     []ClaimedOrder     `protobuf:"bytes,4,rep,name=claimedOrderList,proto3" json:"claimedOrderList"`
	NodeRegistrationList []NodeRegistration `protobuf:"bytes,5,rep,name=nodeRegistrationList,proto3" json:"nodeRegistrationList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_198e3e6486717af4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetOrderList() []Order {
	if m != nil {
		return m.OrderList
	}
	return nil
}

func (m *GenesisState) GetCompletedOrderList() []CompletedOrder {
	if m != nil {
		return m.CompletedOrderList
	}
	return nil
}

func (m *GenesisState) GetClaimedOrderList() []ClaimedOrder {
	if m != nil {
		return m.ClaimedOrderList
	}
	return nil
}

func (m *GenesisState) GetNodeRegistrationList() []NodeRegistration {
	if m != nil {
		return m.NodeRegistrationList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "txlabs.blocklesschain.market.GenesisState")
}

func init() { proto.RegisterFile("market/genesis.proto", fileDescriptor_198e3e6486717af4) }

var fileDescriptor_198e3e6486717af4 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x93, 0xb7, 0x7d, 0x0b, 0xa6, 0x1e, 0x64, 0xed, 0xa1, 0x84, 0xb2, 0x16, 0xf5, 0x50,
	0x44, 0x37, 0x50, 0xef, 0x1e, 0xea, 0xa1, 0x17, 0xb1, 0x52, 0x2f, 0x22, 0x42, 0xd9, 0xa4, 0x4b,
	0xb2, 0x34, 0xc9, 0x86, 0xdd, 0x15, 0xeb, 0xb7, 0xf0, 0x5b, 0xd9, 0x63, 0x8f, 0x9e, 0x44, 0xda,
	0x2f, 0x22, 0xdd, 0x4c, 0xff, 0x58, 0x4b, 0xbc, 0x2d, 0xf3, 0xcc, 0xf3, 0x7b, 0x76, 0x86, 0x71,
	0x6a, 0x09, 0x95, 0x23, 0xa6, 0xbd, 0x90, 0xa5, 0x4c, 0x71, 0x45, 0x32, 0x29, 0xb4, 0x40, 0x0d,
	0x3d, 0x8e, 0xa9, 0xaf, 0x88, 0x1f, 0x8b, 0x60, 0x14, 0x33, 0xa5, 0x82, 0x88, 0xf2, 0x94, 0xe4,
	0xbd, 0x6e, 0x2d, 0x14, 0xa1, 0x30, 0x8d, 0xde, 0xe2, 0x95, 0x7b, 0xdc, 0x43, 0x20, 0x65, 0x54,
	0xd2, 0x04, 0x40, 0x2e, 0x82, 0xa2, 0x90, 0x43, 0x26, 0xa1, 0xd6, 0x80, 0x5a, 0x20, 0x92, 0x2c,
	0x66, 0x9a, 0x0d, 0x07, 0x9b, 0xaa, 0xbb, 0x54, 0x63, 0xca, 0x93, 0x2d, 0x0d, 0x83, 0x96, 0x8a,
	0x21, 0x1b, 0x48, 0x16, 0x72, 0xa5, 0x25, 0xd5, 0x5c, 0xa4, 0xb9, 0x7e, 0xfc, 0x5e, 0x72, 0xf6,
	0xbb, 0xf9, 0x20, 0xf7, 0x9a, 0x6a, 0x86, 0x3a, 0x4e, 0x25, 0xff, 0x4e, 0xdd, 0x6e, 0xda, 0xad,
	0x6a, 0xfb, 0x94, 0x14, 0x0d, 0x46, 0xee, 0x4c, 0x6f, 0xa7, 0x3c, 0xf9, 0x3c, 0xb2, 0xfa, 0xe0,
	0x44, 0x5d, 0x67, 0xcf, 0xfc, 0xe1, 0x86, 0x2b, 0x5d, 0xff, 0xd7, 0x2c, 0xb5, 0xaa, 0xed, 0x93,
	0x62, 0x4c, 0x6f, 0xd1, 0x0e, 0x94, 0xb5, 0x17, 0xf9, 0x0e, 0x5a, 0x8d, 0xdc, 0x5b, 0x11, 0x4b,
	0x86, 0x78, 0x5e, 0x4c, 0xbc, 0xfe, 0xe1, 0x03, 0xf4, 0x0e, 0x1a, 0x7a, 0x72, 0x0e, 0x60, 0x71,
	0xeb, 0x84, 0xb2, 0x49, 0x38, 0xfb, 0x23, 0x61, 0xc3, 0x05, 0xfc, 0x5f, 0x24, 0x14, 0x39, 0xb5,
	0xc5, 0xea, 0xfb, 0x1b, 0x9b, 0x37, 0x09, 0xff, 0x4d, 0x02, 0x29, 0x4e, 0xb8, 0xdd, 0x72, 0x42,
	0xca, 0x4e, 0x62, 0xe7, 0x61, 0x32, 0xc3, 0xf6, 0x74, 0x86, 0xed, 0xaf, 0x19, 0xb6, 0xdf, 0xe6,
	0xd8, 0x9a, 0xce, 0xb1, 0xf5, 0x31, 0xc7, 0xd6, 0xe3, 0x55, 0xc8, 0x75, 0xf4, 0xec, 0x93, 0x40,
	0x24, 0xde, 0x2a, 0x28, 0x65, 0xfa, 0x45, 0xc8, 0x91, 0x27, 0x64, 0x10, 0xb1, 0x25, 0xe8, 0xc2,
	0xc4, 0x7b, 0x63, 0x0f, 0xae, 0x46, 0xbf, 0x66, 0x4c, 0xf9, 0x15, 0x73, 0x2a, 0x97, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0x8a, 0x32, 0x88, 0xf9, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NodeRegistrationList) > 0 {
		for iNdEx := len(m.NodeRegistrationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodeRegistrationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ClaimedOrderList) > 0 {
		for iNdEx := len(m.ClaimedOrderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedOrderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.CompletedOrderList) > 0 {
		for iNdEx := len(m.CompletedOrderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CompletedOrderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OrderList) > 0 {
		for iNdEx := len(m.OrderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.OrderList) > 0 {
		for _, e := range m.OrderList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CompletedOrderList) > 0 {
		for _, e := range m.CompletedOrderList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClaimedOrderList) > 0 {
		for _, e := range m.ClaimedOrderList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NodeRegistrationList) > 0 {
		for _, e := range m.NodeRegistrationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderList = append(m.OrderList, Order{})
			if err := m.OrderList[len(m.OrderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletedOrderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompletedOrderList = append(m.CompletedOrderList, CompletedOrder{})
			if err := m.CompletedOrderList[len(m.CompletedOrderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedOrderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedOrderList = append(m.ClaimedOrderList, ClaimedOrder{})
			if err := m.ClaimedOrderList[len(m.ClaimedOrderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeRegistrationList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeRegistrationList = append(m.NodeRegistrationList, NodeRegistration{})
			if err := m.NodeRegistrationList[len(m.NodeRegistrationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

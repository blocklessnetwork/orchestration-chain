// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/claimed_order.proto

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

type ClaimedOrder struct {
	Index         string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ClaimedBy     string `protobuf:"bytes,3,opt,name=claimedBy,proto3" json:"claimedBy,omitempty"`
	ClaimedHeight string `protobuf:"bytes,4,opt,name=claimedHeight,proto3" json:"claimedHeight,omitempty"`
	Date          string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (m *ClaimedOrder) Reset()         { *m = ClaimedOrder{} }
func (m *ClaimedOrder) String() string { return proto.CompactTextString(m) }
func (*ClaimedOrder) ProtoMessage()    {}
func (*ClaimedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_99afd8324f4fb72e, []int{0}
}
func (m *ClaimedOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimedOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimedOrder.Merge(m, src)
}
func (m *ClaimedOrder) XXX_Size() int {
	return m.Size()
}
func (m *ClaimedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimedOrder proto.InternalMessageInfo

func (m *ClaimedOrder) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ClaimedOrder) GetClaimedBy() string {
	if m != nil {
		return m.ClaimedBy
	}
	return ""
}

func (m *ClaimedOrder) GetClaimedHeight() string {
	if m != nil {
		return m.ClaimedHeight
	}
	return ""
}

func (m *ClaimedOrder) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*ClaimedOrder)(nil), "txlabs.blocklesschain.market.ClaimedOrder")
}

func init() { proto.RegisterFile("market/claimed_order.proto", fileDescriptor_99afd8324f4fb72e) }

var fileDescriptor_99afd8324f4fb72e = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0xd1, 0x4f, 0xce, 0x49, 0xcc, 0xcc, 0x4d, 0x4d, 0x89, 0xcf, 0x2f, 0x4a, 0x49, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0xa9, 0xc8, 0x49, 0x4c, 0x2a, 0xd6, 0x4b,
	0xca, 0xc9, 0x4f, 0xce, 0xce, 0x49, 0x2d, 0x2e, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83, 0xe8,
	0x50, 0xaa, 0xe1, 0xe2, 0x71, 0x86, 0x68, 0xf2, 0x07, 0xe9, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0xcc,
	0x4b, 0x49, 0xad, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x64, 0xb8, 0x38,
	0xa1, 0x46, 0x3b, 0x55, 0x4a, 0x30, 0x83, 0x65, 0x10, 0x02, 0x42, 0x2a, 0x5c, 0xbc, 0x50, 0x8e,
	0x47, 0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0x0b, 0x58, 0x05, 0xaa, 0xa0, 0x90, 0x10, 0x17, 0x4b,
	0x4a, 0x62, 0x49, 0xaa, 0x04, 0x2b, 0x58, 0x12, 0xcc, 0x76, 0xf2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e,
	0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x83, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc,
	0x5c, 0x7d, 0x88, 0x07, 0xf4, 0xe1, 0x1e, 0xd0, 0x05, 0xfb, 0x40, 0xbf, 0x42, 0x1f, 0xea, 0xeb,
	0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x77, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x7d, 0xe5, 0xd5, 0xbc, 0x0c, 0x01, 0x00, 0x00,
}

func (m *ClaimedOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimedOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimedOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintClaimedOrder(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClaimedHeight) > 0 {
		i -= len(m.ClaimedHeight)
		copy(dAtA[i:], m.ClaimedHeight)
		i = encodeVarintClaimedOrder(dAtA, i, uint64(len(m.ClaimedHeight)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClaimedBy) > 0 {
		i -= len(m.ClaimedBy)
		copy(dAtA[i:], m.ClaimedBy)
		i = encodeVarintClaimedOrder(dAtA, i, uint64(len(m.ClaimedBy)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintClaimedOrder(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimedOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimedOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimedOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovClaimedOrder(uint64(l))
	}
	l = len(m.ClaimedBy)
	if l > 0 {
		n += 1 + l + sovClaimedOrder(uint64(l))
	}
	l = len(m.ClaimedHeight)
	if l > 0 {
		n += 1 + l + sovClaimedOrder(uint64(l))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovClaimedOrder(uint64(l))
	}
	return n
}

func sovClaimedOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimedOrder(x uint64) (n int) {
	return sovClaimedOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimedOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimedOrder
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
			return fmt.Errorf("proto: ClaimedOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimedOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimedOrder
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
				return ErrInvalidLengthClaimedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimedOrder
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
				return ErrInvalidLengthClaimedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimedOrder
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
				return ErrInvalidLengthClaimedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimedOrder
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
				return ErrInvalidLengthClaimedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaimedOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimedOrder
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
func skipClaimedOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimedOrder
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
					return 0, ErrIntOverflowClaimedOrder
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
					return 0, ErrIntOverflowClaimedOrder
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
				return 0, ErrInvalidLengthClaimedOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimedOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimedOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimedOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimedOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimedOrder = fmt.Errorf("proto: unexpected end of group")
)

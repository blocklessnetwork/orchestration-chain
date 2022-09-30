// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/node_registration.proto

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

type NodeRegistration struct {
	Index     string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	NodeId    string `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	NodePort  string `protobuf:"bytes,3,opt,name=nodePort,proto3" json:"nodePort,omitempty"`
	NodeIp    string `protobuf:"bytes,4,opt,name=nodeIp,proto3" json:"nodeIp,omitempty"`
	NodeOwner string `protobuf:"bytes,5,opt,name=nodeOwner,proto3" json:"nodeOwner,omitempty"`
	Height    string `protobuf:"bytes,6,opt,name=height,proto3" json:"height,omitempty"`
	Date      string `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (m *NodeRegistration) Reset()         { *m = NodeRegistration{} }
func (m *NodeRegistration) String() string { return proto.CompactTextString(m) }
func (*NodeRegistration) ProtoMessage()    {}
func (*NodeRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_20c16091e909fc71, []int{0}
}
func (m *NodeRegistration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeRegistration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRegistration.Merge(m, src)
}
func (m *NodeRegistration) XXX_Size() int {
	return m.Size()
}
func (m *NodeRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRegistration proto.InternalMessageInfo

func (m *NodeRegistration) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *NodeRegistration) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeRegistration) GetNodePort() string {
	if m != nil {
		return m.NodePort
	}
	return ""
}

func (m *NodeRegistration) GetNodeIp() string {
	if m != nil {
		return m.NodeIp
	}
	return ""
}

func (m *NodeRegistration) GetNodeOwner() string {
	if m != nil {
		return m.NodeOwner
	}
	return ""
}

func (m *NodeRegistration) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

func (m *NodeRegistration) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeRegistration)(nil), "txlabs.blocklesschain.market.NodeRegistration")
}

func init() { proto.RegisterFile("market/node_registration.proto", fileDescriptor_20c16091e909fc71) }

var fileDescriptor_20c16091e909fc71 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0xb3, 0x7a, 0x17, 0xbd, 0xad, 0x64, 0x11, 0x59, 0xe4, 0x58, 0xc4, 0xca, 0xc6, 0xa4,
	0xb0, 0xb7, 0xb0, 0xb3, 0x51, 0xb9, 0x4a, 0x6c, 0x24, 0x7f, 0x86, 0x64, 0x49, 0x6e, 0x27, 0xec,
	0x8e, 0x5c, 0xfc, 0x16, 0x7e, 0x22, 0x6b, 0xcb, 0x2b, 0x2d, 0x25, 0xf9, 0x22, 0x92, 0xcd, 0x19,
	0xaf, 0x7b, 0xef, 0x0d, 0x3f, 0x86, 0xf7, 0xb8, 0x5a, 0x27, 0xb6, 0x02, 0x8a, 0x0d, 0xe6, 0xf0,
	0x6a, 0xa1, 0xd0, 0x8e, 0x6c, 0x42, 0x1a, 0x4d, 0xd4, 0x58, 0x24, 0x14, 0x4b, 0x6a, 0xeb, 0x24,
	0x75, 0x51, 0x5a, 0x63, 0x56, 0xd5, 0xe0, 0x5c, 0x56, 0x26, 0xda, 0x44, 0x23, 0x75, 0xf9, 0xc9,
	0xf8, 0xc9, 0x03, 0xe6, 0xb0, 0xda, 0x03, 0xc5, 0x29, 0x9f, 0x6b, 0x93, 0x43, 0x2b, 0xd9, 0x05,
	0xbb, 0x5a, 0xac, 0x46, 0x23, 0xce, 0x78, 0x38, 0xfc, 0xb8, 0xcf, 0xe5, 0x81, 0x8f, 0x77, 0x4e,
	0x9c, 0xf3, 0xe3, 0x41, 0x3d, 0xa1, 0x25, 0x79, 0xe8, 0x2f, 0x93, 0x9f, 0x98, 0x46, 0xce, 0xf6,
	0x98, 0x46, 0x2c, 0xf9, 0x62, 0x50, 0x8f, 0x1b, 0x03, 0x56, 0xce, 0xfd, 0xe9, 0x3f, 0x18, 0xa8,
	0x12, 0x74, 0x51, 0x92, 0x0c, 0x47, 0x6a, 0x74, 0x42, 0xf0, 0x59, 0x9e, 0x10, 0xc8, 0x23, 0x9f,
	0x7a, 0x7d, 0xf7, 0xfc, 0xd5, 0x29, 0xb6, 0xed, 0x14, 0xfb, 0xe9, 0x14, 0xfb, 0xe8, 0x55, 0xb0,
	0xed, 0x55, 0xf0, 0xdd, 0xab, 0xe0, 0xe5, 0xb6, 0xd0, 0x54, 0xbe, 0xa5, 0x51, 0x86, 0xeb, 0x78,
	0x2a, 0x6f, 0x80, 0x36, 0x68, 0xab, 0x18, 0x6d, 0x56, 0xc2, 0x5f, 0xe1, 0x6b, 0x3f, 0x49, 0xdc,
	0xc6, 0xbb, 0x29, 0xe9, 0xbd, 0x01, 0x97, 0x86, 0x7e, 0xbf, 0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0x77, 0x37, 0xf1, 0x61, 0x01, 0x00, 0x00,
}

func (m *NodeRegistration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeRegistration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeRegistration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintNodeRegistration(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Height) > 0 {
		i -= len(m.Height)
		copy(dAtA[i:], m.Height)
		i = encodeVarintNodeRegistration(dAtA, i, uint64(len(m.Height)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.NodeOwner) > 0 {
		i -= len(m.NodeOwner)
		copy(dAtA[i:], m.NodeOwner)
		i = encodeVarintNodeRegistration(dAtA, i, uint64(len(m.NodeOwner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.NodeIp) > 0 {
		i -= len(m.NodeIp)
		copy(dAtA[i:], m.NodeIp)
		i = encodeVarintNodeRegistration(dAtA, i, uint64(len(m.NodeIp)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NodePort) > 0 {
		i -= len(m.NodePort)
		copy(dAtA[i:], m.NodePort)
		i = encodeVarintNodeRegistration(dAtA, i, uint64(len(m.NodePort)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NodeId) > 0 {
		i -= len(m.NodeId)
		copy(dAtA[i:], m.NodeId)
		i = encodeVarintNodeRegistration(dAtA, i, uint64(len(m.NodeId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintNodeRegistration(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNodeRegistration(dAtA []byte, offset int, v uint64) int {
	offset -= sovNodeRegistration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NodeRegistration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovNodeRegistration(uint64(l))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovNodeRegistration(uint64(l))
	}
	l = len(m.NodePort)
	if l > 0 {
		n += 1 + l + sovNodeRegistration(uint64(l))
	}
	l = len(m.NodeIp)
	if l > 0 {
		n += 1 + l + sovNodeRegistration(uint64(l))
	}
	l = len(m.NodeOwner)
	if l > 0 {
		n += 1 + l + sovNodeRegistration(uint64(l))
	}
	l = len(m.Height)
	if l > 0 {
		n += 1 + l + sovNodeRegistration(uint64(l))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovNodeRegistration(uint64(l))
	}
	return n
}

func sovNodeRegistration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNodeRegistration(x uint64) (n int) {
	return sovNodeRegistration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NodeRegistration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeRegistration
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
			return fmt.Errorf("proto: NodeRegistration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeRegistration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeRegistration
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
				return ErrInvalidLengthNodeRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeRegistration
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
				return ErrInvalidLengthNodeRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeRegistration
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
				return ErrInvalidLengthNodeRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeRegistration
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
				return ErrInvalidLengthNodeRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeRegistration
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
				return ErrInvalidLengthNodeRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeRegistration
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
				return ErrInvalidLengthNodeRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeRegistration
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
				return ErrInvalidLengthNodeRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNodeRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeRegistration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNodeRegistration
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
func skipNodeRegistration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodeRegistration
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
					return 0, ErrIntOverflowNodeRegistration
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
					return 0, ErrIntOverflowNodeRegistration
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
				return 0, ErrInvalidLengthNodeRegistration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNodeRegistration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNodeRegistration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNodeRegistration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodeRegistration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNodeRegistration = fmt.Errorf("proto: unexpected end of group")
)

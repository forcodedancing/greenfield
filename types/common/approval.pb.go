// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/common/approval.proto

package common

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Approval is the signature information returned by the Primary Storage Provider (SP) to the user
// after allowing them to create a bucket or object, which is then used for verification on the chain
// to ensure agreement between the Primary SP and the user.
type Approval struct {
	// expired_height is the block height at which the signature expires.
	ExpiredHeight uint64 `protobuf:"varint,1,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
	// global_virtual_group_family_id is the family id that stored.
	GlobalVirtualGroupFamilyId uint32 `protobuf:"varint,2,opt,name=global_virtual_group_family_id,json=globalVirtualGroupFamilyId,proto3" json:"global_virtual_group_family_id,omitempty"`
	// The signature needs to conform to the EIP 712 specification.
	Sig []byte `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *Approval) Reset()         { *m = Approval{} }
func (m *Approval) String() string { return proto.CompactTextString(m) }
func (*Approval) ProtoMessage()    {}
func (*Approval) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe7c4fc8fb6d4918, []int{0}
}
func (m *Approval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Approval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Approval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Approval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Approval.Merge(m, src)
}
func (m *Approval) XXX_Size() int {
	return m.Size()
}
func (m *Approval) XXX_DiscardUnknown() {
	xxx_messageInfo_Approval.DiscardUnknown(m)
}

var xxx_messageInfo_Approval proto.InternalMessageInfo

func (m *Approval) GetExpiredHeight() uint64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

func (m *Approval) GetGlobalVirtualGroupFamilyId() uint32 {
	if m != nil {
		return m.GlobalVirtualGroupFamilyId
	}
	return 0
}

func (m *Approval) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*Approval)(nil), "greenfield.common.Approval")
}

func init() { proto.RegisterFile("greenfield/common/approval.proto", fileDescriptor_fe7c4fc8fb6d4918) }

var fileDescriptor_fe7c4fc8fb6d4918 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x18, 0x84, 0xb3, 0x9e, 0x88, 0x04, 0x4f, 0x34, 0x58, 0xc4, 0x14, 0x4b, 0x10, 0x84, 0x14, 0x7a,
	0x5b, 0xf8, 0x04, 0x5e, 0x71, 0x6a, 0x9b, 0xc2, 0xc2, 0x26, 0xec, 0x5e, 0x36, 0x7f, 0x16, 0x36,
	0xf9, 0x97, 0xbd, 0xe4, 0x30, 0x4f, 0x60, 0xeb, 0x63, 0x59, 0x5e, 0x69, 0x29, 0xc9, 0x8b, 0x48,
	0xb2, 0x01, 0xed, 0xe6, 0x9f, 0xf9, 0xf8, 0x61, 0xc6, 0x8f, 0xc1, 0x4a, 0x59, 0x17, 0x4a, 0xea,
	0x9c, 0x6d, 0xb1, 0xaa, 0xb0, 0x66, 0xdc, 0x18, 0x8b, 0x7b, 0xae, 0x57, 0xc6, 0x62, 0x83, 0xc1,
	0xe5, 0x1f, 0xb1, 0x72, 0x44, 0x74, 0x05, 0x08, 0x38, 0xa5, 0x6c, 0x54, 0x0e, 0x8c, 0xae, 0x01,
	0x11, 0xb4, 0x64, 0xd3, 0x25, 0xda, 0x82, 0xf1, 0xba, 0x73, 0xd1, 0xcd, 0x07, 0xf1, 0x4f, 0x1f,
	0xe7, 0xb7, 0xc1, 0xad, 0x7f, 0x2e, 0xdf, 0x8d, 0xb2, 0x32, 0xcf, 0x4a, 0xa9, 0xa0, 0x6c, 0x42,
	0x12, 0x93, 0xe4, 0x38, 0x5d, 0xce, 0xee, 0xf3, 0x64, 0x06, 0x6b, 0x9f, 0x82, 0x46, 0xc1, 0x75,
	0xb6, 0x57, 0xb6, 0x69, 0xb9, 0xce, 0xc0, 0x62, 0x6b, 0xb2, 0x82, 0x57, 0x4a, 0x77, 0x99, 0xca,
	0xc3, 0xa3, 0x98, 0x24, 0xcb, 0x34, 0x72, 0xd4, 0xab, 0x83, 0x9e, 0x46, 0x66, 0x33, 0x21, 0x2f,
	0x79, 0x70, 0xe1, 0x2f, 0x76, 0x0a, 0xc2, 0x45, 0x4c, 0x92, 0xb3, 0x74, 0x94, 0xeb, 0xcd, 0x57,
	0x4f, 0xc9, 0xa1, 0xa7, 0xe4, 0xa7, 0xa7, 0xe4, 0x73, 0xa0, 0xde, 0x61, 0xa0, 0xde, 0xf7, 0x40,
	0xbd, 0xb7, 0x3b, 0x50, 0x4d, 0xd9, 0x8a, 0xb1, 0x23, 0x13, 0xb5, 0xb8, 0xdf, 0x96, 0x5c, 0xd5,
	0xec, 0xdf, 0x3c, 0x4d, 0x67, 0xe4, 0x6e, 0x1e, 0x49, 0x9c, 0x4c, 0xc5, 0x1e, 0x7e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xcc, 0x55, 0xc7, 0x9e, 0x40, 0x01, 0x00, 0x00,
}

func (m *Approval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Approval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Approval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintApproval(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GlobalVirtualGroupFamilyId != 0 {
		i = encodeVarintApproval(dAtA, i, uint64(m.GlobalVirtualGroupFamilyId))
		i--
		dAtA[i] = 0x10
	}
	if m.ExpiredHeight != 0 {
		i = encodeVarintApproval(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApproval(dAtA []byte, offset int, v uint64) int {
	offset -= sovApproval(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Approval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpiredHeight != 0 {
		n += 1 + sovApproval(uint64(m.ExpiredHeight))
	}
	if m.GlobalVirtualGroupFamilyId != 0 {
		n += 1 + sovApproval(uint64(m.GlobalVirtualGroupFamilyId))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovApproval(uint64(l))
	}
	return n
}

func sovApproval(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApproval(x uint64) (n int) {
	return sovApproval(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Approval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApproval
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
			return fmt.Errorf("proto: Approval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Approval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredHeight", wireType)
			}
			m.ExpiredHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApproval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiredHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalVirtualGroupFamilyId", wireType)
			}
			m.GlobalVirtualGroupFamilyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApproval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalVirtualGroupFamilyId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApproval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApproval
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApproval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApproval(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApproval
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
func skipApproval(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApproval
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
					return 0, ErrIntOverflowApproval
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
					return 0, ErrIntOverflowApproval
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
				return 0, ErrInvalidLengthApproval
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApproval
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApproval
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApproval        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApproval          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApproval = fmt.Errorf("proto: unexpected end of group")
)

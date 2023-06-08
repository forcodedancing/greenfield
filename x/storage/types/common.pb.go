// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/storage/common.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// SourceType represents the source of resource creation, which can
// from Greenfield native or from a cross-chain transfer from BSC
type SourceType int32

const (
	SOURCE_TYPE_ORIGIN          SourceType = 0
	SOURCE_TYPE_BSC_CROSS_CHAIN SourceType = 1
	SOURCE_TYPE_MIRROR_PENDING  SourceType = 2
)

var SourceType_name = map[int32]string{
	0: "SOURCE_TYPE_ORIGIN",
	1: "SOURCE_TYPE_BSC_CROSS_CHAIN",
	2: "SOURCE_TYPE_MIRROR_PENDING",
}

var SourceType_value = map[string]int32{
	"SOURCE_TYPE_ORIGIN":          0,
	"SOURCE_TYPE_BSC_CROSS_CHAIN": 1,
	"SOURCE_TYPE_MIRROR_PENDING":  2,
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}

func (SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{0}
}

// BucketStatus represents the status of a bucket. After a user successfully
// sends a CreateBucket transaction onto the chain, the status is set to 'Created'.
// When a Discontinue Object transaction is received on chain, the status is set to 'Discontinued'.
type BucketStatus int32

const (
	BUCKET_STATUS_CREATED      BucketStatus = 0
	BUCKET_STATUS_DISCONTINUED BucketStatus = 1
)

var BucketStatus_name = map[int32]string{
	0: "BUCKET_STATUS_CREATED",
	1: "BUCKET_STATUS_DISCONTINUED",
}

var BucketStatus_value = map[string]int32{
	"BUCKET_STATUS_CREATED":      0,
	"BUCKET_STATUS_DISCONTINUED": 1,
}

func (x BucketStatus) String() string {
	return proto.EnumName(BucketStatus_name, int32(x))
}

func (BucketStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{1}
}

// RedundancyType represents the redundancy algorithm type for object data,
// which can be either multi-replica or erasure coding.
type RedundancyType int32

const (
	REDUNDANCY_EC_TYPE      RedundancyType = 0
	REDUNDANCY_REPLICA_TYPE RedundancyType = 1
)

var RedundancyType_name = map[int32]string{
	0: "REDUNDANCY_EC_TYPE",
	1: "REDUNDANCY_REPLICA_TYPE",
}

var RedundancyType_value = map[string]int32{
	"REDUNDANCY_EC_TYPE":      0,
	"REDUNDANCY_REPLICA_TYPE": 1,
}

func (x RedundancyType) String() string {
	return proto.EnumName(RedundancyType_name, int32(x))
}

func (RedundancyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{2}
}

// ObjectStatus represents the creation status of an object. After a user successfully
// sends a CreateObject transaction onto the chain, the status is set to 'Created'.
// After the Primary Service Provider successfully sends a Seal Object transaction onto
// the chain, the status is set to 'Sealed'. When a Discontinue Object transaction is
// received on chain, the status is set to 'Discontinued'.
type ObjectStatus int32

const (
	OBJECT_STATUS_CREATED      ObjectStatus = 0
	OBJECT_STATUS_SEALED       ObjectStatus = 1
	OBJECT_STATUS_DISCONTINUED ObjectStatus = 2
)

var ObjectStatus_name = map[int32]string{
	0: "OBJECT_STATUS_CREATED",
	1: "OBJECT_STATUS_SEALED",
	2: "OBJECT_STATUS_DISCONTINUED",
}

var ObjectStatus_value = map[string]int32{
	"OBJECT_STATUS_CREATED":      0,
	"OBJECT_STATUS_SEALED":       1,
	"OBJECT_STATUS_DISCONTINUED": 2,
}

func (x ObjectStatus) String() string {
	return proto.EnumName(ObjectStatus_name, int32(x))
}

func (ObjectStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{3}
}

// VisibilityType is the resources public status.
type VisibilityType int32

const (
	VISIBILITY_TYPE_UNSPECIFIED VisibilityType = 0
	VISIBILITY_TYPE_PUBLIC_READ VisibilityType = 1
	VISIBILITY_TYPE_PRIVATE     VisibilityType = 2
	// If the bucket Visibility is inherit, it's finally set to private. If the object Visibility is inherit, it's the same as bucket.
	VISIBILITY_TYPE_INHERIT VisibilityType = 3
)

var VisibilityType_name = map[int32]string{
	0: "VISIBILITY_TYPE_UNSPECIFIED",
	1: "VISIBILITY_TYPE_PUBLIC_READ",
	2: "VISIBILITY_TYPE_PRIVATE",
	3: "VISIBILITY_TYPE_INHERIT",
}

var VisibilityType_value = map[string]int32{
	"VISIBILITY_TYPE_UNSPECIFIED": 0,
	"VISIBILITY_TYPE_PUBLIC_READ": 1,
	"VISIBILITY_TYPE_PRIVATE":     2,
	"VISIBILITY_TYPE_INHERIT":     3,
}

func (x VisibilityType) String() string {
	return proto.EnumName(VisibilityType_name, int32(x))
}

func (VisibilityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{4}
}

// SecondarySpSignDoc used to generate seal signature of secondary SP
// If the secondary SP only signs the checksum to declare the object pieces are saved,
// it might be reused by the primary SP to fake it's declaration.
// Then the primary SP can challenge and slash the secondary SP.
// So the id of the object is needed to prevent this.
type SecondarySpSignDoc struct {
	SpAddress string `protobuf:"bytes,1,opt,name=sp_address,json=spAddress,proto3" json:"sp_address,omitempty"`
	ObjectId  Uint   `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3,customtype=Uint" json:"object_id"`
	Checksum  []byte `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (m *SecondarySpSignDoc) Reset()         { *m = SecondarySpSignDoc{} }
func (m *SecondarySpSignDoc) String() string { return proto.CompactTextString(m) }
func (*SecondarySpSignDoc) ProtoMessage()    {}
func (*SecondarySpSignDoc) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eff6c0fa4aaf4c9, []int{0}
}
func (m *SecondarySpSignDoc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecondarySpSignDoc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecondarySpSignDoc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SecondarySpSignDoc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecondarySpSignDoc.Merge(m, src)
}
func (m *SecondarySpSignDoc) XXX_Size() int {
	return m.Size()
}
func (m *SecondarySpSignDoc) XXX_DiscardUnknown() {
	xxx_messageInfo_SecondarySpSignDoc.DiscardUnknown(m)
}

var xxx_messageInfo_SecondarySpSignDoc proto.InternalMessageInfo

func (m *SecondarySpSignDoc) GetSpAddress() string {
	if m != nil {
		return m.SpAddress
	}
	return ""
}

func (m *SecondarySpSignDoc) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func init() {
	proto.RegisterEnum("greenfield.storage.SourceType", SourceType_name, SourceType_value)
	proto.RegisterEnum("greenfield.storage.BucketStatus", BucketStatus_name, BucketStatus_value)
	proto.RegisterEnum("greenfield.storage.RedundancyType", RedundancyType_name, RedundancyType_value)
	proto.RegisterEnum("greenfield.storage.ObjectStatus", ObjectStatus_name, ObjectStatus_value)
	proto.RegisterEnum("greenfield.storage.VisibilityType", VisibilityType_name, VisibilityType_value)
	proto.RegisterType((*SecondarySpSignDoc)(nil), "greenfield.storage.SecondarySpSignDoc")
}

func init() { proto.RegisterFile("greenfield/storage/common.proto", fileDescriptor_4eff6c0fa4aaf4c9) }

var fileDescriptor_4eff6c0fa4aaf4c9 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcf, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x6d, 0x40, 0x9f, 0x60, 0x3e, 0x84, 0xac, 0x11, 0x2d, 0x21, 0x91, 0x1c, 0xd4, 0x15,
	0x42, 0x82, 0x2c, 0xba, 0x68, 0xb7, 0xf6, 0x78, 0x0a, 0x53, 0x52, 0x3b, 0x9a, 0xb1, 0x91, 0xe8,
	0x66, 0x94, 0x8c, 0xa7, 0x66, 0x0a, 0xf1, 0x44, 0xfe, 0x23, 0x35, 0x6f, 0xd0, 0x65, 0xd5, 0x57,
	0xe0, 0x15, 0x78, 0x08, 0x96, 0x88, 0x55, 0xd5, 0x05, 0xaa, 0xe0, 0x45, 0x2a, 0xdb, 0x01, 0x12,
	0x95, 0x5d, 0xee, 0x39, 0x27, 0x67, 0x7e, 0xbe, 0xd2, 0x05, 0xdd, 0x24, 0x93, 0x32, 0xfd, 0xa2,
	0xe4, 0x45, 0xdc, 0xcb, 0x0b, 0x9d, 0x0d, 0x13, 0xd9, 0x13, 0x7a, 0x3c, 0xd6, 0xe9, 0xc1, 0x24,
	0xd3, 0x85, 0x86, 0xf0, 0x39, 0x70, 0x30, 0x0b, 0xb4, 0xb7, 0x85, 0xce, 0xc7, 0x3a, 0xe7, 0x75,
	0xa2, 0xd7, 0x0c, 0x4d, 0xbc, 0xbd, 0x99, 0xe8, 0x44, 0x37, 0x7a, 0xf5, 0xab, 0x51, 0xdf, 0x5c,
	0x9a, 0x00, 0x32, 0x29, 0x74, 0x1a, 0x0f, 0xb3, 0x29, 0x9b, 0x30, 0x95, 0xa4, 0x9e, 0x16, 0xf0,
	0x1d, 0x00, 0xf9, 0x84, 0x0f, 0xe3, 0x38, 0x93, 0x79, 0xde, 0x32, 0x77, 0xcc, 0xdd, 0x35, 0xb7,
	0x75, 0x7b, 0xb5, 0xbf, 0x39, 0xab, 0x74, 0x1a, 0x87, 0x15, 0x99, 0x4a, 0x13, 0xba, 0x96, 0x4f,
	0x66, 0x02, 0x7c, 0x0f, 0xd6, 0xf4, 0xe8, 0xab, 0x14, 0x05, 0x57, 0x71, 0x6b, 0xa9, 0xfe, 0x5f,
	0xe7, 0xfa, 0xae, 0x6b, 0xfc, 0xbe, 0xeb, 0xae, 0x44, 0x2a, 0x2d, 0x6e, 0xaf, 0xf6, 0xff, 0x9f,
	0x75, 0x54, 0x23, 0x5d, 0x6d, 0xd2, 0x24, 0x86, 0x6d, 0xb0, 0x2a, 0xce, 0xa4, 0x38, 0xcf, 0xcb,
	0x71, 0x6b, 0x79, 0xc7, 0xdc, 0x5d, 0xa7, 0x4f, 0xf3, 0xde, 0x39, 0x00, 0x4c, 0x97, 0x99, 0x90,
	0xe1, 0x74, 0x22, 0xe1, 0x6b, 0x00, 0x59, 0x10, 0x51, 0x84, 0x79, 0x78, 0x3a, 0xc0, 0x3c, 0xa0,
	0xe4, 0x90, 0xf8, 0x96, 0x01, 0xbb, 0xa0, 0x33, 0xaf, 0xbb, 0x0c, 0x71, 0x44, 0x03, 0xc6, 0x38,
	0x3a, 0x72, 0x88, 0x6f, 0x99, 0xd0, 0x06, 0xed, 0xf9, 0xc0, 0x27, 0x42, 0x69, 0x40, 0xf9, 0x00,
	0xfb, 0x1e, 0xf1, 0x0f, 0xad, 0xa5, 0xf6, 0xca, 0xf7, 0x4b, 0xdb, 0xd8, 0x0b, 0xc0, 0xba, 0x5b,
	0x8a, 0x73, 0x59, 0xb0, 0x62, 0x58, 0x94, 0x39, 0xdc, 0x06, 0xaf, 0xdc, 0x08, 0x1d, 0xe3, 0x90,
	0xb3, 0xd0, 0x09, 0x23, 0xc6, 0x11, 0xc5, 0x4e, 0x88, 0x3d, 0xcb, 0xa8, 0x0a, 0x17, 0x2d, 0x8f,
	0x30, 0x14, 0xf8, 0x21, 0xf1, 0x23, 0xec, 0x59, 0xe6, 0xac, 0xf0, 0x18, 0x6c, 0x50, 0x19, 0x97,
	0x69, 0x3c, 0x4c, 0xc5, 0xf4, 0xf1, 0x0b, 0x28, 0xf6, 0x22, 0xdf, 0x73, 0x7c, 0x74, 0xca, 0x31,
	0xaa, 0x79, 0x2c, 0x03, 0x76, 0xc0, 0xd6, 0x9c, 0x4e, 0xf1, 0xa0, 0x4f, 0x90, 0xd3, 0x98, 0x8f,
	0x65, 0x0a, 0xac, 0x07, 0xf5, 0xca, 0x9e, 0xe9, 0x02, 0xf7, 0x23, 0x46, 0x2f, 0xd0, 0xb5, 0xc0,
	0xe6, 0xa2, 0xc5, 0xb0, 0xd3, 0xaf, 0xb8, 0x2a, 0xee, 0x45, 0x67, 0x81, 0xfb, 0x71, 0x11, 0x3f,
	0x4d, 0xb0, 0x71, 0xa2, 0x72, 0x35, 0x52, 0x17, 0xaa, 0x68, 0xc0, 0xbb, 0xa0, 0x73, 0x42, 0x18,
	0x71, 0x49, 0x9f, 0x84, 0xa7, 0xcd, 0x16, 0x23, 0x9f, 0x0d, 0x30, 0x22, 0x1f, 0x48, 0xfd, 0xe6,
	0x0b, 0x81, 0x41, 0xe4, 0xf6, 0x09, 0xe2, 0x14, 0x3b, 0xd5, 0xd3, 0x1d, 0xb0, 0xf5, 0x4f, 0x80,
	0x92, 0x13, 0x27, 0xc4, 0xd6, 0xd2, 0x4b, 0x26, 0xf1, 0x8f, 0x30, 0x25, 0xa1, 0xb5, 0xdc, 0x40,
	0xb9, 0xe4, 0xfa, 0xde, 0x36, 0x6f, 0xee, 0x6d, 0xf3, 0xcf, 0xbd, 0x6d, 0xfe, 0x78, 0xb0, 0x8d,
	0x9b, 0x07, 0xdb, 0xf8, 0xf5, 0x60, 0x1b, 0x9f, 0x7b, 0x89, 0x2a, 0xce, 0xca, 0xd1, 0x81, 0xd0,
	0xe3, 0xde, 0x28, 0x1d, 0xed, 0x8b, 0xb3, 0xa1, 0x4a, 0x7b, 0x73, 0x57, 0xf4, 0xed, 0xe9, 0x8e,
	0x8a, 0xe9, 0x44, 0xe6, 0xa3, 0xff, 0xea, 0x13, 0x78, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x74,
	0xb9, 0x56, 0xa4, 0x6a, 0x03, 0x00, 0x00,
}

func (m *SecondarySpSignDoc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecondarySpSignDoc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SecondarySpSignDoc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.ObjectId.Size()
		i -= size
		if _, err := m.ObjectId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCommon(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.SpAddress) > 0 {
		i -= len(m.SpAddress)
		copy(dAtA[i:], m.SpAddress)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.SpAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SecondarySpSignDoc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpAddress)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = m.ObjectId.Size()
	n += 1 + l + sovCommon(uint64(l))
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SecondarySpSignDoc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: SecondarySpSignDoc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecondarySpSignDoc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = append(m.Checksum[:0], dAtA[iNdEx:postIndex]...)
			if m.Checksum == nil {
				m.Checksum = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)

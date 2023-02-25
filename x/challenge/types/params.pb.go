// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/challenge/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	// Challenges which will be emitted in each block, including user triggered or randomly triggered.
	ChallengeCountPerBlock uint64 `protobuf:"varint,1,opt,name=challenge_count_per_block,json=challengeCountPerBlock,proto3" json:"challenge_count_per_block,omitempty" yaml:"challenge_count_per_block"`
	// The count of blocks to stand for the period in which the same storage and object info cannot be slashed again.
	SlashCoolingOffPeriod uint64 `protobuf:"varint,2,opt,name=slash_cooling_off_period,json=slashCoolingOffPeriod,proto3" json:"slash_cooling_off_period,omitempty" yaml:"slash_cooling_off_period"`
	// The slash coin amount will be calculated from the size of object info, and adjusted by this rate.
	SlashAmountSizeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=slash_amount_size_rate,json=slashAmountSizeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_amount_size_rate" yaml:"slash_amount_size_rate"`
	// The minimal slash amount.
	SlashAmountMin github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=slash_amount_min,json=slashAmountMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_amount_min"`
	// The maximum slash amount.
	SlashAmountMax github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=slash_amount_max,json=slashAmountMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_amount_max"`
	// The ratio of slash amount for all validators' rewards.
	RewardValidatorRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=reward_validator_ratio,json=rewardValidatorRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_validator_ratio" yaml:"reward_validator_ratio"`
	// The ratio of slash amount for challenger's rewards.
	RewardChallengerRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=reward_challenger_ratio,json=rewardChallengerRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_challenger_ratio" yaml:"reward_challenger_ratio"`
	// Heartbeat interval defines the frequency of heartbeat based on challenges.
	HeartbeatInterval uint64 `protobuf:"varint,8,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty" yaml:"heartbeat_interval"`
	// The reward to heartbeat transaction submitter will be calculated based on this rate.
	HeartbeatRewardRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=heartbeat_reward_rate,json=heartbeatRewardRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"heartbeat_reward_rate" yaml:"heartbeat_reward_rate"`
	// The reward to heartbeat transaction submitter will be adjusted based on the threshold.
	HeartbeatRewardThreshold github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=heartbeat_reward_threshold,json=heartbeatRewardThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"heartbeat_reward_threshold"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2396367ee53edf57, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetChallengeCountPerBlock() uint64 {
	if m != nil {
		return m.ChallengeCountPerBlock
	}
	return 0
}

func (m *Params) GetSlashCoolingOffPeriod() uint64 {
	if m != nil {
		return m.SlashCoolingOffPeriod
	}
	return 0
}

func (m *Params) GetHeartbeatInterval() uint64 {
	if m != nil {
		return m.HeartbeatInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "bnbchain.greenfield.challenge.Params")
}

func init() { proto.RegisterFile("greenfield/challenge/params.proto", fileDescriptor_2396367ee53edf57) }

var fileDescriptor_2396367ee53edf57 = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x8f, 0xd2, 0x40,
	0x1c, 0xa5, 0x8a, 0x28, 0x73, 0x30, 0xda, 0x5d, 0xb0, 0x10, 0x69, 0xb1, 0x1a, 0xb3, 0x17, 0x20,
	0xc6, 0xdb, 0xc6, 0x8b, 0xe0, 0x85, 0xf8, 0x67, 0x49, 0x35, 0x1e, 0x8c, 0x49, 0x33, 0x6d, 0xa7,
	0xed, 0x64, 0xdb, 0x19, 0x32, 0x9d, 0x5d, 0x59, 0x8e, 0x46, 0x8d, 0x47, 0x8f, 0x1e, 0xfd, 0x10,
	0x7e, 0x88, 0x3d, 0x6e, 0x3c, 0x19, 0x0f, 0x8d, 0x81, 0x6f, 0xc0, 0x27, 0x30, 0x9d, 0x96, 0x6e,
	0x17, 0xdc, 0x03, 0xd9, 0x3d, 0x01, 0xef, 0xbd, 0x79, 0xef, 0xcd, 0xf0, 0xcb, 0x0f, 0xdc, 0xf3,
	0x18, 0x42, 0xc4, 0xc5, 0x28, 0x70, 0x7a, 0xb6, 0x0f, 0x83, 0x00, 0x11, 0x0f, 0xf5, 0xc6, 0x90,
	0xc1, 0x30, 0xea, 0x8e, 0x19, 0xe5, 0x54, 0x6e, 0x59, 0xc4, 0xb2, 0x7d, 0x88, 0x49, 0xf7, 0x54,
	0xdb, 0xcd, 0xb5, 0xcd, 0x86, 0x4d, 0xa3, 0x90, 0x46, 0xa6, 0x10, 0xf7, 0xd2, 0x1f, 0xe9, 0xc9,
	0xe6, 0xb6, 0x47, 0x3d, 0x9a, 0xe2, 0xc9, 0xb7, 0x14, 0xd5, 0xbf, 0x54, 0x41, 0x65, 0x24, 0x02,
	0x64, 0x13, 0x34, 0x72, 0x23, 0xd3, 0xa6, 0x07, 0x84, 0x9b, 0x63, 0xc4, 0x4c, 0x2b, 0xa0, 0xf6,
	0xbe, 0x22, 0xb5, 0xa5, 0x9d, 0x72, 0xff, 0xc1, 0x22, 0xd6, 0xda, 0x47, 0x30, 0x0c, 0x76, 0xf5,
	0x73, 0xa5, 0xba, 0x51, 0xcf, 0xb9, 0x41, 0x42, 0x8d, 0x10, 0xeb, 0x27, 0x84, 0xfc, 0x1e, 0x28,
	0x51, 0x00, 0x23, 0xdf, 0xb4, 0x29, 0x0d, 0x30, 0xf1, 0x4c, 0xea, 0xba, 0xc9, 0x39, 0x4c, 0x1d,
	0xe5, 0x8a, 0xf0, 0xbf, 0xbf, 0x88, 0x35, 0x2d, 0xf5, 0x3f, 0x4f, 0xa9, 0x1b, 0x35, 0x41, 0x0d,
	0x52, 0x66, 0xcf, 0x75, 0x47, 0x02, 0x97, 0x3f, 0x49, 0xa0, 0x9e, 0x1e, 0x82, 0xa1, 0x68, 0x14,
	0xe1, 0x29, 0x32, 0x19, 0xe4, 0x48, 0xb9, 0xda, 0x96, 0x76, 0xaa, 0xfd, 0xbd, 0xe3, 0x58, 0x2b,
	0xfd, 0x89, 0xb5, 0x87, 0x1e, 0xe6, 0xfe, 0x81, 0xd5, 0xb5, 0x69, 0x98, 0xbd, 0x50, 0xf6, 0xd1,
	0x89, 0x9c, 0xfd, 0x1e, 0x3f, 0x1a, 0xa3, 0xa8, 0xfb, 0x0c, 0xd9, 0x8b, 0x58, 0x6b, 0x15, 0xab,
	0xac, 0xba, 0xea, 0xc6, 0x96, 0x20, 0x9e, 0x0a, 0xfc, 0x35, 0x9e, 0x22, 0x03, 0x72, 0x24, 0xbb,
	0xe0, 0xd6, 0x19, 0x7d, 0x88, 0x89, 0x52, 0x16, 0xf9, 0x4f, 0x36, 0xc8, 0x1f, 0x12, 0xfe, 0xeb,
	0x67, 0x07, 0x64, 0x7f, 0xe0, 0x90, 0x70, 0xe3, 0x66, 0x21, 0xec, 0x25, 0x26, 0xeb, 0x39, 0x70,
	0xa2, 0x5c, 0xbb, 0xec, 0x1c, 0x38, 0x91, 0x3f, 0x4b, 0xa0, 0xce, 0xd0, 0x07, 0xc8, 0x1c, 0xf3,
	0x10, 0x06, 0xd8, 0x81, 0x9c, 0xb2, 0xe4, 0xfe, 0x98, 0x2a, 0x95, 0x8b, 0x3d, 0xeb, 0xff, 0x5d,
	0x75, 0x63, 0x3b, 0x25, 0xde, 0x2e, 0x71, 0x23, 0x81, 0xe5, 0xaf, 0x12, 0xb8, 0x93, 0x9d, 0xc8,
	0xc7, 0x6b, 0x59, 0xe4, 0xba, 0x28, 0x32, 0xda, 0xb8, 0x88, 0x7a, 0xa6, 0xc8, 0xaa, 0xad, 0x6e,
	0xd4, 0x52, 0x66, 0x90, 0x13, 0x69, 0x95, 0x17, 0x40, 0xf6, 0x11, 0x64, 0xdc, 0x42, 0x90, 0x9b,
	0x98, 0x70, 0xc4, 0x0e, 0x61, 0xa0, 0xdc, 0x10, 0x13, 0xdc, 0x5a, 0xc4, 0x5a, 0x23, 0xb5, 0x5d,
	0xd7, 0xe8, 0xc6, 0xed, 0x1c, 0x1c, 0x66, 0x98, 0xfc, 0x51, 0x02, 0xb5, 0x53, 0x69, 0xd6, 0x45,
	0x8c, 0x6d, 0x55, 0x5c, 0xeb, 0xd5, 0xc6, 0xd7, 0xba, 0xbb, 0x9a, 0x5f, 0x30, 0xd5, 0x8d, 0xad,
	0x1c, 0x37, 0x04, 0x2c, 0xa6, 0x76, 0x0a, 0x9a, 0x6b, 0x72, 0xee, 0x33, 0x14, 0xf9, 0x34, 0x70,
	0x14, 0x70, 0x09, 0x73, 0xa5, 0xac, 0xc4, 0xbe, 0x59, 0xba, 0xef, 0x96, 0xbf, 0xff, 0xd0, 0x4a,
	0xfd, 0xe7, 0xc7, 0x33, 0x55, 0x3a, 0x99, 0xa9, 0xd2, 0xdf, 0x99, 0x2a, 0x7d, 0x9b, 0xab, 0xa5,
	0x93, 0xb9, 0x5a, 0xfa, 0x3d, 0x57, 0x4b, 0xef, 0x1e, 0x15, 0xf2, 0x2c, 0x62, 0x75, 0xc4, 0xfa,
	0xeb, 0x15, 0x56, 0xe5, 0xa4, 0xb0, 0x2c, 0x45, 0xbc, 0x55, 0x11, 0xcb, 0xed, 0xf1, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x9e, 0xcc, 0x69, 0x8a, 0x51, 0x05, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.HeartbeatRewardThreshold.Size()
		i -= size
		if _, err := m.HeartbeatRewardThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.HeartbeatRewardRate.Size()
		i -= size
		if _, err := m.HeartbeatRewardRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.HeartbeatInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HeartbeatInterval))
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.RewardChallengerRatio.Size()
		i -= size
		if _, err := m.RewardChallengerRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.RewardValidatorRatio.Size()
		i -= size
		if _, err := m.RewardValidatorRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.SlashAmountMax.Size()
		i -= size
		if _, err := m.SlashAmountMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SlashAmountMin.Size()
		i -= size
		if _, err := m.SlashAmountMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SlashAmountSizeRate.Size()
		i -= size
		if _, err := m.SlashAmountSizeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.SlashCoolingOffPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SlashCoolingOffPeriod))
		i--
		dAtA[i] = 0x10
	}
	if m.ChallengeCountPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ChallengeCountPerBlock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChallengeCountPerBlock != 0 {
		n += 1 + sovParams(uint64(m.ChallengeCountPerBlock))
	}
	if m.SlashCoolingOffPeriod != 0 {
		n += 1 + sovParams(uint64(m.SlashCoolingOffPeriod))
	}
	l = m.SlashAmountSizeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashAmountMin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashAmountMax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardValidatorRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.RewardChallengerRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.HeartbeatInterval != 0 {
		n += 1 + sovParams(uint64(m.HeartbeatInterval))
	}
	l = m.HeartbeatRewardRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.HeartbeatRewardThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeCountPerBlock", wireType)
			}
			m.ChallengeCountPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeCountPerBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashCoolingOffPeriod", wireType)
			}
			m.SlashCoolingOffPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashCoolingOffPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountSizeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountSizeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmountMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashAmountMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardValidatorRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardValidatorRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardChallengerRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardChallengerRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatInterval", wireType)
			}
			m.HeartbeatInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeartbeatInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatRewardRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HeartbeatRewardRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatRewardThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HeartbeatRewardThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)

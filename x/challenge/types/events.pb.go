// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/challenge/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type EventStartChallenge struct {
	// The id of challenge, which is generated by blockchain.
	ChallengeId uint64 `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	// The id of object info to be challenged.
	ObjectId uint64 `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	// The segment/piece index of the object info.
	SegmentIndex uint32 `protobuf:"varint,3,opt,name=segment_index,json=segmentIndex,proto3" json:"segment_index,omitempty"`
	// The storage provider to be challenged.
	SpOperatorAddress string `protobuf:"bytes,4,opt,name=sp_operator_address,json=spOperatorAddress,proto3" json:"sp_operator_address,omitempty"`
	// The redundancy index, which comes from the index of storage providers.
	RedundancyIndex int32 `protobuf:"varint,5,opt,name=redundancy_index,json=redundancyIndex,proto3" json:"redundancy_index,omitempty"`
}

func (m *EventStartChallenge) Reset()         { *m = EventStartChallenge{} }
func (m *EventStartChallenge) String() string { return proto.CompactTextString(m) }
func (*EventStartChallenge) ProtoMessage()    {}
func (*EventStartChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9eaa4bfadaa20f8, []int{0}
}
func (m *EventStartChallenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventStartChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventStartChallenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventStartChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStartChallenge.Merge(m, src)
}
func (m *EventStartChallenge) XXX_Size() int {
	return m.Size()
}
func (m *EventStartChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStartChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_EventStartChallenge proto.InternalMessageInfo

func (m *EventStartChallenge) GetChallengeId() uint64 {
	if m != nil {
		return m.ChallengeId
	}
	return 0
}

func (m *EventStartChallenge) GetObjectId() uint64 {
	if m != nil {
		return m.ObjectId
	}
	return 0
}

func (m *EventStartChallenge) GetSegmentIndex() uint32 {
	if m != nil {
		return m.SegmentIndex
	}
	return 0
}

func (m *EventStartChallenge) GetSpOperatorAddress() string {
	if m != nil {
		return m.SpOperatorAddress
	}
	return ""
}

func (m *EventStartChallenge) GetRedundancyIndex() int32 {
	if m != nil {
		return m.RedundancyIndex
	}
	return 0
}

type EventCompleteChallenge struct {
	// The id of challenge.
	ChallengeId uint64 `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	// The result of challenge.
	Result uint32 `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	// The slashed storage provider address.
	SpOperatorAddress string `protobuf:"bytes,3,opt,name=sp_operator_address,json=spOperatorAddress,proto3" json:"sp_operator_address,omitempty"`
	// The slashed amount from the storage provider.
	SlashAmount string `protobuf:"bytes,4,opt,name=slash_amount,json=slashAmount,proto3" json:"slash_amount,omitempty"`
	// The address of challenger.
	ChallengerAddress string `protobuf:"bytes,5,opt,name=challenger_address,json=challengerAddress,proto3" json:"challenger_address,omitempty"`
	// The reward amount to the challenger.
	ChallengerRewardAmount string `protobuf:"bytes,6,opt,name=challenger_reward_amount,json=challengerRewardAmount,proto3" json:"challenger_reward_amount,omitempty"`
	// The submitter of challenge attestation.
	SubmitterAddress string `protobuf:"bytes,7,opt,name=submitter_address,json=submitterAddress,proto3" json:"submitter_address,omitempty"`
	// The reward amount to the submitter.
	SubmitterRewardAmount string `protobuf:"bytes,8,opt,name=submitter_reward_amount,json=submitterRewardAmount,proto3" json:"submitter_reward_amount,omitempty"`
	// The addresses of validators participated in the attestation.
	ValidatorAddresses []string `protobuf:"bytes,9,rep,name=validator_addresses,json=validatorAddresses,proto3" json:"validator_addresses,omitempty"`
	// The reward amount to the each validator.
	ValidatorRewardAmount string `protobuf:"bytes,10,opt,name=validator_reward_amount,json=validatorRewardAmount,proto3" json:"validator_reward_amount,omitempty"`
}

func (m *EventCompleteChallenge) Reset()         { *m = EventCompleteChallenge{} }
func (m *EventCompleteChallenge) String() string { return proto.CompactTextString(m) }
func (*EventCompleteChallenge) ProtoMessage()    {}
func (*EventCompleteChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9eaa4bfadaa20f8, []int{1}
}
func (m *EventCompleteChallenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCompleteChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCompleteChallenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCompleteChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCompleteChallenge.Merge(m, src)
}
func (m *EventCompleteChallenge) XXX_Size() int {
	return m.Size()
}
func (m *EventCompleteChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCompleteChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_EventCompleteChallenge proto.InternalMessageInfo

func (m *EventCompleteChallenge) GetChallengeId() uint64 {
	if m != nil {
		return m.ChallengeId
	}
	return 0
}

func (m *EventCompleteChallenge) GetResult() uint32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *EventCompleteChallenge) GetSpOperatorAddress() string {
	if m != nil {
		return m.SpOperatorAddress
	}
	return ""
}

func (m *EventCompleteChallenge) GetSlashAmount() string {
	if m != nil {
		return m.SlashAmount
	}
	return ""
}

func (m *EventCompleteChallenge) GetChallengerAddress() string {
	if m != nil {
		return m.ChallengerAddress
	}
	return ""
}

func (m *EventCompleteChallenge) GetChallengerRewardAmount() string {
	if m != nil {
		return m.ChallengerRewardAmount
	}
	return ""
}

func (m *EventCompleteChallenge) GetSubmitterAddress() string {
	if m != nil {
		return m.SubmitterAddress
	}
	return ""
}

func (m *EventCompleteChallenge) GetSubmitterRewardAmount() string {
	if m != nil {
		return m.SubmitterRewardAmount
	}
	return ""
}

func (m *EventCompleteChallenge) GetValidatorAddresses() []string {
	if m != nil {
		return m.ValidatorAddresses
	}
	return nil
}

func (m *EventCompleteChallenge) GetValidatorRewardAmount() string {
	if m != nil {
		return m.ValidatorRewardAmount
	}
	return ""
}

type EventChallengeHeartbeat struct {
	// The id of challenge.
	ChallengeId uint64 `protobuf:"varint,1,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
}

func (m *EventChallengeHeartbeat) Reset()         { *m = EventChallengeHeartbeat{} }
func (m *EventChallengeHeartbeat) String() string { return proto.CompactTextString(m) }
func (*EventChallengeHeartbeat) ProtoMessage()    {}
func (*EventChallengeHeartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9eaa4bfadaa20f8, []int{2}
}
func (m *EventChallengeHeartbeat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventChallengeHeartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventChallengeHeartbeat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventChallengeHeartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventChallengeHeartbeat.Merge(m, src)
}
func (m *EventChallengeHeartbeat) XXX_Size() int {
	return m.Size()
}
func (m *EventChallengeHeartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_EventChallengeHeartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_EventChallengeHeartbeat proto.InternalMessageInfo

func (m *EventChallengeHeartbeat) GetChallengeId() uint64 {
	if m != nil {
		return m.ChallengeId
	}
	return 0
}

func init() {
	proto.RegisterType((*EventStartChallenge)(nil), "bnbchain.greenfield.sp.EventStartChallenge")
	proto.RegisterType((*EventCompleteChallenge)(nil), "bnbchain.greenfield.sp.EventCompleteChallenge")
	proto.RegisterType((*EventChallengeHeartbeat)(nil), "bnbchain.greenfield.sp.EventChallengeHeartbeat")
}

func init() { proto.RegisterFile("greenfield/challenge/events.proto", fileDescriptor_e9eaa4bfadaa20f8) }

var fileDescriptor_e9eaa4bfadaa20f8 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x63, 0xf2, 0x43, 0x33, 0x4d, 0x44, 0xeb, 0x40, 0x6a, 0x40, 0xb2, 0x9c, 0xb0, 0x09,
	0x8b, 0xc6, 0x42, 0x48, 0x88, 0x05, 0x9b, 0xb4, 0xaa, 0x68, 0xc4, 0x02, 0xc9, 0xdd, 0xb1, 0xb1,
	0x66, 0x3c, 0x17, 0xc7, 0xc8, 0x9e, 0xb1, 0x66, 0xc6, 0xa5, 0x7d, 0x0b, 0x1e, 0x86, 0x57, 0x40,
	0x62, 0x59, 0xb1, 0x62, 0x89, 0x92, 0x07, 0xe0, 0x15, 0x50, 0xc6, 0xbf, 0x59, 0x40, 0xe9, 0x72,
	0xce, 0x99, 0x7b, 0xcf, 0x77, 0x47, 0x73, 0xd1, 0x24, 0x14, 0x00, 0xec, 0x63, 0x04, 0x31, 0x75,
	0x83, 0x15, 0x8e, 0x63, 0x60, 0x21, 0xb8, 0x70, 0x09, 0x4c, 0xc9, 0x79, 0x2a, 0xb8, 0xe2, 0xe6,
	0x98, 0x30, 0x12, 0xac, 0x70, 0xc4, 0xe6, 0xf5, 0xdd, 0xb9, 0x4c, 0x9f, 0x3c, 0x0e, 0xb8, 0x4c,
	0xb8, 0xf4, 0xf5, 0x2d, 0x37, 0x3f, 0xe4, 0x25, 0xd3, 0xdf, 0x06, 0x1a, 0x9d, 0x6d, 0x7b, 0x5c,
	0x28, 0x2c, 0xd4, 0x69, 0xd9, 0xd7, 0x9c, 0xa0, 0x41, 0x15, 0xe2, 0x47, 0xd4, 0x32, 0x1c, 0x63,
	0xd6, 0xf1, 0xf6, 0x2b, 0x6d, 0x49, 0xcd, 0xa7, 0xa8, 0xcf, 0xc9, 0x27, 0x08, 0xd4, 0xd6, 0xbf,
	0xa7, 0xfd, 0xbd, 0x5c, 0x58, 0x52, 0xf3, 0x19, 0x1a, 0x4a, 0x08, 0x13, 0x60, 0xca, 0x8f, 0x18,
	0x85, 0x2b, 0xab, 0xed, 0x18, 0xb3, 0xa1, 0x37, 0x28, 0xc4, 0xe5, 0x56, 0x33, 0xcf, 0xd1, 0x48,
	0xa6, 0x3e, 0x4f, 0x41, 0x60, 0xc5, 0x85, 0x8f, 0x29, 0x15, 0x20, 0xa5, 0xd5, 0x71, 0x8c, 0x59,
	0xff, 0xc4, 0xfa, 0xf1, 0xf5, 0xf8, 0x61, 0xc1, 0xba, 0xc8, 0x9d, 0x0b, 0x25, 0x22, 0x16, 0x7a,
	0x87, 0x32, 0x7d, 0x5f, 0xd4, 0x14, 0x86, 0xf9, 0x1c, 0x1d, 0x08, 0xa0, 0x19, 0xa3, 0x98, 0x05,
	0xd7, 0x45, 0x62, 0xd7, 0x31, 0x66, 0x5d, 0xef, 0x41, 0xad, 0xeb, 0xd0, 0xe9, 0xb7, 0x0e, 0x1a,
	0xeb, 0x89, 0x4f, 0x79, 0x92, 0xc6, 0xa0, 0xe0, 0x4e, 0x43, 0x8f, 0x51, 0x4f, 0x80, 0xcc, 0x62,
	0xa5, 0x27, 0x1e, 0x7a, 0xc5, 0xe9, 0x6f, 0xa3, 0xb4, 0xef, 0x3e, 0xca, 0x04, 0x0d, 0x64, 0x8c,
	0xe5, 0xca, 0xc7, 0x09, 0xcf, 0x98, 0xca, 0x5f, 0xc3, 0xdb, 0xd7, 0xda, 0x42, 0x4b, 0xe6, 0x5b,
	0x64, 0x56, 0x4c, 0x75, 0x56, 0xf7, 0xb6, 0xac, 0xba, 0xa6, 0xcc, 0x7a, 0x8d, 0xac, 0x46, 0x23,
	0x01, 0x9f, 0xb1, 0xa0, 0x65, 0x6e, 0x4f, 0xe7, 0x8e, 0x6b, 0xdf, 0xd3, 0x76, 0x81, 0x70, 0x86,
	0x0e, 0x65, 0x46, 0x92, 0x48, 0xa9, 0x06, 0xc1, 0xfd, 0x5b, 0x08, 0x0e, 0xaa, 0x92, 0x12, 0xe0,
	0x15, 0x3a, 0xaa, 0xdb, 0xec, 0xe6, 0xef, 0xe9, 0xfc, 0x47, 0x95, 0xbd, 0x13, 0xbf, 0x44, 0xa3,
	0x4b, 0x1c, 0x47, 0xb4, 0xf9, 0xd8, 0x20, 0xad, 0xbe, 0xd3, 0xfe, 0x27, 0x80, 0x59, 0x15, 0x2d,
	0xca, 0x9a, 0x2d, 0x42, 0xdd, 0x6a, 0x17, 0x01, 0xe5, 0x08, 0x95, 0xdd, 0x44, 0x98, 0xbe, 0x41,
	0x47, 0xf9, 0x37, 0x2a, 0x1f, 0xe8, 0x1c, 0xb0, 0x50, 0x04, 0xb0, 0xfa, 0x8f, 0x7f, 0x74, 0xf2,
	0xee, 0xfb, 0xda, 0x36, 0x6e, 0xd6, 0xb6, 0xf1, 0x6b, 0x6d, 0x1b, 0x5f, 0x36, 0x76, 0xeb, 0x66,
	0x63, 0xb7, 0x7e, 0x6e, 0xec, 0xd6, 0x87, 0x17, 0x61, 0xa4, 0x56, 0x19, 0x99, 0x07, 0x3c, 0x71,
	0x09, 0x23, 0xc7, 0x7a, 0xa1, 0xdd, 0xc6, 0xf2, 0x5f, 0x35, 0xd6, 0x5f, 0x5d, 0xa7, 0x20, 0x49,
	0x4f, 0xef, 0xf2, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x65, 0x89, 0x4a, 0x23, 0x04,
	0x00, 0x00,
}

func (m *EventStartChallenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventStartChallenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventStartChallenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RedundancyIndex != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.RedundancyIndex))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SpOperatorAddress) > 0 {
		i -= len(m.SpOperatorAddress)
		copy(dAtA[i:], m.SpOperatorAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SpOperatorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.SegmentIndex != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SegmentIndex))
		i--
		dAtA[i] = 0x18
	}
	if m.ObjectId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ObjectId))
		i--
		dAtA[i] = 0x10
	}
	if m.ChallengeId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChallengeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventCompleteChallenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCompleteChallenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCompleteChallenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorRewardAmount) > 0 {
		i -= len(m.ValidatorRewardAmount)
		copy(dAtA[i:], m.ValidatorRewardAmount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ValidatorRewardAmount)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.ValidatorAddresses) > 0 {
		for iNdEx := len(m.ValidatorAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ValidatorAddresses[iNdEx])
			copy(dAtA[i:], m.ValidatorAddresses[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.ValidatorAddresses[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.SubmitterRewardAmount) > 0 {
		i -= len(m.SubmitterRewardAmount)
		copy(dAtA[i:], m.SubmitterRewardAmount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SubmitterRewardAmount)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SubmitterAddress) > 0 {
		i -= len(m.SubmitterAddress)
		copy(dAtA[i:], m.SubmitterAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SubmitterAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChallengerRewardAmount) > 0 {
		i -= len(m.ChallengerRewardAmount)
		copy(dAtA[i:], m.ChallengerRewardAmount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ChallengerRewardAmount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ChallengerAddress) > 0 {
		i -= len(m.ChallengerAddress)
		copy(dAtA[i:], m.ChallengerAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ChallengerAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SlashAmount) > 0 {
		i -= len(m.SlashAmount)
		copy(dAtA[i:], m.SlashAmount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SlashAmount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SpOperatorAddress) > 0 {
		i -= len(m.SpOperatorAddress)
		copy(dAtA[i:], m.SpOperatorAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.SpOperatorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Result != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	if m.ChallengeId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChallengeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventChallengeHeartbeat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventChallengeHeartbeat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventChallengeHeartbeat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChallengeId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ChallengeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventStartChallenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChallengeId != 0 {
		n += 1 + sovEvents(uint64(m.ChallengeId))
	}
	if m.ObjectId != 0 {
		n += 1 + sovEvents(uint64(m.ObjectId))
	}
	if m.SegmentIndex != 0 {
		n += 1 + sovEvents(uint64(m.SegmentIndex))
	}
	l = len(m.SpOperatorAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.RedundancyIndex != 0 {
		n += 1 + sovEvents(uint64(m.RedundancyIndex))
	}
	return n
}

func (m *EventCompleteChallenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChallengeId != 0 {
		n += 1 + sovEvents(uint64(m.ChallengeId))
	}
	if m.Result != 0 {
		n += 1 + sovEvents(uint64(m.Result))
	}
	l = len(m.SpOperatorAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SlashAmount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ChallengerAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ChallengerRewardAmount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SubmitterAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.SubmitterRewardAmount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.ValidatorAddresses) > 0 {
		for _, s := range m.ValidatorAddresses {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	l = len(m.ValidatorRewardAmount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventChallengeHeartbeat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChallengeId != 0 {
		n += 1 + sovEvents(uint64(m.ChallengeId))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventStartChallenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventStartChallenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventStartChallenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeId", wireType)
			}
			m.ChallengeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectId", wireType)
			}
			m.ObjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObjectId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SegmentIndex", wireType)
			}
			m.SegmentIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SegmentIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpOperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpOperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedundancyIndex", wireType)
			}
			m.RedundancyIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RedundancyIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCompleteChallenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCompleteChallenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCompleteChallenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeId", wireType)
			}
			m.ChallengeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpOperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpOperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengerRewardAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengerRewardAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmitterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitterRewardAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmitterRewardAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddresses = append(m.ValidatorAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorRewardAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorRewardAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventChallengeHeartbeat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventChallengeHeartbeat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventChallengeHeartbeat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeId", wireType)
			}
			m.ChallengeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChallengeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)

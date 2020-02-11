// Code generated by protoc-gen-go. DO NOT EDIT.
// source: x/slashing/internal/types/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// MsgUnjail - struct for unjailing jailed validator
type MsgUnjail struct {
	ValidatorAddr        []byte   `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgUnjail) Reset()         { *m = MsgUnjail{} }
func (m *MsgUnjail) String() string { return proto.CompactTextString(m) }
func (*MsgUnjail) ProtoMessage()    {}
func (*MsgUnjail) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b882c3b0cdd6f57, []int{0}
}

func (m *MsgUnjail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgUnjail.Unmarshal(m, b)
}
func (m *MsgUnjail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgUnjail.Marshal(b, m, deterministic)
}
func (m *MsgUnjail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnjail.Merge(m, src)
}
func (m *MsgUnjail) XXX_Size() int {
	return xxx_messageInfo_MsgUnjail.Size(m)
}
func (m *MsgUnjail) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnjail.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnjail proto.InternalMessageInfo

func (m *MsgUnjail) GetValidatorAddr() []byte {
	if m != nil {
		return m.ValidatorAddr
	}
	return nil
}

// ValidatorSigningInfo defines the signing info for a validator
type ValidatorSigningInfo struct {
	Address              []byte           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	StartHeight          int64            `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	IndexOffset          int64            `protobuf:"varint,3,opt,name=index_offset,json=indexOffset,proto3" json:"index_offset,omitempty"`
	JailedUntil          *types.Timestamp `protobuf:"bytes,4,opt,name=jailed_until,json=jailedUntil,proto3" json:"jailed_until,omitempty"`
	Tombstoned           bool             `protobuf:"varint,5,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	MissedBlocksCounter  int64            `protobuf:"varint,6,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ValidatorSigningInfo) Reset()         { *m = ValidatorSigningInfo{} }
func (m *ValidatorSigningInfo) String() string { return proto.CompactTextString(m) }
func (*ValidatorSigningInfo) ProtoMessage()    {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b882c3b0cdd6f57, []int{1}
}

func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorSigningInfo.Unmarshal(m, b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return xxx_messageInfo_ValidatorSigningInfo.Size(m)
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetIndexOffset() int64 {
	if m != nil {
		return m.IndexOffset
	}
	return 0
}

func (m *ValidatorSigningInfo) GetJailedUntil() *types.Timestamp {
	if m != nil {
		return m.JailedUntil
	}
	return nil
}

func (m *ValidatorSigningInfo) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgUnjail)(nil), "cosmos_sdk.x.slashing.v1.MsgUnjail")
	proto.RegisterType((*ValidatorSigningInfo)(nil), "cosmos_sdk.x.slashing.v1.ValidatorSigningInfo")
}

func init() {
	proto.RegisterFile("x/slashing/internal/types/types.proto", fileDescriptor_2b882c3b0cdd6f57)
}

var fileDescriptor_2b882c3b0cdd6f57 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x26, 0x84, 0xb6, 0xb0, 0x09, 0x95, 0x70, 0xa9, 0xb0, 0x22, 0x24, 0x57, 0x96, 0x40, 0x5c,
	0xba, 0x16, 0x20, 0x2e, 0xbd, 0xe1, 0x72, 0x80, 0x43, 0x85, 0x64, 0xda, 0x1e, 0x38, 0x60, 0xad,
	0xe3, 0xcd, 0x7a, 0xe9, 0x7a, 0x27, 0xf2, 0x4e, 0xaa, 0xe4, 0xc8, 0x1b, 0xf0, 0x58, 0xbc, 0x03,
	0x52, 0x78, 0x07, 0x8e, 0x9c, 0x58, 0xef, 0xc6, 0xc4, 0x20, 0x0e, 0xbd, 0xd8, 0x9e, 0x6f, 0xe6,
	0xfb, 0xbe, 0xf9, 0x31, 0x79, 0xb2, 0x4c, 0x8c, 0x62, 0xa6, 0x92, 0x5a, 0x24, 0x52, 0x23, 0x6f,
	0x34, 0x53, 0x09, 0xae, 0xe6, 0xdc, 0xf8, 0x27, 0x9d, 0x37, 0x80, 0x10, 0x84, 0x53, 0x30, 0x35,
	0x98, 0xdc, 0x94, 0x57, 0x74, 0x49, 0x3b, 0x06, 0xbd, 0x7e, 0x3e, 0x79, 0x8a, 0x95, 0x6c, 0xca,
	0x7c, 0xce, 0x1a, 0x5c, 0x25, 0xae, 0x38, 0x11, 0x20, 0x60, 0xfb, 0xe5, 0x15, 0x26, 0x91, 0x00,
	0x10, 0x8a, 0xfb, 0x92, 0x62, 0x31, 0x4b, 0x50, 0xd6, 0xdc, 0x20, 0xab, 0xe7, 0xbe, 0x20, 0xfe,
	0x32, 0x20, 0xf7, 0xce, 0x8c, 0xb8, 0xd0, 0x9f, 0x99, 0x54, 0x01, 0x92, 0xfd, 0x6b, 0xa6, 0x64,
	0xc9, 0x10, 0x9a, 0x9c, 0x95, 0x65, 0x13, 0x0e, 0x8e, 0x06, 0xcf, 0xc6, 0xe9, 0xd9, 0xcf, 0x75,
	0x14, 0xae, 0x58, 0xad, 0x4e, 0xe2, 0xbf, 0xf3, 0xdc, 0x98, 0xf8, 0xd7, 0x3a, 0x3a, 0x16, 0x12,
	0xab, 0x45, 0x41, 0xa7, 0x50, 0x27, 0xbe, 0xe7, 0xcd, 0xeb, 0xd8, 0xb6, 0xbe, 0x19, 0xe9, 0x92,
	0xa9, 0xd7, 0x9e, 0x91, 0xdd, 0xff, 0x23, 0xd2, 0x22, 0xf1, 0xf7, 0x21, 0x79, 0x78, 0xd9, 0x21,
	0x1f, 0xa4, 0xd0, 0x76, 0xc8, 0x77, 0x7a, 0x06, 0xc1, 0x27, 0xb2, 0xb7, 0x31, 0xd9, 0xf4, 0xf1,
	0xc6, 0xf6, 0xb1, 0xef, 0xfb, 0xe8, 0xb9, 0xd3, 0x1b, 0xb8, 0x9f, 0x82, 0x36, 0x9d, 0x7d, 0x27,
	0x1a, 0x9c, 0x90, 0xb1, 0xdd, 0x45, 0x83, 0x79, 0xc5, 0xa5, 0xa8, 0x30, 0xbc, 0x6d, 0x4d, 0x86,
	0xe9, 0x23, 0x6b, 0x72, 0xe0, 0x4d, 0xfa, 0xd9, 0x38, 0x1b, 0xb9, 0xf0, 0xad, 0x8b, 0x5a, 0xae,
	0xd4, 0x25, 0x5f, 0xe6, 0x30, 0x9b, 0x19, 0x8e, 0xe1, 0xf0, 0x5f, 0x6e, 0x3f, 0x6b, 0xb9, 0x2e,
	0x7c, 0xef, 0x22, 0x3b, 0xd7, 0xb8, 0x5d, 0x37, 0x2f, 0xf3, 0x85, 0x46, 0xa9, 0xc2, 0x3b, 0x96,
	0x3b, 0x7a, 0x31, 0xa1, 0xfe, 0x58, 0xb4, 0x3b, 0x16, 0x3d, 0xef, 0x8e, 0x95, 0x46, 0xdf, 0xd6,
	0xd1, 0xad, 0xad, 0x76, 0x9f, 0x1d, 0x7f, 0xfd, 0x11, 0x0d, 0xb2, 0x91, 0x87, 0x2e, 0x5a, 0x24,
	0x78, 0x45, 0x08, 0x42, 0x5d, 0x18, 0x04, 0xcd, 0xcb, 0x70, 0xc7, 0xaa, 0xdf, 0x4d, 0x0f, 0x2d,
	0xfb, 0x81, 0x67, 0x6f, 0x73, 0x71, 0xd6, 0x2b, 0x0c, 0xce, 0xc9, 0x61, 0x2d, 0x8d, 0xb1, 0xc2,
	0x85, 0x82, 0xe9, 0x95, 0xc9, 0xa7, 0xb0, 0x68, 0x7f, 0xce, 0x70, 0xd7, 0xcd, 0x76, 0x64, 0x15,
	0x1e, 0x7b, 0x85, 0xff, 0x96, 0xc5, 0xd9, 0x81, 0xc7, 0x53, 0x07, 0x9f, 0x7a, 0x34, 0xdd, 0xfb,
	0xb8, 0xe3, 0x4e, 0x50, 0xec, 0xba, 0xb9, 0x5e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x07, 0x25,
	0xe1, 0xf1, 0xfd, 0x02, 0x00, 0x00,
}

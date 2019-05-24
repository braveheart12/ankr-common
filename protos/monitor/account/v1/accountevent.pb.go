// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/v1/accountevent.proto

package account

import (
	fmt "fmt"
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

type AccountEvent struct {
	Address              []string `protobuf:"bytes,1,rep,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountEvent) Reset()         { *m = AccountEvent{} }
func (m *AccountEvent) String() string { return proto.CompactTextString(m) }
func (*AccountEvent) ProtoMessage()    {}
func (*AccountEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_adf0fa52adfff1d5, []int{0}
}

func (m *AccountEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEvent.Unmarshal(m, b)
}
func (m *AccountEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEvent.Marshal(b, m, deterministic)
}
func (m *AccountEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEvent.Merge(m, src)
}
func (m *AccountEvent) XXX_Size() int {
	return xxx_messageInfo_AccountEvent.Size(m)
}
func (m *AccountEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEvent proto.InternalMessageInfo

func (m *AccountEvent) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountEvent)(nil), "account.AccountEvent")
}

func init() { proto.RegisterFile("account/v1/accountevent.proto", fileDescriptor_adf0fa52adfff1d5) }

var fileDescriptor_adf0fa52adfff1d5 = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0xd4, 0x87, 0x32, 0x53, 0xcb, 0x52, 0xf3, 0x4a, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8, 0xa1, 0x62, 0x4a, 0x1a, 0x5c, 0x3c, 0x8e, 0x10, 0xa6, 0x2b,
	0x48, 0x5a, 0x48, 0x82, 0x8b, 0xdd, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81,
	0x59, 0x83, 0x33, 0x08, 0xc6, 0x4d, 0x62, 0x03, 0xeb, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x5e, 0xfa, 0x57, 0x5e, 0x5a, 0x00, 0x00, 0x00,
}
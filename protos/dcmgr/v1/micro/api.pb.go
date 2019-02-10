// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcmgr/v1/micro/api.proto

package dcmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/Ankr-network/dccn-common/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DataCenterListRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenterListRequest) Reset()         { *m = DataCenterListRequest{} }
func (m *DataCenterListRequest) String() string { return proto.CompactTextString(m) }
func (*DataCenterListRequest) ProtoMessage()    {}
func (*DataCenterListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_33e920ed5e7feec4, []int{0}
}
func (m *DataCenterListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterListRequest.Unmarshal(m, b)
}
func (m *DataCenterListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterListRequest.Marshal(b, m, deterministic)
}
func (dst *DataCenterListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterListRequest.Merge(dst, src)
}
func (m *DataCenterListRequest) XXX_Size() int {
	return xxx_messageInfo_DataCenterListRequest.Size(m)
}
func (m *DataCenterListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterListRequest proto.InternalMessageInfo

func (m *DataCenterListRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type DataCenterListResponse struct {
	DcList               []*common.DataCenter `protobuf:"bytes,1,rep,name=dcList,proto3" json:"dcList,omitempty"`
	Error                *common.Error        `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DataCenterListResponse) Reset()         { *m = DataCenterListResponse{} }
func (m *DataCenterListResponse) String() string { return proto.CompactTextString(m) }
func (*DataCenterListResponse) ProtoMessage()    {}
func (*DataCenterListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_33e920ed5e7feec4, []int{1}
}
func (m *DataCenterListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterListResponse.Unmarshal(m, b)
}
func (m *DataCenterListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterListResponse.Marshal(b, m, deterministic)
}
func (dst *DataCenterListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterListResponse.Merge(dst, src)
}
func (m *DataCenterListResponse) XXX_Size() int {
	return xxx_messageInfo_DataCenterListResponse.Size(m)
}
func (m *DataCenterListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterListResponse proto.InternalMessageInfo

func (m *DataCenterListResponse) GetDcList() []*common.DataCenter {
	if m != nil {
		return m.DcList
	}
	return nil
}

func (m *DataCenterListResponse) GetError() *common.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*DataCenterListRequest)(nil), "dcmgr.DataCenterListRequest")
	proto.RegisterType((*DataCenterListResponse)(nil), "dcmgr.DataCenterListResponse")
}

func init() { proto.RegisterFile("dcmgr/v1/micro/api.proto", fileDescriptor_api_33e920ed5e7feec4) }

var fileDescriptor_api_33e920ed5e7feec4 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x49, 0xce, 0x4d,
	0x2f, 0xd2, 0x2f, 0x33, 0xd4, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xcb, 0x48, 0x09, 0x27, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7,
	0xe9, 0x43, 0x28, 0x88, 0x9c, 0x94, 0x10, 0x54, 0x30, 0xb5, 0xa8, 0x28, 0xbf, 0x08, 0x22, 0xa6,
	0x64, 0xc0, 0x25, 0xea, 0x92, 0x58, 0x92, 0xe8, 0x9c, 0x9a, 0x57, 0x92, 0x5a, 0xe4, 0x93, 0x59,
	0x5c, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xce, 0xc5, 0x5e, 0x5a, 0x9c, 0x5a,
	0x14, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x06, 0xe2, 0x7a, 0xa6, 0x28,
	0x95, 0x72, 0x89, 0xa1, 0xeb, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x32, 0xe0, 0x62, 0x4b,
	0x49, 0x06, 0x89, 0x48, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x49, 0xe8, 0x21, 0x5b, 0xaf, 0x87,
	0xd0, 0x15, 0x04, 0x55, 0x27, 0xa4, 0xc9, 0xc5, 0x0a, 0x76, 0x8c, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0x30, 0xaa, 0x06, 0x57, 0x90, 0x54, 0x10, 0x44, 0x85, 0x51, 0x04, 0x17, 0xab, 0x8b,
	0xb3, 0x63, 0x80, 0xa7, 0x90, 0x3f, 0x17, 0x1f, 0xaa, 0xfd, 0x42, 0x32, 0x7a, 0x60, 0x4f, 0xeb,
	0x61, 0xf5, 0x88, 0x94, 0x2c, 0x0e, 0x59, 0x88, 0xa3, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xb6, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xb0, 0x7e, 0xa6, 0x55, 0x01, 0x00, 0x00,
}

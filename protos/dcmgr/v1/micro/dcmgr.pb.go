// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcmgr/v1/micro/dcmgr.proto

package dcmgr

import (
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
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

type DataCenterListResponse struct {
	DcList               []*common.DataCenter `protobuf:"bytes,1,rep,name=dcList,proto3" json:"dcList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DataCenterListResponse) Reset()         { *m = DataCenterListResponse{} }
func (m *DataCenterListResponse) String() string { return proto.CompactTextString(m) }
func (*DataCenterListResponse) ProtoMessage()    {}
func (*DataCenterListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{0}
}

func (m *DataCenterListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterListResponse.Unmarshal(m, b)
}
func (m *DataCenterListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterListResponse.Marshal(b, m, deterministic)
}
func (m *DataCenterListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterListResponse.Merge(m, src)
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

type NetworkInfoResponse struct {
	UserCount            int32    `protobuf:"varint,1,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
	HostCount            int32    `protobuf:"varint,2,opt,name=host_count,json=hostCount,proto3" json:"host_count,omitempty"`
	EnvironmentCount     int32    `protobuf:"varint,3,opt,name=environment_count,json=environmentCount,proto3" json:"environment_count,omitempty"`
	ContainerCount       int32    `protobuf:"varint,4,opt,name=container_count,json=containerCount,proto3" json:"container_count,omitempty"`
	Traffic              int32    `protobuf:"varint,5,opt,name=traffic,proto3" json:"traffic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInfoResponse) Reset()         { *m = NetworkInfoResponse{} }
func (m *NetworkInfoResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInfoResponse) ProtoMessage()    {}
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{1}
}

func (m *NetworkInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInfoResponse.Unmarshal(m, b)
}
func (m *NetworkInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInfoResponse.Marshal(b, m, deterministic)
}
func (m *NetworkInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfoResponse.Merge(m, src)
}
func (m *NetworkInfoResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkInfoResponse.Size(m)
}
func (m *NetworkInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfoResponse proto.InternalMessageInfo

func (m *NetworkInfoResponse) GetUserCount() int32 {
	if m != nil {
		return m.UserCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetHostCount() int32 {
	if m != nil {
		return m.HostCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetEnvironmentCount() int32 {
	if m != nil {
		return m.EnvironmentCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetContainerCount() int32 {
	if m != nil {
		return m.ContainerCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetTraffic() int32 {
	if m != nil {
		return m.Traffic
	}
	return 0
}

type DataCenterLeaderBoardResponse struct {
	List                 []*DataCenterLeaderBoardDetail `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *DataCenterLeaderBoardResponse) Reset()         { *m = DataCenterLeaderBoardResponse{} }
func (m *DataCenterLeaderBoardResponse) String() string { return proto.CompactTextString(m) }
func (*DataCenterLeaderBoardResponse) ProtoMessage()    {}
func (*DataCenterLeaderBoardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{2}
}

func (m *DataCenterLeaderBoardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Unmarshal(m, b)
}
func (m *DataCenterLeaderBoardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Marshal(b, m, deterministic)
}
func (m *DataCenterLeaderBoardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterLeaderBoardResponse.Merge(m, src)
}
func (m *DataCenterLeaderBoardResponse) XXX_Size() int {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Size(m)
}
func (m *DataCenterLeaderBoardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterLeaderBoardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterLeaderBoardResponse proto.InternalMessageInfo

func (m *DataCenterLeaderBoardResponse) GetList() []*DataCenterLeaderBoardDetail {
	if m != nil {
		return m.List
	}
	return nil
}

type DataCenterLeaderBoardDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               float64  `protobuf:"fixed64,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenterLeaderBoardDetail) Reset()         { *m = DataCenterLeaderBoardDetail{} }
func (m *DataCenterLeaderBoardDetail) String() string { return proto.CompactTextString(m) }
func (*DataCenterLeaderBoardDetail) ProtoMessage()    {}
func (*DataCenterLeaderBoardDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_1750e4d9ba65142d, []int{3}
}

func (m *DataCenterLeaderBoardDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Unmarshal(m, b)
}
func (m *DataCenterLeaderBoardDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Marshal(b, m, deterministic)
}
func (m *DataCenterLeaderBoardDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterLeaderBoardDetail.Merge(m, src)
}
func (m *DataCenterLeaderBoardDetail) XXX_Size() int {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Size(m)
}
func (m *DataCenterLeaderBoardDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterLeaderBoardDetail.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterLeaderBoardDetail proto.InternalMessageInfo

func (m *DataCenterLeaderBoardDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenterLeaderBoardDetail) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*DataCenterListResponse)(nil), "dcmgr.DataCenterListResponse")
	proto.RegisterType((*NetworkInfoResponse)(nil), "dcmgr.NetworkInfoResponse")
	proto.RegisterType((*DataCenterLeaderBoardResponse)(nil), "dcmgr.DataCenterLeaderBoardResponse")
	proto.RegisterType((*DataCenterLeaderBoardDetail)(nil), "dcmgr.DataCenterLeaderBoardDetail")
}

func init() { proto.RegisterFile("dcmgr/v1/micro/dcmgr.proto", fileDescriptor_1750e4d9ba65142d) }

var fileDescriptor_1750e4d9ba65142d = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xde, 0xb4, 0xbb, 0xab, 0x7b, 0x0a, 0x55, 0x67, 0xb1, 0x2c, 0x29, 0x2b, 0x32, 0x08, 0x16,
	0x84, 0x8d, 0xae, 0xd0, 0x5b, 0x5d, 0x13, 0x85, 0x88, 0x48, 0x89, 0x8a, 0x97, 0x32, 0x4d, 0xce,
	0xd6, 0xe0, 0xce, 0x0f, 0x93, 0x49, 0xa5, 0xcf, 0xe7, 0xc3, 0xf8, 0x18, 0xca, 0xcc, 0xa4, 0x69,
	0x96, 0xa6, 0xbd, 0xe9, 0x55, 0x72, 0xce, 0xf7, 0x33, 0xcc, 0x77, 0xce, 0x40, 0x58, 0xe4, 0xfc,
	0x4c, 0x47, 0xe7, 0xaf, 0x22, 0x5e, 0xe6, 0x5a, 0x46, 0xae, 0x5c, 0x28, 0x2d, 0x8d, 0x24, 0x23,
	0x57, 0x84, 0xd3, 0x5c, 0x72, 0x2e, 0x45, 0xe4, 0x3f, 0x1e, 0xa3, 0x1f, 0xe1, 0x20, 0x61, 0x86,
	0xc5, 0x28, 0x0c, 0xea, 0x4f, 0x65, 0x65, 0x32, 0xac, 0x94, 0x14, 0x15, 0x92, 0x97, 0x30, 0x2e,
	0x72, 0xdb, 0x99, 0x05, 0x4f, 0x77, 0x8f, 0xf6, 0x96, 0xb3, 0x45, 0x57, 0xb8, 0xb8, 0x52, 0x65,
	0x0d, 0x8f, 0xfe, 0x09, 0x60, 0xfa, 0x19, 0xcd, 0x6f, 0xa9, 0x7f, 0xa5, 0x62, 0x2d, 0x5b, 0xa7,
	0x39, 0x40, 0x5d, 0xa1, 0xfe, 0x91, 0xcb, 0x5a, 0x58, 0xb7, 0xe0, 0x68, 0x94, 0x4d, 0x6c, 0x27,
	0xb6, 0x0d, 0x0b, 0xff, 0x94, 0x95, 0x69, 0xe0, 0x1d, 0x0f, 0xdb, 0x8e, 0x87, 0x5f, 0xc0, 0x23,
	0x14, 0xe7, 0xa5, 0x96, 0x82, 0xa3, 0xb8, 0x64, 0xed, 0x3a, 0xd6, 0xc3, 0x0e, 0xe0, 0xc9, 0xcf,
	0xe1, 0x41, 0x2e, 0x85, 0x61, 0xa5, 0x68, 0xcf, 0x1b, 0x3a, 0xea, 0x7e, 0xdb, 0xf6, 0xc4, 0x19,
	0xdc, 0x33, 0x9a, 0xad, 0xd7, 0x65, 0x3e, 0x1b, 0x39, 0xc2, 0x65, 0x49, 0xbf, 0xc3, 0xbc, 0x93,
	0x08, 0xb2, 0x02, 0xf5, 0x3b, 0xc9, 0x74, 0xd1, 0x5e, 0xe7, 0x18, 0x86, 0x9b, 0xab, 0x58, 0xe8,
	0xc2, 0x47, 0xdd, 0xab, 0x49, 0xd0, 0xb0, 0x72, 0x93, 0x39, 0x3e, 0x4d, 0xe1, 0xf0, 0x16, 0x12,
	0x21, 0x30, 0x14, 0x8c, 0xa3, 0xcb, 0x67, 0x92, 0xb9, 0x7f, 0x72, 0x00, 0x63, 0x51, 0xf3, 0x53,
	0xd4, 0x2e, 0x96, 0x20, 0x6b, 0xaa, 0xe5, 0xbf, 0x00, 0x76, 0x92, 0x98, 0xbc, 0x81, 0x49, 0xac,
	0x91, 0x19, 0x5c, 0x29, 0x45, 0x0e, 0xb7, 0xe7, 0xb3, 0x52, 0x2a, 0x41, 0xb5, 0x91, 0x17, 0x36,
	0x9c, 0x70, 0xba, 0x0d, 0xbe, 0xe7, 0xca, 0x5c, 0xd0, 0x81, 0x35, 0xf8, 0xa6, 0x8a, 0x3b, 0x18,
	0x1c, 0xc3, 0xfd, 0x93, 0x5a, 0x9f, 0x39, 0xfd, 0xf4, 0x9a, 0x3e, 0x4d, 0x6e, 0x3e, 0x78, 0xfc,
	0xc5, 0x30, 0x53, 0x57, 0xa4, 0x8f, 0x10, 0x3e, 0xb9, 0x69, 0xd7, 0xbc, 0x88, 0x0e, 0x96, 0x7f,
	0x03, 0x18, 0x25, 0xf1, 0xea, 0x24, 0x25, 0x1f, 0x60, 0x7f, 0x7b, 0x83, 0xfb, 0x2d, 0xe7, 0xd7,
	0xe7, 0xd4, 0xd9, 0x76, 0x3a, 0x20, 0x5f, 0xe1, 0x71, 0xef, 0x78, 0xfa, 0xed, 0x9e, 0xdd, 0x36,
	0xf6, 0x8e, 0xeb, 0x5b, 0xd8, 0xeb, 0x3c, 0x89, 0x7e, 0xaf, 0xb0, 0xf1, 0xea, 0x79, 0x3b, 0x74,
	0x70, 0x3a, 0x76, 0xd4, 0xd7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x64, 0x60, 0x6f, 0xe2,
	0x03, 0x00, 0x00,
}

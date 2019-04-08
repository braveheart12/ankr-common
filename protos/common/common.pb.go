// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common_proto

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

// Task Events operation code
type DCOperation int32

const (
	DCOperation_TASK_CREATE DCOperation = 0
	DCOperation_TASK_CANCEL DCOperation = 1
	DCOperation_TASK_UPDATE DCOperation = 2
	DCOperation_HEARTBEAT   DCOperation = 3
)

var DCOperation_name = map[int32]string{
	0: "TASK_CREATE",
	1: "TASK_CANCEL",
	2: "TASK_UPDATE",
	3: "HEARTBEAT",
}

var DCOperation_value = map[string]int32{
	"TASK_CREATE": 0,
	"TASK_CANCEL": 1,
	"TASK_UPDATE": 2,
	"HEARTBEAT":   3,
}

func (x DCOperation) String() string {
	return proto.EnumName(DCOperation_name, int32(x))
}

func (DCOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

// Hub app status
type AppStatus int32

const (
	AppStatus_APP_STARTING       AppStatus = 0
	AppStatus_APP_START_SUCCESS  AppStatus = 1
	AppStatus_APP_START_FAILED   AppStatus = 2
	AppStatus_APP_RUNNING        AppStatus = 3
	AppStatus_APP_UPDATING       AppStatus = 4
	AppStatus_APP_UPDATE_SUCCESS AppStatus = 5
	AppStatus_APP_UPDATE_FAILED  AppStatus = 6
	AppStatus_APP_CANCELLING     AppStatus = 7
	AppStatus_APP_CANCELLED      AppStatus = 8
	AppStatus_APP_CANCEL_FAILED  AppStatus = 9
	AppStatus_APP_DONE           AppStatus = 10
)

var AppStatus_name = map[int32]string{
	0:  "APP_STARTING",
	1:  "APP_START_SUCCESS",
	2:  "APP_START_FAILED",
	3:  "APP_RUNNING",
	4:  "APP_UPDATING",
	5:  "APP_UPDATE_SUCCESS",
	6:  "APP_UPDATE_FAILED",
	7:  "APP_CANCELLING",
	8:  "APP_CANCELLED",
	9:  "APP_CANCEL_FAILED",
	10: "APP_DONE",
}

var AppStatus_value = map[string]int32{
	"APP_STARTING":       0,
	"APP_START_SUCCESS":  1,
	"APP_START_FAILED":   2,
	"APP_RUNNING":        3,
	"APP_UPDATING":       4,
	"APP_UPDATE_SUCCESS": 5,
	"APP_UPDATE_FAILED":  6,
	"APP_CANCELLING":     7,
	"APP_CANCELLED":      8,
	"APP_CANCEL_FAILED":  9,
	"APP_DONE":           10,
}

func (x AppStatus) String() string {
	return proto.EnumName(AppStatus_name, int32(x))
}

func (AppStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

// Hub namespace status
type NamespaceStatus int32

const (
	NamespaceStatus_NS_STARTING       NamespaceStatus = 0
	NamespaceStatus_NS_START_SUCCESS  NamespaceStatus = 1
	NamespaceStatus_NS_START_FAILED   NamespaceStatus = 2
	NamespaceStatus_NS_RUNNING        NamespaceStatus = 3
	NamespaceStatus_NS_UPDATING       NamespaceStatus = 4
	NamespaceStatus_NS_UPDATE_SUCCESS NamespaceStatus = 5
	NamespaceStatus_NS_UPDATE_FAILED  NamespaceStatus = 6
	NamespaceStatus_NS_CANCELLING     NamespaceStatus = 7
	NamespaceStatus_NS_CANCELLED      NamespaceStatus = 8
	NamespaceStatus_NS_CANCEL_FAILED  NamespaceStatus = 9
	NamespaceStatus_NS_DONE           NamespaceStatus = 10
)

var NamespaceStatus_name = map[int32]string{
	0:  "NS_STARTING",
	1:  "NS_START_SUCCESS",
	2:  "NS_START_FAILED",
	3:  "NS_RUNNING",
	4:  "NS_UPDATING",
	5:  "NS_UPDATE_SUCCESS",
	6:  "NS_UPDATE_FAILED",
	7:  "NS_CANCELLING",
	8:  "NS_CANCELLED",
	9:  "NS_CANCEL_FAILED",
	10: "NS_DONE",
}

var NamespaceStatus_value = map[string]int32{
	"NS_STARTING":       0,
	"NS_START_SUCCESS":  1,
	"NS_START_FAILED":   2,
	"NS_RUNNING":        3,
	"NS_UPDATING":       4,
	"NS_UPDATE_SUCCESS": 5,
	"NS_UPDATE_FAILED":  6,
	"NS_CANCELLING":     7,
	"NS_CANCELLED":      8,
	"NS_CANCEL_FAILED":  9,
	"NS_DONE":           10,
}

func (x NamespaceStatus) String() string {
	return proto.EnumName(NamespaceStatus_name, int32(x))
}

func (NamespaceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

// Data center status
type DCStatus int32

const (
	DCStatus_AVAILABLE   DCStatus = 0
	DCStatus_UNAVAILABLE DCStatus = 1
)

var DCStatus_name = map[int32]string{
	0: "AVAILABLE",
	1: "UNAVAILABLE",
}

var DCStatus_value = map[string]int32{
	"AVAILABLE":   0,
	"UNAVAILABLE": 1,
}

func (x DCStatus) String() string {
	return proto.EnumName(DCStatus_name, int32(x))
}

func (DCStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

// Emtpy Message
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// App Data structure
type App struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to TypeData:
	//	*App_NamespaceId
	//	*App_Namespace
	TypeData             isApp_TypeData `protobuf_oneof:"type_data"`
	Status               AppStatus      `protobuf:"varint,5,opt,name=status,proto3,enum=common.proto.AppStatus" json:"status,omitempty"`
	Attribute            *AppAttributes `protobuf:"bytes,6,opt,name=attribute,proto3" json:"attribute,omitempty"`
	Uid                  string         `protobuf:"bytes,7,opt,name=uid,proto3" json:"uid,omitempty"`
	Chart                *Chart         `protobuf:"bytes,8,opt,name=chart,proto3" json:"chart,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isApp_TypeData interface {
	isApp_TypeData()
}

type App_NamespaceId struct {
	NamespaceId string `protobuf:"bytes,3,opt,name=namespace_id,json=namespaceId,proto3,oneof"`
}

type App_Namespace struct {
	Namespace *Namespace `protobuf:"bytes,4,opt,name=namespace,proto3,oneof"`
}

func (*App_NamespaceId) isApp_TypeData() {}

func (*App_Namespace) isApp_TypeData() {}

func (m *App) GetTypeData() isApp_TypeData {
	if m != nil {
		return m.TypeData
	}
	return nil
}

func (m *App) GetNamespaceId() string {
	if x, ok := m.GetTypeData().(*App_NamespaceId); ok {
		return x.NamespaceId
	}
	return ""
}

func (m *App) GetNamespace() *Namespace {
	if x, ok := m.GetTypeData().(*App_Namespace); ok {
		return x.Namespace
	}
	return nil
}

func (m *App) GetStatus() AppStatus {
	if m != nil {
		return m.Status
	}
	return AppStatus_APP_STARTING
}

func (m *App) GetAttribute() *AppAttributes {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *App) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *App) GetChart() *Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*App) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*App_NamespaceId)(nil),
		(*App_Namespace)(nil),
	}
}

type AppAttributes struct {
	Hidden               bool     `protobuf:"varint,1,opt,name=hidden,proto3" json:"hidden,omitempty"`
	CreationDate         uint64   `protobuf:"varint,2,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	LastModifiedDate     uint64   `protobuf:"varint,3,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppAttributes) Reset()         { *m = AppAttributes{} }
func (m *AppAttributes) String() string { return proto.CompactTextString(m) }
func (*AppAttributes) ProtoMessage()    {}
func (*AppAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *AppAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppAttributes.Unmarshal(m, b)
}
func (m *AppAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppAttributes.Marshal(b, m, deterministic)
}
func (m *AppAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppAttributes.Merge(m, src)
}
func (m *AppAttributes) XXX_Size() int {
	return xxx_messageInfo_AppAttributes.Size(m)
}
func (m *AppAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_AppAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_AppAttributes proto.InternalMessageInfo

func (m *AppAttributes) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *AppAttributes) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *AppAttributes) GetLastModifiedDate() uint64 {
	if m != nil {
		return m.LastModifiedDate
	}
	return 0
}

type GeoLocation struct {
	Lat                  string   `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  string   `protobuf:"bytes,2,opt,name=lng,proto3" json:"lng,omitempty"`
	Country              string   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoLocation) Reset()         { *m = GeoLocation{} }
func (m *GeoLocation) String() string { return proto.CompactTextString(m) }
func (*GeoLocation) ProtoMessage()    {}
func (*GeoLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *GeoLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoLocation.Unmarshal(m, b)
}
func (m *GeoLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoLocation.Marshal(b, m, deterministic)
}
func (m *GeoLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoLocation.Merge(m, src)
}
func (m *GeoLocation) XXX_Size() int {
	return xxx_messageInfo_GeoLocation.Size(m)
}
func (m *GeoLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoLocation.DiscardUnknown(m)
}

var xxx_messageInfo_GeoLocation proto.InternalMessageInfo

func (m *GeoLocation) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *GeoLocation) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *GeoLocation) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

// Data Center structure
type DataCenter struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GeoLocation          *GeoLocation          `protobuf:"bytes,3,opt,name=geo_location,json=geoLocation,proto3" json:"geo_location,omitempty"`
	Status               DCStatus              `protobuf:"varint,4,opt,name=status,proto3,enum=common.proto.DCStatus" json:"status,omitempty"`
	DcAttributes         *DataCenterAttributes `protobuf:"bytes,5,opt,name=dc_attributes,json=dcAttributes,proto3" json:"dc_attributes,omitempty"`
	DcHeartbeatReport    *DCHeartbeatReport    `protobuf:"bytes,6,opt,name=dc_heartbeat_report,json=dcHeartbeatReport,proto3" json:"dc_heartbeat_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DataCenter) Reset()         { *m = DataCenter{} }
func (m *DataCenter) String() string { return proto.CompactTextString(m) }
func (*DataCenter) ProtoMessage()    {}
func (*DataCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *DataCenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenter.Unmarshal(m, b)
}
func (m *DataCenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenter.Marshal(b, m, deterministic)
}
func (m *DataCenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenter.Merge(m, src)
}
func (m *DataCenter) XXX_Size() int {
	return xxx_messageInfo_DataCenter.Size(m)
}
func (m *DataCenter) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenter.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenter proto.InternalMessageInfo

func (m *DataCenter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataCenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenter) GetGeoLocation() *GeoLocation {
	if m != nil {
		return m.GeoLocation
	}
	return nil
}

func (m *DataCenter) GetStatus() DCStatus {
	if m != nil {
		return m.Status
	}
	return DCStatus_AVAILABLE
}

func (m *DataCenter) GetDcAttributes() *DataCenterAttributes {
	if m != nil {
		return m.DcAttributes
	}
	return nil
}

func (m *DataCenter) GetDcHeartbeatReport() *DCHeartbeatReport {
	if m != nil {
		return m.DcHeartbeatReport
	}
	return nil
}

type DataCenterAttributes struct {
	WalletAddress        string   `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	CreationDate         uint64   `protobuf:"varint,2,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	LastModifiedDate     uint64   `protobuf:"varint,3,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenterAttributes) Reset()         { *m = DataCenterAttributes{} }
func (m *DataCenterAttributes) String() string { return proto.CompactTextString(m) }
func (*DataCenterAttributes) ProtoMessage()    {}
func (*DataCenterAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *DataCenterAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterAttributes.Unmarshal(m, b)
}
func (m *DataCenterAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterAttributes.Marshal(b, m, deterministic)
}
func (m *DataCenterAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterAttributes.Merge(m, src)
}
func (m *DataCenterAttributes) XXX_Size() int {
	return xxx_messageInfo_DataCenterAttributes.Size(m)
}
func (m *DataCenterAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterAttributes proto.InternalMessageInfo

func (m *DataCenterAttributes) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *DataCenterAttributes) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *DataCenterAttributes) GetLastModifiedDate() uint64 {
	if m != nil {
		return m.LastModifiedDate
	}
	return 0
}

type DCHeartbeatReport struct {
	Metrics              string   `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Report               string   `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
	ReportTime           uint64   `protobuf:"varint,3,opt,name=report_time,json=reportTime,proto3" json:"report_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DCHeartbeatReport) Reset()         { *m = DCHeartbeatReport{} }
func (m *DCHeartbeatReport) String() string { return proto.CompactTextString(m) }
func (*DCHeartbeatReport) ProtoMessage()    {}
func (*DCHeartbeatReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *DCHeartbeatReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCHeartbeatReport.Unmarshal(m, b)
}
func (m *DCHeartbeatReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCHeartbeatReport.Marshal(b, m, deterministic)
}
func (m *DCHeartbeatReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCHeartbeatReport.Merge(m, src)
}
func (m *DCHeartbeatReport) XXX_Size() int {
	return xxx_messageInfo_DCHeartbeatReport.Size(m)
}
func (m *DCHeartbeatReport) XXX_DiscardUnknown() {
	xxx_messageInfo_DCHeartbeatReport.DiscardUnknown(m)
}

var xxx_messageInfo_DCHeartbeatReport proto.InternalMessageInfo

func (m *DCHeartbeatReport) GetMetrics() string {
	if m != nil {
		return m.Metrics
	}
	return ""
}

func (m *DCHeartbeatReport) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *DCHeartbeatReport) GetReportTime() uint64 {
	if m != nil {
		return m.ReportTime
	}
	return 0
}

type AppReport struct {
	App                  *App     `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Report               string   `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppReport) Reset()         { *m = AppReport{} }
func (m *AppReport) String() string { return proto.CompactTextString(m) }
func (*AppReport) ProtoMessage()    {}
func (*AppReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

func (m *AppReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppReport.Unmarshal(m, b)
}
func (m *AppReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppReport.Marshal(b, m, deterministic)
}
func (m *AppReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppReport.Merge(m, src)
}
func (m *AppReport) XXX_Size() int {
	return xxx_messageInfo_AppReport.Size(m)
}
func (m *AppReport) XXX_DiscardUnknown() {
	xxx_messageInfo_AppReport.DiscardUnknown(m)
}

var xxx_messageInfo_AppReport proto.InternalMessageInfo

func (m *AppReport) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *AppReport) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

// data center communicate with dc manager
type DCStream struct {
	OpType DCOperation `protobuf:"varint,1,opt,name=op_type,json=opType,proto3,enum=common.proto.DCOperation" json:"op_type,omitempty"`
	// Types that are valid to be assigned to OpPayload:
	//	*DCStream_App
	//	*DCStream_AppReport
	//	*DCStream_DataCenter
	OpPayload            isDCStream_OpPayload `protobuf_oneof:"op_payload"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DCStream) Reset()         { *m = DCStream{} }
func (m *DCStream) String() string { return proto.CompactTextString(m) }
func (*DCStream) ProtoMessage()    {}
func (*DCStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{8}
}

func (m *DCStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCStream.Unmarshal(m, b)
}
func (m *DCStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCStream.Marshal(b, m, deterministic)
}
func (m *DCStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCStream.Merge(m, src)
}
func (m *DCStream) XXX_Size() int {
	return xxx_messageInfo_DCStream.Size(m)
}
func (m *DCStream) XXX_DiscardUnknown() {
	xxx_messageInfo_DCStream.DiscardUnknown(m)
}

var xxx_messageInfo_DCStream proto.InternalMessageInfo

func (m *DCStream) GetOpType() DCOperation {
	if m != nil {
		return m.OpType
	}
	return DCOperation_TASK_CREATE
}

type isDCStream_OpPayload interface {
	isDCStream_OpPayload()
}

type DCStream_App struct {
	App *App `protobuf:"bytes,2,opt,name=app,proto3,oneof"`
}

type DCStream_AppReport struct {
	AppReport *AppReport `protobuf:"bytes,3,opt,name=app_report,json=appReport,proto3,oneof"`
}

type DCStream_DataCenter struct {
	DataCenter *DataCenter `protobuf:"bytes,4,opt,name=data_center,json=dataCenter,proto3,oneof"`
}

func (*DCStream_App) isDCStream_OpPayload() {}

func (*DCStream_AppReport) isDCStream_OpPayload() {}

func (*DCStream_DataCenter) isDCStream_OpPayload() {}

func (m *DCStream) GetOpPayload() isDCStream_OpPayload {
	if m != nil {
		return m.OpPayload
	}
	return nil
}

func (m *DCStream) GetApp() *App {
	if x, ok := m.GetOpPayload().(*DCStream_App); ok {
		return x.App
	}
	return nil
}

func (m *DCStream) GetAppReport() *AppReport {
	if x, ok := m.GetOpPayload().(*DCStream_AppReport); ok {
		return x.AppReport
	}
	return nil
}

func (m *DCStream) GetDataCenter() *DataCenter {
	if x, ok := m.GetOpPayload().(*DCStream_DataCenter); ok {
		return x.DataCenter
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DCStream) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DCStream_App)(nil),
		(*DCStream_AppReport)(nil),
		(*DCStream_DataCenter)(nil),
	}
}

type Chart struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Repo                 string   `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IconUrl              string   `protobuf:"bytes,4,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	LatestVersion        string   `protobuf:"bytes,5,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	LatestAppVersion     string   `protobuf:"bytes,6,opt,name=latest_app_version,json=latestAppVersion,proto3" json:"latest_app_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chart) Reset()         { *m = Chart{} }
func (m *Chart) String() string { return proto.CompactTextString(m) }
func (*Chart) ProtoMessage()    {}
func (*Chart) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{9}
}

func (m *Chart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chart.Unmarshal(m, b)
}
func (m *Chart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chart.Marshal(b, m, deterministic)
}
func (m *Chart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chart.Merge(m, src)
}
func (m *Chart) XXX_Size() int {
	return xxx_messageInfo_Chart.Size(m)
}
func (m *Chart) XXX_DiscardUnknown() {
	xxx_messageInfo_Chart.DiscardUnknown(m)
}

var xxx_messageInfo_Chart proto.InternalMessageInfo

func (m *Chart) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chart) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *Chart) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Chart) GetIconUrl() string {
	if m != nil {
		return m.IconUrl
	}
	return ""
}

func (m *Chart) GetLatestVersion() string {
	if m != nil {
		return m.LatestVersion
	}
	return ""
}

func (m *Chart) GetLatestAppVersion() string {
	if m != nil {
		return m.LatestAppVersion
	}
	return ""
}

type ChartDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Repo                 string   `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	AppVersion           string   `protobuf:"bytes,4,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChartDetail) Reset()         { *m = ChartDetail{} }
func (m *ChartDetail) String() string { return proto.CompactTextString(m) }
func (*ChartDetail) ProtoMessage()    {}
func (*ChartDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{10}
}

func (m *ChartDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChartDetail.Unmarshal(m, b)
}
func (m *ChartDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChartDetail.Marshal(b, m, deterministic)
}
func (m *ChartDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChartDetail.Merge(m, src)
}
func (m *ChartDetail) XXX_Size() int {
	return xxx_messageInfo_ChartDetail.Size(m)
}
func (m *ChartDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ChartDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ChartDetail proto.InternalMessageInfo

func (m *ChartDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChartDetail) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *ChartDetail) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChartDetail) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

type Namespace struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ClusterId            string          `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClusterName          string          `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	CreationDate         uint64          `protobuf:"varint,5,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	CpuLimit             float32         `protobuf:"fixed32,6,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	MemLimit             string          `protobuf:"bytes,7,opt,name=mem_limit,json=memLimit,proto3" json:"mem_limit,omitempty"`
	StorageLimit         string          `protobuf:"bytes,8,opt,name=storage_limit,json=storageLimit,proto3" json:"storage_limit,omitempty"`
	NamespaceStatus      NamespaceStatus `protobuf:"varint,9,opt,name=namespace_status,json=namespaceStatus,proto3,enum=common.proto.NamespaceStatus" json:"namespace_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{11}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Namespace) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Namespace) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *Namespace) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *Namespace) GetCpuLimit() float32 {
	if m != nil {
		return m.CpuLimit
	}
	return 0
}

func (m *Namespace) GetMemLimit() string {
	if m != nil {
		return m.MemLimit
	}
	return ""
}

func (m *Namespace) GetStorageLimit() string {
	if m != nil {
		return m.StorageLimit
	}
	return ""
}

func (m *Namespace) GetNamespaceStatus() NamespaceStatus {
	if m != nil {
		return m.NamespaceStatus
	}
	return NamespaceStatus_NS_STARTING
}

func init() {
	proto.RegisterEnum("common.proto.DCOperation", DCOperation_name, DCOperation_value)
	proto.RegisterEnum("common.proto.AppStatus", AppStatus_name, AppStatus_value)
	proto.RegisterEnum("common.proto.NamespaceStatus", NamespaceStatus_name, NamespaceStatus_value)
	proto.RegisterEnum("common.proto.DCStatus", DCStatus_name, DCStatus_value)
	proto.RegisterType((*Empty)(nil), "common.proto.Empty")
	proto.RegisterType((*App)(nil), "common.proto.App")
	proto.RegisterType((*AppAttributes)(nil), "common.proto.AppAttributes")
	proto.RegisterType((*GeoLocation)(nil), "common.proto.GeoLocation")
	proto.RegisterType((*DataCenter)(nil), "common.proto.DataCenter")
	proto.RegisterType((*DataCenterAttributes)(nil), "common.proto.DataCenterAttributes")
	proto.RegisterType((*DCHeartbeatReport)(nil), "common.proto.DCHeartbeatReport")
	proto.RegisterType((*AppReport)(nil), "common.proto.AppReport")
	proto.RegisterType((*DCStream)(nil), "common.proto.DCStream")
	proto.RegisterType((*Chart)(nil), "common.proto.Chart")
	proto.RegisterType((*ChartDetail)(nil), "common.proto.ChartDetail")
	proto.RegisterType((*Namespace)(nil), "common.proto.Namespace")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdb, 0x6e, 0xe3, 0x44,
	0x18, 0x8e, 0x73, 0xf6, 0xef, 0x24, 0x75, 0x66, 0x77, 0x4b, 0x56, 0xab, 0xd5, 0x96, 0x54, 0x2b,
	0x95, 0x0a, 0x15, 0xa9, 0x5c, 0x00, 0x82, 0x1b, 0x37, 0x31, 0x4d, 0xb5, 0xc1, 0xad, 0x26, 0xe9,
	0xde, 0x5a, 0x53, 0x7b, 0xb6, 0xb5, 0x64, 0xc7, 0x23, 0x7b, 0x02, 0x2a, 0x4f, 0xc1, 0xe3, 0x70,
	0x83, 0x10, 0xcf, 0xc2, 0x0d, 0xe2, 0x29, 0xd0, 0x1c, 0xec, 0x38, 0x69, 0x2b, 0xf5, 0x86, 0x2b,
	0xfb, 0xff, 0xfe, 0xd3, 0x7c, 0xf3, 0x1f, 0x06, 0x7a, 0x41, 0x9a, 0x24, 0xe9, 0xea, 0x84, 0x65,
	0x29, 0x4f, 0xd1, 0x96, 0x34, 0xee, 0x40, 0xcb, 0x4d, 0x18, 0xbf, 0x1f, 0xff, 0x51, 0x87, 0x86,
	0xc3, 0x18, 0x1a, 0x40, 0x3d, 0x0a, 0x47, 0xc6, 0x81, 0x71, 0x64, 0xe2, 0x7a, 0x14, 0x22, 0x04,
	0xcd, 0x15, 0x49, 0xe8, 0xa8, 0x2e, 0x11, 0xf9, 0x8f, 0x0e, 0xa1, 0x27, 0xbe, 0x39, 0x23, 0x01,
	0xf5, 0xa3, 0x70, 0xd4, 0x10, 0xba, 0x59, 0x0d, 0x5b, 0x25, 0x7a, 0x11, 0xa2, 0x6f, 0xc0, 0x2c,
	0xc5, 0x51, 0xf3, 0xc0, 0x38, 0xb2, 0x4e, 0x3f, 0x3b, 0xa9, 0xe6, 0x3e, 0xf1, 0x0a, 0xf5, 0xac,
	0x86, 0x37, 0xb6, 0xe8, 0x2b, 0x68, 0xe7, 0x9c, 0xf0, 0x75, 0x3e, 0x6a, 0x1d, 0x18, 0x47, 0x83,
	0x5d, 0x2f, 0x87, 0xb1, 0x85, 0x54, 0x63, 0x6d, 0x86, 0xbe, 0x03, 0x93, 0x70, 0x9e, 0x45, 0x37,
	0x6b, 0x4e, 0x47, 0x6d, 0x99, 0xe9, 0xcd, 0x03, 0x1f, 0xa7, 0xb0, 0xc8, 0xf1, 0xc6, 0x1a, 0xd9,
	0xd0, 0x58, 0x47, 0xe1, 0xa8, 0x23, 0xc9, 0x89, 0x5f, 0xf4, 0x05, 0xb4, 0x82, 0x3b, 0x92, 0xf1,
	0x51, 0x57, 0x06, 0x7a, 0xb1, 0x1d, 0x68, 0x22, 0x54, 0x58, 0x59, 0x9c, 0x59, 0x60, 0xf2, 0x7b,
	0x46, 0xfd, 0x90, 0x70, 0x32, 0xfe, 0x15, 0xfa, 0x5b, 0x59, 0xd0, 0x3e, 0xb4, 0xef, 0xa2, 0x30,
	0xa4, 0x2b, 0x79, 0x99, 0x5d, 0xac, 0x25, 0x74, 0x08, 0xfd, 0x20, 0xa3, 0x84, 0x47, 0xe9, 0x4a,
	0x78, 0xaa, 0x9b, 0x6d, 0xe2, 0x5e, 0x01, 0x4e, 0x09, 0xa7, 0xe8, 0x4b, 0x40, 0x31, 0xc9, 0xb9,
	0x9f, 0xa4, 0x61, 0xf4, 0x29, 0xa2, 0xa1, 0xb2, 0x6c, 0x48, 0x4b, 0x5b, 0x68, 0x7e, 0xd2, 0x0a,
	0x61, 0x3d, 0xfe, 0x00, 0xd6, 0x39, 0x4d, 0xe7, 0x69, 0x20, 0x03, 0x08, 0x52, 0x31, 0xe1, 0xba,
	0x86, 0xe2, 0x57, 0x22, 0xab, 0x5b, 0x5d, 0x43, 0xf1, 0x8b, 0x46, 0xd0, 0x09, 0xd2, 0xf5, 0x8a,
	0x67, 0xf7, 0xaa, 0x7a, 0xb8, 0x10, 0xc7, 0xbf, 0xd7, 0x01, 0xa6, 0x84, 0x93, 0x09, 0x5d, 0x71,
	0x9a, 0x3d, 0xab, 0x1f, 0x7e, 0x80, 0xde, 0x2d, 0x4d, 0xfd, 0x58, 0x1f, 0x40, 0x46, 0xb4, 0x4e,
	0x5f, 0x6f, 0x5f, 0x5d, 0xe5, 0x84, 0xd8, 0xba, 0xad, 0x1c, 0xf7, 0xa4, 0xac, 0x77, 0x53, 0xd6,
	0x7b, 0x7f, 0xdb, 0x6f, 0x3a, 0xd9, 0x29, 0xf7, 0x39, 0xf4, 0xc3, 0xc0, 0x2f, 0x6b, 0xa8, 0xda,
	0xc4, 0x3a, 0x1d, 0xef, 0xb8, 0x95, 0x14, 0x2a, 0x95, 0xef, 0x85, 0x41, 0xa5, 0x42, 0x97, 0xf0,
	0x22, 0x0c, 0xfc, 0x3b, 0x4a, 0x32, 0x7e, 0x43, 0x09, 0xf7, 0x33, 0xca, 0xd2, 0x8c, 0xeb, 0x0e,
	0x7a, 0xb7, 0x7b, 0x8a, 0x59, 0x61, 0x87, 0xa5, 0x19, 0x1e, 0x86, 0xc1, 0x0e, 0x34, 0xfe, 0xcd,
	0x80, 0x97, 0x8f, 0xe5, 0x45, 0xef, 0x61, 0xf0, 0x0b, 0x89, 0x63, 0xca, 0x7d, 0x12, 0x86, 0x19,
	0xcd, 0x73, 0x7d, 0xa1, 0x7d, 0x85, 0x3a, 0x0a, 0xfc, 0x3f, 0x5a, 0xe3, 0x13, 0x0c, 0x1f, 0x1c,
	0x5d, 0x14, 0x3f, 0xa1, 0x3c, 0x8b, 0x82, 0xe2, 0x1c, 0x85, 0x28, 0x9a, 0x56, 0xdf, 0x82, 0xaa,
	0xaf, 0x96, 0xd0, 0x3b, 0xb0, 0xd4, 0x9f, 0xcf, 0xa3, 0xa4, 0xc8, 0x06, 0x0a, 0x5a, 0x46, 0x09,
	0x1d, 0xcf, 0xc0, 0x74, 0x18, 0xd3, 0xf1, 0x0f, 0xa1, 0x41, 0x18, 0x93, 0xb1, 0xad, 0xd3, 0xe1,
	0x83, 0x51, 0xc4, 0x42, 0xfb, 0x54, 0xaa, 0xf1, 0xbf, 0x06, 0x74, 0x45, 0xcd, 0x33, 0x4a, 0x12,
	0x74, 0x0a, 0x9d, 0x94, 0xf9, 0x62, 0xca, 0x64, 0xb4, 0xc1, 0x6e, 0x53, 0x4d, 0x27, 0x97, 0x8c,
	0x66, 0xaa, 0xa9, 0xda, 0x29, 0x5b, 0xde, 0x33, 0x8a, 0xde, 0xab, 0xec, 0xf5, 0x27, 0xb2, 0xcf,
	0x6a, 0x2a, 0xff, 0xb7, 0x00, 0x84, 0xb1, 0xa2, 0xe8, 0x8d, 0xc7, 0x16, 0x54, 0xc9, 0x48, 0x2c,
	0x28, 0x52, 0xd2, 0xfb, 0x1e, 0x2c, 0x31, 0xf2, 0x7e, 0x20, 0xcb, 0xac, 0x77, 0xdb, 0xe8, 0xa9,
	0xf6, 0x9b, 0xd5, 0x30, 0x84, 0xa5, 0x74, 0xd6, 0x03, 0x48, 0x99, 0xcf, 0xc8, 0x7d, 0x9c, 0x92,
	0x70, 0xfc, 0x97, 0x01, 0x2d, 0xb9, 0x53, 0xca, 0xb9, 0x32, 0x2a, 0x73, 0x85, 0xa0, 0x29, 0x8e,
	0x57, 0xcc, 0x9a, 0xf8, 0x47, 0x07, 0x60, 0x85, 0x34, 0x0f, 0xb2, 0x88, 0x95, 0xa3, 0x66, 0xe2,
	0x2a, 0x84, 0x5e, 0x43, 0x37, 0x0a, 0xd2, 0x95, 0xbf, 0xce, 0x62, 0x79, 0x36, 0x13, 0x77, 0x84,
	0x7c, 0x9d, 0xc5, 0xa2, 0x0f, 0x63, 0xc2, 0x69, 0xce, 0xfd, 0x9f, 0x69, 0x96, 0x0b, 0xff, 0x96,
	0xea, 0x43, 0x85, 0x7e, 0x54, 0xa0, 0x6a, 0x31, 0x69, 0x26, 0x6e, 0xa8, 0x30, 0x6d, 0x4b, 0x53,
	0x5b, 0x69, 0x1c, 0xc6, 0xb4, 0xf5, 0x98, 0x81, 0x25, 0x29, 0x4c, 0x29, 0x27, 0x51, 0xfc, 0x6c,
	0x22, 0x23, 0xe8, 0x14, 0x91, 0xf5, 0x06, 0xd2, 0xa2, 0x68, 0xb6, 0x6a, 0x5e, 0xc5, 0x41, 0x14,
	0xab, 0xc8, 0xf8, 0x67, 0x1d, 0xcc, 0xf2, 0xf1, 0x78, 0xd6, 0x86, 0x7a, 0x0b, 0x10, 0xc4, 0xeb,
	0x9c, 0xd3, 0xac, 0x7c, 0xaf, 0xb0, 0xa9, 0x91, 0x8b, 0x10, 0x7d, 0x0e, 0xbd, 0x42, 0x2d, 0x5d,
	0x55, 0x4a, 0x4b, 0x63, 0x9e, 0x7a, 0xf3, 0x76, 0x66, 0xb3, 0xf5, 0xc8, 0x6c, 0xbe, 0x01, 0x33,
	0x60, 0x6b, 0x3f, 0x8e, 0x92, 0x48, 0xed, 0x91, 0x3a, 0xee, 0x06, 0x6c, 0x3d, 0x17, 0xb2, 0x50,
	0x26, 0x34, 0xd1, 0x4a, 0xf5, 0xe2, 0x74, 0x13, 0x9a, 0x28, 0xe5, 0x21, 0xf4, 0x73, 0x9e, 0x66,
	0xe4, 0x96, 0x6a, 0x83, 0xae, 0x34, 0xe8, 0x69, 0x50, 0x19, 0xcd, 0xc0, 0xde, 0xbc, 0xbb, 0x7a,
	0x67, 0x9a, 0x72, 0x2c, 0xde, 0x3e, 0xf1, 0xb2, 0xea, 0xd5, 0xb9, 0xb7, 0xda, 0x06, 0x8e, 0x2f,
	0xc1, 0xaa, 0x8c, 0x0e, 0xda, 0x03, 0x6b, 0xe9, 0x2c, 0x3e, 0xf8, 0x13, 0xec, 0x3a, 0x4b, 0xd7,
	0xae, 0x6d, 0x00, 0xc7, 0x9b, 0xb8, 0x73, 0xdb, 0x28, 0x81, 0xeb, 0xab, 0xa9, 0xb0, 0xa8, 0xa3,
	0x3e, 0x98, 0x33, 0xd7, 0xc1, 0xcb, 0x33, 0xd7, 0x59, 0xda, 0x8d, 0xe3, 0x7f, 0x0c, 0xb9, 0x00,
	0x54, 0x78, 0x64, 0x43, 0xcf, 0xb9, 0xba, 0xf2, 0x17, 0x4b, 0x07, 0x2f, 0x2f, 0xbc, 0x73, 0xbb,
	0x86, 0x5e, 0xc1, 0xb0, 0x44, 0xfc, 0xc5, 0xf5, 0x64, 0xe2, 0x2e, 0x16, 0xb6, 0x81, 0x5e, 0x82,
	0xbd, 0x81, 0x7f, 0x74, 0x2e, 0xe6, 0xee, 0xd4, 0xae, 0x8b, 0x64, 0x02, 0xc5, 0xd7, 0x9e, 0x27,
	0xbc, 0x1b, 0x45, 0x3c, 0x99, 0x5c, 0x20, 0x4d, 0xb4, 0x0f, 0xa8, 0x44, 0xdc, 0x32, 0x60, 0xab,
	0xc8, 0xa3, 0x71, 0x1d, 0xb1, 0x8d, 0x10, 0x0c, 0x04, 0xac, 0xe8, 0xcc, 0x45, 0x88, 0x0e, 0x1a,
	0x42, 0xbf, 0x82, 0xb9, 0x53, 0xbb, 0x5b, 0x78, 0x2b, 0xa8, 0xf0, 0x36, 0x51, 0x0f, 0xba, 0x02,
	0x9e, 0x5e, 0x7a, 0xae, 0x0d, 0xc7, 0x7f, 0x1b, 0xb0, 0xb7, 0x73, 0xc1, 0xe2, 0xc4, 0xde, 0xa2,
	0xca, 0xf7, 0x25, 0xd8, 0x05, 0x50, 0xa1, 0xfb, 0x02, 0xf6, 0x4a, 0xb4, 0x64, 0x3b, 0x00, 0xf0,
	0x16, 0x15, 0xb2, 0x2a, 0x56, 0x85, 0xeb, 0x2b, 0x18, 0x16, 0x40, 0x95, 0xaa, 0x4a, 0xb1, 0xcb,
	0x74, 0x08, 0x7d, 0x6f, 0xb1, 0x4d, 0xd4, 0x86, 0xde, 0x06, 0x92, 0x3c, 0x95, 0xeb, 0x2e, 0x4d,
	0x0b, 0x3a, 0xde, 0xa2, 0x60, 0x79, 0xac, 0xb6, 0xb0, 0x64, 0xd7, 0x07, 0xd3, 0xf9, 0xe8, 0x5c,
	0xcc, 0x9d, 0xb3, 0xb9, 0x6e, 0x8e, 0x6b, 0x6f, 0x03, 0x18, 0x37, 0x6d, 0xd9, 0x75, 0x5f, 0xff,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x5a, 0xe3, 0xb3, 0x69, 0x0a, 0x00, 0x00,
}

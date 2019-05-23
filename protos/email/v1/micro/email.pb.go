// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email/v1/micro/email.proto

package mail

import (
	fmt "fmt"
	_ "github.com/Ankr-network/dccn-common/protos/common"
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

type EmailType int32

const (
	EmailType_CONFIRM_REGISTRATION EmailType = 0
	EmailType_FORGET_PASSWORD      EmailType = 1
	EmailType_CHANGE_PASSWORD      EmailType = 2
	EmailType_CONFIRM_EMAIL        EmailType = 3
)

var EmailType_name = map[int32]string{
	0: "CONFIRM_REGISTRATION",
	1: "FORGET_PASSWORD",
	2: "CHANGE_PASSWORD",
	3: "CONFIRM_EMAIL",
}

var EmailType_value = map[string]int32{
	"CONFIRM_REGISTRATION": 0,
	"FORGET_PASSWORD":      1,
	"CHANGE_PASSWORD":      2,
	"CONFIRM_EMAIL":        3,
}

func (x EmailType) String() string {
	return proto.EnumName(EmailType_name, int32(x))
}

func (EmailType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a857ab765f9f42f3, []int{0}
}

type ConfirmRegistration struct {
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Code                 string   `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmRegistration) Reset()         { *m = ConfirmRegistration{} }
func (m *ConfirmRegistration) String() string { return proto.CompactTextString(m) }
func (*ConfirmRegistration) ProtoMessage()    {}
func (*ConfirmRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_a857ab765f9f42f3, []int{0}
}

func (m *ConfirmRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmRegistration.Unmarshal(m, b)
}
func (m *ConfirmRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmRegistration.Marshal(b, m, deterministic)
}
func (m *ConfirmRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmRegistration.Merge(m, src)
}
func (m *ConfirmRegistration) XXX_Size() int {
	return xxx_messageInfo_ConfirmRegistration.Size(m)
}
func (m *ConfirmRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmRegistration proto.InternalMessageInfo

func (m *ConfirmRegistration) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ConfirmRegistration) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ConfirmRegistration) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ForgetPassword struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgetPassword) Reset()         { *m = ForgetPassword{} }
func (m *ForgetPassword) String() string { return proto.CompactTextString(m) }
func (*ForgetPassword) ProtoMessage()    {}
func (*ForgetPassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_a857ab765f9f42f3, []int{1}
}

func (m *ForgetPassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgetPassword.Unmarshal(m, b)
}
func (m *ForgetPassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgetPassword.Marshal(b, m, deterministic)
}
func (m *ForgetPassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgetPassword.Merge(m, src)
}
func (m *ForgetPassword) XXX_Size() int {
	return xxx_messageInfo_ForgetPassword.Size(m)
}
func (m *ForgetPassword) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgetPassword.DiscardUnknown(m)
}

var xxx_messageInfo_ForgetPassword proto.InternalMessageInfo

func (m *ForgetPassword) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ForgetPassword) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ChangePassword struct {
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Code                 string   `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePassword) Reset()         { *m = ChangePassword{} }
func (m *ChangePassword) String() string { return proto.CompactTextString(m) }
func (*ChangePassword) ProtoMessage()    {}
func (*ChangePassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_a857ab765f9f42f3, []int{2}
}

func (m *ChangePassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePassword.Unmarshal(m, b)
}
func (m *ChangePassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePassword.Marshal(b, m, deterministic)
}
func (m *ChangePassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePassword.Merge(m, src)
}
func (m *ChangePassword) XXX_Size() int {
	return xxx_messageInfo_ChangePassword.Size(m)
}
func (m *ChangePassword) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePassword.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePassword proto.InternalMessageInfo

func (m *ChangePassword) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ChangePassword) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ChangePassword) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ChangeEmail struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NewEmail             string   `protobuf:"bytes,2,opt,name=new_email,json=newEmail,proto3" json:"new_email,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeEmail) Reset()         { *m = ChangeEmail{} }
func (m *ChangeEmail) String() string { return proto.CompactTextString(m) }
func (*ChangeEmail) ProtoMessage()    {}
func (*ChangeEmail) Descriptor() ([]byte, []int) {
	return fileDescriptor_a857ab765f9f42f3, []int{3}
}

func (m *ChangeEmail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeEmail.Unmarshal(m, b)
}
func (m *ChangeEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeEmail.Marshal(b, m, deterministic)
}
func (m *ChangeEmail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeEmail.Merge(m, src)
}
func (m *ChangeEmail) XXX_Size() int {
	return xxx_messageInfo_ChangeEmail.Size(m)
}
func (m *ChangeEmail) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeEmail.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeEmail proto.InternalMessageInfo

func (m *ChangeEmail) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ChangeEmail) GetNewEmail() string {
	if m != nil {
		return m.NewEmail
	}
	return ""
}

func (m *ChangeEmail) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// MailEvent used to send email
type MailEvent struct {
	Type EmailType `protobuf:"varint,1,opt,name=type,proto3,enum=mail.EmailType" json:"type,omitempty"`
	From string    `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   []string  `protobuf:"bytes,3,rep,name=to,proto3" json:"to,omitempty"`
	// Types that are valid to be assigned to OpMail:
	//	*MailEvent_ConfirmRegistration
	//	*MailEvent_ForgetPassword
	//	*MailEvent_ChangePassword
	//	*MailEvent_ChangeEmail
	OpMail               isMailEvent_OpMail `protobuf_oneof:"op_mail"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MailEvent) Reset()         { *m = MailEvent{} }
func (m *MailEvent) String() string { return proto.CompactTextString(m) }
func (*MailEvent) ProtoMessage()    {}
func (*MailEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a857ab765f9f42f3, []int{4}
}

func (m *MailEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MailEvent.Unmarshal(m, b)
}
func (m *MailEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MailEvent.Marshal(b, m, deterministic)
}
func (m *MailEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MailEvent.Merge(m, src)
}
func (m *MailEvent) XXX_Size() int {
	return xxx_messageInfo_MailEvent.Size(m)
}
func (m *MailEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MailEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MailEvent proto.InternalMessageInfo

func (m *MailEvent) GetType() EmailType {
	if m != nil {
		return m.Type
	}
	return EmailType_CONFIRM_REGISTRATION
}

func (m *MailEvent) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MailEvent) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

type isMailEvent_OpMail interface {
	isMailEvent_OpMail()
}

type MailEvent_ConfirmRegistration struct {
	ConfirmRegistration *ConfirmRegistration `protobuf:"bytes,4,opt,name=confirm_registration,json=confirmRegistration,proto3,oneof"`
}

type MailEvent_ForgetPassword struct {
	ForgetPassword *ForgetPassword `protobuf:"bytes,5,opt,name=forget_password,json=forgetPassword,proto3,oneof"`
}

type MailEvent_ChangePassword struct {
	ChangePassword *ChangePassword `protobuf:"bytes,6,opt,name=change_password,json=changePassword,proto3,oneof"`
}

type MailEvent_ChangeEmail struct {
	ChangeEmail *ChangeEmail `protobuf:"bytes,7,opt,name=change_email,json=changeEmail,proto3,oneof"`
}

func (*MailEvent_ConfirmRegistration) isMailEvent_OpMail() {}

func (*MailEvent_ForgetPassword) isMailEvent_OpMail() {}

func (*MailEvent_ChangePassword) isMailEvent_OpMail() {}

func (*MailEvent_ChangeEmail) isMailEvent_OpMail() {}

func (m *MailEvent) GetOpMail() isMailEvent_OpMail {
	if m != nil {
		return m.OpMail
	}
	return nil
}

func (m *MailEvent) GetConfirmRegistration() *ConfirmRegistration {
	if x, ok := m.GetOpMail().(*MailEvent_ConfirmRegistration); ok {
		return x.ConfirmRegistration
	}
	return nil
}

func (m *MailEvent) GetForgetPassword() *ForgetPassword {
	if x, ok := m.GetOpMail().(*MailEvent_ForgetPassword); ok {
		return x.ForgetPassword
	}
	return nil
}

func (m *MailEvent) GetChangePassword() *ChangePassword {
	if x, ok := m.GetOpMail().(*MailEvent_ChangePassword); ok {
		return x.ChangePassword
	}
	return nil
}

func (m *MailEvent) GetChangeEmail() *ChangeEmail {
	if x, ok := m.GetOpMail().(*MailEvent_ChangeEmail); ok {
		return x.ChangeEmail
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MailEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MailEvent_ConfirmRegistration)(nil),
		(*MailEvent_ForgetPassword)(nil),
		(*MailEvent_ChangePassword)(nil),
		(*MailEvent_ChangeEmail)(nil),
	}
}

func init() {
	proto.RegisterEnum("mail.EmailType", EmailType_name, EmailType_value)
	proto.RegisterType((*ConfirmRegistration)(nil), "mail.ConfirmRegistration")
	proto.RegisterType((*ForgetPassword)(nil), "mail.ForgetPassword")
	proto.RegisterType((*ChangePassword)(nil), "mail.ChangePassword")
	proto.RegisterType((*ChangeEmail)(nil), "mail.ChangeEmail")
	proto.RegisterType((*MailEvent)(nil), "mail.MailEvent")
}

func init() { proto.RegisterFile("email/v1/micro/email.proto", fileDescriptor_a857ab765f9f42f3) }

var fileDescriptor_a857ab765f9f42f3 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x8d, 0xed, 0x40, 0x3c, 0xb4, 0x86, 0xac, 0x91, 0xea, 0x92, 0x4b, 0x44, 0x2f, 0x51,
	0x0f, 0x46, 0xa5, 0x52, 0x0e, 0xbd, 0x44, 0x94, 0x1a, 0xb0, 0x54, 0x4c, 0xb4, 0x20, 0x45, 0xea,
	0xc5, 0x72, 0xec, 0x85, 0x5a, 0x8a, 0xbd, 0x96, 0x71, 0x83, 0xf8, 0xc8, 0xfd, 0x16, 0xd5, 0xce,
	0x12, 0xb0, 0xa5, 0x1c, 0x7b, 0x62, 0xfe, 0xf0, 0x1e, 0xb3, 0xef, 0x27, 0xa0, 0xcf, 0xd2, 0x30,
	0x79, 0x1e, 0xbe, 0x7c, 0x19, 0xa6, 0x49, 0x54, 0xf0, 0x21, 0xb6, 0x4e, 0x5e, 0xf0, 0x92, 0x13,
	0x5d, 0xd4, 0x7d, 0x2b, 0xe2, 0x69, 0xca, 0xb3, 0xa1, 0xfc, 0x90, 0xab, 0x41, 0x00, 0xd6, 0x84,
	0x67, 0x9b, 0xa4, 0x48, 0x29, 0xdb, 0x26, 0xbb, 0xb2, 0x08, 0xcb, 0x84, 0x67, 0xe4, 0x1a, 0x8c,
	0x3f, 0x3b, 0x56, 0x04, 0x59, 0x98, 0x32, 0x5b, 0xbb, 0x69, 0xdc, 0x1a, 0xf4, 0x52, 0x0c, 0xfc,
	0x30, 0x65, 0xe4, 0x03, 0xb4, 0x70, 0x99, 0xc4, 0xb6, 0x8e, 0xab, 0xa6, 0x68, 0xbd, 0x98, 0x10,
	0xd0, 0x23, 0x1e, 0x33, 0xfb, 0x02, 0xa7, 0x58, 0x0f, 0xbe, 0x81, 0x39, 0xe5, 0xc5, 0x96, 0x95,
	0x0f, 0xe1, 0x6e, 0xb7, 0xe7, 0x45, 0x4c, 0x7a, 0x70, 0x81, 0xc7, 0xd9, 0x0d, 0xfc, 0x9a, 0x6c,
	0x4e, 0x5a, 0xb5, 0xa2, 0xfd, 0x05, 0xe6, 0xe4, 0x77, 0x98, 0x6d, 0xd9, 0x49, 0xfb, 0xff, 0xee,
	0x7a, 0x84, 0xb6, 0xf4, 0x76, 0xf1, 0xe7, 0x2b, 0xda, 0x46, 0x4d, 0x7b, 0x0d, 0x46, 0xc6, 0xf6,
	0x81, 0xbc, 0x58, 0x1e, 0x77, 0x99, 0xb1, 0xbd, 0x5b, 0x3b, 0x5a, 0xab, 0x18, 0xff, 0x55, 0xc1,
	0x58, 0x84, 0xc9, 0xb3, 0xfb, 0xc2, 0xb2, 0x92, 0x7c, 0x02, 0xbd, 0x3c, 0xe4, 0x0c, 0x4d, 0xcd,
	0x51, 0xc7, 0x41, 0x2a, 0x28, 0x5e, 0x1f, 0x72, 0x46, 0x71, 0x29, 0x6c, 0x36, 0x05, 0x4f, 0x5f,
	0xdf, 0x2e, 0x6a, 0x62, 0x82, 0x5a, 0x72, 0x5b, 0xbb, 0xd1, 0x6e, 0x0d, 0xaa, 0x96, 0x9c, 0xf8,
	0xd0, 0x8b, 0x24, 0xa8, 0xa0, 0xa8, 0x90, 0xc2, 0x97, 0xb6, 0x47, 0x1f, 0xa5, 0xf1, 0x1b, 0x28,
	0xe7, 0x0a, 0xb5, 0xa2, 0x37, 0x08, 0xdf, 0x43, 0x67, 0x83, 0x5c, 0x82, 0xfc, 0x18, 0x2e, 0xc6,
	0xd3, 0x1e, 0xf5, 0xa4, 0x55, 0x1d, 0xda, 0x5c, 0xa1, 0xe6, 0xa6, 0x8e, 0xf1, 0x1e, 0x3a, 0x11,
	0x06, 0x78, 0x36, 0x68, 0x56, 0x0d, 0xea, 0xe4, 0x84, 0x41, 0x54, 0x67, 0x79, 0x07, 0xef, 0x8e,
	0x06, 0x32, 0xdc, 0x16, 0xaa, 0xaf, 0xaa, 0x6a, 0x0c, 0x6a, 0xae, 0xd0, 0x76, 0x74, 0x6e, 0xbf,
	0x1b, 0xd0, 0xe2, 0x79, 0x20, 0xca, 0xcf, 0x4f, 0x60, 0x9c, 0xb2, 0x24, 0x36, 0xf4, 0x26, 0x4b,
	0x7f, 0xea, 0xd1, 0x45, 0x40, 0xdd, 0x99, 0xb7, 0x5a, 0xd3, 0xf1, 0xda, 0x5b, 0xfa, 0x5d, 0x85,
	0x58, 0xd0, 0x99, 0x2e, 0xe9, 0xcc, 0x5d, 0x07, 0x0f, 0xe3, 0xd5, 0xea, 0x71, 0x49, 0x7f, 0x74,
	0x1b, 0x62, 0x38, 0x99, 0x8f, 0xfd, 0x99, 0x7b, 0x1e, 0xaa, 0xe4, 0x0a, 0xde, 0xbf, 0x7a, 0xb8,
	0x8b, 0xb1, 0xf7, 0xb3, 0xab, 0x8d, 0xee, 0x40, 0x17, 0x38, 0x89, 0x03, 0xfa, 0x8a, 0x65, 0x31,
	0x39, 0x32, 0x3c, 0x21, 0xee, 0x5b, 0x4e, 0xf5, 0x1f, 0xe5, 0xb8, 0x69, 0x5e, 0x1e, 0x06, 0xca,
	0x53, 0x13, 0xdb, 0xaf, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0xa7, 0x1b, 0x7b, 0x99, 0x03,
	0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eid.proto

package twoferrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Inter_Mode int32

const (
	Inter_AUTH Inter_Mode = 0
	Inter_SIGN Inter_Mode = 1
)

var Inter_Mode_name = map[int32]string{
	0: "AUTH",
	1: "SIGN",
}
var Inter_Mode_value = map[string]int32{
	"AUTH": 0,
	"SIGN": 1,
}

func (x Inter_Mode) String() string {
	return proto.EnumName(Inter_Mode_name, int32(x))
}
func (Inter_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{5, 0}
}

type Resp_Status int32

const (
	Resp_STATUS_UNKNOWN     Resp_Status = 0
	Resp_STATUS_PENDING     Resp_Status = 1
	Resp_STATUS_CANCELED    Resp_Status = 2
	Resp_STATUS_RP_CANCELED Resp_Status = 3
	Resp_STATUS_EXPIRED     Resp_Status = 4
	Resp_STATUS_APPROVED    Resp_Status = 5
	Resp_STATUS_REJECTED    Resp_Status = 6
	Resp_STATUS_FAILED      Resp_Status = 7
)

var Resp_Status_name = map[int32]string{
	0: "STATUS_UNKNOWN",
	1: "STATUS_PENDING",
	2: "STATUS_CANCELED",
	3: "STATUS_RP_CANCELED",
	4: "STATUS_EXPIRED",
	5: "STATUS_APPROVED",
	6: "STATUS_REJECTED",
	7: "STATUS_FAILED",
}
var Resp_Status_value = map[string]int32{
	"STATUS_UNKNOWN":     0,
	"STATUS_PENDING":     1,
	"STATUS_CANCELED":    2,
	"STATUS_RP_CANCELED": 3,
	"STATUS_EXPIRED":     4,
	"STATUS_APPROVED":    5,
	"STATUS_REJECTED":    6,
	"STATUS_FAILED":      7,
}

func (x Resp_Status) String() string {
	return proto.EnumName(Resp_Status_name, int32(x))
}
func (Resp_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{7, 0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Providers struct {
	Providers            []*Provider `protobuf:"bytes,1,rep,name=providers" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Providers) Reset()         { *m = Providers{} }
func (m *Providers) String() string { return proto.CompactTextString(m) }
func (*Providers) ProtoMessage()    {}
func (*Providers) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{1}
}
func (m *Providers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Providers.Unmarshal(m, b)
}
func (m *Providers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Providers.Marshal(b, m, deterministic)
}
func (dst *Providers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Providers.Merge(dst, src)
}
func (m *Providers) XXX_Size() int {
	return xxx_messageInfo_Providers.Size(m)
}
func (m *Providers) XXX_DiscardUnknown() {
	xxx_messageInfo_Providers.DiscardUnknown(m)
}

var xxx_messageInfo_Providers proto.InternalMessageInfo

func (m *Providers) GetProviders() []*Provider {
	if m != nil {
		return m.Providers
	}
	return nil
}

type Provider struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{2}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Provider.Unmarshal(m, b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
}
func (dst *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(dst, src)
}
func (m *Provider) XXX_Size() int {
	return xxx_messageInfo_Provider.Size(m)
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type User struct {
	Inferred             bool     `protobuf:"varint,1,opt,name=inferred" json:"inferred,omitempty"`
	Ssn                  string   `protobuf:"bytes,2,opt,name=ssn" json:"ssn,omitempty"`
	SsnCountry           string   `protobuf:"bytes,3,opt,name=ssn_country,json=ssnCountry" json:"ssn_country,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	Ip                   string   `protobuf:"bytes,6,opt,name=ip" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{3}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetInferred() bool {
	if m != nil {
		return m.Inferred
	}
	return false
}

func (m *User) GetSsn() string {
	if m != nil {
		return m.Ssn
	}
	return ""
}

func (m *User) GetSsnCountry() string {
	if m != nil {
		return m.SsnCountry
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type Req struct {
	Provider             *Provider    `protobuf:"bytes,1,opt,name=provider" json:"provider,omitempty"`
	Who                  *User        `protobuf:"bytes,2,opt,name=who" json:"who,omitempty"`
	Payload              *Req_Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{4}
}
func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (dst *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(dst, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetProvider() *Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *Req) GetWho() *User {
	if m != nil {
		return m.Who
	}
	return nil
}

func (m *Req) GetPayload() *Req_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Req_Payload struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req_Payload) Reset()         { *m = Req_Payload{} }
func (m *Req_Payload) String() string { return proto.CompactTextString(m) }
func (*Req_Payload) ProtoMessage()    {}
func (*Req_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{4, 0}
}
func (m *Req_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req_Payload.Unmarshal(m, b)
}
func (m *Req_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req_Payload.Marshal(b, m, deterministic)
}
func (dst *Req_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req_Payload.Merge(dst, src)
}
func (m *Req_Payload) XXX_Size() int {
	return xxx_messageInfo_Req_Payload.Size(m)
}
func (m *Req_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Req_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Req_Payload proto.InternalMessageInfo

func (m *Req_Payload) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Req_Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Inter struct {
	Req                  *Req       `protobuf:"bytes,2,opt,name=req" json:"req,omitempty"`
	Mode                 Inter_Mode `protobuf:"varint,3,opt,name=mode,enum=twoferrpc.Inter_Mode" json:"mode,omitempty"`
	Ref                  string     `protobuf:"bytes,4,opt,name=ref" json:"ref,omitempty"`
	Inferred             string     `protobuf:"bytes,5,opt,name=inferred" json:"inferred,omitempty"`
	QrData               string     `protobuf:"bytes,6,opt,name=qr_data,json=qrData" json:"qr_data,omitempty"`
	Internal             []byte     `protobuf:"bytes,7,opt,name=internal,proto3" json:"internal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Inter) Reset()         { *m = Inter{} }
func (m *Inter) String() string { return proto.CompactTextString(m) }
func (*Inter) ProtoMessage()    {}
func (*Inter) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{5}
}
func (m *Inter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inter.Unmarshal(m, b)
}
func (m *Inter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inter.Marshal(b, m, deterministic)
}
func (dst *Inter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inter.Merge(dst, src)
}
func (m *Inter) XXX_Size() int {
	return xxx_messageInfo_Inter.Size(m)
}
func (m *Inter) XXX_DiscardUnknown() {
	xxx_messageInfo_Inter.DiscardUnknown(m)
}

var xxx_messageInfo_Inter proto.InternalMessageInfo

func (m *Inter) GetReq() *Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *Inter) GetMode() Inter_Mode {
	if m != nil {
		return m.Mode
	}
	return Inter_AUTH
}

func (m *Inter) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *Inter) GetInferred() string {
	if m != nil {
		return m.Inferred
	}
	return ""
}

func (m *Inter) GetQrData() string {
	if m != nil {
		return m.QrData
	}
	return ""
}

func (m *Inter) GetInternal() []byte {
	if m != nil {
		return m.Internal
	}
	return nil
}

type Info struct {
	Inferred             bool     `protobuf:"varint,1,opt,name=inferred" json:"inferred,omitempty"`
	Ssn                  string   `protobuf:"bytes,2,opt,name=ssn" json:"ssn,omitempty"`
	SsnCountry           string   `protobuf:"bytes,3,opt,name=ssn_country,json=ssnCountry" json:"ssn_country,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	Ip                   string   `protobuf:"bytes,6,opt,name=ip" json:"ip,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	Surname              string   `protobuf:"bytes,8,opt,name=surname" json:"surname,omitempty"`
	DateOfBirth          string   `protobuf:"bytes,9,opt,name=date_of_birth,json=dateOfBirth" json:"date_of_birth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{6}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetInferred() bool {
	if m != nil {
		return m.Inferred
	}
	return false
}

func (m *Info) GetSsn() string {
	if m != nil {
		return m.Ssn
	}
	return ""
}

func (m *Info) GetSsnCountry() string {
	if m != nil {
		return m.SsnCountry
	}
	return ""
}

func (m *Info) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Info) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Info) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Info) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Info) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *Info) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

type Resp struct {
	Inter                *Inter      `protobuf:"bytes,2,opt,name=inter" json:"inter,omitempty"`
	Status               Resp_Status `protobuf:"varint,3,opt,name=status,enum=twoferrpc.Resp_Status" json:"status,omitempty"`
	Info                 *User       `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Signature            []byte      `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Extra                []byte      `protobuf:"bytes,6,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_eid_35bf4589973e30f6, []int{7}
}
func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (dst *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(dst, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetInter() *Inter {
	if m != nil {
		return m.Inter
	}
	return nil
}

func (m *Resp) GetStatus() Resp_Status {
	if m != nil {
		return m.Status
	}
	return Resp_STATUS_UNKNOWN
}

func (m *Resp) GetInfo() *User {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Resp) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Resp) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "twoferrpc.Empty")
	proto.RegisterType((*Providers)(nil), "twoferrpc.Providers")
	proto.RegisterType((*Provider)(nil), "twoferrpc.Provider")
	proto.RegisterType((*User)(nil), "twoferrpc.User")
	proto.RegisterType((*Req)(nil), "twoferrpc.Req")
	proto.RegisterType((*Req_Payload)(nil), "twoferrpc.Req.Payload")
	proto.RegisterType((*Inter)(nil), "twoferrpc.Inter")
	proto.RegisterType((*Info)(nil), "twoferrpc.Info")
	proto.RegisterType((*Resp)(nil), "twoferrpc.Resp")
	proto.RegisterEnum("twoferrpc.Inter_Mode", Inter_Mode_name, Inter_Mode_value)
	proto.RegisterEnum("twoferrpc.Resp_Status", Resp_Status_name, Resp_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EID service

type EIDClient interface {
	GetProviders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Providers, error)
	AuthInit(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Inter, error)
	SignInit(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Inter, error)
	Collect(ctx context.Context, in *Inter, opts ...grpc.CallOption) (*Resp, error)
	Peek(ctx context.Context, in *Inter, opts ...grpc.CallOption) (*Resp, error)
	Cancel(ctx context.Context, in *Inter, opts ...grpc.CallOption) (*Empty, error)
}

type eIDClient struct {
	cc *grpc.ClientConn
}

func NewEIDClient(cc *grpc.ClientConn) EIDClient {
	return &eIDClient{cc}
}

func (c *eIDClient) GetProviders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Providers, error) {
	out := new(Providers)
	err := grpc.Invoke(ctx, "/twoferrpc.EID/GetProviders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIDClient) AuthInit(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Inter, error) {
	out := new(Inter)
	err := grpc.Invoke(ctx, "/twoferrpc.EID/AuthInit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIDClient) SignInit(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Inter, error) {
	out := new(Inter)
	err := grpc.Invoke(ctx, "/twoferrpc.EID/SignInit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIDClient) Collect(ctx context.Context, in *Inter, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/twoferrpc.EID/Collect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIDClient) Peek(ctx context.Context, in *Inter, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/twoferrpc.EID/Peek", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIDClient) Cancel(ctx context.Context, in *Inter, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/twoferrpc.EID/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EID service

type EIDServer interface {
	GetProviders(context.Context, *Empty) (*Providers, error)
	AuthInit(context.Context, *Req) (*Inter, error)
	SignInit(context.Context, *Req) (*Inter, error)
	Collect(context.Context, *Inter) (*Resp, error)
	Peek(context.Context, *Inter) (*Resp, error)
	Cancel(context.Context, *Inter) (*Empty, error)
}

func RegisterEIDServer(s *grpc.Server, srv EIDServer) {
	s.RegisterService(&_EID_serviceDesc, srv)
}

func _EID_GetProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIDServer).GetProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twoferrpc.EID/GetProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIDServer).GetProviders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EID_AuthInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIDServer).AuthInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twoferrpc.EID/AuthInit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIDServer).AuthInit(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _EID_SignInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIDServer).SignInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twoferrpc.EID/SignInit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIDServer).SignInit(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _EID_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Inter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIDServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twoferrpc.EID/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIDServer).Collect(ctx, req.(*Inter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EID_Peek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Inter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIDServer).Peek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twoferrpc.EID/Peek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIDServer).Peek(ctx, req.(*Inter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EID_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Inter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIDServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twoferrpc.EID/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIDServer).Cancel(ctx, req.(*Inter))
	}
	return interceptor(ctx, in, info, handler)
}

var _EID_serviceDesc = grpc.ServiceDesc{
	ServiceName: "twoferrpc.EID",
	HandlerType: (*EIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProviders",
			Handler:    _EID_GetProviders_Handler,
		},
		{
			MethodName: "AuthInit",
			Handler:    _EID_AuthInit_Handler,
		},
		{
			MethodName: "SignInit",
			Handler:    _EID_SignInit_Handler,
		},
		{
			MethodName: "Collect",
			Handler:    _EID_Collect_Handler,
		},
		{
			MethodName: "Peek",
			Handler:    _EID_Peek_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _EID_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eid.proto",
}

func init() { proto.RegisterFile("eid.proto", fileDescriptor_eid_35bf4589973e30f6) }

var fileDescriptor_eid_35bf4589973e30f6 = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x8e, 0xda, 0x46,
	0x14, 0xc6, 0xd8, 0x60, 0x38, 0xb0, 0xac, 0x7b, 0x76, 0xbb, 0xb5, 0x50, 0xd5, 0x6e, 0x5d, 0xa9,
	0xda, 0xaa, 0x92, 0xdb, 0xa5, 0x37, 0xbd, 0xaa, 0x44, 0xc1, 0xdd, 0xba, 0x49, 0x58, 0x6b, 0x80,
	0x24, 0x77, 0xc8, 0x8b, 0x87, 0xc5, 0x0a, 0xd8, 0x66, 0x3c, 0x64, 0x77, 0x1f, 0x23, 0xef, 0x11,
	0xe5, 0x11, 0xf2, 0x12, 0x79, 0x84, 0xdc, 0xe4, 0x31, 0xa2, 0x19, 0x1b, 0xf0, 0xfe, 0x44, 0x4a,
	0xee, 0x72, 0x77, 0xce, 0xf7, 0x7d, 0x73, 0xe6, 0xcc, 0x77, 0x0e, 0x06, 0xea, 0x34, 0x0c, 0xec,
	0x84, 0xc5, 0x3c, 0xc6, 0x3a, 0xbf, 0x8a, 0x67, 0x94, 0xb1, 0x64, 0x6a, 0xe9, 0x50, 0x71, 0x96,
	0x09, 0xbf, 0xb1, 0xfe, 0x86, 0xba, 0xc7, 0xe2, 0x97, 0x61, 0x40, 0x59, 0x8a, 0xa7, 0x50, 0x4f,
	0x36, 0x89, 0xa9, 0x1c, 0xab, 0x27, 0x8d, 0xce, 0x81, 0xbd, 0x3d, 0x64, 0x6f, 0x84, 0x64, 0xa7,
	0xb2, 0x7e, 0x80, 0xda, 0x06, 0x46, 0x04, 0x2d, 0xf2, 0x97, 0xd4, 0x54, 0x8e, 0x95, 0x93, 0x3a,
	0x91, 0xb1, 0xf5, 0x4a, 0x01, 0x6d, 0x9c, 0x52, 0x86, 0x6d, 0xa8, 0x85, 0x91, 0x28, 0x44, 0x03,
	0x29, 0xa8, 0x91, 0x6d, 0x8e, 0x06, 0xa8, 0x69, 0x1a, 0x99, 0x65, 0x79, 0x4e, 0x84, 0xf8, 0x23,
	0x34, 0xd2, 0x34, 0x9a, 0x4c, 0xe3, 0x75, 0xc4, 0xd9, 0x8d, 0xa9, 0x4a, 0x06, 0xd2, 0x34, 0xea,
	0x65, 0x08, 0x1e, 0x42, 0x85, 0x2e, 0xfd, 0x70, 0x61, 0x6a, 0x92, 0xca, 0x12, 0x81, 0x26, 0xf3,
	0x38, 0xa2, 0x66, 0x25, 0x43, 0x65, 0x82, 0x2d, 0x28, 0x87, 0x89, 0x59, 0x95, 0x50, 0x39, 0x4c,
	0xac, 0xb7, 0x0a, 0xa8, 0x84, 0xae, 0xf0, 0x77, 0xa8, 0x6d, 0x1e, 0x22, 0x5b, 0xfa, 0xc4, 0x6b,
	0xb7, 0x22, 0xfc, 0x09, 0xd4, 0xab, 0x79, 0x2c, 0xfb, 0x6c, 0x74, 0xf6, 0x0b, 0x5a, 0xf1, 0x42,
	0x22, 0x38, 0xfc, 0x03, 0xf4, 0xc4, 0xbf, 0x59, 0xc4, 0x7e, 0x20, 0x9b, 0x6e, 0x74, 0x8e, 0x0a,
	0x32, 0x42, 0x57, 0xb6, 0x97, 0xb1, 0x64, 0x23, 0x6b, 0x9f, 0x82, 0x9e, 0x63, 0xc2, 0x40, 0x4e,
	0xaf, 0xf9, 0xc6, 0x40, 0x11, 0x0b, 0x2c, 0xf0, 0xb9, 0x2f, 0x2f, 0x6d, 0x12, 0x19, 0x5b, 0xef,
	0x14, 0xa8, 0xb8, 0x11, 0xa7, 0x0c, 0x8f, 0x41, 0x65, 0x74, 0x95, 0x77, 0xd4, 0xba, 0x7d, 0x15,
	0x11, 0x14, 0xfe, 0x0a, 0xda, 0x32, 0x0e, 0xa8, 0xec, 0xa6, 0xd5, 0xf9, 0xb6, 0x20, 0x91, 0x15,
	0xec, 0x27, 0x71, 0x40, 0x89, 0x94, 0x88, 0x31, 0x30, 0x3a, 0xcb, 0x1d, 0x15, 0xe1, 0xad, 0xa1,
	0x65, 0x96, 0xee, 0x86, 0xf6, 0x1d, 0xe8, 0x2b, 0x36, 0x91, 0xbd, 0x65, 0xd6, 0x56, 0x57, 0xac,
	0xef, 0x73, 0x3f, 0x3b, 0xc4, 0x29, 0x8b, 0xfc, 0x85, 0xa9, 0xcb, 0xae, 0xb7, 0xb9, 0xd5, 0x06,
	0x4d, 0x5c, 0x88, 0x35, 0xd0, 0xba, 0xe3, 0xd1, 0x7f, 0x46, 0x49, 0x44, 0x43, 0xf7, 0x6c, 0x60,
	0x28, 0xd6, 0x7b, 0x05, 0x34, 0x37, 0x9a, 0xc5, 0x5f, 0xd1, 0xaa, 0x6c, 0x57, 0x5a, 0xdf, 0xad,
	0x34, 0x9a, 0xa0, 0xa7, 0x6b, 0x26, 0xe1, 0x9a, 0x84, 0x37, 0x29, 0x5a, 0xb0, 0x17, 0xf8, 0x9c,
	0x4e, 0xe2, 0xd9, 0xe4, 0x22, 0x64, 0x7c, 0x6e, 0xd6, 0x25, 0xdf, 0x10, 0xe0, 0xf9, 0xec, 0x1f,
	0x01, 0x59, 0x1f, 0xca, 0xa0, 0x11, 0x9a, 0x26, 0xf8, 0x0b, 0x54, 0xa4, 0x2d, 0xf9, 0xf0, 0x8c,
	0xbb, 0x93, 0x21, 0x19, 0x8d, 0x36, 0x54, 0x53, 0xee, 0xf3, 0x75, 0x9a, 0x8f, 0xf0, 0xf6, 0x42,
	0xa5, 0x89, 0x3d, 0x94, 0x2c, 0xc9, 0x55, 0xf8, 0x33, 0x68, 0x61, 0x34, 0x8b, 0xe5, 0x6b, 0x1f,
	0xd8, 0x52, 0x49, 0xe2, 0xf7, 0x50, 0x4f, 0xc3, 0xcb, 0xc8, 0xe7, 0x6b, 0x96, 0x39, 0xd0, 0x24,
	0x3b, 0x40, 0x3a, 0x76, 0xcd, 0x59, 0x36, 0xd8, 0x26, 0xc9, 0x12, 0xeb, 0x8d, 0x02, 0xd5, 0xec,
	0x2e, 0x44, 0x68, 0x0d, 0x47, 0xdd, 0xd1, 0x78, 0x38, 0x19, 0x0f, 0x1e, 0x0d, 0xce, 0x9f, 0x0d,
	0x8c, 0x52, 0x01, 0xf3, 0x9c, 0x41, 0xdf, 0x1d, 0x9c, 0x19, 0x0a, 0x1e, 0xc0, 0x7e, 0x8e, 0xf5,
	0xba, 0x83, 0x9e, 0xf3, 0xd8, 0xe9, 0x1b, 0x65, 0x3c, 0x02, 0xcc, 0x41, 0xe2, 0xed, 0x70, 0xb5,
	0x50, 0xc0, 0x79, 0xee, 0xb9, 0xc4, 0xe9, 0x1b, 0x5a, 0xa1, 0x40, 0xd7, 0xf3, 0xc8, 0xf9, 0x53,
	0xa7, 0x6f, 0x54, 0x0a, 0x20, 0x71, 0xfe, 0x77, 0x7a, 0x23, 0xa7, 0x6f, 0x54, 0xf1, 0x1b, 0xd8,
	0xcb, 0xc1, 0x7f, 0xbb, 0xae, 0x28, 0xa8, 0x77, 0x5e, 0x97, 0x41, 0x75, 0xdc, 0x3e, 0xfe, 0x05,
	0xcd, 0x33, 0xca, 0x77, 0x9f, 0xb9, 0xa2, 0xd5, 0xf2, 0x2b, 0xd8, 0x3e, 0x7c, 0xe0, 0x77, 0x9f,
	0x5a, 0x25, 0xb4, 0xa1, 0xd6, 0x5d, 0xf3, 0xb9, 0x1b, 0x85, 0x1c, 0xef, 0xfc, 0xba, 0xda, 0xf7,
	0x06, 0x96, 0xe9, 0x87, 0xe1, 0x65, 0xf4, 0x05, 0x7a, 0xbd, 0x17, 0x2f, 0x16, 0x74, 0xca, 0xf1,
	0x1e, 0xdd, 0xde, 0xbf, 0x33, 0x68, 0xab, 0x84, 0xbf, 0x81, 0xe6, 0x51, 0xfa, 0xe2, 0xf3, 0xc4,
	0x36, 0x54, 0x7b, 0x7e, 0x34, 0xa5, 0x8b, 0x07, 0xe4, 0xf7, 0x2c, 0xb0, 0x4a, 0x17, 0x55, 0xf9,
	0x2f, 0xf1, 0xe7, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x5f, 0x9b, 0x00, 0x32, 0x06, 0x00,
	0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/transport_company.proto

package transport

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TransportCompanyType int32

const (
	TransportCompanyType_boxberry   TransportCompanyType = 0
	TransportCompanyType_cdek       TransportCompanyType = 1
	TransportCompanyType_pochta     TransportCompanyType = 2
	TransportCompanyType_ups        TransportCompanyType = 3
	TransportCompanyType_dpd        TransportCompanyType = 4
	TransportCompanyType_dhl        TransportCompanyType = 5
	TransportCompanyType_redexpress TransportCompanyType = 6
)

var TransportCompanyType_name = map[int32]string{
	0: "boxberry",
	1: "cdek",
	2: "pochta",
	3: "ups",
	4: "dpd",
	5: "dhl",
	6: "redexpress",
}

var TransportCompanyType_value = map[string]int32{
	"boxberry":   0,
	"cdek":       1,
	"pochta":     2,
	"ups":        3,
	"dpd":        4,
	"dhl":        5,
	"redexpress": 6,
}

func (x TransportCompanyType) String() string {
	return proto.EnumName(TransportCompanyType_name, int32(x))
}

func (TransportCompanyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_64d0e24739b69be1, []int{0}
}

type TransportCompanyGetResponse struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type                 TransportCompanyType `protobuf:"varint,3,opt,name=type,proto3,enum=transport.TransportCompanyType" json:"type,omitempty"`
	IsActive             bool                 `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Created              string               `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string               `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransportCompanyGetResponse) Reset()         { *m = TransportCompanyGetResponse{} }
func (m *TransportCompanyGetResponse) String() string { return proto.CompactTextString(m) }
func (*TransportCompanyGetResponse) ProtoMessage()    {}
func (*TransportCompanyGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d0e24739b69be1, []int{0}
}

func (m *TransportCompanyGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransportCompanyGetResponse.Unmarshal(m, b)
}
func (m *TransportCompanyGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransportCompanyGetResponse.Marshal(b, m, deterministic)
}
func (m *TransportCompanyGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransportCompanyGetResponse.Merge(m, src)
}
func (m *TransportCompanyGetResponse) XXX_Size() int {
	return xxx_messageInfo_TransportCompanyGetResponse.Size(m)
}
func (m *TransportCompanyGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransportCompanyGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransportCompanyGetResponse proto.InternalMessageInfo

func (m *TransportCompanyGetResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TransportCompanyGetResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TransportCompanyGetResponse) GetType() TransportCompanyType {
	if m != nil {
		return m.Type
	}
	return TransportCompanyType_boxberry
}

func (m *TransportCompanyGetResponse) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *TransportCompanyGetResponse) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *TransportCompanyGetResponse) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

type TransportCompanyID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransportCompanyID) Reset()         { *m = TransportCompanyID{} }
func (m *TransportCompanyID) String() string { return proto.CompactTextString(m) }
func (*TransportCompanyID) ProtoMessage()    {}
func (*TransportCompanyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d0e24739b69be1, []int{1}
}

func (m *TransportCompanyID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransportCompanyID.Unmarshal(m, b)
}
func (m *TransportCompanyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransportCompanyID.Marshal(b, m, deterministic)
}
func (m *TransportCompanyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransportCompanyID.Merge(m, src)
}
func (m *TransportCompanyID) XXX_Size() int {
	return xxx_messageInfo_TransportCompanyID.Size(m)
}
func (m *TransportCompanyID) XXX_DiscardUnknown() {
	xxx_messageInfo_TransportCompanyID.DiscardUnknown(m)
}

var xxx_messageInfo_TransportCompanyID proto.InternalMessageInfo

func (m *TransportCompanyID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TransportCompanyAddParams struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type                 TransportCompanyType `protobuf:"varint,2,opt,name=type,proto3,enum=transport.TransportCompanyType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransportCompanyAddParams) Reset()         { *m = TransportCompanyAddParams{} }
func (m *TransportCompanyAddParams) String() string { return proto.CompactTextString(m) }
func (*TransportCompanyAddParams) ProtoMessage()    {}
func (*TransportCompanyAddParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d0e24739b69be1, []int{2}
}

func (m *TransportCompanyAddParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransportCompanyAddParams.Unmarshal(m, b)
}
func (m *TransportCompanyAddParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransportCompanyAddParams.Marshal(b, m, deterministic)
}
func (m *TransportCompanyAddParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransportCompanyAddParams.Merge(m, src)
}
func (m *TransportCompanyAddParams) XXX_Size() int {
	return xxx_messageInfo_TransportCompanyAddParams.Size(m)
}
func (m *TransportCompanyAddParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TransportCompanyAddParams.DiscardUnknown(m)
}

var xxx_messageInfo_TransportCompanyAddParams proto.InternalMessageInfo

func (m *TransportCompanyAddParams) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TransportCompanyAddParams) GetType() TransportCompanyType {
	if m != nil {
		return m.Type
	}
	return TransportCompanyType_boxberry
}

type TransportCompanyUpdateParams struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type                 TransportCompanyType `protobuf:"varint,3,opt,name=type,proto3,enum=transport.TransportCompanyType" json:"type,omitempty"`
	IsActive             bool                 `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransportCompanyUpdateParams) Reset()         { *m = TransportCompanyUpdateParams{} }
func (m *TransportCompanyUpdateParams) String() string { return proto.CompactTextString(m) }
func (*TransportCompanyUpdateParams) ProtoMessage()    {}
func (*TransportCompanyUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d0e24739b69be1, []int{3}
}

func (m *TransportCompanyUpdateParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransportCompanyUpdateParams.Unmarshal(m, b)
}
func (m *TransportCompanyUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransportCompanyUpdateParams.Marshal(b, m, deterministic)
}
func (m *TransportCompanyUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransportCompanyUpdateParams.Merge(m, src)
}
func (m *TransportCompanyUpdateParams) XXX_Size() int {
	return xxx_messageInfo_TransportCompanyUpdateParams.Size(m)
}
func (m *TransportCompanyUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TransportCompanyUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_TransportCompanyUpdateParams proto.InternalMessageInfo

func (m *TransportCompanyUpdateParams) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TransportCompanyUpdateParams) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TransportCompanyUpdateParams) GetType() TransportCompanyType {
	if m != nil {
		return m.Type
	}
	return TransportCompanyType_boxberry
}

func (m *TransportCompanyUpdateParams) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

type TransportCompanyOkResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransportCompanyOkResponse) Reset()         { *m = TransportCompanyOkResponse{} }
func (m *TransportCompanyOkResponse) String() string { return proto.CompactTextString(m) }
func (*TransportCompanyOkResponse) ProtoMessage()    {}
func (*TransportCompanyOkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d0e24739b69be1, []int{4}
}

func (m *TransportCompanyOkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransportCompanyOkResponse.Unmarshal(m, b)
}
func (m *TransportCompanyOkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransportCompanyOkResponse.Marshal(b, m, deterministic)
}
func (m *TransportCompanyOkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransportCompanyOkResponse.Merge(m, src)
}
func (m *TransportCompanyOkResponse) XXX_Size() int {
	return xxx_messageInfo_TransportCompanyOkResponse.Size(m)
}
func (m *TransportCompanyOkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransportCompanyOkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransportCompanyOkResponse proto.InternalMessageInfo

func (m *TransportCompanyOkResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterEnum("transport.TransportCompanyType", TransportCompanyType_name, TransportCompanyType_value)
	proto.RegisterType((*TransportCompanyGetResponse)(nil), "transport.TransportCompanyGetResponse")
	proto.RegisterType((*TransportCompanyID)(nil), "transport.TransportCompanyID")
	proto.RegisterType((*TransportCompanyAddParams)(nil), "transport.TransportCompanyAddParams")
	proto.RegisterType((*TransportCompanyUpdateParams)(nil), "transport.TransportCompanyUpdateParams")
	proto.RegisterType((*TransportCompanyOkResponse)(nil), "transport.TransportCompanyOkResponse")
}

func init() { proto.RegisterFile("proto/transport_company.proto", fileDescriptor_64d0e24739b69be1) }

var fileDescriptor_64d0e24739b69be1 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x8d, 0xec, 0xc4, 0x75, 0x2e, 0x25, 0x88, 0x4b, 0x1f, 0xbc, 0x74, 0x65, 0x41, 0x74, 0x9b,
	0x19, 0x23, 0x83, 0xf6, 0x0b, 0xb2, 0x05, 0x4a, 0x61, 0x6c, 0xc5, 0x74, 0x2f, 0x7d, 0x29, 0x8e,
	0x75, 0xb7, 0x1a, 0xa7, 0x91, 0x90, 0xd4, 0x51, 0xff, 0xc8, 0xfe, 0x69, 0x3f, 0xb1, 0x6f, 0x19,
	0x95, 0xdb, 0xd4, 0x98, 0xe2, 0x6d, 0x4f, 0x7b, 0xd3, 0xc9, 0x39, 0xba, 0xe7, 0x9e, 0x1c, 0x2c,
	0x38, 0xd0, 0x46, 0x39, 0xf5, 0xce, 0x99, 0x7c, 0x63, 0xb5, 0x32, 0xee, 0xb2, 0x50, 0xd7, 0x3a,
	0xdf, 0xd4, 0x73, 0xff, 0x3b, 0x8e, 0xb7, 0x84, 0xf8, 0xc9, 0x60, 0xff, 0xfc, 0x01, 0x7d, 0x68,
	0x54, 0x27, 0xe4, 0x32, 0xb2, 0x5a, 0x6d, 0x2c, 0xe1, 0x04, 0x82, 0x52, 0x26, 0x6c, 0xc6, 0xd2,
	0x51, 0x16, 0x94, 0x12, 0xf7, 0x60, 0xe4, 0x4a, 0xb7, 0xa6, 0x24, 0x98, 0xb1, 0x74, 0x9c, 0x35,
	0x00, 0x8f, 0x61, 0xe8, 0x6a, 0x4d, 0x49, 0x38, 0x63, 0xe9, 0xe4, 0xe8, 0xc5, 0x7c, 0x3b, 0x7f,
	0xde, 0x9d, 0x7d, 0x5e, 0x6b, 0xca, 0xbc, 0x18, 0xf7, 0x61, 0x5c, 0xda, 0xcb, 0xbc, 0x70, 0xe5,
	0x77, 0x4a, 0x86, 0x33, 0x96, 0xc6, 0x59, 0x5c, 0xda, 0x85, 0xc7, 0x98, 0xc0, 0x4e, 0x61, 0x28,
	0x77, 0x24, 0x93, 0x91, 0x77, 0x7a, 0x80, 0x77, 0xcc, 0x8d, 0x96, 0x9e, 0x89, 0x1a, 0xe6, 0x1e,
	0x8a, 0x43, 0xc0, 0xae, 0xdd, 0xe9, 0xb2, 0x9b, 0x40, 0x7c, 0x85, 0x67, 0x5d, 0xd5, 0x42, 0xca,
	0xb3, 0xdc, 0xe4, 0xd7, 0xf6, 0x31, 0x1e, 0x7b, 0x2a, 0x5e, 0xf0, 0x0f, 0xf1, 0xc4, 0x0f, 0x06,
	0xcf, 0xbb, 0xf4, 0x17, 0xbf, 0xe9, 0xbd, 0xd7, 0x7f, 0xfa, 0x6b, 0xc5, 0x5b, 0x98, 0x76, 0xaf,
	0x7e, 0xae, 0xda, 0x85, 0xab, 0xca, 0x6f, 0x15, 0x67, 0x81, 0xaa, 0xde, 0x10, 0xec, 0x3d, 0x65,
	0x84, 0xbb, 0x10, 0xaf, 0xd4, 0xed, 0x8a, 0x8c, 0xa9, 0xf9, 0x00, 0x63, 0x18, 0x16, 0x92, 0x2a,
	0xce, 0x10, 0x20, 0xd2, 0xaa, 0xb8, 0x72, 0x39, 0x0f, 0x70, 0x07, 0xc2, 0x1b, 0x6d, 0x79, 0x78,
	0x77, 0x90, 0x5a, 0xf2, 0xa1, 0x3f, 0x5c, 0xad, 0xf9, 0x08, 0x27, 0x00, 0x86, 0x24, 0xdd, 0x6a,
	0x43, 0xd6, 0xf2, 0xe8, 0xe8, 0x57, 0x00, 0xbc, 0xeb, 0x83, 0x1f, 0x21, 0x5c, 0x48, 0x89, 0x87,
	0x3d, 0xa1, 0xb7, 0xd5, 0x4d, 0x0f, 0x7a, 0x54, 0xa7, 0x4b, 0x31, 0xc0, 0x4f, 0x10, 0x9e, 0x90,
	0xc3, 0x7e, 0xdd, 0xf4, 0x55, 0x0f, 0xdd, 0xfa, 0x30, 0xc4, 0x00, 0x2f, 0x20, 0x6a, 0xfa, 0xc4,
	0xd7, 0x3d, 0x77, 0xda, 0x95, 0x4f, 0x5f, 0xf6, 0x08, 0x1f, 0x3b, 0x10, 0x03, 0x3c, 0x83, 0x68,
	0x49, 0x6b, 0x72, 0xf4, 0xa7, 0x75, 0xff, 0x76, 0xe2, 0xfb, 0xc9, 0xc5, 0xee, 0xb7, 0xd6, 0x8b,
	0xb0, 0x8a, 0xfc, 0x53, 0x70, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x39, 0xb1, 0xcb, 0x8f, 0x2b,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransportCompanyClient is the client API for TransportCompany service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransportCompanyClient interface {
	Add(ctx context.Context, in *TransportCompanyAddParams, opts ...grpc.CallOption) (*TransportCompanyID, error)
	Get(ctx context.Context, in *TransportCompanyID, opts ...grpc.CallOption) (*TransportCompanyGetResponse, error)
	Update(ctx context.Context, in *TransportCompanyUpdateParams, opts ...grpc.CallOption) (*TransportCompanyOkResponse, error)
	Delete(ctx context.Context, in *TransportCompanyID, opts ...grpc.CallOption) (*TransportCompanyOkResponse, error)
}

type transportCompanyClient struct {
	cc *grpc.ClientConn
}

func NewTransportCompanyClient(cc *grpc.ClientConn) TransportCompanyClient {
	return &transportCompanyClient{cc}
}

func (c *transportCompanyClient) Add(ctx context.Context, in *TransportCompanyAddParams, opts ...grpc.CallOption) (*TransportCompanyID, error) {
	out := new(TransportCompanyID)
	err := c.cc.Invoke(ctx, "/transport.TransportCompany/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyClient) Get(ctx context.Context, in *TransportCompanyID, opts ...grpc.CallOption) (*TransportCompanyGetResponse, error) {
	out := new(TransportCompanyGetResponse)
	err := c.cc.Invoke(ctx, "/transport.TransportCompany/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyClient) Update(ctx context.Context, in *TransportCompanyUpdateParams, opts ...grpc.CallOption) (*TransportCompanyOkResponse, error) {
	out := new(TransportCompanyOkResponse)
	err := c.cc.Invoke(ctx, "/transport.TransportCompany/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyClient) Delete(ctx context.Context, in *TransportCompanyID, opts ...grpc.CallOption) (*TransportCompanyOkResponse, error) {
	out := new(TransportCompanyOkResponse)
	err := c.cc.Invoke(ctx, "/transport.TransportCompany/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportCompanyServer is the server API for TransportCompany service.
type TransportCompanyServer interface {
	Add(context.Context, *TransportCompanyAddParams) (*TransportCompanyID, error)
	Get(context.Context, *TransportCompanyID) (*TransportCompanyGetResponse, error)
	Update(context.Context, *TransportCompanyUpdateParams) (*TransportCompanyOkResponse, error)
	Delete(context.Context, *TransportCompanyID) (*TransportCompanyOkResponse, error)
}

// UnimplementedTransportCompanyServer can be embedded to have forward compatible implementations.
type UnimplementedTransportCompanyServer struct {
}

func (*UnimplementedTransportCompanyServer) Add(ctx context.Context, req *TransportCompanyAddParams) (*TransportCompanyID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTransportCompanyServer) Get(ctx context.Context, req *TransportCompanyID) (*TransportCompanyGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTransportCompanyServer) Update(ctx context.Context, req *TransportCompanyUpdateParams) (*TransportCompanyOkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTransportCompanyServer) Delete(ctx context.Context, req *TransportCompanyID) (*TransportCompanyOkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTransportCompanyServer(s *grpc.Server, srv TransportCompanyServer) {
	s.RegisterService(&_TransportCompany_serviceDesc, srv)
}

func _TransportCompany_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompanyAddParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TransportCompany/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServer).Add(ctx, req.(*TransportCompanyAddParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompany_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompanyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TransportCompany/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServer).Get(ctx, req.(*TransportCompanyID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompany_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompanyUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TransportCompany/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServer).Update(ctx, req.(*TransportCompanyUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportCompany_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransportCompanyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportCompanyServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TransportCompany/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportCompanyServer).Delete(ctx, req.(*TransportCompanyID))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransportCompany_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transport.TransportCompany",
	HandlerType: (*TransportCompanyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TransportCompany_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TransportCompany_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TransportCompany_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TransportCompany_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transport_company.proto",
}

// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/staff.proto

package staff

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Staff service

func NewStaffEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Staff service

type StaffService interface {
	GetLeavingEmployees(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*LeavingList, error)
}

type staffService struct {
	c    client.Client
	name string
}

func NewStaffService(name string, c client.Client) StaffService {
	return &staffService{
		c:    c,
		name: name,
	}
}

func (c *staffService) GetLeavingEmployees(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*LeavingList, error) {
	req := c.c.NewRequest(c.name, "Staff.getLeavingEmployees", in)
	out := new(LeavingList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Staff service

type StaffHandler interface {
	GetLeavingEmployees(context.Context, *empty.Empty, *LeavingList) error
}

func RegisterStaffHandler(s server.Server, hdlr StaffHandler, opts ...server.HandlerOption) error {
	type staff interface {
		GetLeavingEmployees(ctx context.Context, in *empty.Empty, out *LeavingList) error
	}
	type Staff struct {
		staff
	}
	h := &staffHandler{hdlr}
	return s.Handle(s.NewHandler(&Staff{h}, opts...))
}

type staffHandler struct {
	StaffHandler
}

func (h *staffHandler) GetLeavingEmployees(ctx context.Context, in *empty.Empty, out *LeavingList) error {
	return h.StaffHandler.GetLeavingEmployees(ctx, in, out)
}

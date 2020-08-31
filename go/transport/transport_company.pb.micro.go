// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/transport_company.proto

package transport

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Api Endpoints for TransportCompany service

func NewTransportCompanyEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TransportCompany service

type TransportCompanyService interface {
	Add(ctx context.Context, in *TransportCompanyAddParams, opts ...client.CallOption) (*TransportCompanyID, error)
	Get(ctx context.Context, in *TransportCompanyID, opts ...client.CallOption) (*TransportCompanyGetResponse, error)
	Update(ctx context.Context, in *TransportCompanyUpdateParams, opts ...client.CallOption) (*TransportCompanyOkResponse, error)
	Delete(ctx context.Context, in *TransportCompanyID, opts ...client.CallOption) (*TransportCompanyOkResponse, error)
}

type transportCompanyService struct {
	c    client.Client
	name string
}

func NewTransportCompanyService(name string, c client.Client) TransportCompanyService {
	return &transportCompanyService{
		c:    c,
		name: name,
	}
}

func (c *transportCompanyService) Add(ctx context.Context, in *TransportCompanyAddParams, opts ...client.CallOption) (*TransportCompanyID, error) {
	req := c.c.NewRequest(c.name, "TransportCompany.Add", in)
	out := new(TransportCompanyID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyService) Get(ctx context.Context, in *TransportCompanyID, opts ...client.CallOption) (*TransportCompanyGetResponse, error) {
	req := c.c.NewRequest(c.name, "TransportCompany.Get", in)
	out := new(TransportCompanyGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyService) Update(ctx context.Context, in *TransportCompanyUpdateParams, opts ...client.CallOption) (*TransportCompanyOkResponse, error) {
	req := c.c.NewRequest(c.name, "TransportCompany.Update", in)
	out := new(TransportCompanyOkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportCompanyService) Delete(ctx context.Context, in *TransportCompanyID, opts ...client.CallOption) (*TransportCompanyOkResponse, error) {
	req := c.c.NewRequest(c.name, "TransportCompany.Delete", in)
	out := new(TransportCompanyOkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TransportCompany service

type TransportCompanyHandler interface {
	Add(context.Context, *TransportCompanyAddParams, *TransportCompanyID) error
	Get(context.Context, *TransportCompanyID, *TransportCompanyGetResponse) error
	Update(context.Context, *TransportCompanyUpdateParams, *TransportCompanyOkResponse) error
	Delete(context.Context, *TransportCompanyID, *TransportCompanyOkResponse) error
}

func RegisterTransportCompanyHandler(s server.Server, hdlr TransportCompanyHandler, opts ...server.HandlerOption) error {
	type transportCompany interface {
		Add(ctx context.Context, in *TransportCompanyAddParams, out *TransportCompanyID) error
		Get(ctx context.Context, in *TransportCompanyID, out *TransportCompanyGetResponse) error
		Update(ctx context.Context, in *TransportCompanyUpdateParams, out *TransportCompanyOkResponse) error
		Delete(ctx context.Context, in *TransportCompanyID, out *TransportCompanyOkResponse) error
	}
	type TransportCompany struct {
		transportCompany
	}
	h := &transportCompanyHandler{hdlr}
	return s.Handle(s.NewHandler(&TransportCompany{h}, opts...))
}

type transportCompanyHandler struct {
	TransportCompanyHandler
}

func (h *transportCompanyHandler) Add(ctx context.Context, in *TransportCompanyAddParams, out *TransportCompanyID) error {
	return h.TransportCompanyHandler.Add(ctx, in, out)
}

func (h *transportCompanyHandler) Get(ctx context.Context, in *TransportCompanyID, out *TransportCompanyGetResponse) error {
	return h.TransportCompanyHandler.Get(ctx, in, out)
}

func (h *transportCompanyHandler) Update(ctx context.Context, in *TransportCompanyUpdateParams, out *TransportCompanyOkResponse) error {
	return h.TransportCompanyHandler.Update(ctx, in, out)
}

func (h *transportCompanyHandler) Delete(ctx context.Context, in *TransportCompanyID, out *TransportCompanyOkResponse) error {
	return h.TransportCompanyHandler.Delete(ctx, in, out)
}
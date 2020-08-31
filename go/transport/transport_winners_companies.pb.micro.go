// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/transport_winners_companies.proto

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

// Api Endpoints for WinnersCompanies service

func NewWinnersCompaniesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WinnersCompanies service

type WinnersCompaniesService interface {
	Winners(ctx context.Context, in *WinnersParams, opts ...client.CallOption) (*WinnersResponse, error)
}

type winnersCompaniesService struct {
	c    client.Client
	name string
}

func NewWinnersCompaniesService(name string, c client.Client) WinnersCompaniesService {
	return &winnersCompaniesService{
		c:    c,
		name: name,
	}
}

func (c *winnersCompaniesService) Winners(ctx context.Context, in *WinnersParams, opts ...client.CallOption) (*WinnersResponse, error) {
	req := c.c.NewRequest(c.name, "WinnersCompanies.Winners", in)
	out := new(WinnersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WinnersCompanies service

type WinnersCompaniesHandler interface {
	Winners(context.Context, *WinnersParams, *WinnersResponse) error
}

func RegisterWinnersCompaniesHandler(s server.Server, hdlr WinnersCompaniesHandler, opts ...server.HandlerOption) error {
	type winnersCompanies interface {
		Winners(ctx context.Context, in *WinnersParams, out *WinnersResponse) error
	}
	type WinnersCompanies struct {
		winnersCompanies
	}
	h := &winnersCompaniesHandler{hdlr}
	return s.Handle(s.NewHandler(&WinnersCompanies{h}, opts...))
}

type winnersCompaniesHandler struct {
	WinnersCompaniesHandler
}

func (h *winnersCompaniesHandler) Winners(ctx context.Context, in *WinnersParams, out *WinnersResponse) error {
	return h.WinnersCompaniesHandler.Winners(ctx, in, out)
}
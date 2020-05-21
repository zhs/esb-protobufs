// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/mercaux.proto

package mercaux

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

// Api Endpoints for Stocks service

func NewStocksEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{},
	}
}

// Client API for Stocks service

type StocksService interface {
	Get(ctx context.Context, in *StocksGetParams, opts ...client.CallOption) (*StocksGetResponse, error)
}

type stocksService struct {
	c    client.Client
	name string
}

func NewStocksService(name string, c client.Client) StocksService {
	return &stocksService{
		c:    c,
		name: name,
	}
}

func (c *stocksService) Get(ctx context.Context, in *StocksGetParams, opts ...client.CallOption) (*StocksGetResponse, error) {
	req := c.c.NewRequest(c.name, "Stocks.Get", in)
	out := new(StocksGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stocks service

type StocksHandler interface {
	Get(context.Context, *StocksGetParams, *StocksGetResponse) error
}

func RegisterStocksHandler(s server.Server, hdlr StocksHandler, opts ...server.HandlerOption) error {
	type stocks interface {
		Get(ctx context.Context, in *StocksGetParams, out *StocksGetResponse) error
	}
	type Stocks struct {
		stocks
	}
	h := &stocksHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{}))
	return s.Handle(s.NewHandler(&Stocks{h}, opts...))
}

type stocksHandler struct {
	StocksHandler
}

func (h *stocksHandler) Get(ctx context.Context, in *StocksGetParams, out *StocksGetResponse) error {
	return h.StocksHandler.Get(ctx, in, out)
}

// Api Endpoints for Stores service

func NewStoresEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{},
	}
}

// Client API for Stores service

type StoresService interface {
	Get(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*StoresGetResponse, error)
}

type storesService struct {
	c    client.Client
	name string
}

func NewStoresService(name string, c client.Client) StoresService {
	return &storesService{
		c:    c,
		name: name,
	}
}

func (c *storesService) Get(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*StoresGetResponse, error) {
	req := c.c.NewRequest(c.name, "Stores.Get", in)
	out := new(StoresGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stores service

type StoresHandler interface {
	Get(context.Context, *empty.Empty, *StoresGetResponse) error
}

func RegisterStoresHandler(s server.Server, hdlr StoresHandler, opts ...server.HandlerOption) error {
	type stores interface {
		Get(ctx context.Context, in *empty.Empty, out *StoresGetResponse) error
	}
	type Stores struct {
		stores
	}
	h := &storesHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{}))
	return s.Handle(s.NewHandler(&Stores{h}, opts...))
}

type storesHandler struct {
	StoresHandler
}

func (h *storesHandler) Get(ctx context.Context, in *empty.Empty, out *StoresGetResponse) error {
	return h.StoresHandler.Get(ctx, in, out)
}

// Api Endpoints for Sellers service

func NewSellersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{},
	}
}

// Client API for Sellers service

type SellersService interface {
	Get(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*SellersGetResponse, error)
}

type sellersService struct {
	c    client.Client
	name string
}

func NewSellersService(name string, c client.Client) SellersService {
	return &sellersService{
		c:    c,
		name: name,
	}
}

func (c *sellersService) Get(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*SellersGetResponse, error) {
	req := c.c.NewRequest(c.name, "Sellers.Get", in)
	out := new(SellersGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sellers service

type SellersHandler interface {
	Get(context.Context, *empty.Empty, *SellersGetResponse) error
}

func RegisterSellersHandler(s server.Server, hdlr SellersHandler, opts ...server.HandlerOption) error {
	type sellers interface {
		Get(ctx context.Context, in *empty.Empty, out *SellersGetResponse) error
	}
	type Sellers struct {
		sellers
	}
	h := &sellersHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{}))
	return s.Handle(s.NewHandler(&Sellers{h}, opts...))
}

type sellersHandler struct {
	SellersHandler
}

func (h *sellersHandler) Get(ctx context.Context, in *empty.Empty, out *SellersGetResponse) error {
	return h.SellersHandler.Get(ctx, in, out)
}

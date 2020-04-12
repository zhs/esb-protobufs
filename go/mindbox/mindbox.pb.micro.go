// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/mindbox.proto

package mindbox

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Info(ctx context.Context, in *ParamsUser, opts ...client.CallOption) (*ResponseUser, error)
	Orders(ctx context.Context, in *ParamsOrders, opts ...client.CallOption) (*ResponseOrders, error)
	SendOSMICard(ctx context.Context, in *ParamsOSMICard, opts ...client.CallOption) (*ResponseOSMICard, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Info(ctx context.Context, in *ParamsUser, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.name, "User.Info", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Orders(ctx context.Context, in *ParamsOrders, opts ...client.CallOption) (*ResponseOrders, error) {
	req := c.c.NewRequest(c.name, "User.Orders", in)
	out := new(ResponseOrders)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SendOSMICard(ctx context.Context, in *ParamsOSMICard, opts ...client.CallOption) (*ResponseOSMICard, error) {
	req := c.c.NewRequest(c.name, "User.SendOSMICard", in)
	out := new(ResponseOSMICard)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Info(context.Context, *ParamsUser, *ResponseUser) error
	Orders(context.Context, *ParamsOrders, *ResponseOrders) error
	SendOSMICard(context.Context, *ParamsOSMICard, *ResponseOSMICard) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Info(ctx context.Context, in *ParamsUser, out *ResponseUser) error
		Orders(ctx context.Context, in *ParamsOrders, out *ResponseOrders) error
		SendOSMICard(ctx context.Context, in *ParamsOSMICard, out *ResponseOSMICard) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Info(ctx context.Context, in *ParamsUser, out *ResponseUser) error {
	return h.UserHandler.Info(ctx, in, out)
}

func (h *userHandler) Orders(ctx context.Context, in *ParamsOrders, out *ResponseOrders) error {
	return h.UserHandler.Orders(ctx, in, out)
}

func (h *userHandler) SendOSMICard(ctx context.Context, in *ParamsOSMICard, out *ResponseOSMICard) error {
	return h.UserHandler.SendOSMICard(ctx, in, out)
}

// Client API for Mobile service

type MobileService interface {
	InitDevice(ctx context.Context, in *InitDeviceParams, opts ...client.CallOption) (*InitDeviceResponse, error)
	InitClient(ctx context.Context, in *InitClientParams, opts ...client.CallOption) (*InitClientResponse, error)
	Code(ctx context.Context, in *ParamsCode, opts ...client.CallOption) (*ResponseCode, error)
	CheckCode(ctx context.Context, in *ParamsCheckCode, opts ...client.CallOption) (*ResponseCheckCode, error)
	EditUser(ctx context.Context, in *ParamsEditUser, opts ...client.CallOption) (*ResponseEditUser, error)
}

type mobileService struct {
	c    client.Client
	name string
}

func NewMobileService(name string, c client.Client) MobileService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mobile"
	}
	return &mobileService{
		c:    c,
		name: name,
	}
}

func (c *mobileService) InitDevice(ctx context.Context, in *InitDeviceParams, opts ...client.CallOption) (*InitDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "Mobile.InitDevice", in)
	out := new(InitDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) InitClient(ctx context.Context, in *InitClientParams, opts ...client.CallOption) (*InitClientResponse, error) {
	req := c.c.NewRequest(c.name, "Mobile.InitClient", in)
	out := new(InitClientResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) Code(ctx context.Context, in *ParamsCode, opts ...client.CallOption) (*ResponseCode, error) {
	req := c.c.NewRequest(c.name, "Mobile.Code", in)
	out := new(ResponseCode)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) CheckCode(ctx context.Context, in *ParamsCheckCode, opts ...client.CallOption) (*ResponseCheckCode, error) {
	req := c.c.NewRequest(c.name, "Mobile.CheckCode", in)
	out := new(ResponseCheckCode)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) EditUser(ctx context.Context, in *ParamsEditUser, opts ...client.CallOption) (*ResponseEditUser, error) {
	req := c.c.NewRequest(c.name, "Mobile.EditUser", in)
	out := new(ResponseEditUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mobile service

type MobileHandler interface {
	InitDevice(context.Context, *InitDeviceParams, *InitDeviceResponse) error
	InitClient(context.Context, *InitClientParams, *InitClientResponse) error
	Code(context.Context, *ParamsCode, *ResponseCode) error
	CheckCode(context.Context, *ParamsCheckCode, *ResponseCheckCode) error
	EditUser(context.Context, *ParamsEditUser, *ResponseEditUser) error
}

func RegisterMobileHandler(s server.Server, hdlr MobileHandler, opts ...server.HandlerOption) error {
	type mobile interface {
		InitDevice(ctx context.Context, in *InitDeviceParams, out *InitDeviceResponse) error
		InitClient(ctx context.Context, in *InitClientParams, out *InitClientResponse) error
		Code(ctx context.Context, in *ParamsCode, out *ResponseCode) error
		CheckCode(ctx context.Context, in *ParamsCheckCode, out *ResponseCheckCode) error
		EditUser(ctx context.Context, in *ParamsEditUser, out *ResponseEditUser) error
	}
	type Mobile struct {
		mobile
	}
	h := &mobileHandler{hdlr}
	return s.Handle(s.NewHandler(&Mobile{h}, opts...))
}

type mobileHandler struct {
	MobileHandler
}

func (h *mobileHandler) InitDevice(ctx context.Context, in *InitDeviceParams, out *InitDeviceResponse) error {
	return h.MobileHandler.InitDevice(ctx, in, out)
}

func (h *mobileHandler) InitClient(ctx context.Context, in *InitClientParams, out *InitClientResponse) error {
	return h.MobileHandler.InitClient(ctx, in, out)
}

func (h *mobileHandler) Code(ctx context.Context, in *ParamsCode, out *ResponseCode) error {
	return h.MobileHandler.Code(ctx, in, out)
}

func (h *mobileHandler) CheckCode(ctx context.Context, in *ParamsCheckCode, out *ResponseCheckCode) error {
	return h.MobileHandler.CheckCode(ctx, in, out)
}

func (h *mobileHandler) EditUser(ctx context.Context, in *ParamsEditUser, out *ResponseEditUser) error {
	return h.MobileHandler.EditUser(ctx, in, out)
}

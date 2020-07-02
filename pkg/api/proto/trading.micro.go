// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/proto/trading.proto

package tradingAPI

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

// Client API for TradingService service

type TradingService interface {
	GetAccounts(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetGroups(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupResponse, error)
	GetBalances(ctx context.Context, in *BalanceRequest, opts ...client.CallOption) (*BalanceResponse, error)
	GetPositions(ctx context.Context, in *PositionRequest, opts ...client.CallOption) (*PositionResponse, error)
	CreateOrder(ctx context.Context, in *NewOrder, opts ...client.CallOption) (*OrderResponse, error)
	GetOrders(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
}

type tradingService struct {
	c    client.Client
	name string
}

func NewTradingService(name string, c client.Client) TradingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "tradingAPI"
	}
	return &tradingService{
		c:    c,
		name: name,
	}
}

func (c *tradingService) GetAccounts(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "TradingService.GetAccounts", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingService) GetGroups(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupResponse, error) {
	req := c.c.NewRequest(c.name, "TradingService.GetGroups", in)
	out := new(GroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingService) GetBalances(ctx context.Context, in *BalanceRequest, opts ...client.CallOption) (*BalanceResponse, error) {
	req := c.c.NewRequest(c.name, "TradingService.GetBalances", in)
	out := new(BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingService) GetPositions(ctx context.Context, in *PositionRequest, opts ...client.CallOption) (*PositionResponse, error) {
	req := c.c.NewRequest(c.name, "TradingService.GetPositions", in)
	out := new(PositionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingService) CreateOrder(ctx context.Context, in *NewOrder, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "TradingService.CreateOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingService) GetOrders(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "TradingService.GetOrders", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TradingService service

type TradingServiceHandler interface {
	GetAccounts(context.Context, *AccountRequest, *AccountResponse) error
	GetGroups(context.Context, *GroupRequest, *GroupResponse) error
	GetBalances(context.Context, *BalanceRequest, *BalanceResponse) error
	GetPositions(context.Context, *PositionRequest, *PositionResponse) error
	CreateOrder(context.Context, *NewOrder, *OrderResponse) error
	GetOrders(context.Context, *OrderRequest, *OrderResponse) error
}

func RegisterTradingServiceHandler(s server.Server, hdlr TradingServiceHandler, opts ...server.HandlerOption) error {
	type tradingService interface {
		GetAccounts(ctx context.Context, in *AccountRequest, out *AccountResponse) error
		GetGroups(ctx context.Context, in *GroupRequest, out *GroupResponse) error
		GetBalances(ctx context.Context, in *BalanceRequest, out *BalanceResponse) error
		GetPositions(ctx context.Context, in *PositionRequest, out *PositionResponse) error
		CreateOrder(ctx context.Context, in *NewOrder, out *OrderResponse) error
		GetOrders(ctx context.Context, in *OrderRequest, out *OrderResponse) error
	}
	type TradingService struct {
		tradingService
	}
	h := &tradingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TradingService{h}, opts...))
}

type tradingServiceHandler struct {
	TradingServiceHandler
}

func (h *tradingServiceHandler) GetAccounts(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.TradingServiceHandler.GetAccounts(ctx, in, out)
}

func (h *tradingServiceHandler) GetGroups(ctx context.Context, in *GroupRequest, out *GroupResponse) error {
	return h.TradingServiceHandler.GetGroups(ctx, in, out)
}

func (h *tradingServiceHandler) GetBalances(ctx context.Context, in *BalanceRequest, out *BalanceResponse) error {
	return h.TradingServiceHandler.GetBalances(ctx, in, out)
}

func (h *tradingServiceHandler) GetPositions(ctx context.Context, in *PositionRequest, out *PositionResponse) error {
	return h.TradingServiceHandler.GetPositions(ctx, in, out)
}

func (h *tradingServiceHandler) CreateOrder(ctx context.Context, in *NewOrder, out *OrderResponse) error {
	return h.TradingServiceHandler.CreateOrder(ctx, in, out)
}

func (h *tradingServiceHandler) GetOrders(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.TradingServiceHandler.GetOrders(ctx, in, out)
}
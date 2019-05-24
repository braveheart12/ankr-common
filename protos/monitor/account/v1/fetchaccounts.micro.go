// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: account/v1/fetchaccounts.proto

package account

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

// Client API for FetchAccounts service

type FetchAccountsService interface {
	Fetch(ctx context.Context, in *FetchAccountsReq, opts ...client.CallOption) (*FetchAccountsResp, error)
}

type fetchAccountsService struct {
	c    client.Client
	name string
}

func NewFetchAccountsService(name string, c client.Client) FetchAccountsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "account"
	}
	return &fetchAccountsService{
		c:    c,
		name: name,
	}
}

func (c *fetchAccountsService) Fetch(ctx context.Context, in *FetchAccountsReq, opts ...client.CallOption) (*FetchAccountsResp, error) {
	req := c.c.NewRequest(c.name, "FetchAccounts.Fetch", in)
	out := new(FetchAccountsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FetchAccounts service

type FetchAccountsHandler interface {
	Fetch(context.Context, *FetchAccountsReq, *FetchAccountsResp) error
}

func RegisterFetchAccountsHandler(s server.Server, hdlr FetchAccountsHandler, opts ...server.HandlerOption) error {
	type fetchAccounts interface {
		Fetch(ctx context.Context, in *FetchAccountsReq, out *FetchAccountsResp) error
	}
	type FetchAccounts struct {
		fetchAccounts
	}
	h := &fetchAccountsHandler{hdlr}
	return s.Handle(s.NewHandler(&FetchAccounts{h}, opts...))
}

type fetchAccountsHandler struct {
	FetchAccountsHandler
}

func (h *fetchAccountsHandler) Fetch(ctx context.Context, in *FetchAccountsReq, out *FetchAccountsResp) error {
	return h.FetchAccountsHandler.Fetch(ctx, in, out)
}
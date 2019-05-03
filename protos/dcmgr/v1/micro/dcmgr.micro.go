// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: dcmgr/v1/micro/dcmgr.proto

package dcmgr

import (
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
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

// Client API for DC service

type DCService interface {
	CreateApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error)
	UpdateApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error)
	DeleteApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error)
	CreateNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error)
	UpdateNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error)
	DeleteNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error)
	Status(ctx context.Context, in *common.AppID, opts ...client.CallOption) (*common.AppRunStatus, error)
	InitDC(ctx context.Context, in *common.DataCenter, opts ...client.CallOption) (*common.AppResponce, error)
	Overview(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*common.DataCenterStatus, error)
}

type dCService struct {
	c    client.Client
	name string
}

func NewDCService(name string, c client.Client) DCService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dcmgr"
	}
	return &dCService{
		c:    c,
		name: name,
	}
}

func (c *dCService) CreateApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.CreateApp", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) UpdateApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.UpdateApp", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) DeleteApp(ctx context.Context, in *common.AppDeployment, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.DeleteApp", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) CreateNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.CreateNamespace", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) UpdateNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.UpdateNamespace", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) DeleteNamespace(ctx context.Context, in *common.Namespace, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.DeleteNamespace", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) Status(ctx context.Context, in *common.AppID, opts ...client.CallOption) (*common.AppRunStatus, error) {
	req := c.c.NewRequest(c.name, "DC.Status", in)
	out := new(common.AppRunStatus)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) InitDC(ctx context.Context, in *common.DataCenter, opts ...client.CallOption) (*common.AppResponce, error) {
	req := c.c.NewRequest(c.name, "DC.InitDC", in)
	out := new(common.AppResponce)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCService) Overview(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*common.DataCenterStatus, error) {
	req := c.c.NewRequest(c.name, "DC.Overview", in)
	out := new(common.DataCenterStatus)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DC service

type DCHandler interface {
	CreateApp(context.Context, *common.AppDeployment, *common.AppResponce) error
	UpdateApp(context.Context, *common.AppDeployment, *common.AppResponce) error
	DeleteApp(context.Context, *common.AppDeployment, *common.AppResponce) error
	CreateNamespace(context.Context, *common.Namespace, *common.AppResponce) error
	UpdateNamespace(context.Context, *common.Namespace, *common.AppResponce) error
	DeleteNamespace(context.Context, *common.Namespace, *common.AppResponce) error
	Status(context.Context, *common.AppID, *common.AppRunStatus) error
	InitDC(context.Context, *common.DataCenter, *common.AppResponce) error
	Overview(context.Context, *common.Empty, *common.DataCenterStatus) error
}

func RegisterDCHandler(s server.Server, hdlr DCHandler, opts ...server.HandlerOption) error {
	type dC interface {
		CreateApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error
		UpdateApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error
		DeleteApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error
		CreateNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error
		UpdateNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error
		DeleteNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error
		Status(ctx context.Context, in *common.AppID, out *common.AppRunStatus) error
		InitDC(ctx context.Context, in *common.DataCenter, out *common.AppResponce) error
		Overview(ctx context.Context, in *common.Empty, out *common.DataCenterStatus) error
	}
	type DC struct {
		dC
	}
	h := &dCHandler{hdlr}
	return s.Handle(s.NewHandler(&DC{h}, opts...))
}

type dCHandler struct {
	DCHandler
}

func (h *dCHandler) CreateApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error {
	return h.DCHandler.CreateApp(ctx, in, out)
}

func (h *dCHandler) UpdateApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error {
	return h.DCHandler.UpdateApp(ctx, in, out)
}

func (h *dCHandler) DeleteApp(ctx context.Context, in *common.AppDeployment, out *common.AppResponce) error {
	return h.DCHandler.DeleteApp(ctx, in, out)
}

func (h *dCHandler) CreateNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error {
	return h.DCHandler.CreateNamespace(ctx, in, out)
}

func (h *dCHandler) UpdateNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error {
	return h.DCHandler.UpdateNamespace(ctx, in, out)
}

func (h *dCHandler) DeleteNamespace(ctx context.Context, in *common.Namespace, out *common.AppResponce) error {
	return h.DCHandler.DeleteNamespace(ctx, in, out)
}

func (h *dCHandler) Status(ctx context.Context, in *common.AppID, out *common.AppRunStatus) error {
	return h.DCHandler.Status(ctx, in, out)
}

func (h *dCHandler) InitDC(ctx context.Context, in *common.DataCenter, out *common.AppResponce) error {
	return h.DCHandler.InitDC(ctx, in, out)
}

func (h *dCHandler) Overview(ctx context.Context, in *common.Empty, out *common.DataCenterStatus) error {
	return h.DCHandler.Overview(ctx, in, out)
}

// Client API for DCAPI service

type DCAPIService interface {
	DataCenterList(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*DataCenterListResponse, error)
	DataCenterLeaderBoard(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*DataCenterLeaderBoardResponse, error)
	NetworkInfo(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*NetworkInfoResponse, error)
	RegisterDataCenter(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*NetworkInfoResponse, error)
}

type dCAPIService struct {
	c    client.Client
	name string
}

func NewDCAPIService(name string, c client.Client) DCAPIService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dcmgr"
	}
	return &dCAPIService{
		c:    c,
		name: name,
	}
}

func (c *dCAPIService) DataCenterList(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*DataCenterListResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.DataCenterList", in)
	out := new(DataCenterListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIService) DataCenterLeaderBoard(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*DataCenterLeaderBoardResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.DataCenterLeaderBoard", in)
	out := new(DataCenterLeaderBoardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIService) NetworkInfo(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*NetworkInfoResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.NetworkInfo", in)
	out := new(NetworkInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIService) RegisterDataCenter(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*NetworkInfoResponse, error) {
	req := c.c.NewRequest(c.name, "DCAPI.RegisterDataCenter", in)
	out := new(NetworkInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DCAPI service

type DCAPIHandler interface {
	DataCenterList(context.Context, *common.Empty, *DataCenterListResponse) error
	DataCenterLeaderBoard(context.Context, *common.Empty, *DataCenterLeaderBoardResponse) error
	NetworkInfo(context.Context, *common.Empty, *NetworkInfoResponse) error
	RegisterDataCenter(context.Context, *common.Empty, *NetworkInfoResponse) error
}

func RegisterDCAPIHandler(s server.Server, hdlr DCAPIHandler, opts ...server.HandlerOption) error {
	type dCAPI interface {
		DataCenterList(ctx context.Context, in *common.Empty, out *DataCenterListResponse) error
		DataCenterLeaderBoard(ctx context.Context, in *common.Empty, out *DataCenterLeaderBoardResponse) error
		NetworkInfo(ctx context.Context, in *common.Empty, out *NetworkInfoResponse) error
		RegisterDataCenter(ctx context.Context, in *common.Empty, out *NetworkInfoResponse) error
	}
	type DCAPI struct {
		dCAPI
	}
	h := &dCAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&DCAPI{h}, opts...))
}

type dCAPIHandler struct {
	DCAPIHandler
}

func (h *dCAPIHandler) DataCenterList(ctx context.Context, in *common.Empty, out *DataCenterListResponse) error {
	return h.DCAPIHandler.DataCenterList(ctx, in, out)
}

func (h *dCAPIHandler) DataCenterLeaderBoard(ctx context.Context, in *common.Empty, out *DataCenterLeaderBoardResponse) error {
	return h.DCAPIHandler.DataCenterLeaderBoard(ctx, in, out)
}

func (h *dCAPIHandler) NetworkInfo(ctx context.Context, in *common.Empty, out *NetworkInfoResponse) error {
	return h.DCAPIHandler.NetworkInfo(ctx, in, out)
}

func (h *dCAPIHandler) RegisterDataCenter(ctx context.Context, in *common.Empty, out *NetworkInfoResponse) error {
	return h.DCAPIHandler.RegisterDataCenter(ctx, in, out)
}

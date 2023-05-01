// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             (unknown)
// source: ticketing/api/show/v1/show_app.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationShowAppServiceAppGetShow = "/ticketing.api.show.v1.ShowAppService/AppGetShow"
const OperationShowAppServiceAppListShow = "/ticketing.api.show.v1.ShowAppService/AppListShow"
const OperationShowAppServicePlaceShowOrder = "/ticketing.api.show.v1.ShowAppService/PlaceShowOrder"

type ShowAppServiceHTTPServer interface {
	AppGetShow(context.Context, *GetShowRequest) (*Show, error)
	AppListShow(context.Context, *ListShowRequest) (*ListShowReply, error)
	PlaceShowOrder(context.Context, *PlaceShowOrderRequest) (*PlaceShowOrderReply, error)
}

func RegisterShowAppServiceHTTPServer(s *http.Server, srv ShowAppServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/ticketing/app/show/list", _ShowAppService_AppListShow0_HTTP_Handler(srv))
	r.GET("/v1/ticketing/app/show", _ShowAppService_AppListShow1_HTTP_Handler(srv))
	r.GET("/v1/ticketing/app/show/{id}", _ShowAppService_AppGetShow0_HTTP_Handler(srv))
	r.POST("/v1/ticketing/app/show/order", _ShowAppService_PlaceShowOrder0_HTTP_Handler(srv))
}

func _ShowAppService_AppListShow0_HTTP_Handler(srv ShowAppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListShowRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShowAppServiceAppListShow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppListShow(ctx, req.(*ListShowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListShowReply)
		return ctx.Result(200, reply)
	}
}

func _ShowAppService_AppListShow1_HTTP_Handler(srv ShowAppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListShowRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShowAppServiceAppListShow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppListShow(ctx, req.(*ListShowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListShowReply)
		return ctx.Result(200, reply)
	}
}

func _ShowAppService_AppGetShow0_HTTP_Handler(srv ShowAppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetShowRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShowAppServiceAppGetShow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppGetShow(ctx, req.(*GetShowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Show)
		return ctx.Result(200, reply)
	}
}

func _ShowAppService_PlaceShowOrder0_HTTP_Handler(srv ShowAppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlaceShowOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShowAppServicePlaceShowOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PlaceShowOrder(ctx, req.(*PlaceShowOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlaceShowOrderReply)
		return ctx.Result(200, reply)
	}
}

type ShowAppServiceHTTPClient interface {
	AppGetShow(ctx context.Context, req *GetShowRequest, opts ...http.CallOption) (rsp *Show, err error)
	AppListShow(ctx context.Context, req *ListShowRequest, opts ...http.CallOption) (rsp *ListShowReply, err error)
	PlaceShowOrder(ctx context.Context, req *PlaceShowOrderRequest, opts ...http.CallOption) (rsp *PlaceShowOrderReply, err error)
}

type ShowAppServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewShowAppServiceHTTPClient(client *http.Client) ShowAppServiceHTTPClient {
	return &ShowAppServiceHTTPClientImpl{client}
}

func (c *ShowAppServiceHTTPClientImpl) AppGetShow(ctx context.Context, in *GetShowRequest, opts ...http.CallOption) (*Show, error) {
	var out Show
	pattern := "/v1/ticketing/app/show/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShowAppServiceAppGetShow))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShowAppServiceHTTPClientImpl) AppListShow(ctx context.Context, in *ListShowRequest, opts ...http.CallOption) (*ListShowReply, error) {
	var out ListShowReply
	pattern := "/v1/ticketing/app/show"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShowAppServiceAppListShow))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShowAppServiceHTTPClientImpl) PlaceShowOrder(ctx context.Context, in *PlaceShowOrderRequest, opts ...http.CallOption) (*PlaceShowOrderReply, error) {
	var out PlaceShowOrderReply
	pattern := "/v1/ticketing/app/show/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShowAppServicePlaceShowOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
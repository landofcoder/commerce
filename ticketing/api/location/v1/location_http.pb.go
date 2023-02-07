// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             (unknown)
// source: ticketing/api/location/v1/location.proto

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

const OperationLocationServiceCreateLocation = "/ticketing.api.location.v1.LocationService/CreateLocation"
const OperationLocationServiceDeleteLocation = "/ticketing.api.location.v1.LocationService/DeleteLocation"
const OperationLocationServiceGetLocation = "/ticketing.api.location.v1.LocationService/GetLocation"
const OperationLocationServiceListLocation = "/ticketing.api.location.v1.LocationService/ListLocation"
const OperationLocationServiceUpdateLocation = "/ticketing.api.location.v1.LocationService/UpdateLocation"

type LocationServiceHTTPServer interface {
	CreateLocation(context.Context, *CreateLocationRequest) (*Location, error)
	DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationReply, error)
	GetLocation(context.Context, *GetLocationRequest) (*Location, error)
	ListLocation(context.Context, *ListLocationRequest) (*ListLocationReply, error)
	UpdateLocation(context.Context, *UpdateLocationRequest) (*Location, error)
}

func RegisterLocationServiceHTTPServer(s *http.Server, srv LocationServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/ticketing/location/list", _LocationService_ListLocation0_HTTP_Handler(srv))
	r.GET("/v1/ticketing/locations", _LocationService_ListLocation1_HTTP_Handler(srv))
	r.GET("/v1/ticketing/location/{id}", _LocationService_GetLocation0_HTTP_Handler(srv))
	r.POST("/v1/ticketing/location", _LocationService_CreateLocation0_HTTP_Handler(srv))
	r.PATCH("/v1/ticketing/location/{location.id}", _LocationService_UpdateLocation0_HTTP_Handler(srv))
	r.PUT("/v1/ticketing/location/{location.id}", _LocationService_UpdateLocation1_HTTP_Handler(srv))
	r.DELETE("/v1/ticketing/location/{id}", _LocationService_DeleteLocation0_HTTP_Handler(srv))
}

func _LocationService_ListLocation0_HTTP_Handler(srv LocationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListLocationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLocationServiceListLocation)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListLocation(ctx, req.(*ListLocationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListLocationReply)
		return ctx.Result(200, reply)
	}
}

func _LocationService_ListLocation1_HTTP_Handler(srv LocationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListLocationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLocationServiceListLocation)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListLocation(ctx, req.(*ListLocationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListLocationReply)
		return ctx.Result(200, reply)
	}
}

func _LocationService_GetLocation0_HTTP_Handler(srv LocationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLocationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLocationServiceGetLocation)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLocation(ctx, req.(*GetLocationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Location)
		return ctx.Result(200, reply)
	}
}

func _LocationService_CreateLocation0_HTTP_Handler(srv LocationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateLocationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLocationServiceCreateLocation)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateLocation(ctx, req.(*CreateLocationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Location)
		return ctx.Result(200, reply)
	}
}

func _LocationService_UpdateLocation0_HTTP_Handler(srv LocationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateLocationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLocationServiceUpdateLocation)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateLocation(ctx, req.(*UpdateLocationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Location)
		return ctx.Result(200, reply)
	}
}

func _LocationService_UpdateLocation1_HTTP_Handler(srv LocationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateLocationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLocationServiceUpdateLocation)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateLocation(ctx, req.(*UpdateLocationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Location)
		return ctx.Result(200, reply)
	}
}

func _LocationService_DeleteLocation0_HTTP_Handler(srv LocationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteLocationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLocationServiceDeleteLocation)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteLocation(ctx, req.(*DeleteLocationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteLocationReply)
		return ctx.Result(200, reply)
	}
}

type LocationServiceHTTPClient interface {
	CreateLocation(ctx context.Context, req *CreateLocationRequest, opts ...http.CallOption) (rsp *Location, err error)
	DeleteLocation(ctx context.Context, req *DeleteLocationRequest, opts ...http.CallOption) (rsp *DeleteLocationReply, err error)
	GetLocation(ctx context.Context, req *GetLocationRequest, opts ...http.CallOption) (rsp *Location, err error)
	ListLocation(ctx context.Context, req *ListLocationRequest, opts ...http.CallOption) (rsp *ListLocationReply, err error)
	UpdateLocation(ctx context.Context, req *UpdateLocationRequest, opts ...http.CallOption) (rsp *Location, err error)
}

type LocationServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewLocationServiceHTTPClient(client *http.Client) LocationServiceHTTPClient {
	return &LocationServiceHTTPClientImpl{client}
}

func (c *LocationServiceHTTPClientImpl) CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...http.CallOption) (*Location, error) {
	var out Location
	pattern := "/v1/ticketing/location"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLocationServiceCreateLocation))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LocationServiceHTTPClientImpl) DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...http.CallOption) (*DeleteLocationReply, error) {
	var out DeleteLocationReply
	pattern := "/v1/ticketing/location/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLocationServiceDeleteLocation))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LocationServiceHTTPClientImpl) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...http.CallOption) (*Location, error) {
	var out Location
	pattern := "/v1/ticketing/location/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLocationServiceGetLocation))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LocationServiceHTTPClientImpl) ListLocation(ctx context.Context, in *ListLocationRequest, opts ...http.CallOption) (*ListLocationReply, error) {
	var out ListLocationReply
	pattern := "/v1/ticketing/locations"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLocationServiceListLocation))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LocationServiceHTTPClientImpl) UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...http.CallOption) (*Location, error) {
	var out Location
	pattern := "/v1/ticketing/location/{location.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLocationServiceUpdateLocation))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

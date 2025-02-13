// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: order/api/order/v1/order_internal.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OrderInternalService_CreateInternalOrder_FullMethodName     = "/order.api.order.v1.OrderInternalService/CreateInternalOrder"
	OrderInternalService_GetInternalOrder_FullMethodName        = "/order.api.order.v1.OrderInternalService/GetInternalOrder"
	OrderInternalService_InternalOrderPaySuccess_FullMethodName = "/order.api.order.v1.OrderInternalService/InternalOrderPaySuccess"
)

// OrderInternalServiceClient is the client API for OrderInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderInternalServiceClient interface {
	CreateInternalOrder(ctx context.Context, in *CreateInternalOrderRequest, opts ...grpc.CallOption) (*Order, error)
	GetInternalOrder(ctx context.Context, in *GetInternalOrderRequest, opts ...grpc.CallOption) (*Order, error)
	InternalOrderPaySuccess(ctx context.Context, in *InternalOrderPaySuccessRequest, opts ...grpc.CallOption) (*InternalOrderPaySuccessReply, error)
}

type orderInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderInternalServiceClient(cc grpc.ClientConnInterface) OrderInternalServiceClient {
	return &orderInternalServiceClient{cc}
}

func (c *orderInternalServiceClient) CreateInternalOrder(ctx context.Context, in *CreateInternalOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderInternalService_CreateInternalOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderInternalServiceClient) GetInternalOrder(ctx context.Context, in *GetInternalOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderInternalService_GetInternalOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderInternalServiceClient) InternalOrderPaySuccess(ctx context.Context, in *InternalOrderPaySuccessRequest, opts ...grpc.CallOption) (*InternalOrderPaySuccessReply, error) {
	out := new(InternalOrderPaySuccessReply)
	err := c.cc.Invoke(ctx, OrderInternalService_InternalOrderPaySuccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderInternalServiceServer is the server API for OrderInternalService service.
// All implementations should embed UnimplementedOrderInternalServiceServer
// for forward compatibility
type OrderInternalServiceServer interface {
	CreateInternalOrder(context.Context, *CreateInternalOrderRequest) (*Order, error)
	GetInternalOrder(context.Context, *GetInternalOrderRequest) (*Order, error)
	InternalOrderPaySuccess(context.Context, *InternalOrderPaySuccessRequest) (*InternalOrderPaySuccessReply, error)
}

// UnimplementedOrderInternalServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderInternalServiceServer struct {
}

func (UnimplementedOrderInternalServiceServer) CreateInternalOrder(context.Context, *CreateInternalOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInternalOrder not implemented")
}
func (UnimplementedOrderInternalServiceServer) GetInternalOrder(context.Context, *GetInternalOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInternalOrder not implemented")
}
func (UnimplementedOrderInternalServiceServer) InternalOrderPaySuccess(context.Context, *InternalOrderPaySuccessRequest) (*InternalOrderPaySuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InternalOrderPaySuccess not implemented")
}

// UnsafeOrderInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderInternalServiceServer will
// result in compilation errors.
type UnsafeOrderInternalServiceServer interface {
	mustEmbedUnimplementedOrderInternalServiceServer()
}

func RegisterOrderInternalServiceServer(s grpc.ServiceRegistrar, srv OrderInternalServiceServer) {
	s.RegisterService(&OrderInternalService_ServiceDesc, srv)
}

func _OrderInternalService_CreateInternalOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInternalOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderInternalServiceServer).CreateInternalOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderInternalService_CreateInternalOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderInternalServiceServer).CreateInternalOrder(ctx, req.(*CreateInternalOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderInternalService_GetInternalOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInternalOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderInternalServiceServer).GetInternalOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderInternalService_GetInternalOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderInternalServiceServer).GetInternalOrder(ctx, req.(*GetInternalOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderInternalService_InternalOrderPaySuccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalOrderPaySuccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderInternalServiceServer).InternalOrderPaySuccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderInternalService_InternalOrderPaySuccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderInternalServiceServer).InternalOrderPaySuccess(ctx, req.(*InternalOrderPaySuccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderInternalService_ServiceDesc is the grpc.ServiceDesc for OrderInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.api.order.v1.OrderInternalService",
	HandlerType: (*OrderInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInternalOrder",
			Handler:    _OrderInternalService_CreateInternalOrder_Handler,
		},
		{
			MethodName: "GetInternalOrder",
			Handler:    _OrderInternalService_GetInternalOrder_Handler,
		},
		{
			MethodName: "InternalOrderPaySuccess",
			Handler:    _OrderInternalService_InternalOrderPaySuccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/api/order/v1/order_internal.proto",
}

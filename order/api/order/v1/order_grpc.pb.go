// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: order/api/order/v1/order.proto

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
	OrderService_ListOrder_FullMethodName   = "/order.api.order.v1.OrderService/ListOrder"
	OrderService_GetOrder_FullMethodName    = "/order.api.order.v1.OrderService/GetOrder"
	OrderService_CreateOrder_FullMethodName = "/order.api.order.v1.OrderService/CreateOrder"
	OrderService_UpdateOrder_FullMethodName = "/order.api.order.v1.OrderService/UpdateOrder"
	OrderService_DeleteOrder_FullMethodName = "/order.api.order.v1.OrderService/DeleteOrder"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderReply, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderReply, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderReply, error) {
	out := new(ListOrderReply)
	err := c.cc.Invoke(ctx, OrderService_ListOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderReply, error) {
	out := new(DeleteOrderReply)
	err := c.cc.Invoke(ctx, OrderService_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderReply, error)
	GetOrder(context.Context, *GetOrderRequest) (*Order, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*Order, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error)
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) ListOrder(context.Context, *ListOrderRequest) (*ListOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *GetOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServiceServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrder(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.api.order.v1.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOrder",
			Handler:    _OrderService_ListOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderService_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/api/order/v1/order.proto",
}

const (
	OrderAppService_ListAppOrder_FullMethodName   = "/order.api.order.v1.OrderAppService/ListAppOrder"
	OrderAppService_GetAppOrder_FullMethodName    = "/order.api.order.v1.OrderAppService/GetAppOrder"
	OrderAppService_RefundAppOrder_FullMethodName = "/order.api.order.v1.OrderAppService/RefundAppOrder"
)

// OrderAppServiceClient is the client API for OrderAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderAppServiceClient interface {
	ListAppOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderReply, error)
	GetAppOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error)
	RefundAppOrder(ctx context.Context, in *RefundAppOrderRequest, opts ...grpc.CallOption) (*Order, error)
}

type orderAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderAppServiceClient(cc grpc.ClientConnInterface) OrderAppServiceClient {
	return &orderAppServiceClient{cc}
}

func (c *orderAppServiceClient) ListAppOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderReply, error) {
	out := new(ListOrderReply)
	err := c.cc.Invoke(ctx, OrderAppService_ListAppOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAppServiceClient) GetAppOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderAppService_GetAppOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAppServiceClient) RefundAppOrder(ctx context.Context, in *RefundAppOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderAppService_RefundAppOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderAppServiceServer is the server API for OrderAppService service.
// All implementations should embed UnimplementedOrderAppServiceServer
// for forward compatibility
type OrderAppServiceServer interface {
	ListAppOrder(context.Context, *ListOrderRequest) (*ListOrderReply, error)
	GetAppOrder(context.Context, *GetOrderRequest) (*Order, error)
	RefundAppOrder(context.Context, *RefundAppOrderRequest) (*Order, error)
}

// UnimplementedOrderAppServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderAppServiceServer struct {
}

func (UnimplementedOrderAppServiceServer) ListAppOrder(context.Context, *ListOrderRequest) (*ListOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppOrder not implemented")
}
func (UnimplementedOrderAppServiceServer) GetAppOrder(context.Context, *GetOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppOrder not implemented")
}
func (UnimplementedOrderAppServiceServer) RefundAppOrder(context.Context, *RefundAppOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefundAppOrder not implemented")
}

// UnsafeOrderAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderAppServiceServer will
// result in compilation errors.
type UnsafeOrderAppServiceServer interface {
	mustEmbedUnimplementedOrderAppServiceServer()
}

func RegisterOrderAppServiceServer(s grpc.ServiceRegistrar, srv OrderAppServiceServer) {
	s.RegisterService(&OrderAppService_ServiceDesc, srv)
}

func _OrderAppService_ListAppOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAppServiceServer).ListAppOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAppService_ListAppOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAppServiceServer).ListAppOrder(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAppService_GetAppOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAppServiceServer).GetAppOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAppService_GetAppOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAppServiceServer).GetAppOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAppService_RefundAppOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundAppOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAppServiceServer).RefundAppOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAppService_RefundAppOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAppServiceServer).RefundAppOrder(ctx, req.(*RefundAppOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderAppService_ServiceDesc is the grpc.ServiceDesc for OrderAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.api.order.v1.OrderAppService",
	HandlerType: (*OrderAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAppOrder",
			Handler:    _OrderAppService_ListAppOrder_Handler,
		},
		{
			MethodName: "GetAppOrder",
			Handler:    _OrderAppService_GetAppOrder_Handler,
		},
		{
			MethodName: "RefundAppOrder",
			Handler:    _OrderAppService_RefundAppOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/api/order/v1/order.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ticketing/api/category/v1/category.proto

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
	TicketingCategoryService_ListCategory_FullMethodName   = "/ticketing.api.category.v1.TicketingCategoryService/ListCategory"
	TicketingCategoryService_GetCategory_FullMethodName    = "/ticketing.api.category.v1.TicketingCategoryService/GetCategory"
	TicketingCategoryService_CreateCategory_FullMethodName = "/ticketing.api.category.v1.TicketingCategoryService/CreateCategory"
	TicketingCategoryService_UpdateCategory_FullMethodName = "/ticketing.api.category.v1.TicketingCategoryService/UpdateCategory"
	TicketingCategoryService_DeleteCategory_FullMethodName = "/ticketing.api.category.v1.TicketingCategoryService/DeleteCategory"
)

// TicketingCategoryServiceClient is the client API for TicketingCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketingCategoryServiceClient interface {
	ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryReply, error)
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryReply, error)
}

type ticketingCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketingCategoryServiceClient(cc grpc.ClientConnInterface) TicketingCategoryServiceClient {
	return &ticketingCategoryServiceClient{cc}
}

func (c *ticketingCategoryServiceClient) ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryReply, error) {
	out := new(ListCategoryReply)
	err := c.cc.Invoke(ctx, TicketingCategoryService_ListCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingCategoryServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, TicketingCategoryService_GetCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingCategoryServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, TicketingCategoryService_CreateCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingCategoryServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, TicketingCategoryService_UpdateCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketingCategoryServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryReply, error) {
	out := new(DeleteCategoryReply)
	err := c.cc.Invoke(ctx, TicketingCategoryService_DeleteCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketingCategoryServiceServer is the server API for TicketingCategoryService service.
// All implementations should embed UnimplementedTicketingCategoryServiceServer
// for forward compatibility
type TicketingCategoryServiceServer interface {
	ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryReply, error)
	GetCategory(context.Context, *GetCategoryRequest) (*Category, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*Category, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryReply, error)
}

// UnimplementedTicketingCategoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTicketingCategoryServiceServer struct {
}

func (UnimplementedTicketingCategoryServiceServer) ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategory not implemented")
}
func (UnimplementedTicketingCategoryServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedTicketingCategoryServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedTicketingCategoryServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedTicketingCategoryServiceServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}

// UnsafeTicketingCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketingCategoryServiceServer will
// result in compilation errors.
type UnsafeTicketingCategoryServiceServer interface {
	mustEmbedUnimplementedTicketingCategoryServiceServer()
}

func RegisterTicketingCategoryServiceServer(s grpc.ServiceRegistrar, srv TicketingCategoryServiceServer) {
	s.RegisterService(&TicketingCategoryService_ServiceDesc, srv)
}

func _TicketingCategoryService_ListCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingCategoryServiceServer).ListCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingCategoryService_ListCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingCategoryServiceServer).ListCategory(ctx, req.(*ListCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketingCategoryService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingCategoryServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingCategoryService_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingCategoryServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketingCategoryService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingCategoryServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingCategoryService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingCategoryServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketingCategoryService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingCategoryServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingCategoryService_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingCategoryServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketingCategoryService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingCategoryServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingCategoryService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingCategoryServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketingCategoryService_ServiceDesc is the grpc.ServiceDesc for TicketingCategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketingCategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticketing.api.category.v1.TicketingCategoryService",
	HandlerType: (*TicketingCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCategory",
			Handler:    _TicketingCategoryService_ListCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _TicketingCategoryService_GetCategory_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _TicketingCategoryService_CreateCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _TicketingCategoryService_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _TicketingCategoryService_DeleteCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticketing/api/category/v1/category.proto",
}

const (
	TicketingCategoryAppService_ListAppCategory_FullMethodName = "/ticketing.api.category.v1.TicketingCategoryAppService/ListAppCategory"
)

// TicketingCategoryAppServiceClient is the client API for TicketingCategoryAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketingCategoryAppServiceClient interface {
	ListAppCategory(ctx context.Context, in *ListAppCategoryRequest, opts ...grpc.CallOption) (*ListAppCategoryReply, error)
}

type ticketingCategoryAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketingCategoryAppServiceClient(cc grpc.ClientConnInterface) TicketingCategoryAppServiceClient {
	return &ticketingCategoryAppServiceClient{cc}
}

func (c *ticketingCategoryAppServiceClient) ListAppCategory(ctx context.Context, in *ListAppCategoryRequest, opts ...grpc.CallOption) (*ListAppCategoryReply, error) {
	out := new(ListAppCategoryReply)
	err := c.cc.Invoke(ctx, TicketingCategoryAppService_ListAppCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketingCategoryAppServiceServer is the server API for TicketingCategoryAppService service.
// All implementations should embed UnimplementedTicketingCategoryAppServiceServer
// for forward compatibility
type TicketingCategoryAppServiceServer interface {
	ListAppCategory(context.Context, *ListAppCategoryRequest) (*ListAppCategoryReply, error)
}

// UnimplementedTicketingCategoryAppServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTicketingCategoryAppServiceServer struct {
}

func (UnimplementedTicketingCategoryAppServiceServer) ListAppCategory(context.Context, *ListAppCategoryRequest) (*ListAppCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppCategory not implemented")
}

// UnsafeTicketingCategoryAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketingCategoryAppServiceServer will
// result in compilation errors.
type UnsafeTicketingCategoryAppServiceServer interface {
	mustEmbedUnimplementedTicketingCategoryAppServiceServer()
}

func RegisterTicketingCategoryAppServiceServer(s grpc.ServiceRegistrar, srv TicketingCategoryAppServiceServer) {
	s.RegisterService(&TicketingCategoryAppService_ServiceDesc, srv)
}

func _TicketingCategoryAppService_ListAppCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAppCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketingCategoryAppServiceServer).ListAppCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketingCategoryAppService_ListAppCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketingCategoryAppServiceServer).ListAppCategory(ctx, req.(*ListAppCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketingCategoryAppService_ServiceDesc is the grpc.ServiceDesc for TicketingCategoryAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketingCategoryAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticketing.api.category.v1.TicketingCategoryAppService",
	HandlerType: (*TicketingCategoryAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAppCategory",
			Handler:    _TicketingCategoryAppService_ListAppCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticketing/api/category/v1/category.proto",
}

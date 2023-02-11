// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: product/api/product/v1/brand.proto

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

// BrandServiceClient is the client API for BrandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrandServiceClient interface {
	ListBrand(ctx context.Context, in *ListBrandRequest, opts ...grpc.CallOption) (*ListBrandReply, error)
	GetBrand(ctx context.Context, in *GetBrandRequest, opts ...grpc.CallOption) (*Brand, error)
	CreateBrand(ctx context.Context, in *CreateBrandRequest, opts ...grpc.CallOption) (*Brand, error)
	UpdateBrand(ctx context.Context, in *UpdateBrandRequest, opts ...grpc.CallOption) (*Brand, error)
	DeleteBrand(ctx context.Context, in *DeleteBrandRequest, opts ...grpc.CallOption) (*DeleteBrandReply, error)
}

type brandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrandServiceClient(cc grpc.ClientConnInterface) BrandServiceClient {
	return &brandServiceClient{cc}
}

func (c *brandServiceClient) ListBrand(ctx context.Context, in *ListBrandRequest, opts ...grpc.CallOption) (*ListBrandReply, error) {
	out := new(ListBrandReply)
	err := c.cc.Invoke(ctx, "/product.api.product.v1.BrandService/ListBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) GetBrand(ctx context.Context, in *GetBrandRequest, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/product.api.product.v1.BrandService/GetBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) CreateBrand(ctx context.Context, in *CreateBrandRequest, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/product.api.product.v1.BrandService/CreateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) UpdateBrand(ctx context.Context, in *UpdateBrandRequest, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/product.api.product.v1.BrandService/UpdateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) DeleteBrand(ctx context.Context, in *DeleteBrandRequest, opts ...grpc.CallOption) (*DeleteBrandReply, error) {
	out := new(DeleteBrandReply)
	err := c.cc.Invoke(ctx, "/product.api.product.v1.BrandService/DeleteBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrandServiceServer is the server API for BrandService service.
// All implementations should embed UnimplementedBrandServiceServer
// for forward compatibility
type BrandServiceServer interface {
	ListBrand(context.Context, *ListBrandRequest) (*ListBrandReply, error)
	GetBrand(context.Context, *GetBrandRequest) (*Brand, error)
	CreateBrand(context.Context, *CreateBrandRequest) (*Brand, error)
	UpdateBrand(context.Context, *UpdateBrandRequest) (*Brand, error)
	DeleteBrand(context.Context, *DeleteBrandRequest) (*DeleteBrandReply, error)
}

// UnimplementedBrandServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBrandServiceServer struct {
}

func (UnimplementedBrandServiceServer) ListBrand(context.Context, *ListBrandRequest) (*ListBrandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBrand not implemented")
}
func (UnimplementedBrandServiceServer) GetBrand(context.Context, *GetBrandRequest) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrand not implemented")
}
func (UnimplementedBrandServiceServer) CreateBrand(context.Context, *CreateBrandRequest) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (UnimplementedBrandServiceServer) UpdateBrand(context.Context, *UpdateBrandRequest) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (UnimplementedBrandServiceServer) DeleteBrand(context.Context, *DeleteBrandRequest) (*DeleteBrandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}

// UnsafeBrandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrandServiceServer will
// result in compilation errors.
type UnsafeBrandServiceServer interface {
	mustEmbedUnimplementedBrandServiceServer()
}

func RegisterBrandServiceServer(s grpc.ServiceRegistrar, srv BrandServiceServer) {
	s.RegisterService(&BrandService_ServiceDesc, srv)
}

func _BrandService_ListBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).ListBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.api.product.v1.BrandService/ListBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).ListBrand(ctx, req.(*ListBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_GetBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).GetBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.api.product.v1.BrandService/GetBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).GetBrand(ctx, req.(*GetBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_CreateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).CreateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.api.product.v1.BrandService/CreateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).CreateBrand(ctx, req.(*CreateBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_UpdateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).UpdateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.api.product.v1.BrandService/UpdateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).UpdateBrand(ctx, req.(*UpdateBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.api.product.v1.BrandService/DeleteBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).DeleteBrand(ctx, req.(*DeleteBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BrandService_ServiceDesc is the grpc.ServiceDesc for BrandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.api.product.v1.BrandService",
	HandlerType: (*BrandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBrand",
			Handler:    _BrandService_ListBrand_Handler,
		},
		{
			MethodName: "GetBrand",
			Handler:    _BrandService_GetBrand_Handler,
		},
		{
			MethodName: "CreateBrand",
			Handler:    _BrandService_CreateBrand_Handler,
		},
		{
			MethodName: "UpdateBrand",
			Handler:    _BrandService_UpdateBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _BrandService_DeleteBrand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/api/product/v1/brand.proto",
}
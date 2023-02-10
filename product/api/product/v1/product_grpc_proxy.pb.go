// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: product/api/product/v1/product.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ ProductServiceServer = (*productServiceClientProxy)(nil)

const GrpcOperationProductServiceListProduct = "/product.api.product.v1.ProductService/ListProduct"
const GrpcOperationProductServiceGetProduct = "/product.api.product.v1.ProductService/GetProduct"
const GrpcOperationProductServiceCreateProduct = "/product.api.product.v1.ProductService/CreateProduct"
const GrpcOperationProductServiceUpdateProduct = "/product.api.product.v1.ProductService/UpdateProduct"
const GrpcOperationProductServiceDeleteProduct = "/product.api.product.v1.ProductService/DeleteProduct"

// productServiceClientProxy is the proxy to turn ProductService client to server interface.
type productServiceClientProxy struct {
	cc ProductServiceClient
}

func NewProductServiceClientProxy(cc ProductServiceClient) ProductServiceServer {
	return &productServiceClientProxy{cc}
}

func (c *productServiceClientProxy) ListProduct(ctx context.Context, in *ListProductRequest) (*ListProductReply, error) {
	return c.cc.ListProduct(ctx, in)
}
func (c *productServiceClientProxy) GetProduct(ctx context.Context, in *GetProductRequest) (*Product, error) {
	return c.cc.GetProduct(ctx, in)
}
func (c *productServiceClientProxy) CreateProduct(ctx context.Context, in *CreateProductRequest) (*Product, error) {
	return c.cc.CreateProduct(ctx, in)
}
func (c *productServiceClientProxy) UpdateProduct(ctx context.Context, in *UpdateProductRequest) (*Product, error) {
	return c.cc.UpdateProduct(ctx, in)
}
func (c *productServiceClientProxy) DeleteProduct(ctx context.Context, in *DeleteProductRequest) (*DeleteProductReply, error) {
	return c.cc.DeleteProduct(ctx, in)
}

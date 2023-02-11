// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: product/api/category/v1/category.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ ProductCategoryServiceServer = (*productCategoryServiceClientProxy)(nil)

const GrpcOperationProductCategoryServiceListCategory = "/product.api.category.v1.ProductCategoryService/ListCategory"
const GrpcOperationProductCategoryServiceGetCategory = "/product.api.category.v1.ProductCategoryService/GetCategory"
const GrpcOperationProductCategoryServiceCreateCategory = "/product.api.category.v1.ProductCategoryService/CreateCategory"
const GrpcOperationProductCategoryServiceUpdateCategory = "/product.api.category.v1.ProductCategoryService/UpdateCategory"
const GrpcOperationProductCategoryServiceDeleteCategory = "/product.api.category.v1.ProductCategoryService/DeleteCategory"

// productCategoryServiceClientProxy is the proxy to turn ProductCategoryService client to server interface.
type productCategoryServiceClientProxy struct {
	cc ProductCategoryServiceClient
}

func NewProductCategoryServiceClientProxy(cc ProductCategoryServiceClient) ProductCategoryServiceServer {
	return &productCategoryServiceClientProxy{cc}
}

func (c *productCategoryServiceClientProxy) ListCategory(ctx context.Context, in *ListCategoryRequest) (*ListCategoryReply, error) {
	return c.cc.ListCategory(ctx, in)
}
func (c *productCategoryServiceClientProxy) GetCategory(ctx context.Context, in *GetCategoryRequest) (*Category, error) {
	return c.cc.GetCategory(ctx, in)
}
func (c *productCategoryServiceClientProxy) CreateCategory(ctx context.Context, in *CreateCategoryRequest) (*Category, error) {
	return c.cc.CreateCategory(ctx, in)
}
func (c *productCategoryServiceClientProxy) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest) (*Category, error) {
	return c.cc.UpdateCategory(ctx, in)
}
func (c *productCategoryServiceClientProxy) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest) (*DeleteCategoryReply, error) {
	return c.cc.DeleteCategory(ctx, in)
}
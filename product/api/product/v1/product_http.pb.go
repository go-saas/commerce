// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             (unknown)
// source: product/api/product/v1/product.proto

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

const OperationProductServiceCreateProduct = "/product.api.product.v1.ProductService/CreateProduct"
const OperationProductServiceDeleteProduct = "/product.api.product.v1.ProductService/DeleteProduct"
const OperationProductServiceGetProduct = "/product.api.product.v1.ProductService/GetProduct"
const OperationProductServiceListProduct = "/product.api.product.v1.ProductService/ListProduct"
const OperationProductServiceUpdateProduct = "/product.api.product.v1.ProductService/UpdateProduct"

type ProductServiceHTTPServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*Product, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductReply, error)
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	ListProduct(context.Context, *ListProductRequest) (*ListProductReply, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*Product, error)
}

func RegisterProductServiceHTTPServer(s *http.Server, srv ProductServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/product/product/list", _ProductService_ListProduct0_HTTP_Handler(srv))
	r.GET("/v1/product/product", _ProductService_ListProduct1_HTTP_Handler(srv))
	r.GET("/v1/product/product/{id}", _ProductService_GetProduct0_HTTP_Handler(srv))
	r.POST("/v1/product/product", _ProductService_CreateProduct0_HTTP_Handler(srv))
	r.PATCH("/v1/product/product/{product.id}", _ProductService_UpdateProduct0_HTTP_Handler(srv))
	r.PUT("/v1/product/product/{product.id}", _ProductService_UpdateProduct1_HTTP_Handler(srv))
	r.DELETE("/v1/product/product/{id}", _ProductService_DeleteProduct0_HTTP_Handler(srv))
}

func _ProductService_ListProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListProductRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceListProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListProduct(ctx, req.(*ListProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListProductReply)
		return ctx.Result(200, reply)
	}
}

func _ProductService_ListProduct1_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListProductRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceListProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListProduct(ctx, req.(*ListProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListProductReply)
		return ctx.Result(200, reply)
	}
}

func _ProductService_GetProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProductRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceGetProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProduct(ctx, req.(*GetProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Product)
		return ctx.Result(200, reply)
	}
}

func _ProductService_CreateProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateProductRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceCreateProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateProduct(ctx, req.(*CreateProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Product)
		return ctx.Result(200, reply)
	}
}

func _ProductService_UpdateProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProductRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceUpdateProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProduct(ctx, req.(*UpdateProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Product)
		return ctx.Result(200, reply)
	}
}

func _ProductService_UpdateProduct1_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProductRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceUpdateProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProduct(ctx, req.(*UpdateProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Product)
		return ctx.Result(200, reply)
	}
}

func _ProductService_DeleteProduct0_HTTP_Handler(srv ProductServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteProductRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductServiceDeleteProduct)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProduct(ctx, req.(*DeleteProductRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteProductReply)
		return ctx.Result(200, reply)
	}
}

type ProductServiceHTTPClient interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest, opts ...http.CallOption) (rsp *Product, err error)
	DeleteProduct(ctx context.Context, req *DeleteProductRequest, opts ...http.CallOption) (rsp *DeleteProductReply, err error)
	GetProduct(ctx context.Context, req *GetProductRequest, opts ...http.CallOption) (rsp *Product, err error)
	ListProduct(ctx context.Context, req *ListProductRequest, opts ...http.CallOption) (rsp *ListProductReply, err error)
	UpdateProduct(ctx context.Context, req *UpdateProductRequest, opts ...http.CallOption) (rsp *Product, err error)
}

type ProductServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewProductServiceHTTPClient(client *http.Client) ProductServiceHTTPClient {
	return &ProductServiceHTTPClientImpl{client}
}

func (c *ProductServiceHTTPClientImpl) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...http.CallOption) (*Product, error) {
	var out Product
	pattern := "/v1/product/product"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductServiceCreateProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProductServiceHTTPClientImpl) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...http.CallOption) (*DeleteProductReply, error) {
	var out DeleteProductReply
	pattern := "/v1/product/product/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductServiceDeleteProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProductServiceHTTPClientImpl) GetProduct(ctx context.Context, in *GetProductRequest, opts ...http.CallOption) (*Product, error) {
	var out Product
	pattern := "/v1/product/product/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductServiceGetProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProductServiceHTTPClientImpl) ListProduct(ctx context.Context, in *ListProductRequest, opts ...http.CallOption) (*ListProductReply, error) {
	var out ListProductReply
	pattern := "/v1/product/product"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductServiceListProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProductServiceHTTPClientImpl) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...http.CallOption) (*Product, error) {
	var out Product
	pattern := "/v1/product/product/{product.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductServiceUpdateProduct))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             (unknown)
// source: ticketing/api/category/v1/category.proto

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

const OperationTicketingCategoryServiceCreateCategory = "/ticketing.api.category.v1.TicketingCategoryService/CreateCategory"
const OperationTicketingCategoryServiceDeleteCategory = "/ticketing.api.category.v1.TicketingCategoryService/DeleteCategory"
const OperationTicketingCategoryServiceGetCategory = "/ticketing.api.category.v1.TicketingCategoryService/GetCategory"
const OperationTicketingCategoryServiceListCategory = "/ticketing.api.category.v1.TicketingCategoryService/ListCategory"
const OperationTicketingCategoryServiceUpdateCategory = "/ticketing.api.category.v1.TicketingCategoryService/UpdateCategory"

type TicketingCategoryServiceHTTPServer interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryReply, error)
	GetCategory(context.Context, *GetCategoryRequest) (*Category, error)
	ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryReply, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*Category, error)
}

func RegisterTicketingCategoryServiceHTTPServer(s *http.Server, srv TicketingCategoryServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/ticketing/category/list", _TicketingCategoryService_ListCategory0_HTTP_Handler(srv))
	r.GET("/v1/ticketing/category", _TicketingCategoryService_ListCategory1_HTTP_Handler(srv))
	r.GET("/v1/ticketing/category/{key}", _TicketingCategoryService_GetCategory0_HTTP_Handler(srv))
	r.POST("/v1/ticketing/category", _TicketingCategoryService_CreateCategory0_HTTP_Handler(srv))
	r.PATCH("/v1/ticketing/category/{category.key}", _TicketingCategoryService_UpdateCategory0_HTTP_Handler(srv))
	r.PUT("/v1/ticketing/category/{category.key}", _TicketingCategoryService_UpdateCategory1_HTTP_Handler(srv))
	r.DELETE("/v1/ticketing/category/{key}", _TicketingCategoryService_DeleteCategory0_HTTP_Handler(srv))
}

func _TicketingCategoryService_ListCategory0_HTTP_Handler(srv TicketingCategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryServiceListCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*ListCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCategoryReply)
		return ctx.Result(200, reply)
	}
}

func _TicketingCategoryService_ListCategory1_HTTP_Handler(srv TicketingCategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryServiceListCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*ListCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCategoryReply)
		return ctx.Result(200, reply)
	}
}

func _TicketingCategoryService_GetCategory0_HTTP_Handler(srv TicketingCategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryServiceGetCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCategory(ctx, req.(*GetCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Category)
		return ctx.Result(200, reply)
	}
}

func _TicketingCategoryService_CreateCategory0_HTTP_Handler(srv TicketingCategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryServiceCreateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCategory(ctx, req.(*CreateCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Category)
		return ctx.Result(200, reply)
	}
}

func _TicketingCategoryService_UpdateCategory0_HTTP_Handler(srv TicketingCategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryServiceUpdateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCategory(ctx, req.(*UpdateCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Category)
		return ctx.Result(200, reply)
	}
}

func _TicketingCategoryService_UpdateCategory1_HTTP_Handler(srv TicketingCategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryServiceUpdateCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCategory(ctx, req.(*UpdateCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Category)
		return ctx.Result(200, reply)
	}
}

func _TicketingCategoryService_DeleteCategory0_HTTP_Handler(srv TicketingCategoryServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryServiceDeleteCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCategory(ctx, req.(*DeleteCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCategoryReply)
		return ctx.Result(200, reply)
	}
}

type TicketingCategoryServiceHTTPClient interface {
	CreateCategory(ctx context.Context, req *CreateCategoryRequest, opts ...http.CallOption) (rsp *Category, err error)
	DeleteCategory(ctx context.Context, req *DeleteCategoryRequest, opts ...http.CallOption) (rsp *DeleteCategoryReply, err error)
	GetCategory(ctx context.Context, req *GetCategoryRequest, opts ...http.CallOption) (rsp *Category, err error)
	ListCategory(ctx context.Context, req *ListCategoryRequest, opts ...http.CallOption) (rsp *ListCategoryReply, err error)
	UpdateCategory(ctx context.Context, req *UpdateCategoryRequest, opts ...http.CallOption) (rsp *Category, err error)
}

type TicketingCategoryServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTicketingCategoryServiceHTTPClient(client *http.Client) TicketingCategoryServiceHTTPClient {
	return &TicketingCategoryServiceHTTPClientImpl{client}
}

func (c *TicketingCategoryServiceHTTPClientImpl) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...http.CallOption) (*Category, error) {
	var out Category
	pattern := "/v1/ticketing/category"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTicketingCategoryServiceCreateCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TicketingCategoryServiceHTTPClientImpl) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...http.CallOption) (*DeleteCategoryReply, error) {
	var out DeleteCategoryReply
	pattern := "/v1/ticketing/category/{key}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTicketingCategoryServiceDeleteCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TicketingCategoryServiceHTTPClientImpl) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...http.CallOption) (*Category, error) {
	var out Category
	pattern := "/v1/ticketing/category/{key}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTicketingCategoryServiceGetCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TicketingCategoryServiceHTTPClientImpl) ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...http.CallOption) (*ListCategoryReply, error) {
	var out ListCategoryReply
	pattern := "/v1/ticketing/category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTicketingCategoryServiceListCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TicketingCategoryServiceHTTPClientImpl) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...http.CallOption) (*Category, error) {
	var out Category
	pattern := "/v1/ticketing/category/{category.key}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTicketingCategoryServiceUpdateCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationTicketingCategoryAppServiceListAppCategory = "/ticketing.api.category.v1.TicketingCategoryAppService/ListAppCategory"

type TicketingCategoryAppServiceHTTPServer interface {
	ListAppCategory(context.Context, *ListAppCategoryRequest) (*ListAppCategoryReply, error)
}

func RegisterTicketingCategoryAppServiceHTTPServer(s *http.Server, srv TicketingCategoryAppServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/ticketing/app/category/list", _TicketingCategoryAppService_ListAppCategory0_HTTP_Handler(srv))
	r.GET("/v1/ticketing/app/category", _TicketingCategoryAppService_ListAppCategory1_HTTP_Handler(srv))
}

func _TicketingCategoryAppService_ListAppCategory0_HTTP_Handler(srv TicketingCategoryAppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAppCategoryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryAppServiceListAppCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAppCategory(ctx, req.(*ListAppCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAppCategoryReply)
		return ctx.Result(200, reply)
	}
}

func _TicketingCategoryAppService_ListAppCategory1_HTTP_Handler(srv TicketingCategoryAppServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAppCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketingCategoryAppServiceListAppCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAppCategory(ctx, req.(*ListAppCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAppCategoryReply)
		return ctx.Result(200, reply)
	}
}

type TicketingCategoryAppServiceHTTPClient interface {
	ListAppCategory(ctx context.Context, req *ListAppCategoryRequest, opts ...http.CallOption) (rsp *ListAppCategoryReply, err error)
}

type TicketingCategoryAppServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTicketingCategoryAppServiceHTTPClient(client *http.Client) TicketingCategoryAppServiceHTTPClient {
	return &TicketingCategoryAppServiceHTTPClientImpl{client}
}

func (c *TicketingCategoryAppServiceHTTPClientImpl) ListAppCategory(ctx context.Context, in *ListAppCategoryRequest, opts ...http.CallOption) (*ListAppCategoryReply, error) {
	var out ListAppCategoryReply
	pattern := "/v1/ticketing/app/category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTicketingCategoryAppServiceListAppCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

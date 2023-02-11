// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             (unknown)
// source: order/api/order/v1/order.proto

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

const OperationOrderServiceCreateOrder = "/order.api.order.v1.OrderService/CreateOrder"
const OperationOrderServiceDeleteOrder = "/order.api.order.v1.OrderService/DeleteOrder"
const OperationOrderServiceGetOrder = "/order.api.order.v1.OrderService/GetOrder"
const OperationOrderServiceListOrder = "/order.api.order.v1.OrderService/ListOrder"
const OperationOrderServiceUpdateOrder = "/order.api.order.v1.OrderService/UpdateOrder"

type OrderServiceHTTPServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error)
	GetOrder(context.Context, *GetOrderRequest) (*Order, error)
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderReply, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*Order, error)
}

func RegisterOrderServiceHTTPServer(s *http.Server, srv OrderServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/order/order/list", _OrderService_ListOrder0_HTTP_Handler(srv))
	r.GET("/v1/order/order", _OrderService_ListOrder1_HTTP_Handler(srv))
	r.GET("/v1/order/order/{id}", _OrderService_GetOrder0_HTTP_Handler(srv))
	r.POST("/v1/order/order", _OrderService_CreateOrder0_HTTP_Handler(srv))
	r.PATCH("/v1/order/order/{order.id}", _OrderService_UpdateOrder0_HTTP_Handler(srv))
	r.PUT("/v1/order/order/{order.id}", _OrderService_UpdateOrder1_HTTP_Handler(srv))
	r.DELETE("/v1/order/order/{id}", _OrderService_DeleteOrder0_HTTP_Handler(srv))
}

func _OrderService_ListOrder0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceListOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrder(ctx, req.(*ListOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrderReply)
		return ctx.Result(200, reply)
	}
}

func _OrderService_ListOrder1_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceListOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrder(ctx, req.(*ListOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrderReply)
		return ctx.Result(200, reply)
	}
}

func _OrderService_GetOrder0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceGetOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrder(ctx, req.(*GetOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Order)
		return ctx.Result(200, reply)
	}
}

func _OrderService_CreateOrder0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceCreateOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrder(ctx, req.(*CreateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Order)
		return ctx.Result(200, reply)
	}
}

func _OrderService_UpdateOrder0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceUpdateOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOrder(ctx, req.(*UpdateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Order)
		return ctx.Result(200, reply)
	}
}

func _OrderService_UpdateOrder1_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceUpdateOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOrder(ctx, req.(*UpdateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Order)
		return ctx.Result(200, reply)
	}
}

func _OrderService_DeleteOrder0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteOrderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceDeleteOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteOrder(ctx, req.(*DeleteOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteOrderReply)
		return ctx.Result(200, reply)
	}
}

type OrderServiceHTTPClient interface {
	CreateOrder(ctx context.Context, req *CreateOrderRequest, opts ...http.CallOption) (rsp *Order, err error)
	DeleteOrder(ctx context.Context, req *DeleteOrderRequest, opts ...http.CallOption) (rsp *DeleteOrderReply, err error)
	GetOrder(ctx context.Context, req *GetOrderRequest, opts ...http.CallOption) (rsp *Order, err error)
	ListOrder(ctx context.Context, req *ListOrderRequest, opts ...http.CallOption) (rsp *ListOrderReply, err error)
	UpdateOrder(ctx context.Context, req *UpdateOrderRequest, opts ...http.CallOption) (rsp *Order, err error)
}

type OrderServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderServiceHTTPClient(client *http.Client) OrderServiceHTTPClient {
	return &OrderServiceHTTPClientImpl{client}
}

func (c *OrderServiceHTTPClientImpl) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...http.CallOption) (*Order, error) {
	var out Order
	pattern := "/v1/order/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderServiceCreateOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderServiceHTTPClientImpl) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...http.CallOption) (*DeleteOrderReply, error) {
	var out DeleteOrderReply
	pattern := "/v1/order/order/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderServiceDeleteOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderServiceHTTPClientImpl) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...http.CallOption) (*Order, error) {
	var out Order
	pattern := "/v1/order/order/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderServiceGetOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderServiceHTTPClientImpl) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...http.CallOption) (*ListOrderReply, error) {
	var out ListOrderReply
	pattern := "/v1/order/order"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderServiceListOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderServiceHTTPClientImpl) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...http.CallOption) (*Order, error) {
	var out Order
	pattern := "/v1/order/order/{order.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderServiceUpdateOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
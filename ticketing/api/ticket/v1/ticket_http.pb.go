// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             (unknown)
// source: ticketing/api/ticket/v1/ticket.proto

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

const OperationTicketServiceGetTicket = "/ticketing.api.ticket.v1.TicketService/GetTicket"
const OperationTicketServiceListTicket = "/ticketing.api.ticket.v1.TicketService/ListTicket"

type TicketServiceHTTPServer interface {
	GetTicket(context.Context, *GetTicketRequest) (*Ticket, error)
	ListTicket(context.Context, *ListTicketRequest) (*ListTicketReply, error)
}

func RegisterTicketServiceHTTPServer(s *http.Server, srv TicketServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/ticketing/ticket/list", _TicketService_ListTicket0_HTTP_Handler(srv))
	r.GET("/v1/ticketing/ticket", _TicketService_ListTicket1_HTTP_Handler(srv))
	r.GET("/v1/ticketing/ticket/{id}", _TicketService_GetTicket0_HTTP_Handler(srv))
}

func _TicketService_ListTicket0_HTTP_Handler(srv TicketServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTicketRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketServiceListTicket)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTicket(ctx, req.(*ListTicketRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTicketReply)
		return ctx.Result(200, reply)
	}
}

func _TicketService_ListTicket1_HTTP_Handler(srv TicketServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTicketRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketServiceListTicket)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTicket(ctx, req.(*ListTicketRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTicketReply)
		return ctx.Result(200, reply)
	}
}

func _TicketService_GetTicket0_HTTP_Handler(srv TicketServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTicketRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTicketServiceGetTicket)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTicket(ctx, req.(*GetTicketRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Ticket)
		return ctx.Result(200, reply)
	}
}

type TicketServiceHTTPClient interface {
	GetTicket(ctx context.Context, req *GetTicketRequest, opts ...http.CallOption) (rsp *Ticket, err error)
	ListTicket(ctx context.Context, req *ListTicketRequest, opts ...http.CallOption) (rsp *ListTicketReply, err error)
}

type TicketServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTicketServiceHTTPClient(client *http.Client) TicketServiceHTTPClient {
	return &TicketServiceHTTPClientImpl{client}
}

func (c *TicketServiceHTTPClientImpl) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...http.CallOption) (*Ticket, error) {
	var out Ticket
	pattern := "/v1/ticketing/ticket/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTicketServiceGetTicket))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TicketServiceHTTPClientImpl) ListTicket(ctx context.Context, in *ListTicketRequest, opts ...http.CallOption) (*ListTicketReply, error) {
	var out ListTicketReply
	pattern := "/v1/ticketing/ticket"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTicketServiceListTicket))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

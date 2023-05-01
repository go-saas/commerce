// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: ticketing/api/ticket/v1/ticket_app.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ TicketAppServiceServer = (*ticketAppServiceClientProxy)(nil)

const GrpcOperationTicketAppServiceListMyTicket = "/ticketing.api.ticket.v1.TicketAppService/ListMyTicket"

// ticketAppServiceClientProxy is the proxy to turn TicketAppService client to server interface.
type ticketAppServiceClientProxy struct {
	cc TicketAppServiceClient
}

func NewTicketAppServiceClientProxy(cc TicketAppServiceClient) TicketAppServiceServer {
	return &ticketAppServiceClientProxy{cc}
}

func (c *ticketAppServiceClientProxy) ListMyTicket(ctx context.Context, in *ListTicketRequest) (*ListTicketReply, error) {
	return c.cc.ListMyTicket(ctx, in)
}

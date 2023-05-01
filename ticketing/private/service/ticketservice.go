package service

import (
	"context"
	"github.com/go-saas/kit/pkg/authz/authz"

	pb "github.com/go-saas/commerce/ticketing/api/ticket/v1"
)

type TicketServiceService struct {
	auth authz.Service
}

var _ pb.TicketServiceServer = (*TicketServiceService)(nil)
var _ pb.TicketAppServiceServer = (*TicketServiceService)(nil)

func NewTicketServiceService(auth authz.Service) *TicketServiceService {
	return &TicketServiceService{auth: auth}
}

func (s *TicketServiceService) ListTicket(ctx context.Context, req *pb.ListTicketRequest) (*pb.ListTicketReply, error) {
	return &pb.ListTicketReply{}, nil
}

func (s *TicketServiceService) ListMyTicket(ctx context.Context, req *pb.ListTicketRequest) (*pb.ListTicketReply, error) {
	return &pb.ListTicketReply{}, nil
}

func (s *TicketServiceService) GetTicket(ctx context.Context, req *pb.GetTicketRequest) (*pb.Ticket, error) {
	return &pb.Ticket{}, nil
}
func (s *TicketServiceService) CreateTicket(ctx context.Context, req *pb.CreateTicketRequest) (*pb.Ticket, error) {
	return &pb.Ticket{}, nil
}
func (s *TicketServiceService) UpdateTicket(ctx context.Context, req *pb.UpdateTicketRequest) (*pb.Ticket, error) {
	return &pb.Ticket{}, nil
}
func (s *TicketServiceService) DeleteTicket(ctx context.Context, req *pb.DeleteTicketRequest) (*pb.DeleteTicketReply, error) {
	return &pb.DeleteTicketReply{}, nil
}

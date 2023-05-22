package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-saas/commerce/ticketing/api"
	v1 "github.com/go-saas/commerce/ticketing/api/activity/v1"
	v12 "github.com/go-saas/commerce/ticketing/api/show/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/kit/pkg/authn"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/go-saas/kit/pkg/query"
	"github.com/go-saas/kit/pkg/utils"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/go-saas/commerce/ticketing/api/ticket/v1"
)

type TicketServiceService struct {
	auth        authz.Service
	repo        biz.TicketRepo
	activitySrv *ActivityService
	showSrv     *ShowService
}

var _ pb.TicketServiceServer = (*TicketServiceService)(nil)
var _ pb.TicketAppServiceServer = (*TicketServiceService)(nil)

func NewTicketServiceService(auth authz.Service, ticketRepo biz.TicketRepo, activitySrv *ActivityService, showSrv *ShowService) *TicketServiceService {
	return &TicketServiceService{auth: auth, repo: ticketRepo, activitySrv: activitySrv, showSrv: showSrv}
}

func (s *TicketServiceService) ListTicket(ctx context.Context, req *pb.ListTicketRequest) (*pb.ListTicketReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceTicket, "*"), authz.ReadAction); err != nil {
		return nil, err
	}
	ret := &pb.ListTicketReply{}

	totalCount, filterCount, err := s.repo.Count(ctx, req)
	ret.TotalSize = int32(totalCount)
	ret.FilterSize = int32(filterCount)

	if err != nil {
		return ret, err
	}
	items, err := s.repo.List(ctx, req)
	if err != nil {
		return ret, err
	}
	rItems := lo.Map(items, func(g *biz.Ticket, _ int) *pb.Ticket {
		b := &pb.Ticket{}
		s.MapBizTicket2Pb(ctx, g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}

func (s *TicketServiceService) ListMyTicket(ctx context.Context, req *pb.ListTicketRequest) (*pb.ListTicketReply, error) {
	userInfo, err := authn.ErrIfUnauthenticated(ctx)
	if err != nil {
		return nil, err
	}
	ret := &pb.ListTicketReply{}
	if req.Filter == nil {
		req.Filter = &pb.TicketFilter{}
	}
	req.Filter.UserId = &query.StringFilterOperation{Eq: &wrapperspb.StringValue{Value: userInfo.GetId()}}
	cursorRet, err := s.repo.ListCursor(ctx, req)
	if err != nil {
		return nil, err
	}
	ret.NextBeforePageToken = cursorRet.Before
	ret.NextAfterPageToken = cursorRet.After

	if err != nil {
		return ret, err
	}
	items := cursorRet.Items
	rItems := lo.Map(items, func(g *biz.Ticket, _ int) *pb.Ticket {
		b := &pb.Ticket{}
		s.MapBizTicket2Pb(ctx, g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}

func (s *TicketServiceService) GetTicket(ctx context.Context, req *pb.GetTicketRequest) (*pb.Ticket, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceTicket, req.GetId()), authz.ReadAction); err != nil {
		return nil, err
	}

	g, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	res := &pb.Ticket{}
	s.MapBizTicket2Pb(ctx, g, res)
	return res, nil
}

func (s *TicketServiceService) MapBizTicket2Pb(ctx context.Context, a *biz.Ticket, b *pb.Ticket) {
	b.Id = a.ID
	b.CreatedAt = utils.Time2Timepb(&a.CreatedAt)

	activity := &v1.Activity{}
	s.activitySrv.MapBizActivity2Pb(ctx, &a.Activity, activity)
	b.Activity = activity

	show := &v12.Show{}
	s.showSrv.MapBizShow2Pb(ctx, &a.Show, show)
	b.Show = show

	b.Status = a.Status
	b.OrderId = a.OrderID
	b.UserId = a.UserID
	b.Notice = a.Notice
}

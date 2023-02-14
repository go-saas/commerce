package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-saas/commerce/pkg/price"
	"github.com/go-saas/commerce/ticketing/api"
	v13 "github.com/go-saas/commerce/ticketing/api/activity/v1"
	v1 "github.com/go-saas/commerce/ticketing/api/location/v1"
	pb "github.com/go-saas/commerce/ticketing/api/show/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/go-saas/kit/pkg/utils"
	"github.com/google/uuid"
	"github.com/goxiaoy/vfs"
	"github.com/samber/lo"
)

type ShowService struct {
	repo         biz.ShowRepo
	hallRepo     biz.HallRepo
	auth         authz.Service
	blob         vfs.Blob
	mediaRepo    biz.TicketingMediaRepo
	categoryRepo biz.TicketingCategoryRepo
	activitySrv  *ActivityService
	locationSrv  *LocationService
}

var _ pb.ShowServiceServer = (*ShowService)(nil)

func NewShowService(
	repo biz.ShowRepo,
	hallRepo biz.HallRepo,
	auth authz.Service,
	blob vfs.Blob,
	mediaRepo biz.TicketingMediaRepo,
	categoryRepo biz.TicketingCategoryRepo,
	activitySrv *ActivityService,
	locationSrv *LocationService,
) *ShowService {
	return &ShowService{
		repo:         repo,
		hallRepo:     hallRepo,
		auth:         auth,
		blob:         blob,
		mediaRepo:    mediaRepo,
		categoryRepo: categoryRepo,
		activitySrv:  activitySrv,
		locationSrv:  locationSrv,
	}
}

func (s *ShowService) ListShow(ctx context.Context, req *pb.ListShowRequest) (*pb.ListShowReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceShow, "*"), authz.ReadAction); err != nil {
		return nil, err
	}
	ret := &pb.ListShowReply{}

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
	rItems := lo.Map(items, func(g *biz.Show, _ int) *pb.Show {
		b := &pb.Show{}
		s.MapBizShow2Pb(ctx, g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}

func (s *ShowService) GetShow(ctx context.Context, req *pb.GetShowRequest) (*pb.Show, error) {

	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceShow, req.GetId()), authz.ReadAction); err != nil {
		return nil, err
	}

	g, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	res := &pb.Show{}
	s.MapBizShow2Pb(ctx, g, res)
	return res, nil
}
func (s *ShowService) CreateShow(ctx context.Context, req *pb.CreateShowRequest) (*pb.Show, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceShow, "*"), authz.WriteAction); err != nil {
		return nil, err
	}
	e := &biz.Show{}
	err := s.MapCreatePbShow2Biz(ctx, req, e)
	if err != nil {
		return nil, err
	}
	err = s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	res := &pb.Show{}
	s.MapBizShow2Pb(ctx, e, res)
	return res, nil
}
func (s *ShowService) UpdateShow(ctx context.Context, req *pb.UpdateShowRequest) (*pb.Show, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceShow, req.Show.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Show.Id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	err = s.MapUpdatePbShow2Biz(ctx, req.Show, g)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Update(ctx, g.ID.String(), g, nil); err != nil {
		return nil, err
	}
	res := &pb.Show{}
	s.MapBizShow2Pb(ctx, g, res)
	return res, nil
}

func (s *ShowService) DeleteShow(ctx context.Context, req *pb.DeleteShowRequest) (*pb.DeleteShowReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceShow, req.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	err = s.repo.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteShowReply{Id: g.ID.String()}, nil
}

func (s *ShowService) MapBizShow2Pb(ctx context.Context, a *biz.Show, b *pb.Show) {
	b.Id = a.ID.String()

	b.CreatedAt = utils.Time2Timepb(&a.CreatedAt)

	b.ActivityId = a.ActivityID
	if a.Activity != nil {
		b.Activity = &v13.Activity{}
		s.activitySrv.MapBizActivity2Pb(ctx, a.Activity, b.Activity)
	}

	b.StartTime = utils.Time2Timepb(&a.StartTime)
	b.EndTime = utils.Time2Timepb(&a.EndTime)

	b.LocationId = a.LocationID
	if a.Location != nil {
		b.Location = &v1.Location{}
		s.locationSrv.MapBizLocation2Pb(ctx, a.Location, b.Location)
	}
	b.HallId = a.HallID
	if a.Hall != nil {
		b.Hall = mapBizHall2Pb(ctx, a.Hall)
	}
	b.SalesTypes = lo.Map(a.SalesTypes, func(item biz.ShowSalesType, _ int) *pb.ShowSalesType {
		r := &pb.ShowSalesType{}
		mapBizShowSalesType2Pb(ctx, &item, r)
		return r
	})
	b.Notice = a.Notice
	//TODO seats
}

func mapBizShowSalesType2Pb(ctx context.Context, a *biz.ShowSalesType, b *pb.ShowSalesType) {
	b.Id = a.ID.String()
	b.Name = a.Name
	b.SeatGroup = mapBizSeatGroup2Pb(ctx, a.SeatGroup)
	b.SaleableFrom = utils.Time2Timepb(a.SaleableFrom)
	b.SaleableTo = utils.Time2Timepb(a.SaleableTo)
	b.Price = a.Price.ToInfoPb()
}

func mapPbShowSalesType2Biz(ctx context.Context, a *pb.UpdateShowSalesType, b *biz.ShowSalesType) error {
	if len(a.Id) > 0 {
		b.ID = uuid.MustParse(a.Id)
	}
	b.Name = a.Name
	if len(a.SeatGroupId) > 0 {
		b.SeatGroupID = &a.SeatGroupId
	}

	b.SaleableFrom = utils.Timepb2Time(a.SaleableFrom)
	b.SaleableTo = utils.Timepb2Time(a.SaleableTo)

	p, err := price.NewInfoFromPb(a.Price)
	if err != nil {
		return err
	}
	b.Price = p
	return nil
}

func (s *ShowService) MapUpdatePbShow2Biz(ctx context.Context, a *pb.UpdateShow, b *biz.Show) error {

	b.StartTime = *utils.Timepb2Time(a.StartTime)
	b.EndTime = *utils.Timepb2Time(a.EndTime)

	b.Notice = a.Notice

	var sts []biz.ShowSalesType
	for _, salesType := range a.SalesTypes {
		r := biz.ShowSalesType{}
		err := mapPbShowSalesType2Biz(ctx, salesType, &r)
		if err != nil {
			return err
		}
		sts = append(sts, r)
	}
	b.SalesTypes = sts
	return nil

}
func (s *ShowService) MapCreatePbShow2Biz(ctx context.Context, a *pb.CreateShowRequest, b *biz.Show) error {
	b.ActivityID = a.ActivityId
	b.StartTime = *utils.Timepb2Time(a.StartTime)
	b.EndTime = *utils.Timepb2Time(a.EndTime)
	b.LocationID = a.LocationId
	b.HallID = a.HallId

	b.Notice = a.Notice

	var sts []biz.ShowSalesType
	for _, salesType := range a.SalesTypes {
		r := biz.ShowSalesType{}
		err := mapPbShowSalesType2Biz(ctx, salesType, &r)
		if err != nil {
			return err
		}
		sts = append(sts, r)
	}
	b.SalesTypes = sts
	return nil
}

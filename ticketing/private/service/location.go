package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-saas/commerce/ticketing/api"
	pb "github.com/go-saas/commerce/ticketing/api/location/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LocationService struct {
	pb.UnimplementedLocationServiceServer
	repo biz.LocationRepo
	auth authz.Service
}

func NewLocationService(repo biz.LocationRepo, auth authz.Service) *LocationService {
	return &LocationService{repo: repo, auth: auth}
}

func (s *LocationService) ListLocation(ctx context.Context, req *pb.ListLocationRequest) (*pb.ListLocationReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, "*"), authz.ReadAction); err != nil {
		return nil, err
	}
	ret := &pb.ListLocationReply{}

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
	rItems := lo.Map(items, func(g *biz.Location, _ int) *pb.Location {
		b := &pb.Location{}
		MapBizLocation2Pb(g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}
func (s *LocationService) GetLocation(ctx context.Context, req *pb.GetLocationRequest) (*pb.Location, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, req.GetId()), authz.ReadAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	res := &pb.Location{}
	MapBizLocation2Pb(g, res)
	return res, nil
}
func (s *LocationService) CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.Location, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, "*"), authz.WriteAction); err != nil {
		return nil, err
	}
	e := &biz.Location{}
	MapCreatePbLocation2Biz(req, e)
	err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	res := &pb.Location{}
	MapBizLocation2Pb(e, res)
	return res, nil
}
func (s *LocationService) UpdateLocation(ctx context.Context, req *pb.UpdateLocationRequest) (*pb.Location, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, req.Location.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Location.Id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	MapUpdatePbLocation2Biz(req.Location, g)
	if err := s.repo.Update(ctx, g.ID.String(), g, nil); err != nil {
		return nil, err
	}
	res := &pb.Location{}
	MapBizLocation2Pb(g, res)
	return res, nil
}

func (s *LocationService) DeleteLocation(ctx context.Context, req *pb.DeleteLocationRequest) (*pb.DeleteLocationReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, req.Id), authz.WriteAction); err != nil {
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
	return &pb.DeleteLocationReply{Id: g.ID.String(), Name: g.Name}, nil
}
func MapBizLocation2Pb(a *biz.Location, b *pb.Location) {
	b.Id = a.ID.String()
	b.Name = a.Name
	b.CreatedAt = timestamppb.New(a.CreatedAt)
}

func MapUpdatePbLocation2Biz(a *pb.UpdateLocation, b *biz.Location) {
	b.Name = a.Name
}
func MapCreatePbLocation2Biz(a *pb.CreateLocationRequest, b *biz.Location) {
	b.Name = a.Name
}

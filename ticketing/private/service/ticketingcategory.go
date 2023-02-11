package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-saas/commerce/ticketing/api"
	pb "github.com/go-saas/commerce/ticketing/api/category/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/samber/lo"
)

type TicketingCategoryService struct {
	pb.UnimplementedTicketingCategoryServiceServer
	repo biz.TicketingCategoryRepo
	auth authz.Service
}

func NewTicketingCategoryService(repo biz.TicketingCategoryRepo,
	auth authz.Service) *TicketingCategoryService {
	return &TicketingCategoryService{repo: repo, auth: auth}
}

func (s *TicketingCategoryService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceCategory, "*"), authz.ReadAction); err != nil {
		return nil, err
	}
	ret := &pb.ListCategoryReply{}

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
	rItems := lo.Map(items, func(g *biz.TicketingCategory, _ int) *pb.Category {
		b := &pb.Category{}
		s.MapBizCategory2Pb(ctx, g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}

func (s *TicketingCategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {

	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceCategory, req.Key), authz.ReadAction); err != nil {
		return nil, err
	}

	g, err := s.repo.Get(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	res := &pb.Category{}
	s.MapBizCategory2Pb(ctx, g, res)
	return res, nil
}

func (s *TicketingCategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceCategory, "*"), authz.WriteAction); err != nil {
		return nil, err
	}
	e := &biz.TicketingCategory{}
	s.MapCreatePbCategory2Biz(req, e)
	err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	res := &pb.Category{}
	s.MapBizCategory2Pb(ctx, e, res)
	return res, nil
}
func (s *TicketingCategoryService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.Category, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceCategory, req.Category.Key), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Category.Key)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	s.MapUpdatePbCategory2Biz(req.Category, g)
	if err := s.repo.Update(ctx, g.Key, g, nil); err != nil {
		return nil, err
	}

	res := &pb.Category{}
	s.MapBizCategory2Pb(ctx, g, res)
	return res, nil
}

func (s *TicketingCategoryService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceCategory, req.Key), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	err = s.repo.Delete(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCategoryReply{Key: g.Key, Name: g.Name}, nil
}

func (s *TicketingCategoryService) MapBizCategory2Pb(ctx context.Context, a *biz.TicketingCategory, b *pb.Category) {
	b.Key = a.Key
	b.Name = a.Name
	b.Path = a.Path
	if a.ParentID != nil {
		b.Parent = *a.ParentID
	}
}

func (s *TicketingCategoryService) MapCreatePbCategory2Biz(a *pb.CreateCategoryRequest, b *biz.TicketingCategory) {
	b.Key = a.Key
	b.Name = a.Name
	if len(a.Parent) > 0 {
		b.ParentID = &a.Parent
	}
}

func (s *TicketingCategoryService) MapUpdatePbCategory2Biz(a *pb.UpdateCategory, b *biz.TicketingCategory) {
	b.Name = a.Name
	if len(a.Parent) > 0 {
		b.ParentID = &a.Parent
	}
}

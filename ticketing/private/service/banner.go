package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-saas/commerce/ticketing/api"
	pb "github.com/go-saas/commerce/ticketing/api/banner/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/goxiaoy/vfs"
	"github.com/samber/lo"
)

type TicketingBannerServiceService struct {
	repo biz.BannerRepo
	auth authz.Service
	blob vfs.Blob
	*UploadService
}

func NewTicketingBannerServiceService(repo biz.BannerRepo, auth authz.Service, blob vfs.Blob, upload *UploadService) *TicketingBannerServiceService {
	return &TicketingBannerServiceService{repo: repo, auth: auth, blob: blob, UploadService: upload}
}

var _ pb.TicketingBannerServiceServer = (*TicketingBannerServiceService)(nil)
var _ pb.TicketingAppBannerServiceServer = (*TicketingBannerServiceService)(nil)

func (s *TicketingBannerServiceService) ListBanner(ctx context.Context, req *pb.ListBannerRequest) (*pb.ListBannerReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceBanner, "*"), authz.ReadAction); err != nil {
		return nil, err
	}
	ret := &pb.ListBannerReply{}

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
	rItems := lo.Map(items, func(g *biz.Banner, _ int) *pb.Banner {
		b := &pb.Banner{}
		MapBizBanner2Pb(ctx, s.blob, g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}

func (s *TicketingBannerServiceService) ListAppBanner(ctx context.Context, req *pb.ListAppBannerRequest) (*pb.ListAppBannerReply, error) {
	return &pb.ListAppBannerReply{}, nil
}

func (s *TicketingBannerServiceService) CreateBanner(ctx context.Context, req *pb.CreateBannerRequest) (*pb.Banner, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceBanner, "*"), authz.WriteAction); err != nil {
		return nil, err
	}
	e := biz.NewBanner()
	s.MapCreatePbBanner2Biz(req, e)
	err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	res := &pb.Banner{}
	MapBizBanner2Pb(ctx, s.blob, e, res)
	return res, nil
}
func (s *TicketingBannerServiceService) UpdateBanner(ctx context.Context, req *pb.UpdateBannerRequest) (*pb.Banner, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceBanner, req.Banner.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Banner.Id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	s.MapUpdatePbBanner2Biz(req.Banner, g)
	if err := s.repo.Update(ctx, req.Banner.Id, g, nil); err != nil {
		return nil, err
	}

	res := &pb.Banner{}
	MapBizBanner2Pb(ctx, s.blob, g, res)
	return res, nil
}
func (s *TicketingBannerServiceService) DeleteBanner(ctx context.Context, req *pb.DeleteBannerRequest) (*pb.DeleteBannerReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceBanner, req.Id), authz.WriteAction); err != nil {
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
	return &pb.DeleteBannerReply{Id: req.Id}, nil
}

func MapBizBanner2Pb(ctx context.Context, blob vfs.Blob, a *biz.Banner, b *pb.Banner) {
	b.RefType = a.RefType
	b.RefId = a.RefID
	b.Status = a.Status
	b.Id = a.ID.String()
	b.MainPic = mapBizMedia2Pb(ctx, blob, a.MainPic)
}

func (s *TicketingBannerServiceService) MapCreatePbBanner2Biz(a *pb.CreateBannerRequest, b *biz.Banner) {
	b.RefType = a.RefType
	b.RefID = a.RefId
	b.MainPic = mapPbMedia2Biz(a.MainPic)
}

func (s *TicketingBannerServiceService) MapUpdatePbBanner2Biz(a *pb.UpdateBanner, b *biz.Banner) {
	b.RefType = a.RefType
	b.RefID = a.RefId
	b.MainPic = mapPbMedia2Biz(a.MainPic)
	b.Status = a.Status
}

func (s *TicketingBannerServiceService) UploadMedias(ctx http.Context) error {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceBanner, "*"), authz.WriteAction); err != nil {
		return err
	}
	return s.upload(ctx, biz.BannerMediaPath)
}

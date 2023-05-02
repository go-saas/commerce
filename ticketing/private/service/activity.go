package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-saas/commerce/ticketing/api"
	pb "github.com/go-saas/commerce/ticketing/api/activity/v1"
	v12 "github.com/go-saas/commerce/ticketing/api/category/v1"
	v1 "github.com/go-saas/commerce/ticketing/api/location/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/go-saas/kit/pkg/utils"
	"github.com/goxiaoy/vfs"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/durationpb"
)

type ActivityService struct {
	repo         biz.ActivityRepo
	hallRepo     biz.HallRepo
	auth         authz.Service
	blob         vfs.Blob
	mediaRepo    biz.TicketingMediaRepo
	categoryRepo biz.TicketingCategoryRepo
	*UploadService
}

var _ pb.ActivityServiceServer = (*ActivityService)(nil)

func NewActivityService(
	upload *UploadService,
	repo biz.ActivityRepo,
	hallRepo biz.HallRepo,
	auth authz.Service,
	blob vfs.Blob,
	mediaRepo biz.TicketingMediaRepo,
	categoryRepo biz.TicketingCategoryRepo,
) *ActivityService {
	return &ActivityService{repo: repo, hallRepo: hallRepo, auth: auth, blob: blob, mediaRepo: mediaRepo, categoryRepo: categoryRepo, UploadService: upload}
}

func (s *ActivityService) ListActivity(ctx context.Context, req *pb.ListActivityRequest) (*pb.ListActivityReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceActivity, "*"), authz.ReadAction); err != nil {
		return nil, err
	}
	ret := &pb.ListActivityReply{}

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
	rItems := lo.Map(items, func(g *biz.Activity, _ int) *pb.Activity {
		b := &pb.Activity{}
		s.MapBizActivity2Pb(ctx, g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}

func (s *ActivityService) GetActivity(ctx context.Context, req *pb.GetActivityRequest) (*pb.Activity, error) {

	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceActivity, req.GetId()), authz.ReadAction); err != nil {
		return nil, err
	}

	g, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	res := &pb.Activity{}
	s.MapBizActivity2Pb(ctx, g, res)
	return res, nil
}
func (s *ActivityService) CreateActivity(ctx context.Context, req *pb.CreateActivityRequest) (*pb.Activity, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceActivity, "*"), authz.WriteAction); err != nil {
		return nil, err
	}
	e := &biz.Activity{}
	err := s.MapCreatePbActivity2Biz(ctx, req, e)
	if err != nil {
		return nil, err
	}
	err = s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	res := &pb.Activity{}
	s.MapBizActivity2Pb(ctx, e, res)
	return res, nil
}
func (s *ActivityService) UpdateActivity(ctx context.Context, req *pb.UpdateActivityRequest) (*pb.Activity, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceActivity, req.Activity.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Activity.Id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	err = s.MapUpdatePbActivity2Biz(ctx, req.Activity, g)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Update(ctx, g.ID.String(), g, nil); err != nil {
		return nil, err
	}
	res := &pb.Activity{}
	s.MapBizActivity2Pb(ctx, g, res)
	return res, nil
}

func (s *ActivityService) DeleteActivity(ctx context.Context, req *pb.DeleteActivityRequest) (*pb.DeleteActivityReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceActivity, req.Id), authz.WriteAction); err != nil {
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
	return &pb.DeleteActivityReply{Id: g.ID.String(), Name: g.Name}, nil
}

func (s *ActivityService) UploadMedias(ctx http.Context) error {
	return s.upload(ctx, biz.ActivityMediaPath, func(ctx context.Context) error {
		_, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceActivity, "*"), authz.WriteAction)
		return err
	})
}

func (s *ActivityService) MapBizActivity2Pb(ctx context.Context, a *biz.Activity, b *pb.Activity) {
	b.Id = a.ID.String()
	b.Name = a.Name

	b.Medias = lo.Map(a.Medias, func(item biz.TicketingMedia, _ int) *v1.TicketingMedia {
		return mapBizMedia2Pb(ctx, s.blob, &item)
	})
	b.Desc = a.Desc
	b.ShortDesc = a.ShortDesc
	b.Content = utils.Map2Structpb(a.Content)
	b.SeatSelectable = a.SeatSelectable
	b.Status = a.Status

	b.Duration = durationpb.New(a.Duration)
	b.Categories = lo.Map(a.Categories, func(item biz.TicketingCategory, _ int) *v12.Category {
		r := &v12.Category{}
		MapBizCategory2Pb(ctx, &item, r)
		return r
	})
}

func (s *ActivityService) MapUpdatePbActivity2Biz(ctx context.Context, a *pb.UpdateActivity, b *biz.Activity) error {
	b.Name = a.Name

	b.Medias = lo.Map(a.Medias, func(item *v1.TicketingMedia, _ int) biz.TicketingMedia {
		return *mapPbMedia2Biz(item)
	})
	b.Desc = a.Desc
	b.ShortDesc = a.ShortDesc
	b.Content = utils.Structpb2Map(a.Content)
	b.SeatSelectable = a.SeatSelectable

	if a.Duration != nil {
		b.Duration = a.Duration.AsDuration()
	}
	if len(a.CategoryKeys) > 0 {
		c, err := s.categoryRepo.FindByKeys(ctx, a.CategoryKeys)
		if err != nil {
			return err
		}
		b.Categories = c
	}
	return nil

}
func (s *ActivityService) MapCreatePbActivity2Biz(ctx context.Context, a *pb.CreateActivityRequest, b *biz.Activity) error {
	b.Name = a.Name

	b.Medias = lo.Map(a.Medias, func(item *v1.TicketingMedia, _ int) biz.TicketingMedia {
		return *mapPbMedia2Biz(item)
	})
	b.Desc = a.Desc
	b.ShortDesc = a.ShortDesc
	b.Content = utils.Structpb2Map(a.Content)
	b.SeatSelectable = a.SeatSelectable

	if a.Duration != nil {
		b.Duration = a.Duration.AsDuration()
	}
	if len(a.CategoryKeys) > 0 {
		c, err := s.categoryRepo.FindByKeys(ctx, a.CategoryKeys)
		if err != nil {
			return err
		}
		b.Categories = c
	}
	return nil
}

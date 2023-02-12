package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-saas/commerce/ticketing/api"
	pb "github.com/go-saas/commerce/ticketing/api/location/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/go-saas/kit/pkg/blob"
	"github.com/go-saas/kit/pkg/utils"
	"github.com/go-saas/lbs"
	"github.com/google/uuid"
	"github.com/goxiaoy/vfs"
	"github.com/samber/lo"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

type LocationService struct {
	repo      biz.LocationRepo
	hallRepo  biz.HallRepo
	auth      authz.Service
	blob      vfs.Blob
	mediaRepo biz.TicketingMediaRepo
}

var _ pb.LocationServiceServer = (*LocationService)(nil)

func NewLocationService(repo biz.LocationRepo, hallRepo biz.HallRepo, auth authz.Service, blob vfs.Blob, mediaRepo biz.TicketingMediaRepo) *LocationService {
	return &LocationService{repo: repo, hallRepo: hallRepo, auth: auth, blob: blob, mediaRepo: mediaRepo}
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
		s.MapBizLocation2Pb(ctx, g, b)
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
	s.MapBizLocation2Pb(ctx, g, res)
	return res, nil
}
func (s *LocationService) CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.Location, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, "*"), authz.WriteAction); err != nil {
		return nil, err
	}
	e := &biz.Location{}
	s.MapCreatePbLocation2Biz(req, e)
	err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	res := &pb.Location{}
	s.MapBizLocation2Pb(ctx, e, res)
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

	s.MapUpdatePbLocation2Biz(req.Location, g)
	if err := s.repo.Update(ctx, g.ID.String(), g, nil); err != nil {
		return nil, err
	}
	res := &pb.Location{}
	s.MapBizLocation2Pb(ctx, g, res)
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

func (s *LocationService) GetLocationHalls(ctx context.Context, req *pb.GetLocationHallsRequest) (*pb.GetLocationHallsReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, req.Id), authz.ReadAction); err != nil {
		return nil, err
	}
	halls, err := s.repo.ListHalls(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	ret := &pb.GetLocationHallsReply{}
	ret.Halls = lo.Map(halls, func(item biz.Hall, _ int) *pb.LocationHall {
		return mapBizHall2Pb(ctx, &item)
	})

	return ret, nil
}

func (s *LocationService) CreateLocationHall(ctx context.Context, req *pb.CreateLocationHallRequest) (*pb.CreateLocationHallReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, req.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	hallEntity := &biz.Hall{}
	mapPbHall2Biz(ctx, req.Hall, hallEntity)
	err = s.repo.CreateHall(ctx, g, hallEntity)
	if err != nil {
		return nil, err
	}
	return &pb.CreateLocationHallReply{}, err
}

func (s *LocationService) UpdateLocationHall(ctx context.Context, req *pb.UpdateLocationHallRequest) (*pb.UpdateLocationHallReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, req.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	hall, err := s.hallRepo.Get(ctx, req.Hall.Id)
	if err != nil {
		return nil, err
	}
	if hall.LocationID != g.ID.String() {
		return nil, errors.NotFound("", "")
	}
	mapPbHall2Biz(ctx, req.Hall, hall)
	err = s.hallRepo.Update(ctx, req.Hall.Id, hall, nil)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateLocationHallReply{}, nil

}

func (s *LocationService) DeleteLocationHall(ctx context.Context, req *pb.DeleteLocationHallRequest) (*pb.DeleteLocationHallReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, req.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	hall, err := s.hallRepo.Get(ctx, req.HallId)
	if err != nil {
		return nil, err
	}
	if hall.LocationID != g.ID.String() {
		return nil, errors.NotFound("", "")
	}
	err = s.hallRepo.Delete(ctx, req.HallId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteLocationHallReply{}, nil
}

func (s *LocationService) UploadLogo(ctx http.Context) error {
	return s.upload(ctx, biz.LocationLogoPath)
}

func (s *LocationService) UploadMedias(ctx http.Context) error {
	return s.upload(ctx, biz.LocationMediaPath)
}

func (s *LocationService) UploadLegalDocs(ctx http.Context) error {
	return s.upload(ctx, biz.LocationLegalDocumentsPath)
}

func (s *LocationService) upload(ctx http.Context, basePath string) error {
	req := ctx.Request()
	//TODO do not know why should read form file first ...
	if _, _, err := req.FormFile("file"); err != nil {
		return err
	}

	h := ctx.Middleware(func(ctx context.Context, _ interface{}) (interface{}, error) {
		if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourceLocation, "*"), authz.WriteAction); err != nil {
			return nil, err
		}
		file, handle, err := req.FormFile("file")
		if err != nil {
			return nil, err
		}
		defer file.Close()
		fileName := handle.Filename
		ext := filepath.Ext(fileName)
		normalizedName := fmt.Sprintf("%s/%s%s", basePath, uuid.New().String(), ext)

		err = s.blob.MkdirAll(basePath, 0755)
		if err != nil {
			return nil, err
		}
		f, err := s.blob.OpenFile(normalizedName, os.O_WRONLY|os.O_CREATE, 0o666)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			return nil, err
		}
		err = s.mediaRepo.Create(ctx, &biz.TicketingMedia{
			ID:       normalizedName,
			MimeType: mime.TypeByExtension(ext),
			Usage:    "location",
			Name:     strings.TrimSuffix(fileName, ext),
		})
		if err != nil {
			return nil, err
		}

		url, _ := s.blob.PublicUrl(ctx, normalizedName)
		return &blob.BlobFile{
			Id:   normalizedName,
			Name: strings.TrimSuffix(fileName, ext),
			Url:  url.URL,
		}, nil
	})
	out, err := h(ctx, nil)
	if err != nil {
		return err
	}
	return ctx.Result(201, out)
}

func (s *LocationService) MapBizLocation2Pb(ctx context.Context, a *biz.Location, b *pb.Location) {
	b.Id = a.ID.String()
	b.Name = a.Name
	b.Logo = mapBizMedia2Pb(ctx, s.blob, a.Logo)
	b.Medias = lo.Map(a.Medias, func(item biz.TicketingMedia, _ int) *pb.TicketingMedia {
		return mapBizMedia2Pb(ctx, s.blob, &item)
	})
	b.Desc = a.Desc
	b.ShortDesc = a.ShortDesc
	b.Content = utils.Map2Structpb(a.Content)
	addr, _ := a.Address.ToPb()
	b.Address = addr
	b.LegalDocs = lo.Map(a.LegalDocs, func(item biz.TicketingMedia, _ int) *pb.TicketingMedia {
		return mapBizMedia2Pb(ctx, s.blob, &item)
	})
	b.PublicContact = &pb.ContactInfo{}
	b.PublicContact.Phone = a.PublicContact.Phone
	b.PublicContact.Email = a.PublicContact.Email
	b.Rating = int32(a.Rating)
}

func (s *LocationService) MapUpdatePbLocation2Biz(a *pb.UpdateLocation, b *biz.Location) {
	b.Name = a.Name
	b.Logo = mapPbMedia2Biz(a.Logo)
	b.Medias = lo.Map(a.Medias, func(item *pb.TicketingMedia, _ int) biz.TicketingMedia {
		return *mapPbMedia2Biz(item)
	})
	b.Desc = a.Desc
	b.ShortDesc = a.ShortDesc
	b.Content = utils.Structpb2Map(a.Content)
	if a.Address != nil {
		addr, _ := lbs.NewAddressEntityFromPb(a.Address)
		b.Address = *addr
	}

	b.LegalDocs = lo.Map(a.LegalDocs, func(item *pb.TicketingMedia, _ int) biz.TicketingMedia {
		return *mapPbMedia2Biz(item)
	})
	if a.PublicContact != nil {
		b.PublicContact.Phone = a.PublicContact.Phone
		b.PublicContact.Email = a.PublicContact.Email
	}

}
func (s *LocationService) MapCreatePbLocation2Biz(a *pb.CreateLocationRequest, b *biz.Location) {
	b.Name = a.Name
	b.Logo = mapPbMedia2Biz(a.Logo)
	b.Medias = lo.Map(a.Medias, func(item *pb.TicketingMedia, _ int) biz.TicketingMedia {
		return *mapPbMedia2Biz(item)
	})
	b.Desc = a.Desc
	b.ShortDesc = a.ShortDesc
	b.Content = utils.Structpb2Map(a.Content)
	if a.Address != nil {
		addr, _ := lbs.NewAddressEntityFromPb(a.Address)
		b.Address = *addr
	}

	b.LegalDocs = lo.Map(a.LegalDocs, func(item *pb.TicketingMedia, _ int) biz.TicketingMedia {
		return *mapPbMedia2Biz(item)
	})
	if a.PublicContact != nil {
		b.PublicContact.Phone = a.PublicContact.Phone
		b.PublicContact.Email = a.PublicContact.Email
	}
}

func mapBizMedia2Pb(ctx context.Context, v vfs.Blob, a *biz.TicketingMedia) *pb.TicketingMedia {
	if a == nil {
		return nil
	}
	b := &pb.TicketingMedia{}
	mapMedia(ctx, v, a, b)
	return b
}

func mapPbMedia2Biz(a *pb.TicketingMedia) *biz.TicketingMedia {
	if a == nil {
		return nil
	}
	return &biz.TicketingMedia{
		ID:       a.Id,
		Type:     a.Type,
		MimeType: a.MimeType,
		Name:     a.Name,
	}
}

func mapMedia(ctx context.Context, v vfs.Blob, a *biz.TicketingMedia, b *pb.TicketingMedia) {
	b.Id = a.ID
	b.Type = a.Type
	b.MimeType = a.MimeType
	b.Name = a.Name
	url, _ := v.PublicUrl(ctx, b.Id)
	b.Url = url.URL
}

func mapBizHall2Pb(ctx context.Context, a *biz.Hall) *pb.LocationHall {
	b := &pb.LocationHall{}
	b.Id = a.ID.String()
	b.Name = a.Name
	b.Tags = a.Tags
	b.Capacity = int32(a.Capacity)
	return b
}

func mapPbHall2Biz(ctx context.Context, a *pb.UpdateLocationHall, b *biz.Hall) {
	b.Name = a.Name
	b.Tags = a.Tags
	b.Capacity = int(a.Capacity)
	//TODO seat groups seat

}

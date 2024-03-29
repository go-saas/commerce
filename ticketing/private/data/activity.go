package data

import (
	"context"
	"fmt"
	v1 "github.com/go-saas/commerce/ticketing/api/activity/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	sgorm "github.com/go-saas/saas/gorm"
	"github.com/goxiaoy/go-eventbus"
	"gorm.io/gorm"
)

type TicketingMediaRepo struct {
	*kitgorm.Repo[biz.TicketingMedia, string, interface{}]
}

func NewTicketingMediaRepo(dbProvider sgorm.DbProvider, eventbus *eventbus.EventBus) biz.TicketingMediaRepo {
	res := &TicketingMediaRepo{}
	res.Repo = kitgorm.NewRepo[biz.TicketingMedia, string, interface{}](dbProvider, eventbus, res)
	return res
}

func (c *TicketingMediaRepo) GetDb(ctx context.Context) *gorm.DB {
	return GetDb(ctx, c.Repo.DbProvider)
}

type ActivityRepo struct {
	*kitgorm.Repo[biz.Activity, string, v1.ListActivityRequest]
}

func NewActivityRepo(dbProvider sgorm.DbProvider, eventbus *eventbus.EventBus) biz.ActivityRepo {
	res := &ActivityRepo{}
	res.Repo = kitgorm.NewRepo[biz.Activity, string, v1.ListActivityRequest](dbProvider, eventbus, res)
	return res
}

func (c *ActivityRepo) GetDb(ctx context.Context) *gorm.DB {
	return GetDb(ctx, c.Repo.DbProvider)
}
func (c *ActivityRepo) DefaultSorting() []string {
	return []string{"created_at"}
}

// BuildDetailScope preload relations
func (c *ActivityRepo) BuildDetailScope(withDetail bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Preload("Categories").Preload("MainPic")
		if withDetail {
			db = db.Preload("Medias")
		}
		return db
	}
}

// BuildFilterScope filter
func (c *ActivityRepo) BuildFilterScope(q *v1.ListActivityRequest) func(db *gorm.DB) *gorm.DB {
	search := q.Search
	filter := q.Filter
	return func(db *gorm.DB) *gorm.DB {
		ret := db

		if search != "" {
			ret = ret.Where("name like ?", fmt.Sprintf("%%%v%%", search))
		}
		if filter == nil {
			return ret
		}

		if filter.Name != nil {
			ret = ret.Scopes(kitgorm.BuildStringFilter("`name`", filter.Name))
		}
		if filter.IsRecommend != nil {
			ret = ret.Scopes(kitgorm.BuildBooleanFilter("`is_recommend`", filter.IsRecommend))
		}
		return ret
	}
}
func (c *ActivityRepo) UpdateAssociation(ctx context.Context, entity *biz.Activity) error {
	if entity.Categories != nil {
		if err := c.GetDb(ctx).Model(entity).Session(&gorm.Session{FullSaveAssociations: true}).Association("Categories").Replace(entity.Categories); err != nil {
			return err
		}
	} else {
		if err := c.GetDb(ctx).Model(entity).Association("Categories").Clear(); err != nil {
			return err
		}
	}
	if entity.Medias != nil {
		if err := c.GetDb(ctx).Model(entity).Session(&gorm.Session{FullSaveAssociations: true}).Association("Medias").Replace(entity.Medias); err != nil {
			return err
		}
	} else {
		if err := c.GetDb(ctx).Model(entity).Association("Medias").Clear(); err != nil {
			return err
		}
	}
	return nil
}

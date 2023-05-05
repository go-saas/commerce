package data

import (
	"context"
	"fmt"
	v1 "github.com/go-saas/commerce/ticketing/api/show/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	sgorm "github.com/go-saas/saas/gorm"
	"github.com/goxiaoy/go-eventbus"
	"gorm.io/gorm"
)

type ShowRepo struct {
	*kitgorm.Repo[biz.Show, string, v1.ListShowRequest]
}

func NewShowRepo(dbProvider sgorm.DbProvider, eventbus *eventbus.EventBus) biz.ShowRepo {
	res := &ShowRepo{}
	res.Repo = kitgorm.NewRepo[biz.Show, string, v1.ListShowRequest](dbProvider, eventbus, res)
	return res
}

func (c *ShowRepo) GetDb(ctx context.Context) *gorm.DB {
	return GetDb(ctx, c.Repo.DbProvider)
}

func (c *ShowRepo) DefaultSorting() []string {
	return []string{"created_at"}
}

// BuildDetailScope preload relations
func (c *ShowRepo) BuildDetailScope(withDetail bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Preload("Activity").Preload("Activity.Categories").Preload("Activity.MainPic").Preload("Location").Preload("Hall").Preload("MainPic")
		if withDetail {
			db = db.Preload("SalesTypes").Preload("Seats")
		}
		return db
	}
}

// BuildFilterScope filter
func (c *ShowRepo) BuildFilterScope(q *v1.ListShowRequest) func(db *gorm.DB) *gorm.DB {
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
			ret = ret.Scopes(kitgorm.BuildStringFilter("`is_recommend`", filter.Name))
		}
		return ret
	}
}

func (c *ShowRepo) UpdateAssociation(ctx context.Context, entity *biz.Show) error {
	if entity.SalesTypes != nil {
		if err := c.GetDb(ctx).Model(entity).Session(&gorm.Session{FullSaveAssociations: true}).Association("SalesTypes").Replace(entity.SalesTypes); err != nil {
			return err
		}
	} else {
		if err := c.GetDb(ctx).Model(entity).Association("SalesTypes").Clear(); err != nil {
			return err
		}
	}
	return nil
}

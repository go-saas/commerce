package data

import (
	"context"
	"fmt"
	v1 "github.com/go-saas/commerce/ticketing/api/location/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	sgorm "github.com/go-saas/saas/gorm"
	"github.com/goxiaoy/go-eventbus"
	"gorm.io/gorm"
)

type LocationRepo struct {
	*kitgorm.Repo[biz.Location, string, v1.ListLocationRequest]
}

func NewLocationRepo(dbProvider sgorm.DbProvider, eventbus *eventbus.EventBus) biz.LocationRepo {
	res := &LocationRepo{}
	res.Repo = kitgorm.NewRepo[biz.Location, string, v1.ListLocationRequest](dbProvider, eventbus, res)
	return res
}

func (c *LocationRepo) GetDb(ctx context.Context) *gorm.DB {
	return GetDb(ctx, c.Repo.DbProvider)
}

// BuildDetailScope preload relations
func (c *LocationRepo) BuildDetailScope(withDetail bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// BuildFilterScope filter
func (c *LocationRepo) BuildFilterScope(q *v1.ListLocationRequest) func(db *gorm.DB) *gorm.DB {
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
		if filter.Id != nil {
			ret = ret.Scopes(kitgorm.BuildStringFilter("`id`", filter.Id))
		}
		if filter.Name != nil {
			ret = ret.Scopes(kitgorm.BuildStringFilter("`name`", filter.Name))
		}
		return ret
	}
}

func (c *LocationRepo) DefaultSorting() []string {
	return []string{"created_at"}
}

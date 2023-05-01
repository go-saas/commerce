package data

import (
	"context"
	"fmt"
	v12 "github.com/go-saas/commerce/ticketing/api/banner/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	sgorm "github.com/go-saas/saas/gorm"
	"github.com/goxiaoy/go-eventbus"
	"gorm.io/gorm"
)

type BannerRepo struct {
	*kitgorm.Repo[biz.Banner, string, v12.ListBannerRequest]
}

func NewBannerRepo(dbProvider sgorm.DbProvider, eventbus *eventbus.EventBus) biz.BannerRepo {
	res := &BannerRepo{}
	res.Repo = kitgorm.NewRepo[biz.Banner, string, v12.ListBannerRequest](dbProvider, eventbus, res)
	return res
}

func (c *BannerRepo) GetDb(ctx context.Context) *gorm.DB {
	return GetDb(ctx, c.Repo.DbProvider)
}

func (c *BannerRepo) BuildFilterScope(q *v12.ListBannerRequest) func(db *gorm.DB) *gorm.DB {
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

		if filter.Status != nil {
			ret = ret.Scopes(kitgorm.BuildStringFilter("`status`", filter.Status))
		}
		return ret
	}
}

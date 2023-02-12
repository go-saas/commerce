package data

import (
	"context"
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

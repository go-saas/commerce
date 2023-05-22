package data

import (
	"context"
	v1 "github.com/go-saas/commerce/ticketing/api/ticket/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	sgorm "github.com/go-saas/saas/gorm"
	"github.com/goxiaoy/go-eventbus"
	"gorm.io/gorm"
)

type TicketRepo struct {
	*kitgorm.Repo[biz.Ticket, string, v1.ListTicketRequest]
}

func NewTicketRepo(dbProvider sgorm.DbProvider, eventbus *eventbus.EventBus) biz.TicketRepo {
	res := &TicketRepo{}
	res.Repo = kitgorm.NewRepo[biz.Ticket, string, v1.ListTicketRequest](dbProvider, eventbus, res)
	return res
}

func (c *TicketRepo) GetDb(ctx context.Context) *gorm.DB {
	return GetDb(ctx, c.Repo.DbProvider)
}

func (c *TicketRepo) DefaultSorting() []string {
	return []string{"created_at"}
}

// BuildDetailScope preload relations
func (c *TicketRepo) BuildDetailScope(withDetail bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//db = db.Preload("Activity").Preload("Activity.Categories").Preload("Activity.MainPic").
		//	Preload("Location").Preload("Hall").Preload("MainPic").
		//	Preload("SalesTypes")
		//if withDetail {
		//	db = db.Preload("Seats")
		//}
		return db
	}
}

// BuildFilterScope filter
func (c *TicketRepo) BuildFilterScope(q *v1.ListTicketRequest) func(db *gorm.DB) *gorm.DB {
	filter := q.Filter
	return func(db *gorm.DB) *gorm.DB {
		ret := db
		if filter == nil {
			return ret
		}
		return ret
	}
}
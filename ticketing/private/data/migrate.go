package data

import (
	"context"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/saas/seed"
	"gorm.io/gorm"
)

type Migrate struct {
	data *Data
}

func NewMigrate(data *Data) *Migrate {
	return &Migrate{
		data: data,
	}
}
func (m *Migrate) Seed(ctx context.Context, sCtx *seed.Context) error {
	db := GetDb(ctx, m.data.DbProvider)
	return migrateDb(db)
}

func migrateDb(db *gorm.DB) error {
	return db.AutoMigrate(
		&biz.Activity{}, &biz.TicketingMedia{}, &biz.TicketingCategory{},
		&biz.Location{}, &biz.Hall{},
		&biz.Show{}, &biz.SeatGroup{}, &biz.Seat{}, &biz.ShowSeat{}, &biz.ShowSalesType{},
		&biz.Ticket{},
		&biz.Banner{})
}

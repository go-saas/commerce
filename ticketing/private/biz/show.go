package biz

import (
	v1 "github.com/go-saas/commerce/ticketing/api/show/v1"
	"github.com/go-saas/kit/pkg/data"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/kit/pkg/price"
	"time"
)

type Show struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel

	ActivityID string
	Activity   *Activity `gorm:"foreignKey:ActivityID"`

	StartTime time.Time
	EndTime   time.Time

	LocationID string
	Location   *Location `gorm:"foreignKey:LocationID"`

	MainPic   *TicketingMedia `gorm:"foreignKey:MainPicID"`
	MainPicID *string

	HallID string
	Hall   *Hall `gorm:"foreignKey:HallID"`

	SalesTypes []ShowSalesType `gorm:"foreignKey:ShowID"`

	Seats []ShowSeat `gorm:"foreignKey:ShowID"`

	Notice string

	IsRecommend bool
}

type ShowSalesType struct {
	kitgorm.UIDBase
	Name string

	price.Saleable

	ShowID string

	SeatGroupID *string
	SeatGroup   *SeatGroup `gorm:"foreignKey:SeatGroupID"`
}

type ShowSeat struct {
	kitgorm.UIDBase

	Seat   Seat `gorm:"foreignKey:SeatID"`
	SeatID string

	ShowID string

	Available bool

	TicketID *string
}

type ShowRepo interface {
	data.Repo[Show, string, v1.ListShowRequest]
}

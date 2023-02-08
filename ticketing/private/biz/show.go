package biz

import (
	"github.com/go-saas/commerce/pkg/price"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"time"
)

type Show struct {
	kitgorm.UIDBase

	ActivityID string
	Activity   Activity `gorm:"foreignKey:ActivityID"`

	StartTime time.Time
	EndTime   time.Time

	LocationID string
	Location   Location `gorm:"foreignKey:LocationID"`

	Halls []Hall `gorm:"many2many:show_halls;"`

	SalesType []ActivitySalesType `gorm:"foreignKey:ShowID"`

	Seats []ShowSeat `gorm:"foreignKey:ShowID"`
}

type ActivitySalesType struct {
	kitgorm.UIDBase
	Name string

	price.Saleable

	ShowID string

	SeatGroupID *string
	SeatGroup   *SeatGroup `gorm:"foreignKey:SeatGroupID"`
}

type ShowSeat struct {
	kitgorm.UIDBase

	Seat   Seat `gorm:"foreignKey:Seat"`
	SeatID string

	ShowID string

	Available bool

	TicketID *string
}

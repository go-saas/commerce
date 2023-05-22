package biz

import (
	v1 "github.com/go-saas/commerce/ticketing/api/ticket/v1"
	"github.com/go-saas/kit/pkg/data"
	"github.com/go-saas/kit/pkg/gorm"
	"github.com/lithammer/shortuuid/v3"
	"time"
)

type Ticket struct {
	gorm.AuditedModel
	ID     string `gorm:"type:char(36);primaryKey;" json:"id"`
	UserID string `gorm:"type:char(36);index:;"`

	LocationID string
	Location   Location `gorm:"foreignKey:LocationID"`

	HallID string
	Hall   Hall `gorm:"foreignKey:HallID"`

	SeatGroupID *string
	SeatGroup   *SeatGroup `gorm:"foreignKey:SeatGroupID"`

	SeatID *string
	Seat   Seat `gorm:"foreignKey:SeatID"`

	ShowSeatID *string
	ShowSeat   ShowSeat `gorm:"foreignKey:ShowSeatID"`

	ActivityID string
	Activity   Activity `gorm:"foreignKey:ActivityID"`

	ShowID string
	Show   Show `gorm:"foreignKey:ShowID"`

	ShowSalesTypeID string
	ShowSalesType   ShowSalesType `gorm:"foreignKey:ShowSalesTypeID"`
	Status          string

	OrderID  string      `gorm:"type:char(36);index:;"`
	Contact  ContactInfo `gorm:"embedded;embeddedPrefix:contact_"`
	PaidTime time.Time

	Notice string `gorm:"comment:注意事项"`
}

const (
	TicketStatusValid     = "VALID"
	TicketStatusUsed      = "USED"
	TicketStatusCancelled = "CANCELLED"
	TicketStatusExpired   = "EXPIRED"
	TicketStatusRefunded  = "REFUNDED"
)

func NewTicket(userId string) *Ticket {
	ret := &Ticket{ID: shortuuid.New()}
	return ret
}

type TicketRepo interface {
	data.Repo[Ticket, string, v1.ListTicketRequest]
}

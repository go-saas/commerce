package biz

import (
	"github.com/lithammer/shortuuid/v3"
	"time"
)

type Ticket struct {
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
	TicketStatusValid     = "valid"
	TicketStatusUsed      = "used"
	TicketStatusCancelled = "cancelled"
	TicketStatusExpired   = "expired"
)

func NewTicket() *Ticket {
	ret := &Ticket{ID: shortuuid.New()}

	return ret
}

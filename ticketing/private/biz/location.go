package biz

import (
	"context"
	v1 "github.com/go-saas/commerce/ticketing/api/location/v1"
	"github.com/go-saas/kit/pkg/data"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/lbs"
	sgorm "github.com/go-saas/saas/gorm"
)

type Location struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel
	sgorm.MultiTenancy

	Name string `json:"name"`

	LogoID *string         `json:"logo"`
	Logo   *TicketingMedia `gorm:"foreignKey:LogoID"`

	Medias []TicketingMedia `gorm:"polymorphic:Owner;polymorphicValue:location"`

	Desc      string
	ShortDesc string
	Content   data.JSONMap

	Address lbs.AddressEntity `json:"address" gorm:"embedded"`

	LegalDocs []TicketingMedia `gorm:"polymorphic:Owner;polymorphicValue:location_legal_docs"`

	PublicContact ContactInfo `gorm:"embedded;embeddedPrefix:public_contact_"`

	Rating int `gorm:"comment:评分，1-5"`

	Halls []Hall `gorm:"foreignKey:LocationID"`

	Status string
}

const (
	LocationStatusPublic = "public"
	LocationStatusClosed = "closed"
)

type Hall struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel
	Name string `json:"name"`

	LocationID string
	Location   Location

	Tags string

	Capacity int `gorm:"comment:容量"`

	Seats []Seat `gorm:"foreignKey:HallID;references:ID"`

	SeatGroups []SeatGroup `gorm:"foreignKey:HallID;references:ID"`
}

type SeatGroup struct {
	kitgorm.UIDBase
	Name   string
	HallID string
}

type Seat struct {
	kitgorm.UIDBase
	Row int
	Col int

	Group   *SeatGroup `gorm:"foreignKey:GroupID"`
	GroupID *string

	HallID string
}

type ContactInfo struct {
	Phone string
	Email string
}

type LocationRepo interface {
	data.Repo[Location, string, v1.ListLocationRequest]
	ListHalls(ctx context.Context, id string) ([]Hall, error)
	CreateHall(ctx context.Context, location *Location, entity *Hall) error
}

type HallRepo interface {
	data.Repo[Hall, string, interface{}]
}

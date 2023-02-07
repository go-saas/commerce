package biz

import (
	v1 "github.com/go-saas/commerce/ticketing/api/location/v1"
	"github.com/go-saas/kit/pkg/data"
	"github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/lbs"
)

type Location struct {
	gorm.UIDBase
	gorm.AuditedModel
	Name string `json:"name"`

	LogoID string         `json:"logo"`
	Logo   TicketingMedia `gorm:"foreignKey:LogoID"`

	Medias []TicketingMedia `gorm:"polymorphic:Owner;polymorphicValue:location"`

	Desc      string
	ShortDesc string
	Content   data.JSONMap

	Address lbs.AddressEntity `json:"address" gorm:"embedded"`

	LegalDocuments string `json:"legalDocuments"`

	PublicContact ContactInfo `gorm:"embedded;embeddedPrefix:public_contact_"`

	Rating int `gorm:"comment:评分，1-5"`

	Halls []Hall `gorm:"foreignKey:LocationID"`
}

type Hall struct {
	gorm.UIDBase
	gorm.AuditedModel
	Name string `json:"name"`

	LocationID string
	Location   Location

	Tags string

	Capacity int `gorm:"comment:容量"`
}

type ContactInfo struct {
	Phone string
	Email string
}

type LocationRepo interface {
	data.Repo[Location, string, v1.ListLocationRequest]
}

package biz

import (
	"github.com/go-saas/commerce/pkg/sortable"
	"github.com/go-saas/kit/pkg/data"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"time"
)

type Activity struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel

	Name       string
	Desc       string
	ShortDesc  string
	Content    data.JSONMap
	Medias     []TicketingMedia `gorm:"polymorphic:Owner;polymorphicValue:activity"`
	Type       string
	Categories []TicketingCategory `gorm:"many2many:activity_categories;"`
	Status     string

	SeatSelectable bool `gorm:"comment:是否可以选座"`

	Duration time.Duration `gorm:"comment:单次活动预估时间"`
}

const (
	ActivityStatusPlanned = "planned"
	ActivityStatusOngoing = "ongoing"
	ActivityStatusEnded   = "ended"
)

type TicketingMedia struct {
	ID string `gorm:"primaryKey;size:128"`

	OwnerID string
	// OwnerType product/product_sku
	OwnerType string

	Type      string
	MimeType  string
	Usage     string
	Name      string
	Reference string
	sortable.Embed
}

type TicketingMediaRepo interface {
	data.Repo[TicketingMedia, string, interface{}]
}

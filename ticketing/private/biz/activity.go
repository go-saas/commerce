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
	Title     string
	Reference string
	sortable.Embed
}

// TicketingCategory represents some Teaser infos for TicketingCategory
type TicketingCategory struct {
	// Key the identifier of the TicketingCategory
	Key string `gorm:"primaryKey;size:128"`
	// The Path (root to leaf) for this TicketingCategory - separated by "/"
	Path string
	// Name is the speaking name of the category
	Name     string
	ParentID *string
	// Parent is an optional link to parent teaser
	Parent *TicketingCategory `gorm:"foreignKey:ParentID;references:key"`
}

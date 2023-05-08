package biz

import (
	v1 "github.com/go-saas/commerce/ticketing/api/activity/v1"
	"github.com/go-saas/kit/pkg/data"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/kit/pkg/sortable"
	"time"
)

type Activity struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel

	Name      string
	Desc      string
	ShortDesc string
	Content   data.JSONMap

	MainPic   *TicketingMedia `gorm:"foreignKey:MainPicID"`
	MainPicID *string
	Medias    []TicketingMedia `gorm:"polymorphic:Owner;polymorphicValue:activity"`

	Type       string
	Categories []TicketingCategory `gorm:"many2many:activity_categories;"`
	Status     string

	SeatSelectable bool `gorm:"comment:是否可以选座"`

	Duration time.Duration `gorm:"comment:单次活动预估时间"`

	IsRecommend bool

	Notice string
}

const (
	ActivityStatusPlanned = "planned"
	ActivityStatusOngoing = "ongoing"
	ActivityStatusEnded   = "ended"
	ActivityStatusDowned  = "downed"
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
	*sortable.Embed
}

type TicketingMediaRepo interface {
	data.Repo[TicketingMedia, string, interface{}]
}

type ActivityRepo interface {
	data.Repo[Activity, string, v1.ListActivityRequest]
}

package biz

import (
	"context"
	v12 "github.com/go-saas/commerce/ticketing/api/category/v1"
	"github.com/go-saas/kit/pkg/data"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
)

// TicketingCategory represents some Teaser infos for TicketingCategory
type TicketingCategory struct {
	// Key the identifier of the TicketingCategory
	Key string `gorm:"primaryKey;size:128"`
	kitgorm.AuditedModel
	// The Path (root to leaf) for this TicketingCategory - separated by "/"
	Path string
	// Name is the speaking name of the category
	Name     string
	ParentID *string
	// Parent is an optional link to parent teaser
	Parent *TicketingCategory `gorm:"foreignKey:ParentID;references:key"`
}

type TicketingCategoryRepo interface {
	data.Repo[TicketingCategory, string, v12.ListCategoryRequest]
	FindAllChildren(ctx context.Context, entity *TicketingCategory) ([]*TicketingCategory, error)
	FindByKeys(ctx context.Context, cKeys []string) ([]TicketingCategory, error)
}

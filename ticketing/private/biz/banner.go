package biz

import (
	v12 "github.com/go-saas/commerce/ticketing/api/banner/v1"
	"github.com/go-saas/kit/pkg/data"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
)

type Banner struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel
	RefType string
	RefID   string
	Status  string `gorm:"type:char(36);index:,;comment:UNPUBLISHED,PUBLISHED"`

	MainPic   *TicketingMedia `gorm:"foreignKey:MainPicID"`
	MainPicID string
}

func NewBanner() *Banner {
	return &Banner{Status: "UNPUBLISHED"}
}

type BannerRepo interface {
	data.Repo[Banner, string, v12.ListBannerRequest]
}

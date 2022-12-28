package biz

import "github.com/go-saas/kit/pkg/gorm"

type Brand struct {
	gorm.UIDBase
	Code        string
	Name        string
	Logo        string
	Url         string
	Description string
}

type BrandMultilingual struct {
	gorm.UIDBase
	BrandId string
	Brand   Brand

	Name        string
	Url         string
	Description string
}

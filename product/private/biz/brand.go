package biz

import (
	"github.com/go-saas/commerce/pkg/multilingual"
	"github.com/go-saas/kit/pkg/gorm"
	"github.com/samber/lo"
)

type Brand struct {
	gorm.UIDBase
	Code        string
	Name        string
	Logo        string
	Url         string
	Description string

	Trans []*BrandTrans
}

type BrandTrans struct {
	gorm.UIDBase
	multilingual.Embed

	BrandId string
	Brand   Brand

	Name        string
	Url         string
	Description string
}

var _ multilingual.Multilingual = (*Brand)(nil)

func (b *Brand) GetTranslations() []interface{} {
	return lo.Map(b.Trans, func(item *BrandTrans, _ int) interface{} {
		return item
	})
}

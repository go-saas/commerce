package biz

import "github.com/go-saas/kit/pkg/gorm"

// ProductSku sku
type ProductSku struct {
	gorm.UIDBase
	ProductId string
	Product   Product

	Title string

	MainPic Media
	Medias  []Media

	Price PriceInfo `gorm:"embedded;embeddedPrefix:price_"`

	Stock []Stock

	Barcode string `gorm:"comment:商品条码"`
}

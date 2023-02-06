package biz

import kitgorm "github.com/go-saas/kit/pkg/gorm"

// ProductSku sku
type ProductSku struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel

	ProductId string
	Product   Product

	Title string

	MainPic Media   `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`
	Medias  []Media `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`

	Price PriceInfo `gorm:"embedded;embeddedPrefix:price_"`

	NeedShipping bool    `gorm:"comment:是否需要邮寄"`
	Stock        []Stock `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`

	Barcode string `gorm:"comment:商品条码"`
}

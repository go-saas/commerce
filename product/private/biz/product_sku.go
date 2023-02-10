package biz

import (
	"github.com/go-saas/commerce/pkg/price"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
)

// ProductSku sku
type ProductSku struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel

	ProductId string
	Product   Product

	Title string

	MainPic ProductMedia   `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`
	Medias  []ProductMedia `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`

	Price price.Info `gorm:"embedded;embeddedPrefix:price_"`

	Stock []Stock `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`

	Barcode string `gorm:"comment:商品条码"`
}

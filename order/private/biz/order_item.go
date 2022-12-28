package biz

import (
	"github.com/go-saas/commerce/pkg/price"
	"github.com/go-saas/commerce/product/x"
	"github.com/lithammer/shortuuid/v3"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID                 string `gorm:"type:char(36)" json:"id"`
	Qty                float64
	SinglePrice        price.Price `gorm:"embedded;embeddedPrefix:single_price_"`
	SinglePriceInclTax price.Price `gorm:"embedded;embeddedPrefix:single_price_incl_tax_"`
	RowTotal           price.Price `gorm:"embedded;embeddedPrefix:row_total_"`
	TaxAmount          price.Price `gorm:"embedded;embeddedPrefix:tax_amount_"`
	RowTotalInclTax    price.Price `gorm:"embedded;embeddedPrefix:row_total_incl_tax_"`

	Name         string
	Price        price.Price `gorm:"embedded;embeddedPrefix:price_"`
	PriceInclTax price.Price

	SourceID string

	ProductId x.ProductWithSnapshotID `gorm:"embedded;"`
}

func (u *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if len(u.ID) == 0 {
		u.ID = shortuuid.New()
	}
	return nil
}

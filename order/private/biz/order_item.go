package biz

import (
	"github.com/go-saas/kit/pkg/data"
	"github.com/go-saas/kit/pkg/price"
	"github.com/lithammer/shortuuid/v3"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID string `gorm:"type:char(36)" json:"id"`

	Qty                float64
	SinglePrice        price.Price `gorm:"embedded;embeddedPrefix:single_price_"`
	SinglePriceInclTax price.Price `gorm:"embedded;embeddedPrefix:single_price_incl_tax_"`

	RowTotal        price.Price `gorm:"embedded;embeddedPrefix:row_total_"`
	TaxAmount       price.Price `gorm:"embedded;embeddedPrefix:tax_amount_"`
	RowTotalInclTax price.Price `gorm:"embedded;embeddedPrefix:row_total_incl_tax_"`

	Name         string
	Price        price.Price `gorm:"embedded;embeddedPrefix:price_"`
	PriceInclTax price.Price `gorm:"embedded;embeddedPrefix:price_incl_tax_"`

	SourceID string

	ProductType string
	ProductId   ProductWithSnapshotID `gorm:"embedded;"`

	OrderID string

	BizPayload data.JSONMap
}

type ProductWithSnapshotID struct {
	ProductID      string `gorm:"type:char(36);index:,"`
	ProductVersion int64  `gorm:"type:char(36);index:,"`
	ProductType    string `gorm:"size:128;index:,"`
}

func (u *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if len(u.ID) == 0 {
		u.ID = shortuuid.New()
	}
	return nil
}

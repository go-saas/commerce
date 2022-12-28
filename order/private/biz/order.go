package biz

import (
	v1 "github.com/go-saas/commerce/order/api/post/v1"
	"github.com/go-saas/commerce/pkg/price"
	"github.com/go-saas/kit/pkg/data"
	sgorm "github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/lbs"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID string `gorm:"type:char(36)" json:"id"`

	Status string

	TotalPrice price.Price `gorm:"embedded;embeddedPrefix:total_price_"`
	PaidPrice  price.Price `gorm:"embedded;embeddedPrefix:paid_price_"`

	ShippingAddr *lbs.AddressEntity `gorm:"embedded;embeddedPrefix:shipping_addr_"`
	BillingAddr  lbs.AddressEntity  `gorm:"embedded;embeddedPrefix:billing_addr_"`

	sgorm.AuditedModel
	PaidTime *time.Time

	CustomerId string `gorm:"type:char(36);index:,"`

	Extra data.JSONMap
}

type OrderRepo interface {
	data.Repo[Order, string, v1.ListPostRequest]
}

func (u *Order) BeforeCreate(tx *gorm.DB) error {
	if len(u.ID) == 0 {
		u.ID = ksuid.New().String()
	}
	return nil
}

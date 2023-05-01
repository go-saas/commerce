package biz

import (
	v1 "github.com/go-saas/commerce/order/api/order/v1"
	"github.com/go-saas/kit/pkg/data"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/kit/pkg/price"
	"github.com/go-saas/lbs"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID string `gorm:"type:char(36)" json:"id"`
	kitgorm.AuditedModel

	Status string

	TotalPrice price.Price `gorm:"embedded;embeddedPrefix:total_price_"`

	PaidPrice   price.Price `gorm:"embedded;embeddedPrefix:paid_price_"`
	PaidTime    *time.Time
	PayWay      string
	PayWayExtra data.JSONMap

	ShippingAddr lbs.AddressEntity `gorm:"embedded;embeddedPrefix:shipping_addr_"`
	BillingAddr  lbs.AddressEntity `gorm:"embedded;embeddedPrefix:billing_addr_"`

	CustomerID string `gorm:"type:char(36);index:,;comment:一般等于用户ID"`

	Extra data.JSONMap

	Items []OrderItem `gorm:"foreignKey:OrderID;references:ID"`
}

type OrderRepo interface {
	data.Repo[Order, string, v1.ListOrderRequest]
}

func (u *Order) BeforeCreate(tx *gorm.DB) error {
	if len(u.ID) == 0 {
		u.ID = ksuid.New().String()
	}
	return nil
}

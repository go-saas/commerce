package price

import "time"

// Saleable embed struct for saleable item
type Saleable struct {
	IsSaleable   bool
	SaleableFrom *time.Time
	SaleableTo   *time.Time

	Price Info `gorm:"embedded;embeddedPrefix:price_"`
}

// Info embed struct for holding price info
type Info struct {
	Default      Price `gorm:"embedded;embeddedPrefix:default_"`
	Discounted   Price `gorm:"embedded;embeddedPrefix:discounted_"`
	DiscountText string

	DenyMoreDiscounts bool
}

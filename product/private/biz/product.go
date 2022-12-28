package biz

import (
	"github.com/go-saas/commerce/pkg/price"
	"github.com/go-saas/commerce/pkg/sortable"
	v1 "github.com/go-saas/commerce/product/api/post/v1"
	"github.com/go-saas/kit/pkg/data"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/saas/gorm"
	concurrency "github.com/goxiaoy/gorm-concurrency"
	"time"
)

// Product SPU
type Product struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel
	concurrency.Version
	gorm.MultiTenancy

	Title            string `gorm:"comment:商品标题"`
	ShortDescription string `gorm:"comment:商品简述"`
	Description      string `gorm:"comment:商品描述"`

	MainPic Media
	Medias  []Media

	Badges []Badge `gorm:"comment:商品徽章"`

	VisibleFrom *time.Time
	VisibleTo   *time.Time

	IsNew bool

	Categories   []Category
	MainCategory *Category `gorm:"comment:商品主要分类"`

	Barcode string `gorm:"comment:商品条码"`
	Model   string `gorm:"comment:商品型号"`

	BrandId *string
	Brand   *Brand

	Saleable
	IsGiveaway bool

	Attributes []ProductAttribute

	MultiSku bool
	Sku      []ProductSku
}

type Media struct {
	Type      string
	MimeType  string
	Usage     string
	Title     string
	Reference string
}

type Badge struct {
	kitgorm.UIDBase
	Code  string
	Label string
}

type KeyWord struct {
	kitgorm.UIDBase
	Product *Product
	Text    string
	Refer   string
}

// Category represents some Teaser infos for Category
type Category struct {
	// Code the identifier of the Category
	Code string
	// The Path (root to leaf) for this Category - separated by "/"
	Path string
	// Name is the speaking name of the category
	Name string
	// Parent is an optional link to parent teaser
	Parent *Category
}

type ProductRepo interface {
	data.Repo[Product, string, v1.ListPostRequest]
}

type Saleable struct {
	IsSaleable   bool
	SaleableFrom *time.Time
	SaleableTo   *time.Time

	NeedShipping bool `gorm:"comment:是否需要邮寄"`

	Stock   []Stock
	Keyword []KeyWord `gorm:"comment:商品关键字"`

	Price PriceInfo `gorm:"embedded;embeddedPrefix:price_"`
}

type PriceInfo struct {
	Default      price.Price `gorm:"embedded;embeddedPrefix:default_"`
	Discounted   price.Price `gorm:"embedded;embeddedPrefix:discounted_"`
	DiscountText string

	DenyMoreDiscounts bool
	CampaignRules     []CampaignRules
}

// PriceContext defines the scope in which the price was calculated
type PriceContext struct {
}

type CampaignRules struct {
	kitgorm.UIDBase
	Rule      string
	PriceInfo PriceInfo
}

type ProductAttribute struct {
	kitgorm.UIDBase
	ProductId string
	Product   Product
	Title     string
	*sortable.Embed
}

// Stock holds data with product availability info
type Stock struct {
	InStock      bool
	Level        string
	Amount       int
	DeliveryCode string
}

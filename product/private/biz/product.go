package biz

import (
	"github.com/go-saas/commerce/pkg/price"
	"github.com/go-saas/commerce/pkg/sortable"
	v1 "github.com/go-saas/commerce/product/api/product/v1"
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

	Title     string `gorm:"comment:商品标题"`
	ShortDesc string `gorm:"comment:商品简述"`
	Desc      string `gorm:"comment:商品描述"`

	MainPic Media   `gorm:"polymorphic:Owner;polymorphicValue:product"`
	Medias  []Media `gorm:"polymorphic:Owner;polymorphicValue:product"`

	Badges []Badge `gorm:"comment:商品徽章"`

	VisibleFrom *time.Time
	VisibleTo   *time.Time

	IsNew bool

	Categories     []Category `gorm:"many2many:product_categories;"`
	MainCategoryID *string
	MainCategory   *Category `gorm:"foreignKey:MainCategoryID;comment:商品主要分类"`

	Barcode string `gorm:"comment:商品条码"`
	Model   string `gorm:"comment:商品型号"`

	BrandId *string
	Brand   *Brand

	Saleable

	IsGiveaway bool `gorm:"comment:是否赠品"`

	Attributes []ProductAttribute

	MultiSku bool `gorm:"comment:是否多SKU产品,只能创建时修改"`

	Sku []ProductSku

	// CampaignRules
	CampaignRules []CampaignRule `gorm:"polymorphic:Owner;polymorphicValue:product"`

	NeedShipping bool    `gorm:"comment:是否需要邮寄"`
	Stock        []Stock `gorm:"polymorphic:Owner;polymorphicValue:product"`
}

type Media struct {
	kitgorm.UIDBase

	OwnerID string
	// OwnerType product/product_sku
	OwnerType string

	Type      string
	MimeType  string
	Usage     string
	Title     string
	Reference string
	sortable.Embed
}

type Badge struct {
	kitgorm.UIDBase
	ProductId string
	Product   *Product

	Code  string
	Label string
}

type KeyWord struct {
	kitgorm.UIDBase
	ProductId string
	Product   *Product

	Text  string
	Refer string
}

// Category represents some Teaser infos for Category
type Category struct {
	// Code the identifier of the Category
	Code string `gorm:"primaryKey;size:128"`
	// The Path (root to leaf) for this Category - separated by "/"
	Path string
	// Name is the speaking name of the category
	Name     string
	ParentID *string
	// Parent is an optional link to parent teaser
	Parent *Category `gorm:"foreignKey:ParentID;references:code"`
}

type ProductRepo interface {
	data.Repo[Product, string, v1.ListProductRequest]
}

// Saleable embed struct for saleable item
type Saleable struct {
	IsSaleable   bool
	SaleableFrom *time.Time
	SaleableTo   *time.Time

	Keyword []KeyWord `gorm:"comment:商品关键字"`

	Price PriceInfo `gorm:"embedded;embeddedPrefix:price_"`
}

// PriceInfo embed struct for holding price info
type PriceInfo struct {
	Default      price.Price `gorm:"embedded;embeddedPrefix:default_"`
	Discounted   price.Price `gorm:"embedded;embeddedPrefix:discounted_"`
	DiscountText string

	DenyMoreDiscounts bool
}

type CampaignRule struct {
	kitgorm.UIDBase
	OwnerID string
	// OwnerType product/product_sku
	OwnerType string
	Rule      string
	Extra     data.JSONMap
}

// PriceContext defines the scope in which the price was calculated
type PriceContext struct {
}

// ProductAttribute TODO how to add custom attribute
type ProductAttribute struct {
	kitgorm.UIDBase
	ProductId string
	Product   *Product

	Title string
	*sortable.Embed
}

// Stock holds data with product availability info
type Stock struct {
	kitgorm.UIDBase
	OwnerID string
	// OwnerType product/product_sku
	OwnerType    string
	InStock      bool
	Level        string
	Amount       int
	DeliveryCode string
}

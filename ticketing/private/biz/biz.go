package biz

import (
	"github.com/go-saas/kit/pkg/dal"
	kitdi "github.com/go-saas/kit/pkg/di"
)

// ProviderSet is biz providers.
var ProviderSet = kitdi.NewSet(NewPostSeeder)

const (
	LocationLogoPath           = "ticketing/location/logo"
	LocationMediaPath          = "ticketing/location/m"
	LocationLegalDocumentsPath = "ticketing/location/l"

	ActivityMediaPath = "ticketing/activity/m"
	ShowMediaPath     = "ticketing/show/m"
	BannerMediaPath   = "ticketing/banner"
)

const OrderTypeShow = "SHOW"
const ConnName dal.ConnName = "ticketing"

package biz

import (
	kitdi "github.com/go-saas/kit/pkg/di"
)

// ProviderSet is biz providers.
var ProviderSet = kitdi.NewSet(NewPostSeeder)

const (
	LocationLogoPath           = "ticketing/location/logo"
	LocationMediaPath          = "ticketing/location/m"
	LocationLegalDocumentsPath = "ticketing/location/l"
)
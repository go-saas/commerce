package api

import (
	_ "embed"
	"github.com/go-saas/kit/pkg/authz/authz"
)

const (
	ResourceLocation                        = "ticketing.location"
	ResourceShow                            = "ticketing.show"
	ResourceTicket                          = "ticketing.ticket"
	ResourceActivity                        = "ticketing.activity"
	ResourceCategory                        = "ticketing.category"
	ResourceBanner                          = "ticketing.banner"
	ActionActivityRecommend authz.ActionStr = "recommend"
)

//go:embed permission.yaml
var permission []byte

func init() {
	authz.LoadFromYaml(permission)
}

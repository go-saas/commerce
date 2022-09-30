package server

import (
	_ "embed"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"order/api"
	api2 "github.com/go-saas/kit/pkg/api"
	sapi "github.com/go-saas/kit/pkg/api"
	"github.com/go-saas/kit/pkg/authn/jwt"
	conf2 "github.com/go-saas/kit/pkg/conf"
	"github.com/go-saas/kit/pkg/localize"
	"github.com/go-saas/kit/pkg/server"
	"github.com/go-saas/saas"
	shttp "github.com/go-saas/saas/http"
	uow2 "github.com/go-saas/uow"
	kuow "github.com/go-saas/uow/kratos"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf2.Services,
	sCfg *conf2.Security,
	tokenizer jwt.Tokenizer,
	ts saas.TenantStore,
	uowMgr uow2.Manager,
	reqDecoder khttp.DecodeRequestFunc,
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	mOpt *shttp.WebMultiTenancyOption,
	apiOpt *api2.Option,
	validator sapi.TrustedContextValidator,
	register []server.HttpServiceRegister,
	logger log.Logger) *khttp.Server {
	var opts []khttp.ServerOption
	opts = server.PatchHttpOpts(logger, opts, api.ServiceName, c, sCfg, reqDecoder, resEncoder, errEncoder)
	m := []middleware.Middleware{
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		metrics.Server(),
		localize.I18N(),
		validate.Validator(),
		jwt.ServerExtractAndAuth(tokenizer, logger),
		sapi.ServerPropagation(apiOpt, validator, logger),
		server.Saas(mOpt, ts, validator),
		kuow.Uow(uowMgr, kuow.WithLogger(logger)),
	}
	opts = append(opts, []khttp.ServerOption{
		khttp.Middleware(
			m...,
		),
	}...)
	srv := khttp.NewServer(opts...)
	server.ChainHttpServiceRegister(register...).Register(srv, m...)
	return srv
}

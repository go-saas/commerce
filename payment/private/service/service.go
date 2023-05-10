package service

import (
	_ "embed"
	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/go-saas/commerce/payment/api/gateway/v1"
	"github.com/go-saas/commerce/payment/private/conf"
	kitdi "github.com/go-saas/kit/pkg/di"
	kitgrpc "github.com/go-saas/kit/pkg/server/grpc"
	kithttp "github.com/go-saas/kit/pkg/server/http"
	"github.com/stripe/stripe-go/v74/client"
	"net/http"
)

//go:embed openapi/api.swagger.json
var spec []byte

// ProviderSet is service providers.
var ProviderSet = kitdi.NewSet(
	NewGrpcServerRegister,
	NewHttpServerRegister,
	NewPaymentService,
	NewStripeClient,
)

func NewHttpServerRegister(
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	paymentSrv *PaymentService) kithttp.ServiceRegister {
	return kithttp.ServiceRegisterFunc(func(srv *khttp.Server, middleware ...middleware.Middleware) {

		v1.RegisterPaymentGatewayServiceHTTPServer(srv, paymentSrv)
		v1.RegisterStripePaymentGatewayServiceHTTPServer(srv, paymentSrv)
		swaggerRouter := chi.NewRouter()
		swaggerRouter.Use(
			kithttp.MiddlewareConvert(errEncoder, middleware...))
		const apiPrefix = "/v1/payment/dev/swagger"
		swaggerRouter.Handle(apiPrefix+"*", http.StripPrefix(apiPrefix, swaggerui.Handler(spec)))
	})
}

func NewGrpcServerRegister(
	paymentSrv *PaymentService) kitgrpc.ServiceRegister {
	return kitgrpc.ServiceRegisterFunc(func(srv *grpc.Server, middleware ...middleware.Middleware) {
		v1.RegisterPaymentGatewayServiceServer(srv, paymentSrv)
		v1.RegisterStripePaymentGatewayServiceServer(srv, paymentSrv)
	})
}

func NewStripeClient(c *conf.PaymentConf) *client.API {
	sc := &client.API{}
	sc.Init(c.GetMethodOrDefault("").Stripe.PrivateKey, nil)
	return sc
}

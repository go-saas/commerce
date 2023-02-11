package service

import (
	_ "embed"
	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/go-saas/commerce/ticketing/api/category/v1"
	v12 "github.com/go-saas/commerce/ticketing/api/location/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	kitdi "github.com/go-saas/kit/pkg/di"
	kitgrpc "github.com/go-saas/kit/pkg/server/grpc"
	kithttp "github.com/go-saas/kit/pkg/server/http"
	"github.com/goxiaoy/vfs"
	"net/http"
)

//go:embed openapi/api.swagger.json
var spec []byte

// ProviderSet is service providers.
var ProviderSet = kitdi.NewSet(
	NewGrpcServerRegister,
	NewHttpServerRegister,
	NewLocationService,
	NewTicketingCategoryService,
)

func NewHttpServerRegister(
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	vfs vfs.Blob,
	post *LocationService,
	category *TicketingCategoryService) kithttp.ServiceRegister {
	return kithttp.ServiceRegisterFunc(func(srv *khttp.Server, middleware ...middleware.Middleware) {
		v12.RegisterLocationServiceHTTPServer(srv, post)
		v1.RegisterTicketingCategoryServiceHTTPServer(srv, category)
		kithttp.MountBlob(srv, "", biz.LocationLogoPath, vfs)
		kithttp.MountBlob(srv, "", biz.LocationMediaPath, vfs)
		kithttp.MountBlob(srv, "", biz.LocationLegalDocumentsPath, vfs)

		swaggerRouter := chi.NewRouter()
		swaggerRouter.Use(
			kithttp.MiddlewareConvert(errEncoder, middleware...))
		const apiPrefix = "/v1/server/dev/swagger"
		swaggerRouter.Handle(apiPrefix+"*", http.StripPrefix(apiPrefix, swaggerui.Handler(spec)))
	})
}

func NewGrpcServerRegister(post *LocationService, category *TicketingCategoryService) kitgrpc.ServiceRegister {
	return kitgrpc.ServiceRegisterFunc(func(srv *grpc.Server, middleware ...middleware.Middleware) {
		v12.RegisterLocationServiceServer(srv, post)
		v1.RegisterTicketingCategoryServiceServer(srv, category)
	})
}

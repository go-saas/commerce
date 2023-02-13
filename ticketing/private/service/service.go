package service

import (
	_ "embed"
	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	v13 "github.com/go-saas/commerce/ticketing/api/activity/v1"
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
	NewActivityService,
)

func NewHttpServerRegister(
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	vfs vfs.Blob,
	location *LocationService,
	category *TicketingCategoryService,
	activity *ActivityService) kithttp.ServiceRegister {
	return kithttp.ServiceRegisterFunc(func(srv *khttp.Server, middleware ...middleware.Middleware) {
		v12.RegisterLocationServiceHTTPServer(srv, location)
		v1.RegisterTicketingCategoryServiceHTTPServer(srv, category)
		v13.RegisterActivityServiceHTTPServer(srv, activity)
		kithttp.MountBlob(srv, "", biz.LocationLogoPath, vfs)
		kithttp.MountBlob(srv, "", biz.LocationMediaPath, vfs)
		kithttp.MountBlob(srv, "", biz.LocationLegalDocumentsPath, vfs)

		route := srv.Route("/")

		route.POST("/v1/ticketing/location/logo", location.UploadLogo)
		route.POST("/v1/ticketing/location/media", location.UploadMedias)
		route.POST("/v1/ticketing/location/legal-docs", location.UploadLegalDocs)

		route.POST("/v1/ticketing/activity/media", activity.UploadMedias)

		swaggerRouter := chi.NewRouter()
		swaggerRouter.Use(
			kithttp.MiddlewareConvert(errEncoder, middleware...))
		const apiPrefix = "/v1/server/dev/swagger"
		swaggerRouter.Handle(apiPrefix+"*", http.StripPrefix(apiPrefix, swaggerui.Handler(spec)))
	})
}

func NewGrpcServerRegister(locationSrv *LocationService, category *TicketingCategoryService, activity *ActivityService) kitgrpc.ServiceRegister {
	return kitgrpc.ServiceRegisterFunc(func(srv *grpc.Server, middleware ...middleware.Middleware) {
		v12.RegisterLocationServiceServer(srv, locationSrv)
		v1.RegisterTicketingCategoryServiceServer(srv, category)
		v13.RegisterActivityServiceServer(srv, activity)
	})
}

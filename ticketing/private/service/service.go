package service

import (
	_ "embed"
	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	v13 "github.com/go-saas/commerce/ticketing/api/activity/v1"
	v16 "github.com/go-saas/commerce/ticketing/api/banner/v1"
	v1 "github.com/go-saas/commerce/ticketing/api/category/v1"
	v12 "github.com/go-saas/commerce/ticketing/api/location/v1"
	v14 "github.com/go-saas/commerce/ticketing/api/show/v1"
	v15 "github.com/go-saas/commerce/ticketing/api/ticket/v1"
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
	NewShowService,
	NewTicketServiceService,
	NewTicketingBannerServiceService,
	NewUploadService,
)

func NewHttpServerRegister(
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	vfs vfs.Blob,
	location *LocationService,
	category *TicketingCategoryService,
	activity *ActivityService,
	show *ShowService,
	ticket *TicketServiceService,
	banner *TicketingBannerServiceService,
) kithttp.ServiceRegister {
	return kithttp.ServiceRegisterFunc(func(srv *khttp.Server, middleware ...middleware.Middleware) {
		v12.RegisterLocationServiceHTTPServer(srv, location)

		v1.RegisterTicketingCategoryServiceHTTPServer(srv, category)
		v1.RegisterTicketingCategoryAppServiceHTTPServer(srv, category)

		v13.RegisterActivityServiceHTTPServer(srv, activity)
		v13.RegisterActivityAppServiceHTTPServer(srv, activity)

		v14.RegisterShowServiceHTTPServer(srv, show)
		v14.RegisterShowAppServiceHTTPServer(srv, show)

		v15.RegisterTicketServiceHTTPServer(srv, ticket)
		v15.RegisterTicketAppServiceHTTPServer(srv, ticket)

		v16.RegisterTicketingBannerServiceHTTPServer(srv, banner)
		v16.RegisterTicketingAppBannerServiceHTTPServer(srv, banner)
		route := srv.Route("/")

		route.POST("/v1/ticketing/location/logo", location.UploadLogo)
		route.POST("/v1/ticketing/location/media", location.UploadMedias)
		route.POST("/v1/ticketing/location/legal-docs", location.UploadLegalDocs)

		route.POST("/v1/ticketing/activity/media", activity.UploadMedias)
		route.POST("/v1/ticketing/banner/media", banner.UploadMedias)
		route.POST("/v1/ticketing/show/media", show.UploadMedias)

		kithttp.MountBlob(srv, "", biz.LocationLogoPath, vfs)
		kithttp.MountBlob(srv, "", biz.LocationMediaPath, vfs)
		kithttp.MountBlob(srv, "", biz.LocationLegalDocumentsPath, vfs)
		kithttp.MountBlob(srv, "", biz.ActivityMediaPath, vfs)
		kithttp.MountBlob(srv, "", biz.BannerMediaPath, vfs)
		kithttp.MountBlob(srv, "", biz.ShowMediaPath, vfs)
		swaggerRouter := chi.NewRouter()
		swaggerRouter.Use(
			kithttp.MiddlewareConvert(errEncoder, middleware...))
		const apiPrefix = "/v1/ticketing/dev/swagger"
		swaggerRouter.Handle(apiPrefix+"*", http.StripPrefix(apiPrefix, swaggerui.Handler(spec)))
	})
}

func NewGrpcServerRegister(
	locationSrv *LocationService,
	category *TicketingCategoryService,
	activity *ActivityService,
	show *ShowService,
	ticket *TicketServiceService,
	banner *TicketingBannerServiceService,
) kitgrpc.ServiceRegister {
	return kitgrpc.ServiceRegisterFunc(func(srv *grpc.Server, middleware ...middleware.Middleware) {
		v12.RegisterLocationServiceServer(srv, locationSrv)
		v1.RegisterTicketingCategoryServiceServer(srv, category)
		v1.RegisterTicketingCategoryAppServiceServer(srv, category)

		v13.RegisterActivityServiceServer(srv, activity)
		v13.RegisterActivityAppServiceServer(srv, activity)

		v14.RegisterShowServiceServer(srv, show)
		v14.RegisterShowAppServiceServer(srv, show)

		v15.RegisterTicketServiceServer(srv, ticket)
		v15.RegisterTicketAppServiceServer(srv, ticket)

		v16.RegisterTicketingBannerServiceServer(srv, banner)
		v16.RegisterTicketingAppBannerServiceServer(srv, banner)
	})
}

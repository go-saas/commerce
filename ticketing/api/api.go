package api

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	grpc2 "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/go-saas/commerce/ticketing/api/location/v1"
	_ "github.com/go-saas/commerce/ticketing/i18n"
	"github.com/go-saas/kit/pkg/api"
	"github.com/go-saas/kit/pkg/conf"
	kitdi "github.com/go-saas/kit/pkg/di"
	"google.golang.org/grpc"
)

type GrpcConn grpc.ClientConnInterface
type HttpClient *http.Client

const ServiceName = "ticketing"

func NewGrpcConn(client *conf.Client, services *conf.Services, dis registry.Discovery, opt *api.Option, tokenMgr api.TokenManager, logger log.Logger, opts []grpc2.ClientOption) (GrpcConn, func()) {
	return api.NewGrpcConn(client, ServiceName, services, dis, opt, tokenMgr, logger, opts)
}

func NewHttpClient(client *conf.Client, services *conf.Services, dis registry.Discovery, opt *api.Option, tokenMgr api.TokenManager, logger log.Logger, opts []http.ClientOption) (HttpClient, func()) {
	return api.NewHttpClient(client, ServiceName, services, dis, opt, tokenMgr, logger, opts)
}

var GrpcProviderSet = kitdi.NewSet(NewGrpcConn, NewLocationGrpcClient)
var HttpProviderSet = kitdi.NewSet(NewHttpClient, NewLocationHttpClient)

func NewLocationGrpcClient(conn GrpcConn) v1.LocationServiceServer {
	return v1.NewLocationServiceClientProxy(v1.NewLocationServiceClient(conn))
}

func NewLocationHttpClient(http HttpClient) v1.LocationServiceHTTPClient {
	return v1.NewLocationServiceHTTPClient(http)
}

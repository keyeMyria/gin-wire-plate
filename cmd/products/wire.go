//go:build wireinject
// +build wireinject

package main

import (
	"gin-wire-plate/internal/app/products"
	"gin-wire-plate/internal/app/products/controllers"
	"gin-wire-plate/internal/app/products/grpcclients"
	"gin-wire-plate/internal/app/products/services"
	"gin-wire-plate/internal/pkg/app"
	"gin-wire-plate/internal/pkg/config"
	"gin-wire-plate/internal/pkg/consul"
	"gin-wire-plate/internal/pkg/jaeger"
	"gin-wire-plate/internal/pkg/log"
	"gin-wire-plate/internal/pkg/transports/grpc"
	"gin-wire-plate/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	grpcclients.ProviderSet,
	controllers.ProviderSet,
	services.ProviderSet,
	products.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}

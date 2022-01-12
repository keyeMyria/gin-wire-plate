//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wahello/gin-wire-plate/internal/app/products"
	"github.com/wahello/gin-wire-plate/internal/app/products/controllers"
	"github.com/wahello/gin-wire-plate/internal/app/products/grpcclients"
	"github.com/wahello/gin-wire-plate/internal/app/products/services"
	"github.com/wahello/gin-wire-plate/internal/pkg/app"
	"github.com/wahello/gin-wire-plate/internal/pkg/config"
	"github.com/wahello/gin-wire-plate/internal/pkg/consul"
	"github.com/wahello/gin-wire-plate/internal/pkg/jaeger"
	"github.com/wahello/gin-wire-plate/internal/pkg/log"
	"github.com/wahello/gin-wire-plate/internal/pkg/transports/grpc"
	"github.com/wahello/gin-wire-plate/internal/pkg/transports/http"
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

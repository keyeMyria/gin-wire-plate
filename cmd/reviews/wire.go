//go:build wireinject
// +build wireinject

package main

import (
	"gin-wire-plate/internal/app/reviews"
	"gin-wire-plate/internal/app/reviews/controllers"
	"gin-wire-plate/internal/app/reviews/grpcservers"
	"gin-wire-plate/internal/app/reviews/repositories"
	"gin-wire-plate/internal/app/reviews/services"
	"gin-wire-plate/internal/pkg/app"
	"gin-wire-plate/internal/pkg/config"
	"gin-wire-plate/internal/pkg/consul"
	"gin-wire-plate/internal/pkg/database"
	"gin-wire-plate/internal/pkg/jaeger"
	"gin-wire-plate/internal/pkg/log"
	"gin-wire-plate/internal/pkg/transports/grpc"
	"gin-wire-plate/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	reviews.ProviderSet,
	repositories.ProviderSet,
	controllers.ProviderSet,
	grpcservers.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}

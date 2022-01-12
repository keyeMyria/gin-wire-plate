//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wahello/gin-wire-plate/internal/app/ratings"
	"github.com/wahello/gin-wire-plate/internal/app/ratings/controllers"
	"github.com/wahello/gin-wire-plate/internal/app/ratings/grpcservers"
	"github.com/wahello/gin-wire-plate/internal/app/ratings/repositories"
	"github.com/wahello/gin-wire-plate/internal/app/ratings/services"
	"github.com/wahello/gin-wire-plate/internal/pkg/app"
	"github.com/wahello/gin-wire-plate/internal/pkg/config"
	"github.com/wahello/gin-wire-plate/internal/pkg/consul"
	"github.com/wahello/gin-wire-plate/internal/pkg/jaeger"

	"github.com/wahello/gin-wire-plate/internal/pkg/database"
	"github.com/wahello/gin-wire-plate/internal/pkg/log"
	"github.com/wahello/gin-wire-plate/internal/pkg/transports/grpc"
	"github.com/wahello/gin-wire-plate/internal/pkg/transports/http"
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
	ratings.ProviderSet,
	repositories.ProviderSet,
	controllers.ProviderSet,
	grpcservers.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}

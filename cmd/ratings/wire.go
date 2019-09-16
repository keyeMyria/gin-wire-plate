// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/internal/app/ratings"
	"github.com/keyeMyria/gin-wire-plate/internal/app/ratings/controllers"
	"github.com/keyeMyria/gin-wire-plate/internal/app/ratings/grpcservers"
	"github.com/keyeMyria/gin-wire-plate/internal/app/ratings/services"
	"github.com/keyeMyria/gin-wire-plate/internal/app/ratings/repositories"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/app"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/config"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/consul"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/jaeger"

	"github.com/keyeMyria/gin-wire-plate/internal/pkg/database"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/log"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/transports/grpc"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/transports/http"
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

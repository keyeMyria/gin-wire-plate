// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/internal/app/reviews"
	"github.com/keyeMyria/gin-wire-plate/internal/app/reviews/controllers"
	"github.com/keyeMyria/gin-wire-plate/internal/app/reviews/grpcservers"
	"github.com/keyeMyria/gin-wire-plate/internal/app/reviews/services"
	"github.com/keyeMyria/gin-wire-plate/internal/app/reviews/repositories"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/app"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/config"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/consul"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/database"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/jaeger"
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
	reviews.ProviderSet,
	repositories.ProviderSet,
	controllers.ProviderSet,
	grpcservers.ProviderSet,
)


func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}

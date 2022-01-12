// +build wireinject

package grpcservers

import (
	"github.com/google/wire"
	"github.com/wahello/gin-wire-plate/internal/pkg/config"
	"github.com/wahello/gin-wire-plate/internal/pkg/database"
	"github.com/wahello/gin-wire-plate/internal/pkg/log"
	"github.com/wahello/gin-wire-plate/internal/app/ratings/services"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateRatingsServer(cf string, service services.RatingsService) (*RatingsServer, error) {
	panic(wire.Build(testProviderSet))
}
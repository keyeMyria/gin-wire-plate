// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/config"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/database"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/log"
	"github.com/keyeMyria/gin-wire-plate/internal/app/ratings/repositories"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateRatingsService(cf string, sto repositories.RatingsRepository) (RatingsService, error) {
	panic(wire.Build(testProviderSet))
}

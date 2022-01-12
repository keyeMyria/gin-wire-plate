//go:build wireinject
// +build wireinject

package controllers

import (
	"gin-wire-plate/internal/app/ratings/repositories"
	"gin-wire-plate/internal/app/ratings/services"
	"gin-wire-plate/internal/pkg/config"
	"gin-wire-plate/internal/pkg/database"
	"gin-wire-plate/internal/pkg/log"

	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)

func CreateRatingsController(cf string, sto repositories.RatingsRepository) (*RatingsController, error) {
	panic(wire.Build(testProviderSet))
}

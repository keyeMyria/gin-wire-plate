//go:build wireinject
// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/wahello/gin-wire-plate/internal/app/ratings/repositories"
	"github.com/wahello/gin-wire-plate/internal/app/ratings/services"
	"github.com/wahello/gin-wire-plate/internal/pkg/config"
	"github.com/wahello/gin-wire-plate/internal/pkg/database"
	"github.com/wahello/gin-wire-plate/internal/pkg/log"
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

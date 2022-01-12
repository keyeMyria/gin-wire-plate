//go:build wireinject
// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/wahello/gin-wire-plate/internal/app/details/repositories"
	"github.com/wahello/gin-wire-plate/internal/app/details/services"
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

func CreateDetailsController(cf string, sto repositories.DetailsRepository) (*DetailsController, error) {
	panic(wire.Build(testProviderSet))
}

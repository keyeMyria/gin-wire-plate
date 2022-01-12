//go:build wireinject
// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/wahello/gin-wire-plate/internal/app/details/repositories"
	"github.com/wahello/gin-wire-plate/internal/pkg/config"
	"github.com/wahello/gin-wire-plate/internal/pkg/database"
	"github.com/wahello/gin-wire-plate/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateDetailsService(cf string, sto repositories.DetailsRepository) (DetailsService, error) {
	panic(wire.Build(testProviderSet))
}

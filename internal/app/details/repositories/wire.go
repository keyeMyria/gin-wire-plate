//go:build wireinject
// +build wireinject

package repositories

import (
	"github.com/google/wire"

	"gin-wire-plate/internal/pkg/config"
	"gin-wire-plate/internal/pkg/database"
	"gin-wire-plate/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateDetailRepository(f string) (DetailsRepository, error) {
	panic(wire.Build(testProviderSet))
}

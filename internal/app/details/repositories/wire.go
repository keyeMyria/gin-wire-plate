//go:build wireinject
// +build wireinject

package repositories

import (
	"github.com/google/wire"
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

func CreateDetailRepository(f string) (DetailsRepository, error) {
	panic(wire.Build(testProviderSet))
}

//go:build wireinject
// +build wireinject

package repositories

import (
	"gin-wire-plate/internal/pkg/config"
	"gin-wire-plate/internal/pkg/database"
	"gin-wire-plate/internal/pkg/log"

	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateReviewRepository(f string) (ReviewsRepository, error) {
	panic(wire.Build(testProviderSet))
}

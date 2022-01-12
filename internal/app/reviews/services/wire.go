//go:build wireinject
// +build wireinject

package services

import (
	"gin-wire-plate/internal/app/reviews/repositories"
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

func CreateReviewsService(cf string, sto repositories.ReviewsRepository) (ReviewsService, error) {
	panic(wire.Build(testProviderSet))
}

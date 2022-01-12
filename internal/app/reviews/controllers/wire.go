//go:build wireinject
// +build wireinject

package controllers

import (
	"gin-wire-plate/internal/app/reviews/repositories"
	"gin-wire-plate/internal/app/reviews/services"
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

func CreateReviewsController(cf string, sto repositories.ReviewsRepository) (*ReviewsController, error) {
	panic(wire.Build(testProviderSet))
}

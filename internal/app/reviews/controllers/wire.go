//go:build wireinject
// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/wahello/gin-wire-plate/internal/app/reviews/repositories"
	"github.com/wahello/gin-wire-plate/internal/app/reviews/services"
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

func CreateReviewsController(cf string, sto repositories.ReviewsRepository) (*ReviewsController, error) {
	panic(wire.Build(testProviderSet))
}

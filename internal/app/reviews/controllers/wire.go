// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/config"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/database"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/log"
	"github.com/keyeMyria/gin-wire-plate/internal/app/reviews/services"
	"github.com/keyeMyria/gin-wire-plate/internal/app/reviews/repositories"
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

//go:build wireinject
// +build wireinject

package grpcservers

import (
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
	ProviderSet,
)

func CreateReviewsServer(cf string, service services.ReviewsService) (*ReviewsServer, error) {
	panic(wire.Build(testProviderSet))
}

//go:build wireinject
// +build wireinject

package services

import (
	"gin-wire-plate/api/proto"
	"gin-wire-plate/internal/pkg/config"
	"gin-wire-plate/internal/pkg/log"

	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	ProviderSet,
)

func CreateProductsService(cf string,
	detailsSvc proto.DetailsClient,
	ratingsSvc proto.RatingsClient,
	reviewsSvc proto.ReviewsClient) (ProductsService, error) {
	panic(wire.Build(testProviderSet))
}

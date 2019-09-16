// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/api/proto"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/config"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/log"
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

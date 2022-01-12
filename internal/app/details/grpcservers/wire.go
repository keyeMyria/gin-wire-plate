//go:build wireinject
// +build wireinject

package grpcservers

import (
	"github.com/google/wire"

	"gin-wire-plate/internal/app/details/services"
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

func CreateDetailsServer(cf string, service services.DetailsService) (*DetailsServer, error) {
	panic(wire.Build(testProviderSet))
}

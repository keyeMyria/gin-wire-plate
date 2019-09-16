// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/config"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/database"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/log"
	"github.com/keyeMyria/gin-wire-plate/internal/app/details/repositories"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateDetailsService(cf string, sto repositories.DetailsRepository) (DetailsService, error) {
	panic(wire.Build(testProviderSet))
}

package reviews

import (
	"gin-wire-plate/internal/pkg/app"
	"gin-wire-plate/internal/pkg/transports/grpc"
	"gin-wire-plate/internal/pkg/transports/http"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Options struct {
	Name string
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, hs *http.Server, gs *grpc.Server) (*app.Application, error) {
	a, err := app.New(o.Name, logger, app.GrpcServerOption(gs), app.HttpServerOption(hs))

	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)

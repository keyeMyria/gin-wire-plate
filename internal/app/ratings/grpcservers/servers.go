package grpcservers

import (
	"gin-wire-plate/api/proto"
	"gin-wire-plate/internal/pkg/transports/grpc"

	"github.com/google/wire"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	ps *RatingsServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterRatingsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewRatingsServer, CreateInitServersFn)

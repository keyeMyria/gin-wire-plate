package grpcservers

import (
	"gin-wire-plate/api/proto"
	"gin-wire-plate/internal/pkg/transports/grpc"
	"github.com/google/wire"
	stdgrpc "google.golang.org/grpc"
)

func CreateInitServersFn(
	ps *DetailsServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterDetailsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewDetailsServer, CreateInitServersFn)

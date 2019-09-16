package grpcservers

import (
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/api/proto"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/transports/grpc"
	stdgrpc "google.golang.org/grpc"
)



func CreateInitServersFn(
	ps *ReviewsServer,
) grpc.InitServers {
	return func(s *stdgrpc.Server) {
		proto.RegisterReviewsServer(s, ps)
	}
}

var ProviderSet = wire.NewSet(NewReviewsServer, CreateInitServersFn)
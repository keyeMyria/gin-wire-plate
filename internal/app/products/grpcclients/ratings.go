package grpcclients

import (
	"gin-wire-plate/api/pb"
	"gin-wire-plate/internal/pkg/transports/grpc"

	"github.com/pkg/errors"
)

func NewRatingsClient(client *grpc.Client) (pb.RatingsClient, error) {
	conn, err := client.Dial("Ratings")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := pb.NewRatingsClient(conn)

	return c, nil
}

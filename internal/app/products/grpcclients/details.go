package grpcclients

import (
	"gin-wire-plate/api/pb"
	"gin-wire-plate/internal/pkg/transports/grpc"

	"github.com/pkg/errors"
)

func NewDetailsClient(client *grpc.Client) (pb.DetailsClient, error) {
	conn, err := client.Dial("Details")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := pb.NewDetailsClient(conn)

	return c, nil
}

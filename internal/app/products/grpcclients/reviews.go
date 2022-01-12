package grpcclients

import (
	"gin-wire-plate/api/proto"
	"gin-wire-plate/internal/pkg/transports/grpc"

	"github.com/pkg/errors"
)

func NewReviewsClient(client *grpc.Client) (proto.ReviewsClient, error) {
	conn, err := client.Dial("Reviews")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewReviewsClient(conn)

	return c, nil
}

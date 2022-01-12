package grpcclients

import (
	"gin-wire-plate/api/pb"
	"gin-wire-plate/internal/pkg/transports/grpc"

	"github.com/pkg/errors"
)

func NewReviewsClient(client *grpc.Client) (pb.ReviewsClient, error) {
	conn, err := client.Dial("Reviews")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := pb.NewReviewsClient(conn)

	return c, nil
}

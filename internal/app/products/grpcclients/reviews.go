package grpcclients

import (
	"github.com/pkg/errors"
	"github.com/wahello/gin-wire-plate/api/proto"
	"github.com/wahello/gin-wire-plate/internal/pkg/transports/grpc"
)

func NewReviewsClient(client *grpc.Client) (proto.ReviewsClient, error) {
	conn, err := client.Dial("Reviews")
	if err != nil {
		return nil, errors.Wrap(err, "detail client dial error")
	}
	c := proto.NewReviewsClient(conn)

	return c, nil
}

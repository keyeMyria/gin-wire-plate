package grpcclients

import (
	"github.com/pkg/errors"
	"github.com/keyeMyria/gin-wire-plate/api/proto"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/transports/grpc"
)

func NewDetailsClient(client *grpc.Client) (proto.DetailsClient, error) {
	conn,err := client.Dial("Details")
	if err != nil {
		return nil,errors.Wrap(err,"detail client dial error")
	}
	c := proto.NewDetailsClient(conn)

	return c, nil
}

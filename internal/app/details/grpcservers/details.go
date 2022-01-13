package grpcservers

import (
	"context"

	"gin-wire-plate/api/pb"
	"gin-wire-plate/internal/app/details/services"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type DetailsServer struct {
	logger  *zap.Logger
	service services.DetailsService
}

func NewDetailsServer(logger *zap.Logger, ps services.DetailsService) (*DetailsServer, error) {
	return &DetailsServer{
		logger:  logger,
		service: ps,
	}, nil
}

func (s *DetailsServer) Get(ctx context.Context, req *pb.GetDetailRequest) (*pb.Detail, error) {
	p, err := s.service.Get(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "details grpc service get detail error")
	}
	ct := timestamppb.New(p.CreatedTime)

	resp := &pb.Detail{
		Id:          uint64(p.ID),
		Name:        p.Name,
		Price:       p.Price,
		CreatedTime: ct,
	}

	return resp, nil
}

package grpcservers

import (
	"context"
	"gin-wire-plate/api/proto"
	"gin-wire-plate/internal/app/ratings/services"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type RatingsServer struct {
	logger  *zap.Logger
	service services.RatingsService
}

func NewRatingsServer(logger *zap.Logger, ps services.RatingsService) (*RatingsServer, error) {
	return &RatingsServer{
		logger:  logger,
		service: ps,
	}, nil
}

func (s *RatingsServer) Get(ctx context.Context, req *proto.GetRatingRequest) (*proto.Rating, error) {
	r, err := s.service.Get(req.ProductID)
	if err != nil {
		return nil, errors.Wrap(err, "product grpc service get rating error")
	}
	ut, err := ptypes.TimestampProto(r.UpdatedTime)
	if err != nil {
		return nil, errors.Wrap(err, "convert create time error")
	}

	resp := &proto.Rating{
		Id:          uint64(r.ID),
		ProductID:   r.ProductID,
		Score:       r.Score,
		UpdatedTime: ut,
	}

	return resp, nil
}

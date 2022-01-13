package grpcservers

import (
	"context"

	"gin-wire-plate/api/pb"
	"gin-wire-plate/internal/app/ratings/services"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

func (s *RatingsServer) Get(ctx context.Context, req *pb.GetRatingRequest) (*pb.Rating, error) {
	r, err := s.service.Get(req.ProductID)
	if err != nil {
		return nil, errors.Wrap(err, "product grpc service get rating error")
	}
	ut := timestamppb.New(r.UpdatedTime)

	resp := &pb.Rating{
		Id:          uint64(r.ID),
		ProductID:   r.ProductID,
		Score:       r.Score,
		UpdatedTime: ut,
	}

	return resp, nil
}

package grpcservers

import (
	"context"
	"gin-wire-plate/api/pb"
	"gin-wire-plate/internal/app/reviews/services"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type ReviewsServer struct {
	logger  *zap.Logger
	service services.ReviewsService
}

func NewReviewsServer(logger *zap.Logger, ps services.ReviewsService) (*ReviewsServer, error) {
	return &ReviewsServer{
		logger:  logger,
		service: ps,
	}, nil
}

func (s *ReviewsServer) Query(ctx context.Context, req *pb.QueryReviewsRequest) (*pb.QueryReviewsResponse, error) {
	rs, err := s.service.Query(req.ProductID)
	if err != nil {
		return nil, errors.Wrap(err, "reviews grpc service get reviews error")
	}

	resp := &pb.QueryReviewsResponse{
		Reviews: make([]*pb.Review, 0, len(rs)),
	}
	for _, r := range rs {
		ct := timestamppb.New(r.CreatedTime)

		pr := &pb.Review{
			Id:          uint64(r.ID),
			ProductID:   r.ProductID,
			Message:     r.Message,
			CreatedTime: ct,
		}

		resp.Reviews = append(resp.Reviews, pr)
	}

	return resp, nil
}

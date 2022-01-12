package services

import (
	"gin-wire-plate/internal/app/ratings/repositories"
	"gin-wire-plate/internal/pkg/models"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type RatingsService interface {
	Get(ID uint64) (*models.Rating, error)
}

type DefaultRatingsService struct {
	logger     *zap.Logger
	Repository repositories.RatingsRepository
}

func NewRatingService(logger *zap.Logger, Repository repositories.RatingsRepository) RatingsService {
	return &DefaultRatingsService{
		logger:     logger.With(zap.String("type", "DefaultRatingsService")),
		Repository: Repository,
	}
}

func (s *DefaultRatingsService) Get(productID uint64) (p *models.Rating, err error) {
	if p, err = s.Repository.Get(productID); err != nil {
		return nil, errors.Wrap(err, "get rating error")
	}

	return
}

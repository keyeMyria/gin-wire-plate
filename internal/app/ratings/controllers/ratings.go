package controllers

import (
	"gin-wire-plate/internal/app/ratings/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RatingsController struct {
	logger  *zap.Logger
	service services.RatingsService
}

func NewRatingsController(logger *zap.Logger, s services.RatingsService) *RatingsController {
	return &RatingsController{
		logger:  logger,
		service: s,
	}
}

func (pc *RatingsController) Get(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("productID"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p, err := pc.service.Get(ID)
	if err != nil {
		pc.logger.Error("get rating by productID error", zap.Error(err))
		c.String(http.StatusInternalServerError, "%+v", err)
		return
	}

	c.JSON(http.StatusOK, p)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wahello/gin-wire-plate/internal/app/details/services"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type DetailsController struct {
	logger  *zap.Logger
	service services.DetailsService
}

func NewDetailsController(logger *zap.Logger, s services.DetailsService) *DetailsController {
	return &DetailsController{
		logger:  logger,
		service: s,
	}
}

func (pc *DetailsController) Get(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p, err := pc.service.Get(ID)
	if err != nil {
		pc.logger.Error("get product by id error", zap.Error(err))
		c.String(http.StatusInternalServerError, "%+v", err)
		return
	}

	c.JSON(http.StatusOK, p)
}

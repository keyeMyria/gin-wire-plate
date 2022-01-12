package controllers

import (
	"gin-wire-plate/internal/app/products/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductsController struct {
	logger  *zap.Logger
	service services.ProductsService
}

func NewProductsController(logger *zap.Logger, s services.ProductsService) *ProductsController {
	return &ProductsController{
		logger:  logger,
		service: s,
	}
}

func (pc *ProductsController) Get(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p, err := pc.service.Get(c.Request.Context(), ID)
	if err != nil {
		pc.logger.Error("get product by id error", zap.Error(err))
		c.String(http.StatusInternalServerError, "%+v", err)
		return
	}

	c.JSON(http.StatusOK, p)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/transports/http"
)

func CreateInitControllersFn(
	pc *ProductsController,
) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/product/:id", pc.Get)
	}
}

var ProviderSet = wire.NewSet(NewProductsController, CreateInitControllersFn)

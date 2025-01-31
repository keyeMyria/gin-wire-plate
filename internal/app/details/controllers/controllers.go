package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"gin-wire-plate/internal/pkg/transports/http"
)

func CreateInitControllersFn(
	pc *DetailsController,
) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/detail/:id", pc.Get)
	}
}

var ProviderSet = wire.NewSet(NewDetailsController, CreateInitControllersFn)

package router

import (
	"go-challenge/internal/application/usecase"
	"go-challenge/internal/infrastructure/query"
	"go-challenge/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "OK"})
	})

	activeEVSELocationQueryServiceGorm := query.NewActiveEVSELocationQueryServiceGorm()
	activeEVSELocationUseCase := usecase.NewActiveEVSELocationUseCase(activeEVSELocationQueryServiceGorm)
	activeEVSELocationController := controllers.NewActiveEVSELocationController(activeEVSELocationUseCase)

	r.GET("/api/locations", activeEVSELocationController.FetchActiveEVSELocations)

	return r
}

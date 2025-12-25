package router

import (
	"go-challenge/database"
	"go-challenge/handler"
	"go-challenge/repository"
	"go-challenge/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "OK"})
	})

	repo := repository.NewLocationRepository(database.GetDB())
	svc := service.NewLocationService(repo)
	locationHandler := handler.NewLocationHandler(svc)

	api := r.Group("/api")
	{
		api.GET("/locations", locationHandler.GetLocations)
	}

	return r
}

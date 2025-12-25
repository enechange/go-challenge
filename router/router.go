package router

import (
	"go-challenge/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "OK"})
	})

	locationHandler := handler.NewLocationHandler()
	api := r.Group("/api")
	{
		api.GET("/locations", locationHandler.GetLocations)
	}

	return r
}

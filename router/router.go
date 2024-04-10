package router

import (
	"go-challenge/models"
	"go-challenge/repositories"
	"go-challenge/services"
	"net/http"

	"github.com/gin-gonic/gin"
)
var locationService services.ILocationService

func InitLocationService() {
	locationService = &services.LocationService{
		LocationRepository: &repositories.LocationRepository{},
	}
}

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "OK"})
	})
	r.GET("/api/locations", func(c *gin.Context) {
		var locationQuery models.LocationSearchParameters
		if err := c.ShouldBindQuery(&locationQuery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		locations, err := locationService.GetNearbyEvses(float32(locationQuery.Longitude), float32(locationQuery.Latitude), locationQuery.Radius)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, locations)
	})

	return r
}

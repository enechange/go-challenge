package router

import (
	"go-challenge/repositories"
	"go-challenge/services"
	"go-challenge/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
var locationService interfaces.ILocationService

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
		longitudeStr := c.Query("longitude")
		latitudeStr := c.Query("latitude")
		radiusStr := c.DefaultQuery("radius", "100")

		// Convert long/lat strings to float
		userLongitude, err := strconv.ParseFloat(longitudeStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
			return
		}

		userLatitude, err := strconv.ParseFloat(latitudeStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
			return
		}

		radius, err := strconv.Atoi(radiusStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid radius"})
			return
		}

		locations, err := locationService.GetNearbyEvses(float32(userLongitude), float32(userLatitude), radius)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, locations)
	})

	return r
}

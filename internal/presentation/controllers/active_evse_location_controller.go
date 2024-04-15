package controllers

import (
	"fmt"
	"go-challenge/internal/application/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationRequest struct {
	Latitude  string `form:"latitude" binding:"required,latitude"`
	Longitude string `form:"longitude" binding:"required,longitude"`
	Radius    int    `form:"radius" binding:"omitempty,min=1" default:"100"`
}

type ActiveEVSELocationController struct {
	LocationUseCase usecase.ActiveEVSELocationUseCaseInterface
}

func NewActiveEVSELocationController(locationUseCase usecase.ActiveEVSELocationUseCaseInterface,
) *ActiveEVSELocationController {
	return &ActiveEVSELocationController{LocationUseCase: locationUseCase}
}

func (lc *ActiveEVSELocationController) FetchActiveEVSELocations(c *gin.Context) {
	var requestParams LocationRequest
	if err := c.ShouldBindQuery(&requestParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please check your input. " + err.Error()})
		return
	}

	floatLatitude, floatLongitude, err := lc.parseFloatCoordinates(requestParams.Latitude, requestParams.Longitude)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	locations, err := lc.LocationUseCase.FindLocationsWithActiveEVSE(
		c.Request.Context(),
		floatLatitude,
		floatLongitude,
		requestParams.Radius,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An error occurred while processing your request. Please try again later.",
		})
		return
	}

	if len(locations) == 0 {
		c.JSON(http.StatusOK, gin.H{"locations": locations, "message": "No locations found matching your criteria."})
	} else {
		c.JSON(http.StatusOK, gin.H{"locations": locations})
	}
}

func (lc *ActiveEVSELocationController) parseFloatCoordinates(latStr, longStr string) (float64, float64, error) {
	floatLatitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("please enter latitude in decimal format")
	}
	floatLongitude, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("please enter longitude in decimal format")
	}
	return floatLatitude, floatLongitude, nil
}

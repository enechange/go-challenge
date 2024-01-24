package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/diskymg/go-challenge/diskymg/lib"
	"github.com/diskymg/go-challenge/diskymg/models"
	"github.com/diskymg/go-challenge/diskymg/ocpi"
)

// GetLocations implements all the handlers in the ServerInterface
func (o *OcpiApi) GetLocations(c *gin.Context, params ocpi.GetLocationsParams) {
	rows, err := models.FindLocations(params.DateFrom, params.DateTo, params.Offset, params.Limit)
	if err != nil {
		c.JSON(http.StatusOK, ocpi.LocationsResponse{
			Data:          nil,
			StatusCode:    3000,
			StatusMessage: lib.Ptr("Server errors"),
			Timestamp:     time.Now(),
		})
		return
	}

	dests := make([]ocpi.Location, len(rows))
	for i, row := range rows {
		dests[i] = row.ToOcpi()
	}
	c.JSON(http.StatusOK, ocpi.LocationsResponse{
		Data:          &dests,
		StatusCode:    1000,
		StatusMessage: lib.Ptr("Success"),
		Timestamp:     time.Now(),
	})
}

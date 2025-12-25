package handler

import (
	"net/http"
	"strconv"
	"time"

	"go-challenge/service"

	"github.com/gin-gonic/gin"
)

const defaultRadius = 100.0

type LocationHandler struct {
	service service.LocationService
}

func NewLocationHandler(svc service.LocationService) *LocationHandler {
	return &LocationHandler{service: svc}
}

type GeoLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type EVSEResponse struct {
	UID    string `json:"uid"`
	Status string `json:"status"`
}

type LocationResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name,omitempty"`
	Address     string         `json:"address"`
	Coordinates GeoLocation    `json:"coordinates"`
	EVSEs       []EVSEResponse `json:"evses,omitempty"`
}

func (h *LocationHandler) GetLocations(c *gin.Context) {
	latStr := c.Query("latitude")
	lngStr := c.Query("longitude")
	radiusStr := c.Query("radius")
	dateFromStr := c.Query("date_from")
	dateToStr := c.Query("date_to")

	if latStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "latitude is required"})
		return
	}
	if lngStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "longitude is required"})
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid latitude format"})
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid longitude format"})
		return
	}

	radius := defaultRadius
	if radiusStr != "" {
		r, err := strconv.ParseFloat(radiusStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid radius format"})
			return
		}
		radius = r
	}

	var dateFrom, dateTo *time.Time
	if dateFromStr != "" {
		t, err := time.Parse(time.RFC3339, dateFromStr)
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05", dateFromStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_from format"})
				return
			}
		}
		dateFrom = &t
	}
	if dateToStr != "" {
		t, err := time.Parse(time.RFC3339, dateToStr)
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05", dateToStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_to format"})
				return
			}
		}
		dateTo = &t
	}

	results, err := h.service.SearchLocations(c.Request.Context(), service.SearchParams{
		Latitude:  lat,
		Longitude: lng,
		Radius:    radius,
		DateFrom:  dateFrom,
		DateTo:    dateTo,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]LocationResponse, len(results))
	for i, loc := range results {
		evses := make([]EVSEResponse, len(loc.EVSEs))
		for j, e := range loc.EVSEs {
			evses[j] = EVSEResponse{
				UID:    e.UID,
				Status: e.Status,
			}
		}

		var name string
		if loc.Name != nil {
			name = *loc.Name
		}

		response[i] = LocationResponse{
			ID:      strconv.Itoa(int(loc.ID)),
			Name:    name,
			Address: loc.Address,
			Coordinates: GeoLocation{
				Latitude:  loc.Latitude,
				Longitude: loc.Longitude,
			},
			EVSEs: evses,
		}
	}

	c.JSON(http.StatusOK, response)
}

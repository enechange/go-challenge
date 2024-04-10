// services/location_service.go
package services

import (
	"encoding/json"
	"go-challenge/constants"
	"go-challenge/models"
	"go-challenge/repositories"
)

type ILocationService interface {
	GetNearbyEvses(userLongitude float32, userLatitude float32, radius int) ([]models.LocationResponse, error)
}

type LocationService struct {
	LocationRepository repositories.ILocationRepository
}

var _ ILocationService = &LocationService{}

func (s *LocationService) GetNearbyEvses(userLongitude float32, userLatitude float32, radius int) ([]models.LocationResponse, error) {
	radius = radius * 1000 // set radius to km
	var locations []models.Location
	locations, err := s.LocationRepository.GetNearbyFromRadius(userLongitude, userLatitude, radius)
	if err != nil {
		return nil, err
	}

	var locationResponses []models.LocationResponse
	for _, location := range locations {
		var evseDbResults []models.Evse
		var evseResponses []models.EvseResponse
		if err := json.Unmarshal([]byte(location.Evses), &evseDbResults); err != nil {
			return nil, err
		}
		// Map status string from db value of int
		for _, evseDbResult := range evseDbResults {
			statusString := constants.Status(evseDbResult.Status)
			evseResponses = append(evseResponses, models.EvseResponse{
				Uid:    evseDbResult.Uid,
				Status: statusString.String(),
			})
		}

		if len(evseDbResults) > 0 {
			locationResponses = append(locationResponses, models.LocationResponse{
				ID:      location.ID,
				Name:    location.Name,
				Address: location.Address,
				Coordinates: models.GeoLocation{
					Longitude: location.Longitude,
					Latitude:  location.Latitude,
				},
				Evses: evseResponses,
			})
		}
	}
	return locationResponses, nil
}

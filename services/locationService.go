// services/location_service.go
package services

import (
	"encoding/json"
	"fmt"
	"go-challenge/constants"
	"go-challenge/interfaces"
	"go-challenge/models"
)

type LocationService struct {
	LocationRepository interfaces.ILocationRepository
}

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
			statusStr, ok := constants.StatusMap[evseDbResult.Status]
			if !ok {
				return nil, fmt.Errorf("invalid status code for: %s", evseDbResult.Uid)
			}
			evseResponses = append(evseResponses, models.EvseResponse{
				Uid:    evseDbResult.Uid,
				Status: statusStr,
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

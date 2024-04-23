package usecase

import (
	"go-challenge/internal/application/dto"
	"go-challenge/internal/domain"
	"strconv"
)

func ConvertQueryResultToLocations(
	availableEvseLocations []dto.AvailableEVSELocation,
) ([]domain.Location, error) {
	locationsMap := make(map[string]*domain.Location)

	for _, availableEvseLocation := range availableEvseLocations {
		if loc, exists := locationsMap[availableEvseLocation.ID]; !exists {
			latitude := strconv.FormatFloat(availableEvseLocation.Latitude, 'f', 10, 64)
			longitude := strconv.FormatFloat(availableEvseLocation.Longitude, 'f', 10, 64)
			loc = &domain.Location{
				ID:          availableEvseLocation.ID,
				Name:        availableEvseLocation.Name,
				Address:     availableEvseLocation.Address,
				Coordinates: domain.GeoLocation{Latitude: latitude, Longitude: longitude},
				EVSES:       []domain.EVSE{},
			}
			locationsMap[availableEvseLocation.ID] = loc
		}
		locationsMap[availableEvseLocation.ID].EVSES = append(
			locationsMap[availableEvseLocation.ID].EVSES,
			domain.EVSE{
				UID:    availableEvseLocation.UID,
				Status: availableEvseLocation.Status,
			},
		)
	}

	var locations []domain.Location
	for _, loc := range locationsMap {
		locations = append(locations, *loc)
	}
	return locations, nil
}

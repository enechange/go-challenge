package usecase

import (
	"context"
	appQuery "go-challenge/internal/application/query"
	"go-challenge/internal/domain"
	"strconv"
)

type ActiveEVSELocationUseCaseInterface interface {
	FindLocationsWithActiveEVSE(
		ctx context.Context,
		latitude, longitude float64,
		radius int,
	) ([]domain.Location, error)
}

type ActiveEVSELocationUseCase struct {
	repo appQuery.ActiveEVSELocationQueryService
}

func NewActiveEVSELocationUseCase(
	repo appQuery.ActiveEVSELocationQueryService,
) *ActiveEVSELocationUseCase {
	return &ActiveEVSELocationUseCase{repo: repo}
}

func (uc *ActiveEVSELocationUseCase) FindLocationsWithActiveEVSE(
	ctx context.Context,
	latitude, longitude float64,
	radius int,
) ([]domain.Location, error) {
	queryResults, err := uc.repo.FindLocationsWithActiveEVSE(ctx, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	return ConvertQueryResultToLocations(queryResults)
}

func ConvertQueryResultToLocations(
	availableEvseLocations []domain.AvailableEVSELocation,
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
				Status: domain.Status(availableEvseLocation.Status),
			},
		)
	}

	var locations []domain.Location
	for _, loc := range locationsMap {
		locations = append(locations, *loc)
	}
	return locations, nil
}

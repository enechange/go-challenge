package service

import (
	"context"
	"math"
	"strconv"
	"time"

	"go-challenge/repository"
)

type SearchParams struct {
	Latitude  float64
	Longitude float64
	Radius    float64
	DateFrom  *time.Time
	DateTo    *time.Time
}

type LocationResult struct {
	ID        int32
	Name      *string
	Address   string
	Latitude  string
	Longitude string
	EVSEs     []EVSEResult
}

type EVSEResult struct {
	UID    string
	Status string
}

type LocationService interface {
	SearchLocations(ctx context.Context, params SearchParams) ([]LocationResult, error)
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) SearchLocations(ctx context.Context, params SearchParams) ([]LocationResult, error) {
	locations, err := s.repo.GetLocationsByDateRange(ctx, params.DateFrom, params.DateTo)
	if err != nil {
		return nil, err
	}

	allEVSEs, err := s.repo.GetAllEVSEs(ctx)
	if err != nil {
		return nil, err
	}

	evseMap := make(map[int32][]repository.EVSEWithLocationID)
	for _, e := range allEVSEs {
		evseMap[e.LocationID] = append(evseMap[e.LocationID], e)
	}

	var results []LocationResult
	for _, loc := range locations {
		lat, _ := strconv.ParseFloat(loc.Latitude, 64)
		lng, _ := strconv.ParseFloat(loc.Longitude, 64)

		distance := haversine(params.Latitude, params.Longitude, lat, lng)
		if distance > params.Radius {
			continue
		}

		locEVSEs := evseMap[loc.ID]
		evseResults := make([]EVSEResult, len(locEVSEs))
		for i, e := range locEVSEs {
			evseResults[i] = EVSEResult{
				UID:    e.UID,
				Status: statusToString(e.Status),
			}
		}

		results = append(results, LocationResult{
			ID:        loc.ID,
			Name:      loc.Name,
			Address:   loc.Address,
			Latitude:  loc.Latitude,
			Longitude: loc.Longitude,
			EVSEs:     evseResults,
		})
	}

	return results, nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371.0

	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	deltaLat := (lat2 - lat1) * math.Pi / 180
	deltaLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func statusToString(status int32) string {
	switch status {
	case 1:
		return "AVAILABLE"
	case 2:
		return "BLOCKED"
	case 3:
		return "CHARGING"
	case 4:
		return "INOPERATIVE"
	case 5:
		return "OUTOFORDER"
	case 6:
		return "PLANNED"
	case 7:
		return "REMOVED"
	case 8:
		return "RESERVED"
	default:
		return "UNKNOWN"
	}
}

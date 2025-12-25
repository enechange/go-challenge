package repository

import (
	"context"
	"time"

	"go-challenge/database"
)

type Location struct {
	ID          int32
	Name        *string
	Address     string
	Latitude    string
	Longitude   string
	LastUpdated time.Time
}

type EVSE struct {
	UID    string
	Status int32
}

type LocationWithEVSEs struct {
	Location
	EVSEs []EVSE
}

type LocationRepository interface {
	GetAllLocations(ctx context.Context) ([]Location, error)
	GetLocationsByDateRange(ctx context.Context, dateFrom, dateTo *time.Time) ([]Location, error)
	GetEVSEsByLocationID(ctx context.Context, locationID int32) ([]EVSE, error)
}

type locationRepository struct {
	queries *database.Queries
}

func NewLocationRepository(db database.DBTX) LocationRepository {
	return &locationRepository{
		queries: database.New(db),
	}
}

func (r *locationRepository) GetAllLocations(ctx context.Context) ([]Location, error) {
	dbLocations, err := r.queries.GetAllLocations(ctx)
	if err != nil {
		return nil, err
	}

	locations := make([]Location, len(dbLocations))
	for i, loc := range dbLocations {
		locations[i] = toLocation(loc)
	}
	return locations, nil
}

func (r *locationRepository) GetLocationsByDateRange(ctx context.Context, dateFrom, dateTo *time.Time) ([]Location, error) {
	var dbLocations []database.Location
	var err error

	switch {
	case dateFrom != nil && dateTo != nil:
		dbLocations, err = r.queries.GetLocationsByDateRange(ctx, database.GetLocationsByDateRangeParams{
			LastUpdated:   *dateFrom,
			LastUpdated_2: *dateTo,
		})
	case dateFrom != nil:
		dbLocations, err = r.queries.GetLocationsByDateFrom(ctx, *dateFrom)
	case dateTo != nil:
		dbLocations, err = r.queries.GetLocationsByDateTo(ctx, *dateTo)
	default:
		dbLocations, err = r.queries.GetAllLocations(ctx)
	}

	if err != nil {
		return nil, err
	}

	locations := make([]Location, len(dbLocations))
	for i, loc := range dbLocations {
		locations[i] = toLocation(loc)
	}
	return locations, nil
}

func (r *locationRepository) GetEVSEsByLocationID(ctx context.Context, locationID int32) ([]EVSE, error) {
	dbEVSEs, err := r.queries.GetEVSEsByLocationID(ctx, locationID)
	if err != nil {
		return nil, err
	}

	evses := make([]EVSE, len(dbEVSEs))
	for i, e := range dbEVSEs {
		evses[i] = EVSE{
			UID:    e.Uid,
			Status: e.Status,
		}
	}
	return evses, nil
}

func toLocation(loc database.Location) Location {
	var name *string
	if loc.Name.Valid {
		name = &loc.Name.String
	}
	return Location{
		ID:          loc.ID,
		Name:        name,
		Address:     loc.Address,
		Latitude:    loc.Latitude,
		Longitude:   loc.Longitude,
		LastUpdated: loc.LastUpdated,
	}
}

package query

import (
	"context"
	"go-challenge/internal/domain"
	"go-challenge/internal/infrastructure/database"

	"gorm.io/gorm"
)

type ActiveEVSELocationQueryServiceGorm struct {
	db *gorm.DB
}

func NewActiveEVSELocationQueryServiceGorm() *ActiveEVSELocationQueryServiceGorm {
	return &ActiveEVSELocationQueryServiceGorm{
		db: database.GetDB(),
	}
}

func (s *ActiveEVSELocationQueryServiceGorm) FindLocationsWithActiveEVSE(
	ctx context.Context,
	latitude, longitude float64,
	radius int,
) ([]domain.AvailableEVSELocation, error) {

	var availableEVSELocations []domain.AvailableEVSELocation
	radiusInMeters := radius * 1000

	err := s.db.Table("locations l").
		Select(
			"l.id, l.name, l.address, "+
				"CAST(l.latitude AS DECIMAL(10, 6)) AS latitude, "+
				"CAST(l.longitude AS DECIMAL(10, 6)) AS longitude, "+
				"e.uid, e.status",
		).
		Joins("INNER JOIN evses e ON l.id = e.location_id").
		Where(
			"e.status = ? AND "+
				"ST_Distance_Sphere("+
				"point(CAST(l.longitude AS DECIMAL(10, 6)), CAST(l.latitude  AS DECIMAL(10, 6))), "+
				"point(?, ?)) <= ?",
			domain.Available, longitude, latitude, radiusInMeters,
		).
		Scan(&availableEVSELocations).Error

	if err != nil {
		return nil, err
	}

	return availableEVSELocations, nil
}

package repositories

import (
	"go-challenge/database"
	"go-challenge/models"
)

type LocationRepository struct{}

func (l *LocationRepository) GetNearbyFromRadius(userLongitude float32, userLatitude float32, radius int) ([]models.Location, error) {
	var locations []models.Location

	database.GetDB().Table("locations l").
		Select("l.id, l.name, l.address, ST_X(l.coordinates) as longitude, ST_Y(l.coordinates) as latitude, ST_Distance_Sphere(l.coordinates, POINT(?, ?)) AS distance, JSON_ARRAYAGG(JSON_OBJECT('uid', e.uid, 'status', e.status)) AS evses", userLongitude, userLatitude).
		Joins("LEFT JOIN evses e ON l.id = e.locationId").
		Group("l.id").
		Having("distance < ? AND MAX(e.uid IS NOT NULL)", radius).
		Order("distance").
		Scan(&locations)

	return locations, nil
}

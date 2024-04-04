package interfaces

import "go-challenge/models"

type ILocationRepository interface {
	GetNearbyFromRadius(userLongitude float32, userLatitude float32, radius int) ([]models.Location, error)
}

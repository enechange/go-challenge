package interfaces

import "go-challenge/models"

type ILocationService interface {
	GetNearbyEvses(userLongitude float32, userLatitude float32, radius int) ([]models.LocationResponse, error)
}
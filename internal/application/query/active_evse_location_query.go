package query

import (
	"context"
	"go-challenge/internal/domain"
)

type ActiveEVSELocationQueryService interface {
	FindLocationsWithActiveEVSE(
		ctx context.Context,
		latitude, longitude float64,
		radius int,
	) (
		[]domain.AvailableEVSELocation,
		error,
	)
}

package query

import (
	"context"
	"go-challenge/internal/application/dto"
)

type ActiveEVSELocationQueryService interface {
	FindLocationsWithActiveEVSE(
		ctx context.Context,
		latitude, longitude float64,
		radius int,
	) (
		[]dto.AvailableEVSELocation,
		error,
	)
}

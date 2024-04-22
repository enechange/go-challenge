package usecase

import (
	"context"
	appQuery "go-challenge/internal/application/query"
	"go-challenge/internal/domain"
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

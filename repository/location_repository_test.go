package repository_test

import (
	"context"
	"testing"

	"go-challenge/config"
	"go-challenge/database"
	"go-challenge/repository"
)

func TestGetAllLocations(t *testing.T) {
	config.Init()
	database.Init()
	defer database.Close()

	repo := repository.NewLocationRepository(database.GetDB())
	ctx := context.Background()

	locations, err := repo.GetAllLocations(ctx)
	if err != nil {
		t.Fatalf("GetAllLocations failed: %v", err)
	}

	if len(locations) != 100 {
		t.Errorf("expected 100 locations, got %d", len(locations))
	}
}

func TestGetEVSEsByLocationID(t *testing.T) {
	config.Init()
	database.Init()
	defer database.Close()

	repo := repository.NewLocationRepository(database.GetDB())
	ctx := context.Background()

	evses, err := repo.GetEVSEsByLocationID(ctx, 1)
	if err != nil {
		t.Fatalf("GetEVSEsByLocationID failed: %v", err)
	}

	if len(evses) == 0 {
		t.Error("expected at least one EVSE for location 1")
	}
}

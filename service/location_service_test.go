package service_test

import (
	"context"
	"testing"
	"time"

	"go-challenge/repository"
	"go-challenge/service"
)

type mockLocationRepository struct {
	locations []repository.Location
	evses     map[int32][]repository.EVSE
}

func (m *mockLocationRepository) GetAllLocations(ctx context.Context) ([]repository.Location, error) {
	return m.locations, nil
}

func (m *mockLocationRepository) GetLocationsByDateRange(ctx context.Context, dateFrom, dateTo *time.Time) ([]repository.Location, error) {
	return m.locations, nil
}

func (m *mockLocationRepository) GetEVSEsByLocationID(ctx context.Context, locationID int32) ([]repository.EVSE, error) {
	return m.evses[locationID], nil
}

func (m *mockLocationRepository) GetAllEVSEs(ctx context.Context) ([]repository.EVSEWithLocationID, error) {
	var result []repository.EVSEWithLocationID
	for locationID, evses := range m.evses {
		for _, e := range evses {
			result = append(result, repository.EVSEWithLocationID{
				LocationID: locationID,
				UID:        e.UID,
				Status:     e.Status,
			})
		}
	}
	return result, nil
}

func TestSearchLocations_FiltersByRadius(t *testing.T) {
	name1 := "Location 1"
	name2 := "Location 2"

	mockRepo := &mockLocationRepository{
		locations: []repository.Location{
			{ID: 1, Name: &name1, Address: "Address 1", Latitude: "43.0", Longitude: "141.0", LastUpdated: time.Now()},
			{ID: 2, Name: &name2, Address: "Address 2", Latitude: "35.0", Longitude: "139.0", LastUpdated: time.Now()}, // 遠い
		},
		evses: map[int32][]repository.EVSE{
			1: {{UID: "EVSE1", Status: 1}},
			2: {{UID: "EVSE2", Status: 1}},
		},
	}

	svc := service.NewLocationService(mockRepo)
	ctx := context.Background()

	results, err := svc.SearchLocations(ctx, service.SearchParams{
		Latitude:  43.0,
		Longitude: 141.0,
		Radius:    100, // 100km
	})

	if err != nil {
		t.Fatalf("SearchLocations failed: %v", err)
	}

	// Location 1 は範囲内、Location 2 は範囲外（約900km離れている）
	if len(results) != 1 {
		t.Errorf("expected 1 location, got %d", len(results))
	}

	if len(results) > 0 && results[0].ID != 1 {
		t.Errorf("expected location ID 1, got %d", results[0].ID)
	}
}

func TestSearchLocations_IncludesEVSEs(t *testing.T) {
	name1 := "Location 1"

	mockRepo := &mockLocationRepository{
		locations: []repository.Location{
			{ID: 1, Name: &name1, Address: "Address 1", Latitude: "43.0", Longitude: "141.0", LastUpdated: time.Now()},
		},
		evses: map[int32][]repository.EVSE{
			1: {
				{UID: "EVSE1", Status: 1},
				{UID: "EVSE2", Status: 3},
			},
		},
	}

	svc := service.NewLocationService(mockRepo)
	ctx := context.Background()

	results, err := svc.SearchLocations(ctx, service.SearchParams{
		Latitude:  43.0,
		Longitude: 141.0,
		Radius:    100,
	})

	if err != nil {
		t.Fatalf("SearchLocations failed: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 location, got %d", len(results))
	}

	if len(results[0].EVSEs) != 2 {
		t.Errorf("expected 2 EVSEs, got %d", len(results[0].EVSEs))
	}
}

func TestSearchLocations_ConvertsStatusToString(t *testing.T) {
	name1 := "Location 1"

	mockRepo := &mockLocationRepository{
		locations: []repository.Location{
			{ID: 1, Name: &name1, Address: "Address 1", Latitude: "43.0", Longitude: "141.0", LastUpdated: time.Now()},
		},
		evses: map[int32][]repository.EVSE{
			1: {{UID: "EVSE1", Status: 1}},
		},
	}

	svc := service.NewLocationService(mockRepo)
	ctx := context.Background()

	results, err := svc.SearchLocations(ctx, service.SearchParams{
		Latitude:  43.0,
		Longitude: 141.0,
		Radius:    100,
	})

	if err != nil {
		t.Fatalf("SearchLocations failed: %v", err)
	}

	if results[0].EVSEs[0].Status != "AVAILABLE" {
		t.Errorf("expected status AVAILABLE, got %s", results[0].EVSEs[0].Status)
	}
}

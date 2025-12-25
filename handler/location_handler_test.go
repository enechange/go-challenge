package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-challenge/config"
	"go-challenge/database"
	"go-challenge/router"
)

func init() {
	config.Init()
	database.Init()
}

func TestGetLocations_ReturnsOK(t *testing.T) {
	r := router.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/locations?latitude=43.0&longitude=141.0", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestGetLocations_MissingLatitude_ReturnsBadRequest(t *testing.T) {
	r := router.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/locations?longitude=141.0", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestGetLocations_MissingLongitude_ReturnsBadRequest(t *testing.T) {
	r := router.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/locations?latitude=43.0", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestGetLocations_InvalidLatitude_ReturnsBadRequest(t *testing.T) {
	r := router.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/locations?latitude=invalid&longitude=141.0", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestGetLocations_DefaultRadius(t *testing.T) {
	r := router.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/locations?latitude=43.0&longitude=141.0", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestGetLocations_ReturnsJSONArray(t *testing.T) {
	r := router.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/locations?latitude=43.0&longitude=141.0", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	var result []map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Errorf("response is not valid JSON array: %v", err)
	}
}

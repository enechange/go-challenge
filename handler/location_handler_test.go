package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-challenge/router"
)

func TestGetLocations_ReturnsOK(t *testing.T) {
	r := router.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/locations?latitude=43.0&longitude=141.0", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

package controllers

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/diskymg/go-challenge/diskymg/ocpi"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
)

func TestOcpiApi_GetLocations(t *testing.T) {
	r := initialize()
	defer finalize()

	t.Run("Get all locations", func(t *testing.T) {
		response := testutil.NewRequest().Get("/ocpi/diskymg/2.2/locations").WithAcceptJson().GoWithHTTPHandler(t, r)
		assert.Equal(t, http.StatusOK, response.Recorder.Code)

		var result ocpi.LocationsResponse
		err := json.NewDecoder(response.Recorder.Body).Decode(&result)
		assert.NoError(t, err, "error getting response", err)
		assert.Equal(t, 3, len(*result.Data))
	})

	t.Run("Filter locations by last_updated", func(t *testing.T) {
		// date_from は閾値を含む、date_to は閾値を含まない
		response := testutil.NewRequest().Get("/ocpi/diskymg/2.2/locations?date_from=2021-04-01T00%3A00%3A00Z&date_to=2023-04-01T00%3A00%3A00Z").WithAcceptJson().GoWithHTTPHandler(t, r)
		assert.Equal(t, http.StatusOK, response.Recorder.Code)

		var result ocpi.LocationsResponse
		err := json.NewDecoder(response.Recorder.Body).Decode(&result)
		assert.NoError(t, err, "error getting response", err)
		assert.Equal(t, 2, len(*result.Data))
		assert.Equal(t, "Gent Zuid", *(*result.Data)[0].Name)
		assert.Equal(t, "ihomer", *(*result.Data)[1].Name)
	})

	t.Run("offset 1 & limit 2", func(t *testing.T) {
		response := testutil.NewRequest().Get("/ocpi/diskymg/2.2/locations?offset=1&limit=2").WithAcceptJson().GoWithHTTPHandler(t, r)
		assert.Equal(t, http.StatusOK, response.Recorder.Code)

		var result ocpi.LocationsResponse
		err := json.NewDecoder(response.Recorder.Body).Decode(&result)
		assert.NoError(t, err, "error getting response", err)
		assert.Equal(t, 2, len(*result.Data))
		assert.Equal(t, "ihomer", *(*result.Data)[0].Name)
		assert.Equal(t, "Water State", *(*result.Data)[1].Name)
	})
}

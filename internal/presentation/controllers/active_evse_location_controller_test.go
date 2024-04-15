package controllers

import (
	"errors"
	"go-challenge/internal/application/usecase/mock"
	"go-challenge/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestFetchActiveEVSELocations(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCase := mock.NewMockIActiveEVSELocationUseCase(mockCtrl)
	controller := NewActiveEVSELocationController(mockUseCase)
	router := setupRouter()
	router.GET("/api/locations", controller.FetchActiveEVSELocations)

	locationName := "Location 1"
	tests := []struct {
		name         string
		setupRequest func() *http.Request
		expectedCode int
		expectedBody string
		setupMock    func()
	}{
		{
			name: "valid request",
			setupRequest: func() *http.Request {
				req, _ := http.NewRequest("GET", "/api/locations?latitude=34.0522&longitude=118.2437&radius=10", nil)
				return req
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"locations":[{"ID":"1","Name":"Location 1","Address":"Example Address","Coordinates":{"Latitude":"34.0522","Longitude":"118.2437"},"EVSES":[{"UID":"Evse1","Status":1}]}]}`,
			setupMock: func() {
				mockUseCase.EXPECT().FindLocationsWithActiveEVSE(gomock.Any(), 34.0522, 118.2437, 10).Return([]domain.Location{
					{
						ID:          "1",
						Name:        &locationName,
						Address:     "Example Address",
						Coordinates: domain.GeoLocation{Latitude: "34.0522", Longitude: "118.2437"},
						EVSES:       []domain.EVSE{{UID: "Evse1", Status: domain.Available}},
					},
				}, nil)
			},
		},
		{
			name: "invalid latitude format",
			setupRequest: func() *http.Request {
				req, _ := http.NewRequest("GET", "/api/locations?latitude=abc&longitude=118.2437&radius=10", nil)
				return req
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Please check your input. Key: 'LocationRequest.Latitude' Error:Field validation for 'Latitude' failed on the 'latitude' tag"}`,
			setupMock:    func() {},
		},
		{
			name: "use case error",
			setupRequest: func() *http.Request {
				req, _ := http.NewRequest("GET", "/api/locations?latitude=34.0522&longitude=118.2437&radius=10", nil)
				return req
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"An error occurred while processing your request. Please try again later."}`,
			setupMock: func() {
				mockUseCase.EXPECT().FindLocationsWithActiveEVSE(gomock.Any(), 34.0522, 118.2437, 10).Return(nil, errors.New("internal error"))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, tc.setupRequest())
			assert.Equal(t, tc.expectedCode, recorder.Code)
			assert.JSONEq(t, tc.expectedBody, recorder.Body.String())
		})
	}
}

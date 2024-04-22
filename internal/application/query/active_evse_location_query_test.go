package query

import (
	"context"
	"errors"
	"go-challenge/internal/application/dto"
	"go-challenge/internal/application/query/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFindLocationsWithActiveEVSE(t *testing.T) {
	tests := []struct {
		name              string
		latitude          float64
		longitude         float64
		radius            int
		expectedLocations []dto.AvailableEVSELocation
		expectedError     error
		prepareMockFn     func(m *mock.MockActiveEVSELocationQueryService)
	}{
		{
			name:      "Normal Request",
			latitude:  35.6895,
			longitude: 139.6917,
			radius:    10,
			expectedLocations: []dto.AvailableEVSELocation{
				{ID: "1", Name: new(string), Address: "東京都新宿区...", Latitude: 35.6895, Longitude: 139.6917, UID: "UID123", Status: 1},
			},
			prepareMockFn: func(m *mock.MockActiveEVSELocationQueryService) {
				m.EXPECT().FindLocationsWithActiveEVSE(gomock.Any(), 35.6895, 139.6917, 10).
					Return([]dto.AvailableEVSELocation{{ID: "1", Name: new(string), Address: "東京都新宿区...", Latitude: 35.6895, Longitude: 139.6917, UID: "UID123", Status: 1}}, nil).Times(1)
			},
		},
		{
			name:              "Case that Returns an Error Response",
			latitude:          34.6937,
			longitude:         135.5023,
			radius:            5,
			expectedLocations: nil,
			expectedError:     errors.New("database error"),
			prepareMockFn: func(m *mock.MockActiveEVSELocationQueryService) {
				m.EXPECT().FindLocationsWithActiveEVSE(gomock.Any(), 34.6937, 135.5023, 5).
					Return(nil, errors.New("database error")).Times(1)
			},
		},
		{
			name:              "Invalid Range (Negative Radius)",
			latitude:          35.6895,
			longitude:         139.6917,
			radius:            -10,
			expectedLocations: nil,
			expectedError:     errors.New("invalid radius"),
			prepareMockFn: func(m *mock.MockActiveEVSELocationQueryService) {
				m.EXPECT().FindLocationsWithActiveEVSE(gomock.Any(), 35.6895, 139.6917, -10).
					Return(nil, errors.New("invalid radius")).Times(1)
			},
		},
		{
			name:              "Boundary Value Test (0 km Radius)",
			latitude:          35.6895,
			longitude:         139.6917,
			radius:            0,
			expectedLocations: []dto.AvailableEVSELocation{},
			expectedError:     nil,
			prepareMockFn: func(m *mock.MockActiveEVSELocationQueryService) {
				m.EXPECT().FindLocationsWithActiveEVSE(gomock.Any(), 35.6895, 139.6917, 0).
					Return([]dto.AvailableEVSELocation{}, nil).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock.NewMockActiveEVSELocationQueryService(ctrl)
			tt.prepareMockFn(m)

			locations, err := m.FindLocationsWithActiveEVSE(context.Background(), tt.latitude, tt.longitude, tt.radius)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, len(tt.expectedLocations), len(locations), "Expected number of locations does not match")
			assert.Equal(t, tt.expectedLocations, locations, "Returned locations are not as expected")
		})
	}
}

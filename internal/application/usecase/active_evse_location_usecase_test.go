package usecase

import (
	"context"
	"errors"
	"go-challenge/internal/application/query/mock"
	"go-challenge/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFindLocationsWithActiveEVSE(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockActiveEVSELocationQueryService(ctrl)
	uc := NewActiveEVSELocationUseCase(mockRepo)

	tests := []struct {
		name           string
		latitude       float64
		longitude      float64
		radius         int
		mockSetup      func()
		expectedError  error
		expectedResult []domain.Location
	}{
		{
			name:      "Normal Response",
			latitude:  35.6895,
			longitude: 139.6917,
			radius:    10,
			mockSetup: func() {
				mockRepo.EXPECT().FindLocationsWithActiveEVSE(gomock.Any(), 35.6895, 139.6917, 10).
					Return([]domain.AvailableEVSELocation{
						{ID: "1", Name: new(string), Address: "東京都新宿区", Latitude: 35.6895, Longitude: 139.6917, UID: "UID123", Status: 1},
					}, nil)
			},
			expectedError: nil,
			expectedResult: []domain.Location{
				{ID: "1", Name: new(string), Address: "東京都新宿区", Coordinates: domain.GeoLocation{Latitude: "35.6895000000", Longitude: "139.6917000000"}, EVSES: []domain.EVSE{{UID: "UID123", Status: domain.Status(1)}}},
			},
		},
		{
			name:      "Error Returned from Repository",
			latitude:  34.6937,
			longitude: 135.5022,
			radius:    5,
			mockSetup: func() {
				mockRepo.EXPECT().FindLocationsWithActiveEVSE(gomock.Any(), 34.6937, 135.5022, 5).
					Return(nil, errors.New("database error"))
			},
			expectedError:  errors.New("database error"),
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			result, err := uc.FindLocationsWithActiveEVSE(context.Background(), tt.latitude, tt.longitude, tt.radius)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}
func TestConvertQueryResultToLocations(t *testing.T) {
	tests := []struct {
		name           string
		input          []domain.AvailableEVSELocation
		expectedOutput []domain.Location
		expectedError  error
	}{
		{
			name: "Normal Case",
			input: []domain.AvailableEVSELocation{
				{
					ID:        "1",
					Name:      new(string),
					Address:   "東京都新宿区",
					Latitude:  35.6895,
					Longitude: 139.6917,
					UID:       "UID123",
					Status:    1,
				},
			},
			expectedOutput: []domain.Location{
				{
					ID:          "1",
					Name:        new(string),
					Address:     "東京都新宿区",
					Coordinates: domain.GeoLocation{Latitude: "35.6895000000", Longitude: "139.6917000000"},
					EVSES: []domain.EVSE{
						{
							UID:    "UID123",
							Status: domain.Status(1),
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			name:           "Empty Input",
			input:          []domain.AvailableEVSELocation{},
			expectedOutput: []domain.Location{},
			expectedError:  nil,
		},
		{
			name: "Duplicate ID",
			input: []domain.AvailableEVSELocation{
				{
					ID:        "1",
					Name:      new(string),
					Address:   "東京都新宿区",
					Latitude:  35.6895,
					Longitude: 139.6917,
					UID:       "UID123",
					Status:    1,
				},
				{
					ID:        "1",
					Name:      new(string),
					Address:   "東京都新宿区",
					Latitude:  35.6895,
					Longitude: 139.6917,
					UID:       "UID124",
					Status:    2,
				},
			},
			expectedOutput: []domain.Location{
				{
					ID:          "1",
					Name:        new(string),
					Address:     "東京都新宿区",
					Coordinates: domain.GeoLocation{Latitude: "35.6895000000", Longitude: "139.6917000000"},
					EVSES: []domain.EVSE{
						{
							UID:    "UID123",
							Status: domain.Status(1),
						},
						{
							UID:    "UID124",
							Status: domain.Status(2),
						},
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := ConvertQueryResultToLocations(tt.input)
			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if len(tt.expectedOutput) == 0 {
				assert.Empty(t, output)
			} else {
				assert.Equal(t, tt.expectedOutput, output)
			}
		})
	}
}

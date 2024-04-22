package query

import (
	"context"
	"go-challenge/internal/domain"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestFindLocationsWithActiveEVSE(t *testing.T) {
	gormDB, mock, err := setupMock()
	assert.NoError(t, err)

	service := NewActiveEVSELocationQueryServiceGorm(gormDB)

	testLocation := "日産プリンス札幌販売 月寒支店"
	rows := sqlmock.NewRows([]string{"id", "name", "address", "latitude", "longitude", "uid", "status"}).
		AddRow("11", &testLocation, "北海道札幌市豊平区月寒中央通11-6-37", 43.0239481931252, 141.4013143447731, "CJNET15K400901", domain.Available)

	query := regexp.QuoteMeta(
		`SELECT l.id, l.name, l.address, CAST(l.latitude AS DECIMAL(10, 6)) AS latitude, CAST(l.longitude AS DECIMAL(10, 6)) AS longitude, e.uid, e.status FROM locations l INNER JOIN evses e ON l.id = e.location_id WHERE e.status = ? AND ST_Distance_Sphere(point(CAST(l.longitude AS DECIMAL(10, 6)), CAST(l.latitude AS DECIMAL(10, 6))), point(?, ?)) <= ?`)

	testCases := []struct {
		description    string
		latitude       float64
		longitude      float64
		radius         int
		expectedError  bool
		expectedLength int
	}{
		{
			description:    "Valid coordinates and radius",
			latitude:       35.6895,
			longitude:      139.6917,
			radius:         10,
			expectedError:  false,
			expectedLength: 1,
		},
		{
			description:    "Invalid coordinates",
			latitude:       -91.0,
			longitude:      182.0,
			radius:         10,
			expectedError:  true,
			expectedLength: 0,
		},
		{
			description:    "Negative radius",
			latitude:       35.6895,
			longitude:      139.6917,
			radius:         -1,
			expectedError:  true,
			expectedLength: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.expectedError {
				mock.ExpectQuery(query).
					WithArgs(domain.Available, tc.longitude, tc.latitude, tc.radius*1000).
					WillReturnError(gorm.ErrInvalidData)
			} else {
				mock.ExpectQuery(query).
					WithArgs(domain.Available, tc.longitude, tc.latitude, tc.radius*1000).
					WillReturnRows(rows)
			}

			ctx := context.Background()
			locations, err := service.FindLocationsWithActiveEVSE(ctx, tc.latitude, tc.longitude, tc.radius)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, locations, tc.expectedLength)
				if len(locations) > 0 {
					assert.Equal(t, &testLocation, locations[0].Name)
				}
			}

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err, "there were unfulfilled expectations")
		})
	}
}

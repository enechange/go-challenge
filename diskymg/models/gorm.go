package models

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/diskymg/go-challenge/diskymg/lib"
)

var db *gorm.DB

// Init initializes database
func InitDB() error {
	var err error
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to Open: %w", err)
	}
	return nil
}

func AutoMigrate() error {
	return db.AutoMigrate(Location{}, LocationOpeningTimesRegularHour{}, Evse{}, Connector{})
}

func Seed() error {
	var count int64

	// main.goが実行される度にレコードが生成されないようにする。
	db.Model(&Location{}).Count(&count)
	if count > 0 {
		return nil
	}

	location1 := Location{
		CountryCode:          "BE",
		PartyID:              "BEC",
		Publish:              true,
		Name:                 lib.Ptr("Gent Zuid"),
		Address:              "F.Rooseveltlaan 3A",
		City:                 "Gent",
		Country:              "BEL",
		CoordinatesLatitude:  "51.047599",
		CoordinatesLongitude: "3.729944",
		Evses: []Evse{
			{
				EvseID: lib.Ptr("BE*BEC*E041503001"),
				Status: "AVAILABLE",
				Connectors: []Connector{
					{
						Standard:    "IEC_62196_T2",
						Format:      "CABLE",
						PowerType:   "AC_3_PHASE",
						MaxVoltage:  220,
						MaxAmperage: 16,
						LastUpdated: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				LastUpdated: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				EvseID: lib.Ptr("BE*BEC*E041503002"),
				Status: "RESERVED",
				Connectors: []Connector{
					{
						Standard:    "IEC_62196_T2",
						Format:      "SOCKET",
						PowerType:   "AC_3_PHASE",
						MaxVoltage:  220,
						MaxAmperage: 16,
						LastUpdated: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				LastUpdated: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				EvseID: lib.Ptr("SE*EVC*E000000123"),
				Status: "AVAILABLE",
				Connectors: []Connector{
					{
						Standard:    "IEC_62196_T2",
						Format:      "SOCKET",
						PowerType:   "AC_3_PHASE",
						MaxVoltage:  230,
						MaxAmperage: 32,
						LastUpdated: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				LastUpdated: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		TimeZone:                    "Europe/Brussels",
		OpeningTimesTwentyfourseven: false,
		OpeningTimesRegularHours: []LocationOpeningTimesRegularHour{
			{Weekday: 1, PeriodBegin: "07:00", PeriodEnd: "18:00"},
			{Weekday: 2, PeriodBegin: "07:00", PeriodEnd: "18:00"},
			{Weekday: 3, PeriodBegin: "07:00", PeriodEnd: "18:00"},
			{Weekday: 4, PeriodBegin: "07:00", PeriodEnd: "18:00"},
			{Weekday: 5, PeriodBegin: "07:00", PeriodEnd: "18:00"},
			{Weekday: 6, PeriodBegin: "07:00", PeriodEnd: "18:00"},
			{Weekday: 7, PeriodBegin: "07:00", PeriodEnd: "18:00"},
		},
		Model:       gorm.Model{CreatedAt: time.Date(2001, 4, 1, 0, 0, 0, 0, time.UTC)},
		LastUpdated: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	location2 := Location{
		CountryCode:          "DE",
		PartyID:              "ALF",
		Publish:              true,
		Name:                 lib.Ptr("ihomer"),
		Address:              "Tamboerijn 7",
		City:                 "Etten-Leur",
		Country:              "NLD",
		CoordinatesLatitude:  "51.562787",
		CoordinatesLongitude: "4.638975",
		Evses: []Evse{
			{
				EvseID: lib.Ptr("NL*ALF*E000000001"),
				Status: "AVAILABLE",
				Connectors: []Connector{
					{
						Standard:    "IEC_62196_T2",
						Format:      "SOCKET",
						PowerType:   "AC_3_PHASE",
						MaxVoltage:  220,
						MaxAmperage: 16,
						LastUpdated: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				LastUpdated: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				EvseID: lib.Ptr("NL*ALF*E000000002"),
				Status: "AVAILABLE",
				Connectors: []Connector{
					{
						Standard:    "IEC_62196_T2",
						Format:      "SOCKET",
						PowerType:   "AC_1_PHASE",
						MaxVoltage:  230,
						MaxAmperage: 8,
						LastUpdated: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				LastUpdated: time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		TimeZone:                    "Europe/Amsterdam",
		OpeningTimesTwentyfourseven: true,
		Model:                       gorm.Model{CreatedAt: time.Date(2002, 4, 1, 0, 0, 0, 0, time.UTC)},
		LastUpdated:                 time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	location3 := Location{
		CountryCode:          "NL",
		PartyID:              "ALL",
		Publish:              false,
		Name:                 lib.Ptr("Water State"),
		Address:              "Taco van der Veenplein 12",
		City:                 "Leeuwarden",
		Country:              "DEU",
		CoordinatesLatitude:  "53.213763",
		CoordinatesLongitude: "5.804638",
		Evses: []Evse{
			{
				EvseID: lib.Ptr("NL*ALF*EG00000001"),
				Status: "AVAILABLE",
				Connectors: []Connector{
					{
						Standard:    "IEC_62196_T2",
						Format:      "SOCKET",
						PowerType:   "AC_3_PHASE",
						MaxVoltage:  220,
						MaxAmperage: 16,
						LastUpdated: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				LastUpdated: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		TimeZone:                    "Europe/Berlin",
		OpeningTimesTwentyfourseven: true,
		Model:                       gorm.Model{CreatedAt: time.Date(2003, 4, 1, 0, 0, 0, 0, time.UTC)},
		LastUpdated:                 time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
	}

	if err := db.Create([]Location{location1, location2, location3}).Error; err != nil {
		return fmt.Errorf("failed to Create: %w", err)
	}

	return nil
}

// GetDB returns database connection
func DB() *gorm.DB {
	return db
}

// Close closes database
func CloseDB() {
	sqldb, _ := db.DB()
	sqldb.Close()
}

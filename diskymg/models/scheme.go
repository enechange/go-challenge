package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model

	ID                          string `gorm:"default:gen_random_uuid()"` // db func
	CountryCode                 string
	PartyID                     string
	Publish                     bool
	Name                        *string
	Address                     string
	City                        string
	Country                     string
	CoordinatesLatitude         string
	CoordinatesLongitude        string
	Evses                       []Evse
	TimeZone                    string
	OpeningTimesTwentyfourseven bool
	OpeningTimesRegularHours    []LocationOpeningTimesRegularHour
	LastUpdated                 time.Time
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}

type LocationOpeningTimesRegularHour struct {
	gorm.Model

	ID          string `gorm:"default:gen_random_uuid()"` // db func
	LocationID  string
	Weekday     int
	PeriodBegin string
	PeriodEnd   string
}

type Evse struct {
	gorm.Model

	ID          string `gorm:"default:gen_random_uuid()"` // db func
	LocationID  string
	EvseID      *string
	Status      string
	Connectors  []Connector
	LastUpdated time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Connector struct {
	gorm.Model

	ID          string `gorm:"default:gen_random_uuid()"` // db func
	EvseID      string
	Standard    string
	Format      string
	PowerType   string
	MaxVoltage  int
	MaxAmperage int
	LastUpdated time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

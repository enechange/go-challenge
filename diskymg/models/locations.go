package models

import (
	"fmt"
	"time"

	"github.com/diskymg/go-challenge/diskymg/ocpi"
)

// 必須
// - country_code
// - party_id
// - id
// - publish
// - name
// - address
// - city
// - country
// - coordinates
// - evses
// - time_zone
// - opening_times
//   - twentyfourseven
//   - regular_hours
//
// - last_updated
func (r *Location) ToOcpi() (d ocpi.Location) {
	ocpiEvses := make([]ocpi.EVSE, len(r.Evses))
	for i, evse := range r.Evses {
		ocpiEvses[i] = evse.ToOcpi()
	}
	ocpiRegularHours := make([]ocpi.RegularHours, len(r.OpeningTimesRegularHours))
	for i, ocpiRegularHour := range r.OpeningTimesRegularHours {
		ocpiRegularHours[i] = ocpi.RegularHours{
			Weekday:     ocpiRegularHour.Weekday,
			PeriodBegin: ocpiRegularHour.PeriodBegin,
			PeriodEnd:   ocpiRegularHour.PeriodEnd,
		}
	}
	return ocpi.Location{
		CountryCode: r.CountryCode,
		PartyId:     r.PartyID,
		Id:          r.ID,
		Publish:     r.Publish,
		Name:        r.Name,
		Address:     r.Address,
		City:        r.City,
		Country:     r.Country,
		Coordinates: ocpi.GeoLocation{Latitude: r.CoordinatesLatitude, Longitude: r.CoordinatesLongitude},
		Evses:       &ocpiEvses,
		TimeZone:    r.TimeZone,
		OpeningTimes: &ocpi.Hours{
			Twentyfourseven: r.OpeningTimesTwentyfourseven,
			RegularHours:    &ocpiRegularHours,
		},
		LastUpdated: r.LastUpdated.UTC(),
	}
}

// 必須
// - uid
// - evse_id
// - status
// - connectors
// - last_updated
func (r *Evse) ToOcpi() (d ocpi.EVSE) {
	ocpiConnectors := make([]ocpi.Connector, len(r.Connectors))
	for i, connector := range r.Connectors {
		ocpiConnectors[i] = connector.ToOcpi()
	}
	return ocpi.EVSE{
		Uid:         r.ID,
		EvseId:      r.EvseID,
		Status:      ocpi.Status(r.Status),
		Connectors:  ocpiConnectors,
		LastUpdated: r.LastUpdated.UTC(),
	}
}

// 必須
// - id
// - standard
// - format
// - power_type
// - max_voltage
// - max_amperage
// - last_updated
func (r *Connector) ToOcpi() (d ocpi.Connector) {
	return ocpi.Connector{
		Id:          r.ID,
		Standard:    ocpi.ConnectorType(r.Standard),
		Format:      ocpi.ConnectorFormat(r.Format),
		PowerType:   ocpi.PowerType(r.PowerType),
		MaxVoltage:  r.MaxVoltage,
		MaxAmperage: r.MaxAmperage,
		LastUpdated: r.LastUpdated.UTC(),
	}
}

func FindLocations(dateFrom, dateTo *time.Time, offset, limit *int) ([]*Location, error) {
	var rows []*Location
	query := DB().Where("").Preload("OpeningTimesRegularHours").Preload("Evses.Connectors")
	if dateFrom != nil {
		// Only return Locations that have last_updated after or equal to this Date/Time (inclusive).
		query = query.Where("last_updated >= ?", dateFrom)
	}
	if dateTo != nil {
		// Only return Locations that have last_updated up to this Date/Time, but not including (exclusive).
		query = query.Where("last_updated < ?", dateTo)
	}
	if offset != nil {
		query = query.Offset(*offset)
	}
	if limit != nil {
		query = query.Limit(*limit)
	}
	result := query.Find(&rows)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to Find: %w", result.Error)
	}
	return rows, nil
}

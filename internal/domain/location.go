package domain

type Location struct {
	ID          string
	Name        *string
	Address     string
	Coordinates GeoLocation
	EVSES       []EVSE
}

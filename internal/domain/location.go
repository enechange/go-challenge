package domain

type Location struct {
	ID          string
	Name        *string
	Address     string
	Coordinates GeoLocation
	EVSES       []EVSE
}

type AvailableEVSELocation struct {
	ID        string
	Name      *string
	Address   string
	Latitude  float64
	Longitude float64
	UID       string
	Status    int
}

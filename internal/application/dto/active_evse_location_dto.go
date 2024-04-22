package dto

type AvailableEVSELocation struct {
	ID        string
	Name      *string
	Address   string
	Latitude  float64
	Longitude float64
	UID       string
	Status    int
}

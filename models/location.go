package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	ID        string
	Name      string
	Address   string
	Longitude float64
	Latitude  float64
	Distance  float64
	Evses     string
}

type GeoLocation struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type LocationResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Address     string         `json:"address"`
	Coordinates GeoLocation    `json:"coordinates"`
	Evses       []EvseResponse `json:"evses"`
}

type LocationSearchParameters struct {
    Longitude float64 `form:"longitude" binding:"required"`
    Latitude  float64 `form:"latitude" binding:"required"`
    Radius    int     `form:"radius,default=100"`
}

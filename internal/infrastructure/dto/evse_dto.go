package dto

type EVSE struct {
	LocationID string `gorm:"type:varchar(36);not null;index:idx_location_status,sort:asc"`
	UID        string `gorm:"type:varchar(36);primaryKey"`
	Status     int    `gorm:"type:int;not null;index:idx_location_status,sort:asc"`
}

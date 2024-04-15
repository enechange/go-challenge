package dto

type Location struct {
	ID        string  `gorm:"type:varchar(36);primaryKey"`
	Name      *string `gorm:"type:varchar(255)"`
	Address   string  `gorm:"type:varchar(45);not null"`
	Latitude  string  `gorm:"type:varchar(11);not null"`
	Longitude string  `gorm:"type:varchar(12);not null"`
}

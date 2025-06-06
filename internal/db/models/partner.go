package models

import "time"

type Partner struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	LogoURL   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Offers []Offer `gorm:"foreignKey:PartnerID"`
}

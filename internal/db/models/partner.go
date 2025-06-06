package models

import "time"

type Partner struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	LogoUrl   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Offers []Offer `gorm:"foreignKey:PartnerId"`
}

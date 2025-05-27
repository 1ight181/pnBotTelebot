package models

import "time"

type StatisticsLog struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index"`
	OfferID   uint `gorm:"index"`
	ClickedAt time.Time
	IPAddress string

	User  User
	Offer Offer
}

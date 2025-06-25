package models

import "time"

type StatisticsLog struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint `gorm:"index"`
	OfferId   uint `gorm:"index"`
	ClickedAt time.Time
	IpAddress string

	User  User  `gorm:"foreignKey:UserId;references:Id"`
	Offer Offer `gorm:"foreignKey:OfferId;references:Id"`
}

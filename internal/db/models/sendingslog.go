package models

import "time"

type SendingsLog struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"index"`
	OfferId   uint      `gorm:"index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User  User  `gorm:"foreignKey:UserId;references:Id"`
	Offer Offer `gorm:"foreignKey:OfferId;references:Id"`
}

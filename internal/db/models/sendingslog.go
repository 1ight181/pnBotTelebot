package models

import "time"

type SendingsLog struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"index"`
	OfferId   uint      `gorm:"index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User  User
	Offer Offer
}

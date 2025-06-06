package models

import "time"

type Creative struct {
	Id                        uint `gorm:"primaryKey"`
	PartnerInternalCreativeId string
	OfferId                   uint `gorm:"index"`
	Type                      string
	ResourceUrl               string
	Width                     int
	Height                    int
	CreatedAt                 time.Time `gorm:"autoCreateTime"`
	UpdatedAt                 time.Time `gorm:"autoUpdateTime"`

	Offer Offer
}

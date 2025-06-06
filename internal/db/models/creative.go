package models

import "time"

type Creative struct {
	ID                        uint `gorm:"primaryKey"`
	PartnerInternalCreativeID string
	OfferID                   uint `gorm:"index"`
	Type                      string
	ResourceURL               string
	Width                     int
	Height                    int
	CreatedAt                 time.Time `gorm:"autoCreateTime"`
	UpdatedAt                 time.Time `gorm:"autoUpdateTime"`

	Offer Offer
}

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
	AddedAt                   time.Time
	UpdatedAt                 time.Time

	Offer Offer
}

package models

import (
	"pnBot/internal/db/enums"
	"time"
)

type Offer struct {
	Id                     uint `gorm:"primaryKey"`
	PartnerInternalOfferId string
	Description            string
	Title                  string
	Status                 enums.OfferStatus `gorm:"type:offer_status;index"`
	CategoryId             uint              `gorm:"index"`
	PartnerId              uint              `gorm:"index"`
	TrackingLink           string
	CreatedAt              time.Time `gorm:"autoCreateTime"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime"`
	Payout                 float64

	Category  Category   `gorm:"foreignKey:CategoryId;references:Id"`
	Partner   Partner    `gorm:"foreignKey:PartnerId;references:Id"`
	Creatives []Creative `gorm:"foreignKey:OfferId"`
}

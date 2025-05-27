package models

import (
	enums "pnBot/internal/db/enums"
	"time"
)

type Offer struct {
	ID                     uint `gorm:"primaryKey"`
	PartnerInternalOfferID string
	Description            string
	Title                  string
	Status                 enums.OfferStatus `gorm:"type:offer_status;index"`
	CategoryID             uint              `gorm:"index"`
	PartnerID              uint              `gorm:"index"`
	TrackingLink           string
	AddedAt                time.Time
	UpdatedAt              time.Time
	Payout                 float64

	Category  Category
	Partner   Partner
	Creatives []Creative
}

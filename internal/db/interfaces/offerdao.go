package interfaces

import (
	dbmodels "pnBot/internal/db/models"
	"time"
)

type OfferDao interface {
	GetLastAvailableOffers(userId int64, limit int, offerCooldown time.Time) ([]dbmodels.Offer, error)
	AddSendingLog(userId int64, offerId uint) error
}

package interfaces

import dbmodels "pnBot/internal/db/models"

type OfferDao interface {
	GetNextAvailableOffer(userId int64) (*dbmodels.Offer, error)
	GetLastAvailableOffers(userId int64, limit int) ([]dbmodels.Offer, error)
	AddSendingLog(userId int64, offerId uint) error
}

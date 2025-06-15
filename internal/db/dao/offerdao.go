package dao

import (
	"errors"
	"pnBot/internal/db/enums"
	dberrors "pnBot/internal/db/errors"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"time"

	"gorm.io/gorm"
)

type GormOfferDao struct {
	db *gorm.DB
}

func NewOfferDao(rawDb *gorm.DB) dbifaces.OfferDao {
	return &GormOfferDao{db: rawDb}
}

func (dao *GormOfferDao) AddSendingLog(userId int64, offerId uint) error {
	var user dbmodels.User
	if err := dao.db.Where("tg_id = ?", userId).First(&user).Error; err != nil {
		return err
	}

	var existingLog dbmodels.SendingsLog
	err := dao.db.Where("user_id = ? AND offer_id = ?", user.Id, offerId).First(&existingLog).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logEntry := dbmodels.SendingsLog{
			UserId:  user.Id,
			OfferId: offerId,
		}
		return dao.db.Create(&logEntry).Error
	}

	return dao.db.Model(&existingLog).Update("created_at", time.Now()).Error
}

func (dao *GormOfferDao) GetLastAvailableOffers(userId int64, limit int, offerCooldown time.Time) ([]dbmodels.Offer, error) {
	var offers []dbmodels.Offer
	var user dbmodels.User

	err := dao.db.Preload("PreferredCategories").Where("tg_id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	categoryIDs := make([]uint, 0, len(user.PreferredCategories))
	for _, cat := range user.PreferredCategories {
		categoryIDs = append(categoryIDs, cat.Id)
	}

	if len(categoryIDs) == 0 {
		return nil, nil
	}

	err = dao.db.
		Model(&dbmodels.Offer{}).
		Joins(`LEFT JOIN sendings_logs 
		       ON sendings_logs.offer_id = offers.id 
		      AND sendings_logs.user_id = ? 
		      AND sendings_logs.created_at > NOW() - ? `, user.Id, offerCooldown).
		Where("offers.status = ?", enums.StatusActive).
		Where("offers.category_id IN ?", categoryIDs).
		Order("offers.created_at DESC").
		Limit(limit).
		Preload("Creatives").
		Find(&offers).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dberrors.ErrRecordNotFound
		}
		return nil, err
	}

	return offers, nil
}

package dao

import (
	dbifaces "pnBot/internal/db/interfaces"
	"pnBot/internal/db/models"

	"gorm.io/gorm"
)

type GormUserDao struct {
	db *gorm.DB
}

func NewUserDao(rawDb *gorm.DB) dbifaces.UserDao {
	return &GormUserDao{db: rawDb}
}

func (dao *GormUserDao) GetAllWithBans(search string) ([]models.User, error) {
	var users []models.User

	query := dao.db.Preload("BannedUser")
	if search != "" {
		query = query.Where("fullname ILIKE ? OR username ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Order("id DESC").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *GormUserDao) Delete(id uint) error {
	return dao.db.Delete(&models.User{}, id).Error
}

func (dao *GormUserDao) GetTgIdByUserId(id uint) (int64, error) {
	var user models.User
	if err := dao.db.Select("tg_id").First(&user, id).Error; err != nil {
		return 0, err
	}
	return user.TgId, nil
}

package models

type UserCategory struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	CategoryID uint

	User     User     `gorm:"foreignKey:UserID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}

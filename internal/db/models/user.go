package models

import "time"

type User struct {
	Id           uint  `gorm:"primaryKey"`
	TgId         int64 `gorm:"uniqueIndex;not null"` // Telegram User Id
	ChatId       int64 `gorm:"uniqueIndex;not null"` // Telegram Chat Id
	Username     string
	Fullname     string    //firstnale + lastname
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	IsSubscribed bool

	PreferredCategories []Category `gorm:"many2many:user_categories;"`
	StatisticsLogs      []StatisticsLog
}

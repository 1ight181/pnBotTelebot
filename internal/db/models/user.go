package models

import "time"

type User struct {
	ID           uint  `gorm:"primaryKey"`
	TgID         int64 `gorm:"uniqueIndex;not null"` // Telegram User ID
	ChatID       int64 `gorm:"uniqueIndex;not null"` // Telegram Chat ID
	Username     string
	Name         string    //firstnale + lastname
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	IsSubscribed bool

	PreferredCategories []Category `gorm:"many2many:user_categories;"`
	StatisticsLogs      []StatisticsLog
}

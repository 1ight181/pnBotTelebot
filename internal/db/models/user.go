package models

import "time"

type User struct {
	ID           uint  `gorm:"primaryKey"`
	ChatID       int64 `gorm:"uniqueIndex;not null"` // Telegram Chat ID
	Username     string
	CreatedAt    time.Time
	IsSubscribed bool

	PreferredCategories []Category `gorm:"many2many:user_categories;"`
	StatisticsLogs      []StatisticsLog
}

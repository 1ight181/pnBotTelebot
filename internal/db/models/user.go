package models

import "time"

type User struct {
	Id                    uint  `gorm:"primaryKey"`
	TgId                  int64 `gorm:"uniqueIndex;not null"` // Telegram User Id
	ChatId                int64 `gorm:"uniqueIndex;not null"` // Telegram Chat Id
	Username              string
	Fullname              string
	CreatedAt             time.Time `gorm:"autoCreateTime"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime"`
	IsSubscribed          bool
	NotificationFrequency int `gorm:"default:4"`

	PreferredCategories []Category      `gorm:"many2many:user_categories;"`
	StatisticsLogs      []StatisticsLog `gorm:"foreignKey:UserId"`
	BannedUser          *BannedUser     `gorm:"foreignKey:UserId"`
	SendingsLogs        []SendingsLog   `gorm:"foreignKey:UserId"`
}

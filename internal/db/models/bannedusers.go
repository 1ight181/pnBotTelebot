package models

import "time"

type BannedUser struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	UserID    uint       `gorm:"not null;index"`
	Reason    string     `gorm:"type:text"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	ExpiresAt *time.Time `gorm:"index"`
	CreatedBy string     `gorm:"type:varchar(255)"`

	User User
}

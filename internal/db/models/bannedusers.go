package models

import "time"

type BannedUser struct {
	Id        uint       `gorm:"primaryKey;autoIncrement"`
	UserId    uint       `gorm:"uniqueIndex;not null"`
	Reason    string     `gorm:"type:text"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	ExpiresAt *time.Time `gorm:"index"`
	CreatedBy string     `gorm:"type:varchar(255)"`

	User *User `gorm:"foreignKey:UserId;references:Id"`
}

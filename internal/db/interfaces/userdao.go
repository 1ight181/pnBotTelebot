package interfaces

import (
	"pnBot/internal/db/models"
)

type UserDao interface {
	GetAllWithBans(search string) ([]models.User, error)
	Delete(id uint) error
	GetTgIdByUserId(userId uint) (int64, error)
}

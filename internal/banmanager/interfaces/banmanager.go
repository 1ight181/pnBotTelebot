package interfaces

import "time"

type BanManager interface {
	Ban(userId int64, reason string, duration time.Duration, author string) error
	IsBanned(userId int64) (bool, error)
	Unban(userId int64) error
}

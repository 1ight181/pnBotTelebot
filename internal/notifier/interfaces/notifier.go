package intrefaces

import (
	units "pnBot/internal/notifier/units"
)

type Notifier interface {
	AddUser(userId int64) error
	RemoveUser(userId int64) error
	SetUserFrequency(userId int64, frequency int) error
	GetFrequency(userId int64) (int, error)
	GetFrequencyUnit() (units.FrequencyUnit, error)
	Stop() error
	Start() error
}

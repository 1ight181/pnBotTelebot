package intrefaces

import (
	units "pnBot/internal/notifier/units"
	"time"
)

type Notifier interface {
	AddUser(userId int64, frequency int, offerCooldownDuration time.Duration, frequencyUnit units.FrequencyUnit) error
	RemoveUser(userId int64) error
	SetUserFrequency(userId int64, frequency int, offerCooldownDuration time.Duration, frequencyUnit units.FrequencyUnit) error
}

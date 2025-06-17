package loaders

import (
	conf "pnBot/internal/config/models"
	units "pnBot/internal/notifier/units"
	"strconv"
	"time"
)

func LoadNotifierConfig(notifierConfig conf.Notifier) (time.Duration, int, units.FrequencyUnit) {
	offerCooldownInt, _ := strconv.Atoi(notifierConfig.OfferCooldown)
	offerCooldown := time.Duration(offerCooldownInt)
	defaultFrequency, _ := strconv.Atoi(notifierConfig.DefaultFrequency)
	frequencyUnit, _ := units.ParseFrequencyUnit(notifierConfig.FrequencyUnit)

	return offerCooldown, defaultFrequency, frequencyUnit
}

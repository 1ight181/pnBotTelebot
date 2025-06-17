package models

import (
	"errors"
	units "pnBot/internal/notifier/units"
	"strconv"
)

type Notifier struct {
	OfferCooldown    string `mapstructure:"offer_cooldown"`
	DefaultFrequency string `mapstructure:"defualt_frequency"`
	FrequencyUnit    string `mapstructure:"frequency_unit"`
}

func (n *Notifier) Validate() error {
	if n.OfferCooldown == "" {
		if _, err := strconv.Atoi(n.OfferCooldown); err != nil {
			return errors.New("недопустимое значение для offer_cooldown")
		}
		return errors.New("требуется указание offer_cooldown")
	}
	if n.DefaultFrequency == "" {
		if _, err := strconv.Atoi(n.DefaultFrequency); err != nil {
			return errors.New("недопустимое значение для defualt_frequency")
		}
		return errors.New("требуется указание defualt_frequency")
	}
	if n.FrequencyUnit == "" {
		if _, err := units.ParseFrequencyUnit(n.FrequencyUnit); err != nil {
			return errors.New("недопустимое значение для frequency_unit")
		}
		return errors.New("требуется указание frequency_unit")
	}
	return nil
}

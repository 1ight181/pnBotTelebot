package models

import (
	"errors"
	"strconv"
)

type SpamManager struct {
	MessageLimit  string `mapstructure:"message_limit"`
	Interval      string `mapstructure:"interval"`
	WarnLimit     string `mapstructure:"warn_limit"`
	BanDuration   string `mapstructure:"ban_duration"`
	BanReasonText string `mapstructure:"ban_reason_text"`
	BanAuthor     string `mapstructure:"ban_author"`
}

func (sm *SpamManager) Validate() error {
	if sm.MessageLimit == "" {
		if _, err := strconv.Atoi(sm.MessageLimit); err != nil {
			return errors.New("недопустимое значение лимита сообщений в интервал")
		}
		return errors.New("требуется указание лимита сообщений в интервал")
	}
	if sm.Interval == "" {
		if _, err := strconv.Atoi(sm.Interval); err != nil {
			return errors.New("недопустимое значение длительности интервала в секундах")
		}
		return errors.New("требуется указание длительности интервала в секундах")
	}
	if sm.WarnLimit == "" {
		if _, err := strconv.Atoi(sm.WarnLimit); err != nil {
			return errors.New("недопустимое значение количества предупреждений")
		}
		return errors.New("требуется указание количества предупреждений")
	}
	if sm.BanDuration == "" {
		if _, err := strconv.Atoi(sm.Interval); err != nil {
			return errors.New("недопустимое значение продолжительности бана в часах")
		}
		return errors.New("требуется указание продолжительности бана в часах")
	}
	if sm.BanReasonText == "" {
		return errors.New("требуется указание текста причины бана по спаму")
	}
	if sm.BanAuthor == "" {
		return errors.New("требуется указание, кто выдает бан за спам")
	}
	return nil
}

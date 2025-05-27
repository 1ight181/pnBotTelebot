package callback

import (
	"errors"
	deps "pnBot/internal/bot/processors/dependencies"
	"strings"

	"gopkg.in/telebot.v3"
)

type CallbackProcessor struct {
	dependencies *deps.ProcessorDependencies
}

func New(dependencies *deps.ProcessorDependencies) *CallbackProcessor {
	return &CallbackProcessor{
		dependencies: dependencies,
	}
}

func (p *CallbackProcessor) ProcessCallback(c telebot.Context) error {
	rawData := c.Callback().Data
	data := strings.TrimPrefix(rawData, "\f")

	switch data {
	case "subscribe":
		return p.ProcessSubscribe(c)
	case "unsubscribe":
		return p.ProcessUnsubscribe(c)
	default:
		return errors.New("Получен неизвестный callback: " + data)
	}
}

package command

import (
	"errors"
	deps "pnBot/internal/bot/processors/dependencies"

	"gopkg.in/telebot.v3"
)

type CommandProcessor struct {
	deps *deps.ProcessorDependencies
}

func New(deps *deps.ProcessorDependencies) *CommandProcessor {
	return &CommandProcessor{
		deps: deps,
	}
}

func (p *CommandProcessor) ProcessCommand(c telebot.Context) error {
	data := c.Message().Text

	switch data {
	case "/start":
		return p.ProcessStart(c)
	case "/help":
		return p.ProcessHelp(c)
	default:
		return errors.New("Полученно неизвестное сообщение: " + data)
	}
}

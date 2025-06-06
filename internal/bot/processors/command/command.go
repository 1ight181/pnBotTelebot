package command

import (
	"errors"
	deps "pnBot/internal/bot/processors/dependencies"

	"gopkg.in/telebot.v3"
)

type CommandProcessor struct {
	dependencies *deps.ProcessorDependencies
}

func New(dependencies *deps.ProcessorDependencies) *CommandProcessor {
	return &CommandProcessor{
		dependencies: dependencies,
	}
}

func (cp *CommandProcessor) ProcessCommand(c telebot.Context) error {
	data := c.Message().Text

	switch data {
	case "/start":
		return cp.ProcessStart(c)
	case "/help":
		return cp.ProcessHelp(c)
	default:
		return errors.New("Полученно неизвестное сообщение: " + data)
	}
}

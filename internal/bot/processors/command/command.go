package command

import (
	common "pnBot/internal/bot/processors/common"
	deps "pnBot/internal/bot/processors/dependencies"

	"gopkg.in/telebot.v3"
)

type CommandProcessor struct {
	dependencies *deps.ProcessorDependencies
}

func NewCommandProcessor(dependencies *deps.ProcessorDependencies) *CommandProcessor {
	return &CommandProcessor{
		dependencies: dependencies,
	}
}

func (cp *CommandProcessor) ProcessCommand(c telebot.Context) error {
	if c.Message().Via != nil {
		if c.Message().Via.IsBot {
			return nil
		}
	}
	data := c.Message().Text

	switch data {
	case "/start":
		return cp.ProcessStart(c)
	case "/help":
		return cp.ProcessHelp(c)
	case "/menu":
		return common.ProcessMenu(c, cp.dependencies.TextProvider, cp.dependencies.DbProvider)
	default:
		unknownText := cp.dependencies.TextProvider.GetText("unknown")
		return c.Send(unknownText)
	}
}

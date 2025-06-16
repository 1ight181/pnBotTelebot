package command

import (
	common "pnBot/internal/bot/processors/common"
	deps "pnBot/internal/bot/processors/dependencies"
	"strconv"
	"strings"

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

	processingText := cp.dependencies.TextProvider.GetText("processing")
	message, err := c.Bot().Send(c.Chat(), processingText, &telebot.ReplyMarkup{RemoveKeyboard: true})
	if err != nil {
		return err
	}
	defer func() {
		c.Bot().Delete(message)
	}()

	switch data {
	case "/start":
		return cp.ProcessStart(c)
	case "/help":
		return cp.ProcessHelp(c)
	case "/menu":
		return common.ProcessMenu(c, cp.dependencies.TextProvider, cp.dependencies.DbProvider)
	default:
		everyXHoursButtonText := cp.dependencies.TextProvider.GetButtonText("every_x_hours")
		everyXHoursButtonTextPrefix := strings.Split(everyXHoursButtonText, " ")[0]

		everydayButtonText := cp.dependencies.TextProvider.GetButtonText("everyday")
		everydayButtonTextPrefix := strings.Split(everydayButtonText, " ")[0]
		if strings.HasPrefix(data, everyXHoursButtonTextPrefix) {
			splitedData := strings.Split(data, " ")

			value, err := strconv.Atoi(splitedData[1])
			if err != nil {
				unknownText := cp.dependencies.TextProvider.GetText("unknown")
				return c.Send(unknownText)
			}
			return cp.ProcessHourFrequencySetting(c, value)
		}
		if strings.HasPrefix(data, everydayButtonTextPrefix) {
			return cp.ProcessEverydayFrequencySetting(c)
		}
		unknownText := cp.dependencies.TextProvider.GetText("unknown")
		return c.Send(unknownText)
	}
}

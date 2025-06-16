package command

import (
	"pnBot/internal/bot/processors/keyboards"

	"gopkg.in/telebot.v3"
)

func (cp *CommandProcessor) ProcessHourFrequencySetting(c telebot.Context, value int) error {
	userId := c.Sender().ID

	if err := cp.dependencies.Notifier.SetUserFrequency(userId, value); err != nil {
		return err
	}

	frequencySettedText := cp.dependencies.TextProvider.GetText("frequency_setted")
	removeReplyKeyboard := &telebot.ReplyMarkup{RemoveKeyboard: true}

	if err := c.Send(frequencySettedText, removeReplyKeyboard); err != nil {
		return err
	}
	menuText := cp.dependencies.TextProvider.GetText("menu")
	menuKeyboard := keyboards.GetMenuKeyBoard(cp.dependencies.TextProvider)
	return c.Send(menuText, menuKeyboard)

}

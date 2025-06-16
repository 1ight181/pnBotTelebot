package command

import (
	"pnBot/internal/bot/processors/keyboards"

	"gopkg.in/telebot.v3"
)

func (cp *CommandProcessor) ProcessEverydayFrequencySetting(c telebot.Context) error {
	userId := c.Sender().ID

	frequency := 24

	if err := cp.dependencies.Notifier.SetUserFrequency(userId, frequency); err != nil {
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

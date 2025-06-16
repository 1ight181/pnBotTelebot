package callback

import (
	"pnBot/internal/bot/processors/keyboards"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessFrequencySettings(c telebot.Context) error {
	disabledKeyboard := telebot.ReplyMarkup{}
	processText := cp.dependencies.TextProvider.GetText("process")

	btn := disabledKeyboard.Data(processText, "", "")
	disabledKeyboard.Inline(disabledKeyboard.Row(btn))

	menuText := cp.dependencies.TextProvider.GetText("frequency_settings")
	menuKeyboard := keyboards.GetFrequencyKeyboard(cp.dependencies.TextProvider)

	return c.Send(menuText, menuKeyboard)
}

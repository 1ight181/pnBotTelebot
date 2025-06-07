package callback

import (
	keyboards "pnBot/internal/bot/processors/keyboards"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessApplyFilter(c telebot.Context, data string) error {
	disabledKeyboard := telebot.ReplyMarkup{}
	processText := cp.dependencies.TextProvider.GetText("process")

	btn := disabledKeyboard.Data(processText, "", "")
	disabledKeyboard.Inline(disabledKeyboard.Row(btn))

	err := c.Edit("Подождите, идет обработка", &disabledKeyboard)
	if err != nil {
		return err
	}

	_, selectedCategories, err := cp.parseFilterData(data)
	if err != nil {
		return err
	}

	userId := c.Sender().ID

	if err := cp.addPreferredCategories(userId, selectedCategories); err != nil {
		return err
	}

	menuText := cp.dependencies.TextProvider.GetText("menu")
	menuKeyboard := keyboards.GetMenuKeyBoard(cp.dependencies.TextProvider)

	return c.Edit(menuText, menuKeyboard)
}

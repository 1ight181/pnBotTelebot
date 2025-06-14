package callback

import (
	"pnBot/internal/bot/processors/keyboards"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessFilterSettings(c telebot.Context) error {
	keyboard := telebot.ReplyMarkup{}
	c.Edit("Обработка, подождите...", &keyboard)

	userId := c.Sender().ID

	allCategories, err := cp.getCategories()
	if err != nil {
		return err
	}

	selectedCategories, err := cp.getUserCategories(userId)
	if err != nil {
		return err
	}

	if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
		return err
	}

	categoryFilterText := cp.dependencies.TextProvider.GetText("category_filter")
	categoriesKeyboard := keyboards.GetFilterSettingsKeyboard(allCategories, selectedCategories, cp.dependencies.TextProvider)

	return c.Edit(categoryFilterText, &categoriesKeyboard)
}

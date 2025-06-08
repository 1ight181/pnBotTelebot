package callback

import (
	"pnBot/internal/bot/processors/keyboards"
	sliceutils "pnBot/internal/sliceutils"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessFilterToggle(c telebot.Context, data string) error {
	disabledKeyboard := telebot.ReplyMarkup{}
	processText := cp.dependencies.TextProvider.GetText("process")

	btn := disabledKeyboard.Data(processText, "", "")
	disabledKeyboard.Inline(disabledKeyboard.Row(btn))

	err := c.Edit("Подождите, идет обработка", &disabledKeyboard)
	if err != nil {
		return err
	}

	allCategories, err := cp.getCategories()
	if err != nil {
		return err
	}

	currentToggledCategoryId, selectedCategories, err := cp.parseFilterData(data)
	if err != nil {
		return err
	}

	if sliceutils.In(currentToggledCategoryId, selectedCategories) {
		selectedCategories = sliceutils.RemoveByValue(selectedCategories, currentToggledCategoryId)
	} else {
		selectedCategories = append(selectedCategories, currentToggledCategoryId)
	}

	c.Respond(&telebot.CallbackResponse{
		Text:      "",
		ShowAlert: false,
	})

	categoriesKeyboard := keyboards.GetFilterToggleKeyboard(allCategories, selectedCategories, cp.dependencies.TextProvider)
	categoryFilterText := cp.dependencies.TextProvider.GetText("category_filter")

	return c.Edit(categoryFilterText, &categoriesKeyboard)
}

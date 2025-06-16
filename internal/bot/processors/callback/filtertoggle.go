package callback

import (
	"pnBot/internal/bot/processors/keyboards"
	sliceutils "pnBot/internal/sliceutils"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessFilterToggle(c telebot.Context, data string) error {
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

	if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
		return err
	}

	categoriesKeyboard := keyboards.GetFilterToggleKeyboard(allCategories, selectedCategories, cp.dependencies.TextProvider)
	categoryFilterText := cp.dependencies.TextProvider.GetText("category_filter")

	return c.Edit(categoryFilterText, &categoriesKeyboard)
}

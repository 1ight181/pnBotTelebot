package callback

import (
	sliceutils "pnBot/internal/sliceutils"
	"strings"

	"fmt"

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

	var selectedCategoriesIdsParts []string
	var selectedCategorisIdsComparable []int
	for _, selectedCategory := range selectedCategories {
		selectedCategoriesIdsParts = append(selectedCategoriesIdsParts, fmt.Sprint(selectedCategory.Id))
		selectedCategorisIdsComparable = append(selectedCategorisIdsComparable, int(selectedCategory.Id))
	}
	selectedCategoriesIds := strings.Join(selectedCategoriesIdsParts, ",")

	categoriesKeyboard := telebot.ReplyMarkup{}
	var categoriesButtonRows []telebot.Row
	for _, category := range allCategories {
		categoryName := category.Name
		currentToggledCetegoryId := fmt.Sprint(category.Id)

		var categoryText string
		if sliceutils.In(int(category.Id), selectedCategorisIdsComparable) {
			categoryText = fmt.Sprintf("%s ✅", categoryName)
		} else {
			categoryText = fmt.Sprintf("%s ❌", categoryName)
		}

		categoryButton := categoriesKeyboard.Data(
			categoryText,
			"filter",
			selectedCategoriesIds,
			currentToggledCetegoryId,
		)

		categoriesButtons := categoriesKeyboard.Row(categoryButton)

		categoriesButtonRows = append(categoriesButtonRows, categoriesButtons)
	}

	applyFilterButtonText := cp.dependencies.TextProvider.GetButtonText("apply_filter")

	applyFilterButton := categoriesKeyboard.Data(
		applyFilterButtonText,
		"apply_filter",
		selectedCategoriesIds,
		"",
	)

	applyFilterRow := categoriesKeyboard.Row(applyFilterButton)

	allRows := append(categoriesButtonRows, applyFilterRow)

	categoriesKeyboard.Inline(allRows...)

	categoryFilterText := cp.dependencies.TextProvider.GetText("category_filter")

	c.Respond(&telebot.CallbackResponse{
		Text:      "",
		ShowAlert: false,
	})

	return c.Edit(categoryFilterText, &categoriesKeyboard)
}

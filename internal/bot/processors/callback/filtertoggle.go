package callback

import (
	"fmt"
	"strings"

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

	var selectedCategoriesIdsParts []string
	for _, selectedCategory := range selectedCategories {
		selectedCategoriesIdsParts = append(selectedCategoriesIdsParts, fmt.Sprint(selectedCategory))
	}
	selectedCategoriesIds := strings.Join(selectedCategoriesIdsParts, ",")

	categoriesKeyboard := telebot.ReplyMarkup{}
	var categoriesButtonRows []telebot.Row
	for _, category := range allCategories {
		categoryName := category.Name
		currentToggledCategoryId := fmt.Sprint(category.Id)

		var categoryText string
		if sliceutils.In(int(category.Id), selectedCategories) {
			categoryText = fmt.Sprintf("%s ✅", categoryName)
		} else {
			categoryText = fmt.Sprintf("%s ❌", categoryName)
		}

		categoryButton := categoriesKeyboard.Data(
			categoryText,
			"filter",
			selectedCategoriesIds,
			currentToggledCategoryId,
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

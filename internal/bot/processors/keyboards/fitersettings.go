package keyboards

import (
	"fmt"
	botifaces "pnBot/internal/bot/interfaces"
	"pnBot/internal/db/models"
	"pnBot/internal/sliceutils"
	"strings"

	"gopkg.in/telebot.v3"
)

func GetFilterSettingsKeyboard(
	allCategories []models.Category,
	selectedCategories []models.Category,
	textProvider botifaces.TextProvider,
) telebot.ReplyMarkup {
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

	applyFilterButtonText := textProvider.GetButtonText("apply_filter")

	applyFilterButton := categoriesKeyboard.Data(
		applyFilterButtonText,
		"apply_filter",
		selectedCategoriesIds,
		"",
	)

	applyFilterRow := categoriesKeyboard.Row(applyFilterButton)

	allRows := append(categoriesButtonRows, applyFilterRow)

	categoriesKeyboard.Inline(allRows...)

	return categoriesKeyboard
}

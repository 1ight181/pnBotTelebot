package keyboards

import (
	"fmt"
	botifaces "pnBot/internal/bot/interfaces"
	"pnBot/internal/db/models"
	"pnBot/internal/sliceutils"
	"strings"

	"gopkg.in/telebot.v3"
)

func GetFilterToggleKeyboard(
	allCategories []models.Category,
	selectedCategories []int,
	textProvider botifaces.TextProvider,
) telebot.ReplyMarkup {
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

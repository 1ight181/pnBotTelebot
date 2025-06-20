package category

import (
	ctx "context"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func UpdateCategoryGet(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		backgroundContext := ctx.Background()
		HXRequest := context.Header("HX-Request")
		isHtmx := HXRequest == "true"

		categoryIdStr := context.Query("category_id")
		categoryId, _ := strconv.Atoi(categoryIdStr)

		var categories []dbmodels.Category
		if err := db.Find(backgroundContext, &categories); err != nil {
			return context.Status(500).SendString("Ошибка загрузки категорий")
		}

		var category dbmodels.Category
		if categoryId != 0 {
			if err := db.Find(backgroundContext, &category, categoryId); err != nil {
				return context.Status(404).SendString("Категория не найдена")
			}
		}

		data := map[string]interface{}{
			"Categories": categories,
			"Category":   category,
		}

		if isHtmx {
			if categoryId == 0 {
				return context.SendString("")
			}
			return context.Render(200, "categoryedit", data)
		}

		return context.Render(200, "updatecategory", data)
	}
}

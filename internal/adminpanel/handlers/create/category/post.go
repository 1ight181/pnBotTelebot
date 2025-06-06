package category

import (
	ctx "context"
	"fmt"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
)

func CategoryPost(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		name := context.FormValue("name")
		if name == "" {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Имя категории обязательно</div>")
		}

		newCategory := dbmodels.Category{Name: name}
		if err := db.Create(contextBackground, &newCategory); err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Ошибка при создании категории</div>")
		}

		var categories []dbmodels.Category
		if err := db.Find(contextBackground, &categories); err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Ошибка при загрузке категорий</div>")
		}

		response := fmt.Sprintf(`
			<div class="success-box">Категория "%s" успешно добавлена!</div>

			<select id="category-select" name="category_id" hx-swap-oob="true">
		`, newCategory.Name)

		for _, category := range categories {
			selected := ""
			if category.ID == newCategory.ID {
				selected = "selected"
			}
			response += fmt.Sprintf(`<option value="%d" %s>%s</option>`, category.ID, selected, category.Name)
		}
		response += "</select>"

		return context.Status(200).Type("text/html").SendString(response)
	}
}

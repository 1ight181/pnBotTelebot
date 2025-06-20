package category

import (
	ctx "context"
	"fmt"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func UpdateCategoryPost(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()
		categoryId, err := strconv.Atoi(context.FormValue("category_id"))
		if err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Ошибка при попытке получить id категории </div>")
		}

		name := context.FormValue("name")
		if name == "" {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Имя категории обязательно</div>")
		}

		var category dbmodels.Category
		if err := db.First(contextBackground, &category, "id = ?", categoryId); err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Категория не найдена</div>")
		}

		category.Name = name
		if err := db.Save(contextBackground, &category); err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Ошибка при обновлении категории</div>")
		}

		var categories []dbmodels.Category
		if err := db.Find(contextBackground, &categories); err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Ошибка при загрузке категорий</div>")
		}

		response := `
			<div class="success-box">Категория успешно обновлена!</div>

			<select 
				id="category-select" 
				hx-swap-oob="true" 
				class="input"
				hx-get="/update/categories"
				hx-target="#category-edit-form"
				hx-swap="innerHTML"
				name="category_id"
			>
		`

		response += `<option value="">-- Выберите категорию для редактирования --</option>`

		for _, category := range categories {
			selected := ""
			if category.Id == uint(categoryId) {
				selected = "selected"
			}
			response += fmt.Sprintf(`<option value="%d" %s>%s</option>`, category.Id, selected, category.Name)
		}
		response += "</select>"

		return context.Status(200).Type("text/html").SendString(response)
	}
}

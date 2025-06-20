package category

import (
	ctx "context"
	"fmt"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func DeleteCategoryPost(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		ctxBg := ctx.Background()

		idStr := context.FormValue("category_id")
		if idStr == "" {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Не выбран ID категории</div>`)
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Некорректный ID категории</div>`)
		}

		category := dbmodels.Category{Id: uint(id)}
		if err := db.Delete(ctxBg, &category); err != nil {
			return context.Status(200).Type("text/html").SendString(fmt.Sprintf(`<div class="error-box">Ошибка при удалении категории: %v</div>`, err))
		}

		var categories []dbmodels.Category
		if err := db.Find(ctxBg, &categories); err != nil {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Ошибка при загрузке категорий</div>`)
		}

		response := `<div class="success-box">Категория успешно удалена!</div>
		
		<select class="input" id="category-select" name="category_id" hx-swap-oob="true">`

		response += `<option value="">-- Выберите категорию --</option>`

		for _, c := range categories {
			response += fmt.Sprintf(`<option value="%d">%s</option>`, c.Id, c.Name)
		}
		response += `</select>`

		return context.Status(200).Type("text/html").SendString(response)
	}
}

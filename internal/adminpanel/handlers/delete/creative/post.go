package creative

import (
	ctx "context"
	"fmt"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func DeleteCreativePost(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		ctxBg := ctx.Background()

		idStr := context.FormValue("creative_id")
		if idStr == "" {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Не выбран ID креатива</div>`)
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Некорректный ID креатива</div>`)
		}

		creative := dbmodels.Creative{Id: uint(id)}
		if err := db.Delete(ctxBg, &creative); err != nil {
			return context.Status(200).Type("text/html").SendString(fmt.Sprintf(`<div class="error-box">Ошибка при удалении креатива: %v</div>`, err))
		}

		var creatives []dbmodels.Creative
		if err := db.Find(ctxBg, &creatives); err != nil {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Ошибка при загрузке креативов</div>`)
		}

		response := `<div class="success-box">Креатив успешно удалён!</div>
		
		<select class="input" id="creative-select" name="creative_id" hx-swap-oob="true">`

		response += `<option value="">-- Выберите категорию --</option>`

		for _, c := range creatives {
			response += fmt.Sprintf(`<option value="%d">%s</option>`, c.Id, c.PartnerInternalCreativeId)
		}
		response += `</select>`

		return context.Status(200).Type("text/html").SendString(response)
	}
}

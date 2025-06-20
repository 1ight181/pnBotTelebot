package partner

import (
	ctx "context"
	"fmt"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func DeletePartnerPost(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		ctxBg := ctx.Background()

		idStr := context.FormValue("partner_id")
		if idStr == "" {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Не выбран ID партнёра</div>`)
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Некорректный ID партнёра</div>`)
		}

		partner := dbmodels.Partner{Id: uint(id)}
		if err := db.Delete(ctxBg, &partner); err != nil {
			return context.Status(200).Type("text/html").SendString(fmt.Sprintf(`<div class="error-box">Ошибка при удалении партнёра: %v</div>`, err))
		}

		var partners []dbmodels.Partner
		if err := db.Find(ctxBg, &partners); err != nil {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Ошибка при загрузке партнёров</div>`)
		}

		response := `<div class="success-box">Партнёр успешно удалён!</div>

		<select class="input" id="partner-select" name="partner_id" hx-swap-oob="true">`

		response += `<option value="">-- Выберите партнера --</option>`

		for _, p := range partners {
			response += fmt.Sprintf(`<option value="%d">%s</option>`, p.Id, p.Name)
		}
		response += `</select>`

		return context.Status(200).Type("text/html").SendString(response)
	}
}

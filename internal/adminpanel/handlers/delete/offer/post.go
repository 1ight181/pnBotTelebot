package offer

import (
	ctx "context"
	"fmt"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func DeleteOfferPost(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		ctxBg := ctx.Background()

		idStr := context.FormValue("offer_id")
		if idStr == "" {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Не выбран ID оффера</div>`)
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Некорректный ID оффера</div>`)
		}

		offer := dbmodels.Offer{Id: uint(id)}
		if err := db.Delete(ctxBg, &offer); err != nil {
			return context.Status(200).Type("text/html").SendString(fmt.Sprintf(`<div class="error-box">Ошибка при удалении оффера: %v</div>`, err))
		}

		var offers []dbmodels.Offer
		if err := db.Find(ctxBg, &offers); err != nil {
			return context.Status(200).Type("text/html").SendString(`<div class="error-box">Ошибка при загрузке офферов</div>`)
		}

		response := `<div class="success-box">Оффер успешно удалён!</div>
		
		<select class="input" id="offer-select" name="offer_id" hx-swap-oob="true">`

		response += `<option value="">-- Выберите категорию --</option>`

		for _, o := range offers {
			response += fmt.Sprintf(`<option value="%d">%s</option>`, o.Id, o.Title)
		}
		response += `</select>`

		return context.Status(200).Type("text/html").SendString(response)
	}
}

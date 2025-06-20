package offer

import (
	ctx "context"
	"fmt"
	"strconv"

	adminifaces "pnBot/internal/adminpanel/interfaces"
	enums "pnBot/internal/db/enums"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
)

func UpdateOfferPost(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		offerID, err := strconv.Atoi(context.FormValue("id"))
		if err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Некорректный ID оффера</div>")
		}

		var offer dbmodels.Offer
		if err := db.First(contextBackground, &offer, "id = ?", offerID); err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Оффер не найден</div>")
		}

		partnerIDStr := context.FormValue("partner_id")
		categoryIDStr := context.FormValue("category_id")
		partnerInternalOfferID := context.FormValue("partner_internal_offer_id")
		title := context.FormValue("title")
		description := context.FormValue("description")
		statusStr := context.FormValue("status")
		trackingLink := context.FormValue("tracking_link")
		payoutStr := context.FormValue("payout")

		if partnerIDStr == "" || categoryIDStr == "" || partnerInternalOfferID == "" || title == "" {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Обязательные поля не заполнены</div>")
		}

		partnerID, err := strconv.ParseUint(partnerIDStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Некорректный partner_id</div>")
		}
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Некорректный category_id</div>")
		}
		payout, _ := strconv.ParseFloat(payoutStr, 64)

		offer.PartnerId = uint(partnerID)
		offer.CategoryId = uint(categoryID)
		offer.PartnerInternalOfferId = partnerInternalOfferID
		offer.Title = title
		offer.Description = description
		offer.Status = enums.OfferStatus(statusStr)
		offer.TrackingLink = trackingLink
		offer.Payout = payout

		if err := db.Save(contextBackground, &offer); err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Ошибка при обновлении оффера</div>")
		}

		// Загрузка всех офферов
		var offers []dbmodels.Offer
		if err := db.Find(contextBackground, &offers); err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Ошибка при загрузке офферов</div>")
		}

		// Формирование HTML-ответа
		response := `
			<div class="success-box">Оффер успешно обновлён!</div>

			<select id="offer-select" name="offer_id" hx-swap-oob="true"
				hx-get="/update/offers"
				hx-target="#offer-edit-form"
				hx-swap="innerHTML"
				class="input"
			>`

		response += `<option value="">-- Выберите оффер --</option>`

		for _, offer := range offers {
			selected := ""
			if offer.Id == uint(offerID) {
				selected = "selected"
			}
			response += fmt.Sprintf(`<option value="%d" %s>%s</option>`, offer.Id, selected, offer.Title)
		}
		response += "</select>"

		return context.Status(200).Type("text/html").SendString(response)
	}
}

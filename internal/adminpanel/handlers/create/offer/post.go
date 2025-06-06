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

func OfferPost(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		partnerIDStr := context.FormValue("partner_id")
		categoryIDStr := context.FormValue("category_id")
		partnerInternalOfferID := context.FormValue("partner_internal_offer_id")
		title := context.FormValue("title")
		description := context.FormValue("description")
		statusStr := context.FormValue("status")
		trackingLink := context.FormValue("tracking_link")
		payoutStr := context.FormValue("payout")

		if partnerIDStr == "" || categoryIDStr == "" || partnerInternalOfferID == "" || title == "" {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Обязательные поля не заполнены</div>")
		}

		partnerID, err := strconv.ParseUint(partnerIDStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Некорректный partner_id</div>")
		}
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Некорректный category_id</div>")
		}
		payout, _ := strconv.ParseFloat(payoutStr, 64)

		newOffer := dbmodels.Offer{
			PartnerInternalOfferId: partnerInternalOfferID,
			Title:                  title,
			Description:            description,
			Status:                 enums.OfferStatus(statusStr),
			CategoryId:             uint(categoryID),
			PartnerId:              uint(partnerID),
			TrackingLink:           trackingLink,
			Payout:                 payout,
		}

		if err := db.Create(contextBackground, &newOffer); err != nil {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при создании оффера</div>")
		}

		var offers []dbmodels.Offer
		if err := db.Find(contextBackground, &offers); err != nil {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при загрузке офферов</div>")
		}

		// Формируем HTML-ответ
		response := fmt.Sprintf(`
			<div class="success-box"> Оффер "%s" успешно добавлен! </div>

			<select id="offer-select" name="offer_id" hx-swap-oob="true">
		`, newOffer.Title)

		for _, offer := range offers {
			selected := ""
			if offer.Id == newOffer.Id {
				selected = "selected"
			}
			response += fmt.Sprintf(`<option value="%d" %s>%s</option>`, offer.Id, selected, offer.Title)
		}
		response += "</select>"

		return context.Status(200).Type("text/html").SendString(response)
	}
}

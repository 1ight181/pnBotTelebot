package offer

import (
	ctx "context"
	"fmt"
	"strconv"
	"time"

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
			return context.SendString(400, "Обязательные поля не заполнены")
		}

		partnerID, err := strconv.ParseUint(partnerIDStr, 10, 64)
		if err != nil {
			return context.SendString(400, "Некорректный partner_id")
		}
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			return context.SendString(400, "Некорректный category_id")
		}
		payout, _ := strconv.ParseFloat(payoutStr, 64)

		newOffer := dbmodels.Offer{
			PartnerInternalOfferID: partnerInternalOfferID,
			Title:                  title,
			Description:            description,
			Status:                 enums.OfferStatus(statusStr),
			CategoryID:             uint(categoryID),
			PartnerID:              uint(partnerID),
			TrackingLink:           trackingLink,
			Payout:                 payout,
			AddedAt:                time.Now(),
			UpdatedAt:              time.Now(),
		}

		if err := db.Create(contextBackground, &newOffer); err != nil {
			return context.SendString(500, "Ошибка при создании оффера")
		}

		var offers []dbmodels.Offer
		if err := db.Find(contextBackground, &offers); err != nil {
			return context.SendString(500, "Ошибка при загрузке офферов")
		}

		// Формируем HTML-ответ
		response := fmt.Sprintf(`
			<div id="offer-result" hx-swap-oob="true" style="color:green;">
				Оффер "%s" успешно добавлен!
			</div>

			<select id="offer-select" name="offer_id" hx-swap-oob="true">
		`, newOffer.Title)

		for _, offer := range offers {
			selected := ""
			if offer.ID == newOffer.ID {
				selected = "selected"
			}
			response += fmt.Sprintf(`<option value="%d" %s>%s</option>`, offer.ID, selected, offer.Title)
		}
		response += "</select>"

		return context.SendString(200, response)
	}
}

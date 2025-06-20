package offer

import (
	ctx "context"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func UpdateOfferGet(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		backgroundContext := ctx.Background()
		HXRequest := context.Header("HX-Request")
		isHtmx := HXRequest == "true"

		offerIdStr := context.Query("offer_id")
		offerId, _ := strconv.Atoi(offerIdStr)

		var categories []dbmodels.Category
		if err := db.Find(backgroundContext, &categories); err != nil {
			return context.Status(500).SendString("Ошибка загрузки категорий")
		}

		var partners []dbmodels.Partner
		if err := db.Find(backgroundContext, &partners); err != nil {
			return context.Status(500).SendString("Ошибка загрузки партнёров")
		}

		var offers []dbmodels.Offer
		if err := db.Find(backgroundContext, &offers); err != nil {
			return context.Status(500).SendString("Ошибка загрузки офферов")
		}

		var offer dbmodels.Offer
		if offerId != 0 {
			if err := db.Find(backgroundContext, &offer, offerId); err != nil {
				return context.Status(404).SendString("Оффер не найден")
			}
		}

		data := map[string]interface{}{
			"Categories": categories,
			"Partners":   partners,
			"Offers":     offers,
			"Offer":      offer,
		}

		if isHtmx {
			if offerId == 0 {
				return context.SendString("")
			}
			return context.Render(200, "offeredit", data)
		}

		return context.Render(200, "updateoffer", data)
	}
}

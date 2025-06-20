package partner

import (
	ctx "context"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func UpdatePartnerGet(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		backgroundContext := ctx.Background()
		HXRequest := context.Header("HX-Request")
		isHtmx := HXRequest == "true"

		partnerIdStr := context.Query("partner_id")
		partnerId, _ := strconv.Atoi(partnerIdStr)

		var partners []dbmodels.Partner
		if err := db.Find(backgroundContext, &partners); err != nil {
			return context.Status(500).SendString("Ошибка загрузки партнёров")
		}

		var partner dbmodels.Partner
		if partnerId != 0 {
			if err := db.Find(backgroundContext, &partner, partnerId); err != nil {
				return context.Status(404).SendString("Партнёр не найден")
			}
		}

		data := map[string]interface{}{
			"Partners": partners,
			"Partner":  partner,
		}

		if isHtmx {
			if partnerId == 0 {
				return context.SendString("")
			}
			return context.Render(200, "partneredit", data)
		}

		return context.Render(200, "updatepartner", data)
	}
}

package creative

import (
	ctx "context"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"strconv"
)

func UpdateCreativeGet(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		backgroundContext := ctx.Background()
		HXRequest := context.Header("HX-Request")
		isHtmx := HXRequest == "true"

		creativeIdStr := context.Query("creative_id")
		creativeId, _ := strconv.Atoi(creativeIdStr)

		var creatives []dbmodels.Creative
		if err := db.Find(backgroundContext, &creatives); err != nil {
			return context.Status(500).SendString("Ошибка загрузки креативов")
		}

		var offers []dbmodels.Offer
		if err := db.Find(backgroundContext, &offers); err != nil {
			return context.Status(500).SendString("Ошибка загрузки офферов")
		}

		var creative dbmodels.Creative
		if creativeId != 0 {
			if err := db.Find(backgroundContext, &creative, creativeId); err != nil {
				return context.Status(404).SendString("Креатив не найден")
			}
		}

		data := map[string]interface{}{
			"Creatives": creatives,
			"Offers":    offers,
			"Creative":  creative,
		}

		if isHtmx {
			if creativeId == 0 {
				return context.SendString("")
			}
			return context.Render(200, "creativeedit", data)
		}

		return context.Render(200, "updatecreative", data)
	}
}

package delete

import (
	ctx "context"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
)

func DeleteGet(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		backgroundContext := ctx.Background()

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

		var creatives []dbmodels.Creative
		if err := db.Find(backgroundContext, &creatives); err != nil {
			return context.Status(500).SendString("Ошибка загрузки креативов")
		}

		data := map[string]any{
			"Categories": categories,
			"Partners":   partners,
			"Offers":     offers,
			"Creatives":  creatives,
		}

		return context.Render(200, "deleteform", data)
	}
}

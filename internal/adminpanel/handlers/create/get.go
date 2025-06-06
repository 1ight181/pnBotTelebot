package create

import (
	ctx "context"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
)

func CreateGet(db dbifaces.DataBaseProvider) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		var categories []dbmodels.Category
		backgroundContext := ctx.Background()

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

		data := map[string]any{
			"Categories": categories,
			"Partners":   partners,
			"Offers":     offers,
		}

		return context.Render(200, "createform", data)
	}
}

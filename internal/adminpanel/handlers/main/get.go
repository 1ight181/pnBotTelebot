package adminpanel

import (
	ctx "context"
	"pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
)

type MenuData struct {
	HasCategories bool
	HasPartners   bool
	HasOffers     bool
	HasCreatives  bool
}

func MainGet(db dbifaces.DataBaseProvider) interfaces.HandlerFunc {
	return func(context interfaces.Context) error {
		c := ctx.Background()

		checkExists := func(model any) (bool, error) {
			var count int64
			err := db.Count(c, model, &count)
			return count > 0, err
		}

		hasCategories, _ := checkExists(&dbmodels.Category{})
		hasPartners, _ := checkExists(&dbmodels.Partner{})
		hasOffers, _ := checkExists(&dbmodels.Offer{})
		hasCreatives, _ := checkExists(&dbmodels.Creative{})

		data := MenuData{
			HasCategories: hasCategories,
			HasPartners:   hasPartners,
			HasOffers:     hasOffers,
			HasCreatives:  hasCreatives,
		}

		return context.Render(200, "main", data)
	}
}

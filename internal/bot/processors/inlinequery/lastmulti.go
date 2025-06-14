package inlinequery

import (
	"fmt"
	"pnBot/internal/bot/processors/common"
	"strconv"

	"gopkg.in/telebot.v3"
)

func (iqp *InlineQueryProcessor) ProcessLastMultiple(c telebot.Context, countStr string) error {
	userId := c.Sender().ID

	limit, err := strconv.Atoi(countStr)
	if err != nil || limit <= 0 {
		limit = 5 // дефолтное число
	}
	if limit > 50 {
		limit = 50 // максимум 50
	}

	offers, err := iqp.dependencies.OfferDao.GetLastAvailableOffers(userId, limit)
	if err != nil || len(offers) == 0 {
		title := iqp.dependencies.TextProvider.GetInlineQueryTitle("no_available_offer")
		description := iqp.dependencies.TextProvider.GetInlineQueryDescription("no_available_offer")
		result := &telebot.ArticleResult{
			Title:       title,
			Description: description,
			Text:        common.EscapeMarkdownV2(description),
		}
		result.SetResultID("no_offer")
		return c.Answer(&telebot.QueryResponse{
			Results:   []telebot.Result{result},
			CacheTime: 1,
		})
	}

	results := make([]telebot.Result, 0, len(offers))
	for i, offer := range offers {
		text := fmt.Sprintf("\n\n%s\n\n",
			common.WrapURLsWithPreviousWord(common.EscapeMarkdownV2(offer.Description)),
		)

		result := &telebot.ArticleResult{
			Title:       fmt.Sprintf("Акция #%d:", i+1),
			Description: offer.Title,
			Text:        text,
		}
		result.SetResultID(fmt.Sprintf("last_offer_%d", offer.Id))
		results = append(results, result)
	}

	return c.Answer(&telebot.QueryResponse{
		Results:   results,
		CacheTime: 5,
	})
}

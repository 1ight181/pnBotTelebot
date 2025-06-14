package inlinequery

import (
	"fmt"
	"pnBot/internal/bot/processors/common"

	"gopkg.in/telebot.v3"
)

func (iqp *InlineQueryProcessor) ProcessLastOne(c telebot.Context) error {
	userId := c.Sender().ID

	offers, err := iqp.dependencies.OfferDao.GetLastAvailableOffers(userId, 1)
	if err != nil {
		return err
	}
	if len(offers) == 0 {
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

	offer := offers[0]

	text := fmt.Sprintf("*%s*\n\n%s\n\n",
		common.EscapeMarkdownV2(offer.Title),
		common.WrapURLsWithPreviousWord(common.EscapeMarkdownV2(offer.Description)),
	)

	result := &telebot.ArticleResult{
		Title:       "Актуальная акция!",
		Description: offer.Title,
		Text:        text,
	}
	result.SetResultID(fmt.Sprintf("last_offer_%d", offer.Id))

	return c.Answer(&telebot.QueryResponse{
		Results:   []telebot.Result{result},
		CacheTime: 5,
	})
}

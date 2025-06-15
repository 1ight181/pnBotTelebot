package inlinequery

import (
	"errors"
	"fmt"
	"pnBot/internal/bot/processors/common"
	"strconv"
	"time"

	dberrors "pnBot/internal/db/errors"

	"gopkg.in/telebot.v3"
)

func (iqp *InlineQueryProcessor) ProcessLastMultiple(c telebot.Context, recordCount string) error {
	userId := c.Sender().ID

	limit, err := strconv.Atoi(recordCount)
	if err != nil || limit <= 0 {
		limit = 5
	}
	if limit > 50 {
		limit = 50
	}

	offerCooldown := time.Now().Add(-iqp.dependencies.OfferCooldownDuration)
	offers, err := iqp.dependencies.OfferDao.GetLastAvailableOffers(userId, limit, offerCooldown)
	if errors.Is(err, dberrors.ErrRecordNotFound) {
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
	} else if err != nil {
		errorTitle := iqp.dependencies.TextProvider.GetInlineQueryTitle("error")
		errorDescription := iqp.dependencies.TextProvider.GetInlineQueryDescription("error")
		result := &telebot.ArticleResult{
			Title:       errorTitle,
			Description: errorDescription,
			Text:        common.EscapeMarkdownV2(errorDescription),
		}
		result.SetResultID("error")
		if answerErr := c.Answer(&telebot.QueryResponse{
			Results:   []telebot.Result{result},
			CacheTime: 1,
		}); answerErr != nil {
			return answerErr
		}
		return err
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

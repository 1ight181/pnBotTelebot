package inlinequery

import (
	"errors"
	"fmt"
	"pnBot/internal/bot/processors/common"
	dberrors "pnBot/internal/db/errors"
	"time"

	"gopkg.in/telebot.v3"
)

func (iqp *InlineQueryProcessor) ProcessLastOne(c telebot.Context) error {
	userId := c.Sender().ID

	offerCooldown := time.Now().Add(-iqp.dependencies.OfferCooldownDuration)
	offers, err := iqp.dependencies.OfferDao.GetLastAvailableOffers(userId, 1, offerCooldown)
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

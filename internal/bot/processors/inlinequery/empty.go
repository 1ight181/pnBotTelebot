package inlinequery

import (
	"pnBot/internal/bot/processors/common"

	"gopkg.in/telebot.v3"
)

func (iqp *InlineQueryProcessor) ProcessEmpty(c telebot.Context) error {
	title := iqp.dependencies.TextProvider.GetInlineQueryTitle("empty")
	description := iqp.dependencies.TextProvider.GetInlineQueryDescription("empty")
	result := &telebot.ArticleResult{
		Title:       title,
		Description: description,
		Text:        common.EscapeMarkdownV2(description),
	}
	result.SetResultID("empty_result")

	return c.Answer(&telebot.QueryResponse{
		Results:   []telebot.Result{result},
		CacheTime: 1,
	})
}

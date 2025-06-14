package inlinequery

import (
	"pnBot/internal/bot/processors/common"

	"gopkg.in/telebot.v3"
)

func (iqp *InlineQueryProcessor) ProcessUnknown(c telebot.Context) error {
	title := iqp.dependencies.TextProvider.GetInlineQueryTitle("unknown")
	description := iqp.dependencies.TextProvider.GetInlineQueryDescription("unknown")
	result := &telebot.ArticleResult{
		Title:       title,
		Description: description,
		Text:        common.EscapeMarkdownV2(description),
	}
	result.SetResultID("unknown_inline")

	return c.Answer(&telebot.QueryResponse{
		Results:   []telebot.Result{result},
		CacheTime: 1,
	})
}

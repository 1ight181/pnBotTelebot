package inlinequery

import (
	dependencies "pnBot/internal/bot/processors/dependencies"
	"strings"

	"gopkg.in/telebot.v3"
)

type InlineQueryProcessor struct {
	dependencies *dependencies.ProcessorDependencies
}

func NewInlineQueryProcessor(deps *dependencies.ProcessorDependencies) *InlineQueryProcessor {
	return &InlineQueryProcessor{dependencies: deps}
}

func (iqp *InlineQueryProcessor) ProcessInlineQuery(c telebot.Context) error {
	query := c.Query()
	text := query.Text

	if strings.HasPrefix(text, "last") {
		parts := strings.Fields(text)
		if len(parts) == 1 {
			return iqp.ProcessLastOne(c)
		} else if len(parts) == 2 {
			return iqp.ProcessLastMultiple(c, parts[1])
		}
	}

	switch text {
	case "":
		return iqp.ProcessEmpty(c)
	default:
		return iqp.ProcessUnknown(c)
	}
}

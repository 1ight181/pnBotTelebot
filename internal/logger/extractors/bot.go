package extractors

import (
	"context"
	"pnBot/internal/logger/contextkeys"
)

type BotContextExtractor struct{}

func (e *BotContextExtractor) Extract(ctx context.Context) map[string]interface{} {
	fields := make(map[string]interface{})

	if userID, ok := ctx.Value(contextkeys.UserIDKey).(int64); ok {
		fields["user_id"] = userID
	}
	if chatID, ok := ctx.Value(contextkeys.ChatIDKey).(int64); ok {
		fields["chat_id"] = chatID
	}
	if text, ok := ctx.Value(contextkeys.TextKey).(string); ok {
		fields["text"] = text
	}
	if data, ok := ctx.Value(contextkeys.DataKey).(string); ok {
		fields["data"] = data
	}

	return fields
}

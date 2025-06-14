package middleware

import (
	"context"
	"pnBot/internal/logger/contextkeys"
	loggerifaces "pnBot/internal/logger/interfaces"

	"gopkg.in/telebot.v3"
)

func LogMiddleware(logger loggerifaces.Logger) telebot.MiddlewareFunc {
	return func(nextFunc telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			ctx := context.Background()

			if sender := c.Sender(); sender != nil {
				ctx = context.WithValue(ctx, contextkeys.UserIDKey, sender.ID)
			}

			if chat := c.Chat(); chat != nil {
				ctx = context.WithValue(ctx, contextkeys.ChatIDKey, chat.ID)
			}

			if text := c.Text(); text != "" {
				ctx = context.WithValue(ctx, contextkeys.TextKey, text)
			}

			if data := c.Data(); data != "" {
				ctx = context.WithValue(ctx, contextkeys.DataKey, data)
			}

			contextLogger := logger.WithContext(ctx)
			contextLogger.Infof("Получен update")

			return nextFunc(c)
		}
	}
}

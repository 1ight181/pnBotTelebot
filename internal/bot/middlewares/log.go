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
			ctx := context.WithValue(context.Background(), contextkeys.UserIDKey, c.Sender().ID)
			ctx = context.WithValue(ctx, contextkeys.ChatIDKey, c.Chat().ID)
			ctx = context.WithValue(ctx, contextkeys.TextKey, c.Text())
			ctx = context.WithValue(ctx, contextkeys.DataKey, c.Data())

			contextLogger := logger.WithContext(ctx)

			contextLogger.Infof("Получен update")
			return nextFunc(c)
		}
	}
}

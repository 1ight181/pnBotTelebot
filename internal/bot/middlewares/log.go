package middleware

import (
	ctx "context"
	"pnBot/internal/logger/contextkeys"
	loggerifaces "pnBot/internal/logger/interfaces"

	"gopkg.in/telebot.v3"
)

func LogMiddleware(logger loggerifaces.Logger) telebot.MiddlewareFunc {
	return func(nextFunc telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			context := ctx.Background()

			if sender := c.Sender(); sender != nil {
				context = ctx.WithValue(context, contextkeys.UserIDKey, sender.ID)
			}

			if chat := c.Chat(); chat != nil {
				context = ctx.WithValue(context, contextkeys.ChatIDKey, chat.ID)
			}

			if text := c.Text(); text != "" {
				context = ctx.WithValue(context, contextkeys.TextKey, text)
			}

			if data := c.Data(); data != "" {
				context = ctx.WithValue(context, contextkeys.DataKey, data)
			}

			contextLogger := logger.WithContext(context)
			contextLogger.Infof("Получен update")

			return nextFunc(c)
		}
	}
}

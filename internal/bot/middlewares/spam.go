package middleware

import (
	ctx "context"
	"fmt"
	"pnBot/internal/logger/contextkeys"
	loggerifaces "pnBot/internal/logger/interfaces"
	spamifaces "pnBot/internal/spammanager/interfaces"
	textprovifaces "pnBot/internal/textprovider/interfaces"

	"gopkg.in/telebot.v3"
)

func SpamMiddleware(logger loggerifaces.Logger, spamManager spamifaces.SpamManager, textPovider textprovifaces.TextProvider) telebot.MiddlewareFunc {
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

			isBanned, warned, remaining, err := spamManager.IsAllowed(c.Sender().ID)
			if err != nil {
				contextLogger.Warnf("Ошибка при проверке на спам: %v", err)
				return nil
			}

			if isBanned {
				return c.Send(textPovider.GetText("ban"))
			}

			if warned {
				return c.Send(fmt.Sprintf(textPovider.GetText("warn"), remaining))
			}

			return nextFunc(c)
		}
	}
}

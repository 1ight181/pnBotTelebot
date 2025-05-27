package errorhandler

import (
	"context"
	"pnBot/internal/logger/contextkeys"
	loggerifaces "pnBot/internal/logger/interfaces"

	"gopkg.in/telebot.v3"
)

type ErrorHandler struct {
	logger loggerifaces.Logger
}

func New(logger loggerifaces.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (eh *ErrorHandler) HandleError(err error, c telebot.Context) {
	userId := c.Sender().ID
	chatId := c.Chat().ID

	ctx := context.WithValue(context.Background(), contextkeys.UserIDKey, userId)
	ctx = context.WithValue(ctx, contextkeys.ChatIDKey, chatId)
	ctx = context.WithValue(ctx, contextkeys.TextKey, c.Text())
	ctx = context.WithValue(ctx, contextkeys.DataKey, c.Data())

	contextLogger := eh.logger.WithContext(ctx)

	contextLogger.Warnf("Ошибка: %v", err)

	if c != nil {
		c.Send("Произошла ошибка. Попробуйте позже.")
	}
}

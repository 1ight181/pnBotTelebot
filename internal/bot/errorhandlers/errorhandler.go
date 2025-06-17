package errorhandler

import (
	ctx "context"
	"fmt"
	contextkeys "pnBot/internal/logger/contextkeys"
	loggerifaces "pnBot/internal/logger/interfaces"
	textproviface "pnBot/internal/textprovider/interfaces"

	"gopkg.in/telebot.v3"
)

type ErrorHandler struct {
	logger       loggerifaces.Logger
	textprovider textproviface.TextProvider
}

func NewErrorHandler(logger loggerifaces.Logger, textProvider textproviface.TextProvider) *ErrorHandler {
	return &ErrorHandler{
		logger:       logger,
		textprovider: textProvider,
	}
}

func (eh *ErrorHandler) HandleError(err error, c telebot.Context) {
	if err == nil {
		return
	}
	if c == nil {
		eh.logger.Warnf("Ошибка без контекста: %v", err)
		return
	}

	context := ctx.Background()

	if sender := c.Sender(); sender != nil {
		context = ctx.WithValue(context, contextkeys.UserIDKey, sender.ID)
	}

	if chat := c.Chat(); chat != nil {
		context = ctx.WithValue(context, contextkeys.ChatIDKey, chat.ID)
	}

	context = ctx.WithValue(context, contextkeys.TextKey, c.Text())
	context = ctx.WithValue(context, contextkeys.DataKey, c.Data())

	contextLogger := eh.logger.WithContext(context)
	contextLogger.Warnf("Ошибка: %v", err)

	errorText := eh.textprovider.GetText("error")

	if err := c.Delete(); err != nil {
		contextLogger.Errorf("Не удалось удалить последнее сообщение при отправке сообщения об ошибке: %v", err)
	}

	if sendErr := c.Send(errorText); sendErr != nil {
		contextLogger.Errorf("Не удалось отправить уведомление об ошибке пользователю: %v", sendErr)
	}
}

func (eh *ErrorHandler) ErrorMiddleware() telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			defer func() {
				if r := recover(); r != nil {
					eh.HandleError(fmt.Errorf("panic: %v", r), c)
					_ = c.Respond(&telebot.CallbackResponse{})
				}
			}()

			err := next(c)
			if err != nil {
				eh.HandleError(err, c)

				if c.Callback() != nil {
					_ = c.Respond(&telebot.CallbackResponse{})
				}
			}
			return nil
		}
	}
}

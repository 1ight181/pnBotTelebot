package bot

import (
	"context"
	ifaces "pnBot/internal/bot/interfaces"

	"gopkg.in/telebot.v3"
)

type TelegramBot struct {
	botApi      *telebot.Bot
	handlers    []ifaces.Handler
	middlewares []telebot.MiddlewareFunc
	context     context.Context
}

type TelegramBotOptions struct {
	BotApi      *telebot.Bot
	Handlers    []ifaces.Handler
	Middlewares []telebot.MiddlewareFunc
	Context     context.Context
}

func New(opts TelegramBotOptions) *TelegramBot {
	return &TelegramBot{
		botApi:      opts.BotApi,
		handlers:    opts.Handlers,
		middlewares: opts.Middlewares,
		context:     opts.Context,
	}
}

func (tgb *TelegramBot) Start() {
	tgb.botApi.Use(tgb.middlewares...)

	for _, handler := range tgb.handlers {
		handler.StartHandling(tgb.botApi)
	}

	go tgb.botApi.Start()

	go func() {
		<-tgb.context.Done()
		tgb.botApi.Stop()
	}()
}

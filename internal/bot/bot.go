package bot

import (
	"context"
	ifaces "pnBot/internal/bot/interfaces"

	"gopkg.in/telebot.v3"
)

type TelegramBot struct {
	botAPI   *telebot.Bot
	handlers []ifaces.Handler
	context  context.Context
}

type TelegramBotOptions struct {
	BotAPI   *telebot.Bot
	Handlers []ifaces.Handler
	Context  context.Context
}

func New(opts TelegramBotOptions) *TelegramBot {
	return &TelegramBot{
		botAPI:   opts.BotAPI,
		handlers: opts.Handlers,
		context:  opts.Context,
	}
}

func (tb *TelegramBot) Start() {
	for _, handler := range tb.handlers {
		handler.StartHandling(tb.botAPI)
	}

	go tb.botAPI.Start()

	go func() {
		<-tb.context.Done()
		tb.botAPI.Stop()
	}()
}

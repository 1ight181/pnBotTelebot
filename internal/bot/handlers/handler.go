package handlers

import "gopkg.in/telebot.v3"

type Handler struct {
	endpoint   interface{}
	handleFunc telebot.HandlerFunc
}

func New(
	endpoint interface{},
	handleFunc telebot.HandlerFunc,
) *Handler {
	return &Handler{
		endpoint:   endpoint,
		handleFunc: handleFunc,
	}
}

func (h *Handler) StartHandling(bot *telebot.Bot) {
	bot.Handle(h.endpoint, h.handleFunc)
}

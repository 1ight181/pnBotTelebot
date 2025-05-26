package interfaces

import "gopkg.in/telebot.v3"

type Handler interface {
	StartHandling(bot *telebot.Bot)
}

package middleware

import "gopkg.in/telebot.v3"

func NoopMiddleware(nextFunc telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		return nextFunc(c)
	}
}

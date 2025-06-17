package common

import (
	"pnBot/internal/bot/processors/keyboards"
	dbifaces "pnBot/internal/db/interfaces"
	textproviface "pnBot/internal/textprovider/interfaces"

	"gopkg.in/telebot.v3"
)

func ProcessMenu(c telebot.Context, textProvider textproviface.TextProvider, dbProvider dbifaces.DataBaseProvider) error {
	userId := c.Sender().ID
	isSubscribed, err := IsSubscribed(userId, dbProvider)
	if err != nil {
		return err
	}
	if !isSubscribed {
		notSubscribedText := textProvider.GetText("not_subscribed")
		return c.Send(notSubscribedText, &telebot.ReplyMarkup{RemoveKeyboard: true})
	}

	menuText := textProvider.GetText("menu")

	menuKeyboard := keyboards.GetMenuKeyBoard(textProvider)

	return c.Send(menuText, menuKeyboard)
}

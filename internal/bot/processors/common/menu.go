package common

import (
	botifaces "pnBot/internal/bot/interfaces"
	"pnBot/internal/bot/processors/keyboards"
	dbifaces "pnBot/internal/db/interfaces"

	"gopkg.in/telebot.v3"
)

func ProcessMenu(c telebot.Context, textProvider botifaces.TextProvider, dbProvider dbifaces.DataBaseProvider) error {
	userId := c.Sender().ID
	isSubscribed, err := isSubscribed(userId, dbProvider)
	if err != nil {
		return err
	}
	if !isSubscribed {
		notAllowedText := textProvider.GetText("not_allowed")
		subscribeKeyboard := &telebot.ReplyMarkup{}

		subscribeButtonText := textProvider.GetButtonText("subscribe")

		subscribeButton := subscribeKeyboard.Data(
			subscribeButtonText,
			"subscribe",
		)

		subscribeKeyboard.Inline(
			subscribeKeyboard.Row(subscribeButton),
		)
		return c.Send(notAllowedText, subscribeKeyboard)
	}

	menuText := textProvider.GetText("menu")

	menuKeyboard := keyboards.GetMenuKeyBoard(textProvider)

	return c.Send(menuText, menuKeyboard)
}

package keyboards

import (
	ifaces "pnBot/internal/bot/interfaces"

	"gopkg.in/telebot.v3"
)

func GetSubscribeKeyBoard(textProvider ifaces.TextProvider) *telebot.ReplyMarkup {
	subscribeKeyboard := &telebot.ReplyMarkup{}

	subscribeButtonText := textProvider.GetButtonText("subscribe")

	subscribeButton := subscribeKeyboard.Data(
		subscribeButtonText,
		"subscribe",
	)

	subscribeKeyboard.Inline(
		subscribeKeyboard.Row(subscribeButton),
	)

	return subscribeKeyboard
}

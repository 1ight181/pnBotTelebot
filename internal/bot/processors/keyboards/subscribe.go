package keyboards

import (
	textproviface "pnBot/internal/textprovider/interfaces"

	"gopkg.in/telebot.v3"
)

func GetSubscribeKeyBoard(textProvider textproviface.TextProvider) *telebot.ReplyMarkup {
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

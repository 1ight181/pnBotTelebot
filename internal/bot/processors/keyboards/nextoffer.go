package keyboards

import (
	ifaces "pnBot/internal/bot/interfaces"

	"gopkg.in/telebot.v3"
)

func NextOfferKeyBoard(textProvider ifaces.TextProvider) *telebot.ReplyMarkup {
	nextOfferKeyboard := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	nextOfferButtonText := textProvider.GetButtonText("next_offer")

	nextOfferButton := nextOfferKeyboard.Data(
		nextOfferButtonText,
		"next",
	)

	menuButtonText := textProvider.GetButtonText("menu")
	menuButtonButton := nextOfferKeyboard.Data(
		menuButtonText,
		"menu",
	)

	nextOfferKeyboard.Inline(
		nextOfferKeyboard.Row(nextOfferButton),
		nextOfferKeyboard.Row(menuButtonButton),
	)

	return nextOfferKeyboard
}

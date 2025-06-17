package keyboards

import (
	textproviface "pnBot/internal/textprovider/interfaces"

	"gopkg.in/telebot.v3"
)

func GetNextOfferKeyBoard(textProvider textproviface.TextProvider) *telebot.ReplyMarkup {
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

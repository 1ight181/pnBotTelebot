package keyboards

import (
	ifaces "pnBot/internal/bot/interfaces"

	"gopkg.in/telebot.v3"
)

func GetMenuKeyBoard(textProvider ifaces.TextProvider) *telebot.ReplyMarkup {
	menuKeyboard := &telebot.ReplyMarkup{}

	lastPromoButtonText := textProvider.GetButtonText("last_promo")

	lastPromoButton := menuKeyboard.Data(
		lastPromoButtonText,
		"lastPromoButton",
	)

	filterSettingsButtonText := textProvider.GetButtonText("filter_settings")

	filterSettingsButton := menuKeyboard.Data(
		filterSettingsButtonText,
		"filter_settings",
	)

	unsubscribeButtonText := textProvider.GetButtonText("unsubscribe")

	unsubscribeButton := menuKeyboard.Data(
		unsubscribeButtonText,
		"unsubscribe",
	)

	menuKeyboard.Inline(
		menuKeyboard.Row(lastPromoButton),
		menuKeyboard.Row(filterSettingsButton),
		menuKeyboard.Row(unsubscribeButton),
	)

	return menuKeyboard
}

package keyboards

import (
	textproviface "pnBot/internal/textprovider/interfaces"

	"gopkg.in/telebot.v3"
)

func GetMenuKeyBoard(textProvider textproviface.TextProvider) *telebot.ReplyMarkup {
	menuKeyboard := &telebot.ReplyMarkup{}

	lastPromoButtonText := textProvider.GetButtonText("last_promo")

	lastPromoButton := menuKeyboard.Data(
		lastPromoButtonText,
		"last",
	)

	filterSettingsButtonText := textProvider.GetButtonText("filter_settings")

	filterSettingsButton := menuKeyboard.Data(
		filterSettingsButtonText,
		"filter_settings",
	)

	frequencySettingsButtonText := textProvider.GetButtonText("frequency_settings")

	frequencySettingsButton := menuKeyboard.Data(
		frequencySettingsButtonText,
		"frequency_settings",
	)

	feedbackButtonText := textProvider.GetButtonText("feedback")

	feedbackButton := menuKeyboard.Data(
		feedbackButtonText,
		"feedback",
	)

	bugReprotButtonText := textProvider.GetButtonText("bug_report")

	bugReprotButton := menuKeyboard.Data(
		bugReprotButtonText,
		"bug_report",
	)

	unsubscribeButtonText := textProvider.GetButtonText("unsubscribe")

	unsubscribeButton := menuKeyboard.Data(
		unsubscribeButtonText,
		"unsubscribe",
	)

	menuKeyboard.Inline(
		menuKeyboard.Row(lastPromoButton),
		menuKeyboard.Row(filterSettingsButton),
		menuKeyboard.Row(frequencySettingsButton),
		menuKeyboard.Row(feedbackButton),
		menuKeyboard.Row(bugReprotButton),
		menuKeyboard.Row(unsubscribeButton),
	)

	return menuKeyboard
}

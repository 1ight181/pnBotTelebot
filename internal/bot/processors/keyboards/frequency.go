package keyboards

import (
	"fmt"
	textproviface "pnBot/internal/textprovider/interfaces"

	"gopkg.in/telebot.v3"
)

func GetFrequencyKeyboard(textprovider textproviface.TextProvider) *telebot.ReplyMarkup {
	frequencyKeyboard := &telebot.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	everyXHoursButtonText := textprovider.GetButtonText("every_x_hours")
	everydayButtonText := textprovider.GetButtonText("everyday")

	everyTwoHoursButton := frequencyKeyboard.Text(
		fmt.Sprintf(everyXHoursButtonText, 2),
	)
	everyFourHoursButton := frequencyKeyboard.Text(
		fmt.Sprintf(everyXHoursButtonText, 4),
	)
	everySixHoursButton := frequencyKeyboard.Text(
		fmt.Sprintf(everyXHoursButtonText, 6),
	)
	everyTwelveHoursButton := frequencyKeyboard.Text(
		fmt.Sprintf(everyXHoursButtonText, 12),
	)
	everyDayButton := frequencyKeyboard.Text(
		everydayButtonText,
	)

	frequencyKeyboard.Reply(
		frequencyKeyboard.Row(everyTwoHoursButton),
		frequencyKeyboard.Row(everyFourHoursButton),
		frequencyKeyboard.Row(everySixHoursButton),
		frequencyKeyboard.Row(everyTwelveHoursButton),
		frequencyKeyboard.Row(everyDayButton),
	)

	return frequencyKeyboard
}

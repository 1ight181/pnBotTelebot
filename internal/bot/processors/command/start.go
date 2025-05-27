package command

import "gopkg.in/telebot.v3"

func (p *CommandProcessor) ProcessStart(c telebot.Context) error {
	startKeyboard := &telebot.ReplyMarkup{}

	subscribeButtonText := p.dependencies.TextProvider.GetButtonText("subscribe")

	subscribeButton := startKeyboard.Data(
		subscribeButtonText,
		"subscribe",
	)

	startKeyboard.Inline(
		startKeyboard.Row(subscribeButton),
	)

	startText := p.dependencies.TextProvider.GetText("start")
	return c.Send(startText, startKeyboard)
}

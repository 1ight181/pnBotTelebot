package command

import "gopkg.in/telebot.v3"

func (p *CommandProcessor) ProcessStart(c telebot.Context) error {
	startKeyboard := &telebot.ReplyMarkup{}

	subscribeButtonText := p.deps.TextProvider.GetButtonText("subscribe")

	subscribeButton := startKeyboard.Data(
		subscribeButtonText,
		"subscribe",
	)

	startKeyboard.Inline(
		startKeyboard.Row(subscribeButton),
	)

	p.deps.Logger.Infof("Пользователь %d запустил бота!", c.Sender().ID)
	startText := p.deps.TextProvider.GetText("start")
	return c.Send(startText, startKeyboard)
}

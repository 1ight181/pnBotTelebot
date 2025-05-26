package callback

import "gopkg.in/telebot.v3"

func (p *CallbackProcessor) ProcessUnsubscribe(c telebot.Context) error {
	userID := c.Sender().ID

	p.deps.Logger.Infof("Пользователь %d отписался!", userID)

	unsubscribeText := p.deps.TextProvider.GetCallbackText("unsubscribe")
	return c.Respond(&telebot.CallbackResponse{
		Text:      unsubscribeText,
		ShowAlert: true,
	})
}

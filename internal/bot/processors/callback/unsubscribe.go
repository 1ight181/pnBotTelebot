package callback

import "gopkg.in/telebot.v3"

func (p *CallbackProcessor) ProcessUnsubscribe(c telebot.Context) error {
	unsubscribeText := p.dependencies.TextProvider.GetCallbackText("unsubscribe")
	return c.Respond(&telebot.CallbackResponse{
		Text:      unsubscribeText,
		ShowAlert: true,
	})
}

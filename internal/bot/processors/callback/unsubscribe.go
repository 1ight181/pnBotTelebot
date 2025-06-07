package callback

import (
	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessUnsubscribe(c telebot.Context) error {
	userId := c.Sender().ID

	err := cp.setUnsubscribed(userId)
	if err != nil {
		return err
	}

	unsubscribeText := cp.dependencies.TextProvider.GetCallbackText("unsubscribe")
	c.Respond(&telebot.CallbackResponse{
		Text:      unsubscribeText,
		ShowAlert: false,
	})

	return c.Delete()
}

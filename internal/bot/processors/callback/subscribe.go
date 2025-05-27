package callback

import "gopkg.in/telebot.v3"

func (cp *CallbackProcessor) ProcessSubscribe(c telebot.Context) error {
	subscribeText := cp.dependencies.TextProvider.GetCallbackText("subscribe")
	if err := c.Respond(&telebot.CallbackResponse{
		Text:      subscribeText,
		ShowAlert: false,
	}); err != nil {
		return err
	}

	return c.Edit(c.Message().Text, &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{},
	})
}

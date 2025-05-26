package callback

import "gopkg.in/telebot.v3"

func (cp *CallbackProcessor) ProcessSubscribe(c telebot.Context) error {
	userID := c.Sender().ID

	cp.deps.Logger.Infof("Пользователь %d подписался!", userID)

	subscribeText := cp.deps.TextProvider.GetCallbackText("subscribe")
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

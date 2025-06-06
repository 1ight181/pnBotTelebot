package callback

import (
	ctx "context"
	dbmodels "pnBot/internal/db/models"

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

func (cp *CallbackProcessor) setUnsubscribed(userId int64) error {
	context := ctx.Background()

	where := dbmodels.User{
		TgId: userId,
	}

	return cp.dependencies.DbProvider.Update(context, where, "is_subscribed", false)
}

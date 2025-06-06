package callback

import (
	ctx "context"
	keyboards "pnBot/internal/bot/processors/keyboards"
	dbmodels "pnBot/internal/db/models"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessSubscribe(c telebot.Context) error {
	userId := c.Sender().ID

	isUserAlreadySubscribed, err := cp.setSubscribed(userId)
	if err != nil {
		return err
	}

	var subscribeText string

	if isUserAlreadySubscribed {
		subscribeText = cp.dependencies.TextProvider.GetCallbackText("already_subscribed")
		if err := c.Respond(&telebot.CallbackResponse{
			Text:      subscribeText,
			ShowAlert: false,
		}); err != nil {
			return err
		}
	} else {
		subscribeText = cp.dependencies.TextProvider.GetCallbackText("subscribe")
		if err := c.Respond(&telebot.CallbackResponse{
			Text:      subscribeText,
			ShowAlert: false,
		}); err != nil {
			return err
		}
	}

	c.Delete()

	menuText := cp.dependencies.TextProvider.GetText("menu")
	menuKeyboard := keyboards.GetMenuKeyBoard(cp.dependencies.TextProvider)

	return c.Send(menuText, menuKeyboard)
}

func (cp *CallbackProcessor) setSubscribed(userId int64) (bool, error) {
	context := ctx.Background()
	user := dbmodels.User{}

	where := dbmodels.User{
		TgId: userId,
	}

	if err := cp.dependencies.DbProvider.Find(context, &user, where); err != nil {
		return true, err
	}
	if user.IsSubscribed {
		return true, nil
	}

	return false, cp.dependencies.DbProvider.Update(context, where, "is_subscribed", true)
}

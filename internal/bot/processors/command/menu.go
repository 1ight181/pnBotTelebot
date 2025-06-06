package command

import (
	ctx "context"
	"pnBot/internal/bot/processors/keyboards"
	dbmodels "pnBot/internal/db/models"

	"gopkg.in/telebot.v3"
)

func (cp *CommandProcessor) ProcessMenu(c telebot.Context) error {
	userId := c.Sender().ID
	isSubscribed, err := cp.isSubscribed(userId)
	if err != nil {
		return err
	}
	if !isSubscribed {
		notAllowedText := cp.dependencies.TextProvider.GetText("not_allowed")
		subscribeKeyboard := &telebot.ReplyMarkup{}

		subscribeButtonText := cp.dependencies.TextProvider.GetButtonText("subscribe")

		subscribeButton := subscribeKeyboard.Data(
			subscribeButtonText,
			"subscribe",
		)

		subscribeKeyboard.Inline(
			subscribeKeyboard.Row(subscribeButton),
		)
		return c.Send(notAllowedText, subscribeKeyboard)
	}

	menuText := cp.dependencies.TextProvider.GetText("menu")

	menuKeyboard := keyboards.GetMenuKeyBoard(cp.dependencies.TextProvider)

	return c.Send(menuText, menuKeyboard)
}

func (cp *CommandProcessor) isSubscribed(userId int64) (bool, error) {
	user := dbmodels.User{
		TgId: userId,
	}
	if err := cp.dependencies.DbProvider.Find(ctx.Background(), &user, user); err != nil {
		return false, err
	}

	if user.IsSubscribed {
		return true, nil
	} else {
		return false, nil
	}
}

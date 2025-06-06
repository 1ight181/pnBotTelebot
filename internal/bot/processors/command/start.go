package command

import (
	ctx "context"
	"fmt"
	dbmodels "pnBot/internal/db/models"

	"gopkg.in/telebot.v3"
)

func (cp *CommandProcessor) ProcessStart(c telebot.Context) error {
	userId := c.Sender().ID
	chatId := c.Chat().ID
	username := c.Sender().Username
	fristname := c.Sender().FirstName
	lastname := c.Sender().LastName
	fullname := fmt.Sprintf("%s %s", fristname, lastname)

	user, isUserCreated, err := cp.addUserToDb(userId, chatId, username, fullname)
	if err != nil {
		return err
	}

	var startText string
	if isUserCreated {
		startText = cp.dependencies.TextProvider.GetText("first_start")
	} else {
		if user.IsSubscribed {
			startText = cp.dependencies.TextProvider.GetText("not_first_start_subscribed")
			return c.Send(startText)
		} else {
			startText = cp.dependencies.TextProvider.GetText("not_first_start_not_subscribed")
		}
	}

	startKeyboard := &telebot.ReplyMarkup{}

	subscribeButtonText := cp.dependencies.TextProvider.GetButtonText("subscribe")

	subscribeButton := startKeyboard.Data(
		subscribeButtonText,
		"subscribe",
	)

	startKeyboard.Inline(
		startKeyboard.Row(subscribeButton),
	)

	return c.Send(startText, startKeyboard)
}

func (cp *CommandProcessor) addUserToDb(
	userId int64,
	chatId int64,
	username string,
	name string,
) (dbmodels.User, bool, error) {
	user := dbmodels.User{}
	where := dbmodels.User{TgID: userId}
	defaults := dbmodels.User{
		TgID:         userId,
		ChatID:       chatId,
		Username:     username,
		Fullname:     name,
		IsSubscribed: false,
	}

	ctx := ctx.Background()

	created, err := cp.dependencies.DbProvider.FirstOrCreate(ctx, &user, where, defaults)
	return user, created, err
}

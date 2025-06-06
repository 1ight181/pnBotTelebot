package command

import (
	ctx "context"
	"fmt"
	"pnBot/internal/bot/processors/keyboards"
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
	subscribeKeyboard := keyboards.GetSubscribeKeyBoard(cp.dependencies.TextProvider)

	if isUserCreated {
		startText = cp.dependencies.TextProvider.GetText("first_start")
	} else {
		if user.IsSubscribed {
			startText = cp.dependencies.TextProvider.GetText("not_first_start_subscribed")
			c.Send(startText)

			menuText := cp.dependencies.TextProvider.GetText("menu")
			menuKeyboard := keyboards.GetMenuKeyBoard(cp.dependencies.TextProvider)

			return c.Send(menuText, menuKeyboard)
		} else {
			startText = cp.dependencies.TextProvider.GetText("not_first_start_not_subscribed")
		}
	}

	return c.Send(startText, subscribeKeyboard)
}

func (cp *CommandProcessor) addUserToDb(
	userId int64,
	chatId int64,
	username string,
	name string,
) (dbmodels.User, bool, error) {
	user := dbmodels.User{}
	where := dbmodels.User{TgId: userId}
	defaults := dbmodels.User{
		TgId:         userId,
		ChatId:       chatId,
		Username:     username,
		Fullname:     name,
		IsSubscribed: false,
	}

	ctx := ctx.Background()

	created, err := cp.dependencies.DbProvider.FirstOrCreate(ctx, &user, where, defaults)
	return user, created, err
}

package command

import (
	"context"

	dbmodels "pnBot/internal/db/models"

	"gopkg.in/telebot.v3"
)

func (cp *CommandProcessor) ProcessStart(c telebot.Context) error {
	userId := c.Sender().ID
	chatId := c.Chat().ID
	username := c.Sender().Username
	name := c.Sender().FirstName + c.Sender().LastName

	if err := cp.addUserToDb(userId, chatId, username, name); err != nil {
		return err
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

	startText := cp.dependencies.TextProvider.GetText("start")
	return c.Send(startText, startKeyboard)
}

func (cp *CommandProcessor) addUserToDb(
	userId int64,
	chatId int64,
	username string,
	name string,
) error {
	user := dbmodels.User{
		TgID:         userId,
		ChatID:       chatId,
		Username:     username,
		Name:         name,
		IsSubscribed: false,
	}

	return cp.dependencies.DbProvider.Create(context.Background(), user)
}

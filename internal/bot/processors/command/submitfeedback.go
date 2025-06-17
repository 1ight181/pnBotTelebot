package command

import (
	"fmt"
	"pnBot/internal/bot/processors/keyboards"

	"gopkg.in/telebot.v3"
)

func (cp *CommandProcessor) ProcessSubmitFeedback(c telebot.Context) error {
	userId := c.Sender().ID
	if cp.dependencies.Fsm.Get(userId) != "feedback" {
		return nil
	}

	userFirstName := c.Sender().FirstName
	userSecondName := c.Sender().FirstName
	userFullname := fmt.Sprintf("%s %s", userFirstName, userSecondName)

	username := c.Sender().Username

	feedbackMessageText := c.Message().Text

	subjectFeedbackText := cp.dependencies.TextProvider.GetEmailSubject("feedback")
	bodyFeedbackText := fmt.Sprintf("ID: %d\nFullname: %s\nUsername: %s\n\n%s", userId, userFullname, username, feedbackMessageText)

	cp.dependencies.EmailSender.SendToAdmin(subjectFeedbackText, bodyFeedbackText)
	menuText := cp.dependencies.TextProvider.GetText("menu")
	menuKeyboard := keyboards.GetMenuKeyBoard(cp.dependencies.TextProvider)
	return c.Send(menuText, menuKeyboard)
}

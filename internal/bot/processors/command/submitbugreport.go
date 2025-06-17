package command

import (
	"fmt"
	"pnBot/internal/bot/processors/keyboards"

	"gopkg.in/telebot.v3"
)

func (cp *CommandProcessor) ProcessSubmitBugReport(c telebot.Context) error {
	userId := c.Sender().ID
	if cp.dependencies.Fsm.Get(userId) != "bug_report" {
		return nil
	}

	userFirstName := c.Sender().FirstName
	userSecondName := c.Sender().FirstName
	userFullname := fmt.Sprintf("%s %s", userFirstName, userSecondName)

	username := c.Sender().Username

	bugReportMessageText := c.Message().Text

	subjectBugReportText := cp.dependencies.TextProvider.GetEmailSubject("bug_report")
	bodyBugReportText := fmt.Sprintf("ID: %d\nFullname: %s\nUsername: %s\n\n%s", userId, userFullname, username, bugReportMessageText)

	cp.dependencies.EmailSender.SendToAdmin(subjectBugReportText, bodyBugReportText)
	menuText := cp.dependencies.TextProvider.GetText("menu")
	menuKeyboard := keyboards.GetMenuKeyBoard(cp.dependencies.TextProvider)
	return c.Send(menuText, menuKeyboard)
}

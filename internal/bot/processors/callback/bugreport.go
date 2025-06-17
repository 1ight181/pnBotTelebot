package callback

import (
	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessBugReport(c telebot.Context) error {
	userId := c.Sender().ID
	cp.dependencies.Fsm.Set(userId, "bug_report")
	submitBugReportText := cp.dependencies.TextProvider.GetText("bug_report_hint")

	if err := c.Edit(submitBugReportText); err != nil {
		return err
	}

	return c.Respond(&telebot.CallbackResponse{})
}

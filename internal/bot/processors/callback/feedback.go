package callback

import (
	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessFeedback(c telebot.Context) error {
	userId := c.Sender().ID
	cp.dependencies.Fsm.Set(userId, "feedback")
	submitFeedbackText := cp.dependencies.TextProvider.GetText("feedback_hint")

	if err := c.Edit(submitFeedbackText); err != nil {
		return err
	}

	return c.Respond(&telebot.CallbackResponse{})
}

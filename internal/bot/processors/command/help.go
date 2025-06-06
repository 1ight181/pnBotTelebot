package command

import "gopkg.in/telebot.v3"

func (cp *CommandProcessor) ProcessHelp(c telebot.Context) error {
	helpText := cp.dependencies.TextProvider.GetText("help")
	return c.Send(helpText)
}

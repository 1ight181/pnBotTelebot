package command

import "gopkg.in/telebot.v3"

func (p *CommandProcessor) ProcessHelp(c telebot.Context) error {
	helpText := p.dependencies.TextProvider.GetText("help")
	return c.Send(helpText)
}

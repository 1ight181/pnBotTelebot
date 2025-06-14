package callback

import (
	"pnBot/internal/bot/processors/common"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProccesMenu(c telebot.Context) error {

	err := c.EditCaption(c.Text(), &telebot.SendOptions{
		Entities: c.Entities(),
	})
	if err != nil {
		return err
	}

	if err := common.ProcessMenu(c, cp.dependencies.TextProvider, cp.dependencies.DbProvider); err != nil {
		return err
	}

	return c.Respond(&telebot.CallbackResponse{})
}

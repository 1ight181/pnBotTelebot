package callback

import (
	keyboards "pnBot/internal/bot/processors/keyboards"
	units "pnBot/internal/notifier/units"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProcessSubscribe(c telebot.Context) error {
	userId := c.Sender().ID

	isUserAlreadySubscribed, err := cp.setSubscribed(userId)
	if err != nil {
		return err
	}

	if err := cp.addUserToAllCategories(userId); err != nil {
		return err
	}

	var subscribeText string

	if isUserAlreadySubscribed {
		subscribeText = cp.dependencies.TextProvider.GetCallbackText("already_subscribed")
		if err := c.Respond(&telebot.CallbackResponse{
			Text:      subscribeText,
			ShowAlert: false,
		}); err != nil {
			return err
		}
	} else {
		subscribeText = cp.dependencies.TextProvider.GetCallbackText("subscribe")
		if err := c.Respond(&telebot.CallbackResponse{
			Text:      subscribeText,
			ShowAlert: false,
		}); err != nil {
			return err
		}
	}
	frequency := 4
	frequencyUnit := units.Hours
	cp.dependencies.Notifier.AddUser(userId, frequency, cp.dependencies.OfferCooldownDuration, frequencyUnit)

	menuText := cp.dependencies.TextProvider.GetText("menu")
	menuKeyboard := keyboards.GetMenuKeyBoard(cp.dependencies.TextProvider)

	return c.Edit(menuText, menuKeyboard)
}

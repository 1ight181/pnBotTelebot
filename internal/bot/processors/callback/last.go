package callback

import (
	"fmt"
	"pnBot/internal/bot/processors/common"
	"pnBot/internal/bot/processors/keyboards"
	dbmodels "pnBot/internal/db/models"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProccesLast(c telebot.Context) error {
	userId := c.Sender().ID

	offer, err := cp.dependencies.OfferDao.GetNextAvailableOffer(userId)
	if err != nil {
		return err
	}

	if offer == nil {
		noAvailableOfferText := cp.dependencies.TextProvider.GetText("no_available_offer")
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
		c.Edit(noAvailableOfferText)
		return common.ProcessMenu(c, cp.dependencies.TextProvider, cp.dependencies.DbProvider)
	}

	offerCreatives := offer.Creatives

	var offerCreative dbmodels.Creative
	if len(offerCreatives) != 0 {
		offerCreative = offerCreatives[0]
	}

	offerImageUrl := offerCreative.ResourceUrl
	offerTitle := offer.Title
	offerDescription := offer.Description

	escapedTitle := common.EscapeMarkdownV2(offerTitle)
	escapedDescription := common.EscapeMarkdownV2(offerDescription)

	escapedDescription = common.WrapURLsWithPreviousWord(escapedDescription)

	offerMessage := &telebot.Photo{
		File:    telebot.FromURL(offerImageUrl),
		Caption: fmt.Sprintf("*%s* \n\n%s", escapedTitle, escapedDescription),
	}

	nextOfferKeyboard := keyboards.NextOfferKeyBoard(cp.dependencies.TextProvider)
	if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
		return err
	}

	if err := c.Send(offerMessage, nextOfferKeyboard); err != nil {
		return err
	}

	cp.dependencies.OfferDao.AddSendingLog(userId, offer.Id)
	return nil
}

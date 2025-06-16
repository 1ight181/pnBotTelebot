package callback

import (
	"errors"
	"fmt"
	"pnBot/internal/bot/processors/common"
	"pnBot/internal/bot/processors/keyboards"
	dberrors "pnBot/internal/db/errors"
	dbmodels "pnBot/internal/db/models"
	"time"

	"gopkg.in/telebot.v3"
)

func (cp *CallbackProcessor) ProccesLast(c telebot.Context) error {
	userId := c.Sender().ID

	isSubscribed, err := common.IsSubscribed(userId, cp.dependencies.DbProvider)
	if err != nil {
		return err
	}
	if !isSubscribed {
		notSubscribedText := cp.dependencies.TextProvider.GetText("not_subscribed")
		return c.Send(notSubscribedText)
	}

	offerCooldown := time.Now().Add(-cp.dependencies.OfferCooldownDuration)
	limit := 1
	offers, err := cp.dependencies.OfferDao.GetLastAvailableOffers(userId, limit, offerCooldown)
	if errors.Is(err, dberrors.ErrRecordNotFound) {
		noAvailableOfferText := cp.dependencies.TextProvider.GetText("no_available_offer")
		if err := c.Respond(&telebot.CallbackResponse{}); err != nil {
			return err
		}
		c.Edit(noAvailableOfferText)
		return common.ProcessMenu(c, cp.dependencies.TextProvider, cp.dependencies.DbProvider)
	} else if err != nil {
		return err
	}

	offer := offers[0]

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

	if err := c.Edit(offerMessage, nextOfferKeyboard); err != nil {
		return err
	}

	cp.dependencies.OfferDao.AddSendingLog(userId, offer.Id)
	return nil
}

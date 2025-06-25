package app

import (
	ctx "context"
	banifaces "pnBot/internal/banmanager/interfaces"
	cacheifaces "pnBot/internal/cache/interfaces"
	"pnBot/internal/config/loaders"
	"pnBot/internal/config/models"
	spam "pnBot/internal/spammanager"
	spamifaces "pnBot/internal/spammanager/interfaces"
)

func createSpamManager(context ctx.Context, spamManagerConfig *models.SpamManager, banManager banifaces.BanManager, cacheProvider cacheifaces.CacheProvider) spamifaces.SpamManager {
	messageLimit, interval, warnLimit, banDuration, banReasonText, banAuthor := loaders.LoadSpamManager(*spamManagerConfig)

	spamManagerOptions := spam.SpamManagerOptions{
		CacheProvider: cacheProvider,
		MessageLimit:  messageLimit,
		Interval:      interval,
		WarnLimit:     warnLimit,
		BanDuration:   banDuration,
		BanManager:    banManager,
		BanReasonText: banReasonText,
		BanAuthor:     banAuthor,
	}

	spamManager := spam.NewSpamManager(spamManagerOptions)

	return spamManager
}

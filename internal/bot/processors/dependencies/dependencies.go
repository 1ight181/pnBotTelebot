package deps

import (
	botifaces "pnBot/internal/bot/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	notifierifaces "pnBot/internal/notifier/interfaces"
	"time"
)

type ProcessorDependencies struct {
	TextProvider          botifaces.TextProvider
	DbProvider            dbifaces.DataBaseProvider
	OfferDao              dbifaces.OfferDao
	OfferCooldownDuration time.Duration
	Notifier              notifierifaces.Notifier
}

type ProcessorDependenciesOptions struct {
	TextProvider          botifaces.TextProvider
	DbProvider            dbifaces.DataBaseProvider
	OfferDao              dbifaces.OfferDao
	OfferCooldownDuration time.Duration
	Notifier              notifierifaces.Notifier
}

func NewProcessorDependencies(opts ProcessorDependenciesOptions) *ProcessorDependencies {
	return &ProcessorDependencies{
		TextProvider:          opts.TextProvider,
		DbProvider:            opts.DbProvider,
		OfferDao:              opts.OfferDao,
		OfferCooldownDuration: opts.OfferCooldownDuration,
		Notifier:              opts.Notifier,
	}
}

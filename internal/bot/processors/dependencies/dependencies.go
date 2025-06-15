package deps

import (
	botifaces "pnBot/internal/bot/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	schedulerifaces "pnBot/internal/scheduler/interfaces"
	"time"
)

type ProcessorDependencies struct {
	TextProvider          botifaces.TextProvider
	DbProvider            dbifaces.DataBaseProvider
	OfferDao              dbifaces.OfferDao
	Scheduler             schedulerifaces.Scheduler
	OfferCooldownDuration time.Duration
}

type ProcessorDependenciesOptions struct {
	TextProvider          botifaces.TextProvider
	DbProvider            dbifaces.DataBaseProvider
	OfferDao              dbifaces.OfferDao
	Scheduler             schedulerifaces.Scheduler
	OfferCooldownDuration time.Duration
}

func NewProcessorDependencies(opts ProcessorDependenciesOptions) *ProcessorDependencies {
	return &ProcessorDependencies{
		TextProvider:          opts.TextProvider,
		DbProvider:            opts.DbProvider,
		OfferDao:              opts.OfferDao,
		Scheduler:             opts.Scheduler,
		OfferCooldownDuration: opts.OfferCooldownDuration,
	}
}

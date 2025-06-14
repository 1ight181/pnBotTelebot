package deps

import (
	botifaces "pnBot/internal/bot/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
)

type ProcessorDependencies struct {
	TextProvider botifaces.TextProvider
	DbProvider   dbifaces.DataBaseProvider
	OfferDao     dbifaces.OfferDao
}

type ProcessorDependenciesOptions struct {
	TextProvider botifaces.TextProvider
	DbProvider   dbifaces.DataBaseProvider
	OfferDao     dbifaces.OfferDao
}

func NewProcessorDependencies(opts ProcessorDependenciesOptions) *ProcessorDependencies {
	return &ProcessorDependencies{
		TextProvider: opts.TextProvider,
		DbProvider:   opts.DbProvider,
		OfferDao:     opts.OfferDao,
	}
}

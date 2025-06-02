package deps

import (
	botifaces "pnBot/internal/bot/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
)

type ProcessorDependencies struct {
	TextProvider botifaces.TextProvider
	DbProvider   dbifaces.DataBaseProvider
}

type ProcessorDependenciesOptions struct {
	TextProvider botifaces.TextProvider
	DbProvider   dbifaces.DataBaseProvider
}

func New(opts ProcessorDependenciesOptions) *ProcessorDependencies {
	return &ProcessorDependencies{
		TextProvider: opts.TextProvider,
		DbProvider:   opts.DbProvider,
	}
}

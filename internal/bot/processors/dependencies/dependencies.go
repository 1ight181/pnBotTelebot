package deps

import (
	botifaces "pnBot/internal/bot/interfaces"
)

type ProcessorDependencies struct {
	TextProvider botifaces.TextProvider
}

type ProcessorDependenciesOptions struct {
	TextProvider botifaces.TextProvider
}

func New(opts ProcessorDependenciesOptions) *ProcessorDependencies {
	return &ProcessorDependencies{
		TextProvider: opts.TextProvider,
	}
}

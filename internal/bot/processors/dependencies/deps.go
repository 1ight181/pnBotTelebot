package deps

import (
	botifaces "pnBot/internal/bot/interfaces"
	loggerifaces "pnBot/internal/logger/interfaces"
)

type ProcessorDependencies struct {
	Logger       loggerifaces.Logger
	TextProvider botifaces.TextProvider
}

type ProcessorDependenciesOptions struct {
	Logger       loggerifaces.Logger
	TextProvider botifaces.TextProvider
}

func New(opts ProcessorDependenciesOptions) *ProcessorDependencies {
	return &ProcessorDependencies{
		Logger:       opts.Logger,
		TextProvider: opts.TextProvider,
	}
}

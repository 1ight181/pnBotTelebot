package deps

import (
	dbifaces "pnBot/internal/db/interfaces"
	fsmifaces "pnBot/internal/fsm/interfaces"
	notifierifaces "pnBot/internal/notifier/interfaces"
	emailifaces "pnBot/internal/sender/interfaces"
	textproviface "pnBot/internal/textprovider/interfaces"
	"time"
)

type ProcessorDependencies struct {
	TextProvider          textproviface.TextProvider
	DbProvider            dbifaces.DataBaseProvider
	OfferDao              dbifaces.OfferDao
	OfferCooldownDuration time.Duration
	Notifier              notifierifaces.Notifier
	Fsm                   fsmifaces.Fsm
	EmailSender           emailifaces.EmailSender
}

type ProcessorDependenciesOptions struct {
	TextProvider          textproviface.TextProvider
	DbProvider            dbifaces.DataBaseProvider
	OfferDao              dbifaces.OfferDao
	OfferCooldownDuration time.Duration
	Notifier              notifierifaces.Notifier
	Fsm                   fsmifaces.Fsm
	EmailSender           emailifaces.EmailSender
}

func NewProcessorDependencies(opts ProcessorDependenciesOptions) *ProcessorDependencies {
	return &ProcessorDependencies{
		TextProvider:          opts.TextProvider,
		DbProvider:            opts.DbProvider,
		OfferDao:              opts.OfferDao,
		OfferCooldownDuration: opts.OfferCooldownDuration,
		Notifier:              opts.Notifier,
		Fsm:                   opts.Fsm,
		EmailSender:           opts.EmailSender,
	}
}

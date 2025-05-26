package factories

import (
	logruslogger "pnBot/internal/logger/logruslogger"

	"github.com/sirupsen/logrus"
)

type LogrusLoggerFactory struct {
	Level     logrus.Level
	Formatter logrus.Formatter
}

type LogrusLoggerFactoryOptions struct {
	Level     logrus.Level
	Formatter logrus.Formatter
}

func New(opts LogrusLoggerFactoryOptions) *LogrusLoggerFactory {
	return &LogrusLoggerFactory{
		Level:     opts.Level,
		Formatter: opts.Formatter,
	}
}

type NewModuleLoggerOptions struct {
	BaseLogger *logruslogger.LogrusLogger
	ModuleName string
	Hook       logrus.Hook
}

func (f *LogrusLoggerFactory) NewBaseLogger() *logruslogger.LogrusLogger {
	logger := logrus.New()
	logger.SetLevel(f.Level)
	logger.SetFormatter(f.Formatter)

	return &logruslogger.LogrusLogger{
		Entry: logrus.NewEntry(logger),
	}
}

func (f *LogrusLoggerFactory) NewLoggerWithContext(opts NewModuleLoggerOptions) *logruslogger.LogrusLogger {
	entry := opts.BaseLogger.Entry.WithField("module", opts.ModuleName)
	if opts.Hook != nil {
		entry.Logger.AddHook(opts.Hook)
	}

	return &logruslogger.LogrusLogger{
		Entry: entry,
	}
}

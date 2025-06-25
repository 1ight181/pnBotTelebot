package app

import (
	loggeriface "pnBot/internal/logger/interfaces"
	"pnBot/internal/logger/logruslogger"
	logrusfac "pnBot/internal/logger/logruslogger/factories"

	"github.com/sirupsen/logrus"
)

func createLoggerFactory() loggeriface.LoggerFactory[logrusfac.NewModuleLoggerOptions, logruslogger.LogrusLogger] {
	loggerFactoryOptions := logrusfac.LogrusLoggerFactoryOptions{
		Level: logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     true,
		},
	}
	return logrusfac.New(loggerFactoryOptions)
}

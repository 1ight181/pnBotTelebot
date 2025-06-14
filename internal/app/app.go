package app

import (
	ctx "context"
	"os"
	"os/signal"
	viperprov "pnBot/internal/config/providers/viper"
	"pnBot/internal/logger/extractors"
	loggerfactory "pnBot/internal/logger/logruslogger/factories"
	"pnBot/internal/logger/logruslogger/hooks"
	"syscall"
	"time"
)

func Run() {
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	context, cancel := ctx.WithCancel(ctx.Background())

	loggerFactory := CreateLoggerFactory()
	baseLogger := loggerFactory.NewBaseLogger()

	botContextHook := hooks.ContextHook{
		Extractor: &extractors.BotContextExtractor{},
	}

	botLoggerOptions := loggerfactory.NewModuleLoggerOptions{
		BaseLogger: baseLogger,
		ModuleName: "TELEGRAM_BOT",
		Hook:       &botContextHook,
	}

	dbLoggerOptions := loggerfactory.NewModuleLoggerOptions{
		BaseLogger: baseLogger,
		ModuleName: "DATA_BASE",
		Hook:       nil,
	}

	adminPanelLoggerOptions := loggerfactory.NewModuleLoggerOptions{
		BaseLogger: baseLogger,
		ModuleName: "ADMIN_PANEL",
		Hook:       nil,
	}

	botLogger := loggerFactory.NewLoggerWithContext(botLoggerOptions)
	dbLogger := loggerFactory.NewLoggerWithContext(dbLoggerOptions)
	adminPanelLogger := loggerFactory.NewLoggerWithContext(adminPanelLoggerOptions)

	appConfigOptions := AppConfigOptions{
		Provider:    &viperprov.ViperConfigProvider{},
		FileName:    "config",
		FileType:    "yaml",
		EnvVar:      "PNBOT_CONFIG_PATH",
		DefaultPath: "../configs",
	}

	config, err := loadAppConfig(appConfigOptions)
	if err != nil {
		baseLogger.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	dbProvider, offerDao := CreateDataBase(config.DataBase, dbLogger, context)

	StartBot(&config.Bot, botLogger, dbProvider, offerDao, context)

	StartAdminPanel(config.AdminPanel, config.ImageUploader, dbProvider, adminPanelLogger, context)

	<-stopSignal
	baseLogger.Info("Получен сигнал завершения, остановка...")

	cancel()

	shutdownContext, shutdownCancel := ctx.WithTimeout(ctx.Background(), 3*time.Second)
	defer shutdownCancel()

	<-shutdownContext.Done()
}

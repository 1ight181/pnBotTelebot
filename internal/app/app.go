package app

import (
	ctx "context"
	"os"
	"os/signal"
	"pnBot/internal/banmanager"
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

	loggerFactory := createLoggerFactory()
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

	redisClientLoggerOptions := loggerfactory.NewModuleLoggerOptions{
		BaseLogger: baseLogger,
		ModuleName: "REDIS_CLIENT",
		Hook:       nil,
	}

	botLogger := loggerFactory.NewLoggerWithContext(botLoggerOptions)
	dbLogger := loggerFactory.NewLoggerWithContext(dbLoggerOptions)
	adminPanelLogger := loggerFactory.NewLoggerWithContext(adminPanelLoggerOptions)
	redisClientLogger := loggerFactory.NewLoggerWithContext(redisClientLoggerOptions)

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

	dbProvider, offerDao, userDao := createDataBase(context, config.DataBase, dbLogger)

	redisClient := createRedisClient(context, &config.Cache, redisClientLogger)

	banManager := banmanager.NewBanManager(context, dbProvider, redisClient)
	spamManager := createSpamManager(context, &config.SpamManager, banManager, redisClient)

	startBot(context, &config.Bot, &config.Notifier, &config.Smtp, botLogger, dbProvider, offerDao, spamManager)

	startAdminPanel(context, config.AdminPanel, config.ImageUploader, dbProvider, adminPanelLogger, userDao, banManager, redisClient)

	<-stopSignal
	baseLogger.Info("Получен сигнал завершения, остановка...")

	cancel()

	shutdownContext, shutdownCancel := ctx.WithTimeout(ctx.Background(), 3*time.Second)
	defer shutdownCancel()

	<-shutdownContext.Done()
}

package app

import (
	ctx "context"
	"fmt"
	"time"

	cs "pnBot/internal/scheduler/schedulers/cron"

	c "github.com/robfig/cron/v3"
	"gopkg.in/telebot.v3"

	tb "pnBot/internal/bot"
	errorhandler "pnBot/internal/bot/errorhandlers"
	handlers "pnBot/internal/bot/handlers"
	ifaces "pnBot/internal/bot/interfaces"
	middleware "pnBot/internal/bot/middlewares"
	callback "pnBot/internal/bot/processors/callback"
	command "pnBot/internal/bot/processors/command"
	deps "pnBot/internal/bot/processors/dependencies"
	"pnBot/internal/bot/processors/inlinequery"
	loaders "pnBot/internal/config/loaders"
	models "pnBot/internal/config/models"

	fsm "pnBot/internal/fsm/inmemory"
	tgnotifier "pnBot/internal/notifier/telegram"
	email "pnBot/internal/sender/email"

	dbifaces "pnBot/internal/db/interfaces"
	loggerifaces "pnBot/internal/logger/interfaces"
)

func StartBot(botConfig *models.Bot, notifierConfig *models.Notifier, smtpConfig *models.Smtp, logger loggerifaces.Logger, dbProvider dbifaces.DataBaseProvider, offerDao dbifaces.OfferDao, context ctx.Context) {
	token, isDebug, port, host, webhookUrl := loaders.LoadBotConfig(*botConfig)

	address := fmt.Sprintf("%s:%s", host, port)

	var poller telebot.Poller

	if isDebug {
		poller = &telebot.LongPoller{
			Timeout: time.Second * 5,
		}
	} else {
		poller = &telebot.Webhook{
			Listen: address,
			Endpoint: &telebot.WebhookEndpoint{
				PublicURL: webhookUrl,
			},
		}
	}

	textProvider := CreateTextProvider()

	errorhandler := errorhandler.NewErrorHandler(
		logger,
		textProvider,
	)

	pref := telebot.Settings{
		Token:     token,
		Poller:    poller,
		ParseMode: telebot.ModeMarkdownV2,
		OnError:   errorhandler.HandleError,
	}

	botApi, err := telebot.NewBot(pref)
	if err != nil {
		logger.Fatalf("Ошибка при создании telebot: %v", err)
	}

	cronScheduler := cs.NewCronScheduler(
		c.New(),
		make(map[int]c.EntryID),
	)

	offerCooldown, defaultFrequency, frequencyUnit := loaders.LoadNotifierConfig(*notifierConfig)

	telegramNotifierOptions := tgnotifier.TelegramNotifierOptions{
		DbProvider:            dbProvider,
		OfferDao:              offerDao,
		Scheduler:             cronScheduler,
		Logger:                logger,
		Bot:                   botApi,
		DefaultFrequency:      defaultFrequency,
		FrequencyUnit:         frequencyUnit,
		OfferCooldownDuration: offerCooldown * time.Hour,
	}

	telegramNotifier := tgnotifier.NewTelegramNotifier(telegramNotifierOptions)
	if err := telegramNotifier.Start(); err != nil {
		logger.Fatalf("Ошибка при запуске TelegramNotifier: %v", err)
	}
	logger.Info("TelegramNotifier запущен")

	go func() {
		<-context.Done()
		if err := telegramNotifier.Stop(); err != nil {
			logger.Errorf("Ошибка при завершении работы TelegramNotifier: %v", err)
		}
		logger.Info("TelegramNotifier завершил работу")
	}()

	fsm := fsm.NewInMemoryStateManager()

	host, port, from, password, to := loaders.LoadSmtpConfig(*smtpConfig)

	emailSenderOptions := email.SmtpEmailSenderOptions{
		Host:       host,
		Port:       port,
		From:       from,
		Password:   password,
		AdminEmail: to,
	}

	emailSender := email.NewSmptEmailSender(emailSenderOptions)

	dependenciesOptions := deps.ProcessorDependenciesOptions{
		TextProvider: textProvider,
		DbProvider:   dbProvider,
		OfferDao:     offerDao,
		Notifier:     telegramNotifier,
		Fsm:          fsm,
		EmailSender:  emailSender,
	}

	dependencies := deps.NewProcessorDependencies(dependenciesOptions)

	commandProcessor := command.NewCommandProcessor(dependencies)
	commandHandler := handlers.NewHandler(
		telebot.OnText,
		commandProcessor.ProcessCommand,
	)

	callbackProcessor := callback.NewCallbackProcessor(dependencies)
	callbackHandler := handlers.NewHandler(
		telebot.OnCallback,
		callbackProcessor.ProcessCallback,
	)

	inlineQueryProcessor := inlinequery.NewInlineQueryProcessor(dependencies)
	inlineQueryHandler := handlers.NewHandler(
		telebot.OnQuery,
		inlineQueryProcessor.ProcessInlineQuery,
	)

	handlers := []ifaces.Handler{
		commandHandler,
		callbackHandler,
		inlineQueryHandler,
	}

	middlewares := []telebot.MiddlewareFunc{
		middleware.LogMiddleware(logger),
		errorhandler.ErrorMiddleware(),
	}

	botOptions := tb.TelegramBotOptions{
		BotApi:      botApi,
		Handlers:    handlers,
		Middlewares: middlewares,
		Context:     context,
		Logger:      logger,
	}

	bot := tb.New(botOptions)

	bot.Start()
	logger.Infof("Бот запущен и слушает на %s", address)
}

package app

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/telebot.v3"

	tb "pnBot/internal/bot"
	errorhandler "pnBot/internal/bot/errorhandlers"
	handlers "pnBot/internal/bot/handlers"
	ifaces "pnBot/internal/bot/interfaces"
	middleware "pnBot/internal/bot/middlewares"
	callback "pnBot/internal/bot/processors/callback"
	command "pnBot/internal/bot/processors/command"
	deps "pnBot/internal/bot/processors/dependencies"
	loaders "pnBot/internal/config/loaders"
	models "pnBot/internal/config/models"
	dbifaces "pnBot/internal/db/interfaces"
	loggerifaces "pnBot/internal/logger/interfaces"
)

func StartBot(botConfig *models.Bot, logger loggerifaces.Logger, dbProvider dbifaces.DataBaseProvider, ctx context.Context) {
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

	errorhandler := errorhandler.New(logger)

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

	textProvider := CreateTextProvider()

	dependenciesOptions := deps.ProcessorDependenciesOptions{
		TextProvider: textProvider,
		DbProvider:   dbProvider,
	}

	dependencies := deps.New(dependenciesOptions)

	commandProcessor := command.New(dependencies)
	commandHandler := handlers.New(
		telebot.OnText,
		commandProcessor.ProcessCommand,
	)

	callbackProcessor := callback.New(dependencies)
	callbackHandler := handlers.New(
		telebot.OnCallback,
		callbackProcessor.ProcessCallback,
	)

	handlers := []ifaces.Handler{
		commandHandler,
		callbackHandler,
	}

	middlewares := []telebot.MiddlewareFunc{
		middleware.LogMiddleware(logger),
	}

	botOptions := tb.TelegramBotOptions{
		BotApi:      botApi,
		Handlers:    handlers,
		Middlewares: middlewares,
		Context:     ctx,
		Logger:      logger,
	}

	bot := tb.New(botOptions)

	bot.Start()
	logger.Infof("Бот запущен и слушает на %s", address)
}

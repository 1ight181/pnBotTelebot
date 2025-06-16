package telegram

import (
	"context"
	"errors"
	"fmt"
	"pnBot/internal/bot/processors/common"
	dberrors "pnBot/internal/db/errors"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	loggerifaces "pnBot/internal/logger/interfaces"
	units "pnBot/internal/notifier/units"
	schedulerifaces "pnBot/internal/scheduler/interfaces"
	"strconv"
	"sync"
	"time"

	"gopkg.in/telebot.v3"
)

type TelegramNotifier struct {
	mu                    sync.Mutex
	userFrequency         map[int64]int
	userJobs              map[int64]int
	dbProvider            dbifaces.DataBaseProvider
	offerDao              dbifaces.OfferDao
	scheduler             schedulerifaces.Scheduler
	logger                loggerifaces.Logger
	bot                   *telebot.Bot
	defaultFrequency      int
	frequencyUnit         units.FrequencyUnit
	offerCooldownDuration time.Duration
}

type TelegramNotifierOptions struct {
	DbProvider            dbifaces.DataBaseProvider
	OfferDao              dbifaces.OfferDao
	Scheduler             schedulerifaces.Scheduler
	Logger                loggerifaces.Logger
	Bot                   *telebot.Bot
	DefaultFrequency      int
	FrequencyUnit         units.FrequencyUnit
	OfferCooldownDuration time.Duration
}

func NewTelegramNotifier(opts TelegramNotifierOptions) *TelegramNotifier {
	return &TelegramNotifier{
		userFrequency:         make(map[int64]int),
		userJobs:              make(map[int64]int),
		dbProvider:            opts.DbProvider,
		offerDao:              opts.OfferDao,
		scheduler:             opts.Scheduler,
		logger:                opts.Logger,
		bot:                   opts.Bot,
		defaultFrequency:      opts.DefaultFrequency,
		frequencyUnit:         opts.FrequencyUnit,
		offerCooldownDuration: opts.OfferCooldownDuration,
	}
}

func (tn *TelegramNotifier) Start() error {
	if err := tn.getSubscribedUsers(); err != nil {
		return fmt.Errorf("ошибка при получении подписанных пользователей: %v", err)
	}
	tn.scheduler.Start()
	return nil
}

func (tn *TelegramNotifier) getSubscribedUsers() error {
	var users []dbmodels.User
	if err := tn.dbProvider.Find(context.Background(), &users, "is_subscribed = ?", true); err != nil {
		return fmt.Errorf("ошибка при получении подписанных пользователей: %v", err)
	}

	for _, user := range users {
		if err := tn.AddUser(user.TgId); err != nil {
			return fmt.Errorf("ошибка при добавлении пользователя с ID %d: %v", user.TgId, err)
		}
	}

	return nil
}

func (tn *TelegramNotifier) getNewOfferAndSendToUser(userId int64) error {
	offerCooldown := time.Now().Add(-tn.offerCooldownDuration)
	limit := 1
	offers, err := tn.offerDao.GetLastAvailableOffers(userId, limit, offerCooldown)
	if errors.Is(err, dberrors.ErrRecordNotFound) {
		return nil
	} else if err != nil {
		return fmt.Errorf("ошибка при получении последней доступной акции для пользователя %d: %v", userId, err)
	}

	offer := offers[0]

	offerCreatives := offer.Creatives

	var offerCreative dbmodels.Creative
	if len(offerCreatives) != 0 {
		offerCreative = offerCreatives[0]
	}

	offerImageUrl := offerCreative.ResourceUrl
	offerTitle := offer.Title
	offerDescription := offer.Description

	escapedTitle := common.EscapeMarkdownV2(offerTitle)
	escapedDescription := common.EscapeMarkdownV2(offerDescription)

	escapedDescription = common.WrapURLsWithPreviousWord(escapedDescription)

	offerMessage := &telebot.Photo{
		File:    telebot.FromURL(offerImageUrl),
		Caption: fmt.Sprintf("*%s* \n\n%s", escapedTitle, escapedDescription),
	}

	recipient := &telebot.User{ID: userId}

	_, err = tn.bot.Send(recipient, offerMessage)
	if err != nil {
		return fmt.Errorf("ошибка отправки сообщения пользователю %d: %w", userId, err)
	}
	return nil
}

func (tn *TelegramNotifier) AddUser(userId int64) error {
	tn.mu.Lock()
	defer tn.mu.Unlock()

	jobId, err := tn.scheduleUserJob(userId, tn.defaultFrequency)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении задачи для пользователя с ID %d: %v", userId, err)
	}

	tn.userJobs[userId] = jobId
	tn.userFrequency[userId] = tn.defaultFrequency

	return nil
}

func (tn *TelegramNotifier) RemoveUser(userId int64) error {
	tn.mu.Lock()
	defer tn.mu.Unlock()

	if _, exists := tn.userFrequency[userId]; !exists {
		return fmt.Errorf("пользователь с ID %d не найден в списке рассылок", userId)
	}

	if jobId, exists := tn.userJobs[userId]; exists {
		if err := tn.scheduler.RemoveJob(jobId); err != nil {
			return fmt.Errorf("ошибка при удалении задачи для пользователя с ID %d: %v", userId, err)
		}
		delete(tn.userJobs, userId)
		delete(tn.userFrequency, userId)
	}

	return nil
}

func (tn *TelegramNotifier) SetUserFrequency(userId int64, frequency int) error {
	tn.mu.Lock()
	defer tn.mu.Unlock()

	if _, exists := tn.userFrequency[userId]; !exists {
		return fmt.Errorf("пользователь с ID %d не найден в списке рассылок", userId)
	}

	if frequency <= 0 {
		return fmt.Errorf("частота рассылки не может быть отрицательной для пользователя с ID %d", userId)
	}

	if jobId, exists := tn.userJobs[userId]; exists {
		if err := tn.scheduler.RemoveJob(jobId); err != nil {
			return fmt.Errorf("ошибка при удалении задачи для пользователя с ID %d: %v", userId, err)
		}
	}

	user := dbmodels.User{
		TgId: userId,
	}

	if err := tn.dbProvider.Update(context.Background(), user, "notification_frequency", frequency); err != nil {
		return fmt.Errorf("ошибка при обновлении частоты рассылки для пользователя с ID %d: %v", userId, err)
	}

	jobId, err := tn.scheduleUserJob(userId, frequency)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении задачи для пользователя с ID %d: %v", userId, err)
	}

	tn.userJobs[userId] = jobId
	tn.userFrequency[userId] = frequency

	return nil
}

func (tn *TelegramNotifier) GetFrequency(userId int64) (int, error) {
	tn.mu.Lock()
	defer tn.mu.Unlock()

	if frequency, exists := tn.userFrequency[userId]; exists {
		return frequency, nil
	}

	return 0, fmt.Errorf("пользователь с ID %d не найден в списке рассылок", userId)
}

func (tn *TelegramNotifier) GetFrequencyUnit() (units.FrequencyUnit, error) {
	return tn.frequencyUnit, nil
}

func (tn *TelegramNotifier) scheduleUserJob(userId int64, frequency int) (int, error) {
	return tn.scheduler.AddJob("@every "+strconv.Itoa(frequency)+tn.frequencyUnit.String(), func() {
		if err := tn.getNewOfferAndSendToUser(userId); err != nil {
			tn.logger.Errorf(err.Error())
		}
	})
}

func (tn *TelegramNotifier) Stop() error {
	tn.mu.Lock()
	defer tn.mu.Unlock()

	for userId, jobId := range tn.userJobs {
		if err := tn.scheduler.RemoveJob(jobId); err != nil {
			return fmt.Errorf("ошибка при удалении задачи для пользователя %d: %v", userId, err)
		}
	}

	tn.scheduler.Stop()
	return nil
}

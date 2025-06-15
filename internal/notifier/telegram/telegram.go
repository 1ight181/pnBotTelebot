package telegram

import (
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
	mu            sync.Mutex
	userFrequency map[int64]int
	userJobs      map[int64]int
	dbProvider    dbifaces.DataBaseProvider
	offerDao      dbifaces.OfferDao
	scheduler     schedulerifaces.Scheduler
	logger        loggerifaces.Logger
	bot           *telebot.Bot
}

type TelegramNotifierOptions struct {
	DbProvider dbifaces.DataBaseProvider
	OfferDao   dbifaces.OfferDao
	Scheduler  schedulerifaces.Scheduler
	Logger     loggerifaces.Logger
	Bot        *telebot.Bot
}

func NewTelegramNotifier(opts TelegramNotifierOptions) *TelegramNotifier {
	return &TelegramNotifier{
		userFrequency: make(map[int64]int),
		userJobs:      make(map[int64]int),
		dbProvider:    opts.DbProvider,
		offerDao:      opts.OfferDao,
		scheduler:     opts.Scheduler,
		logger:        opts.Logger,
		bot:           opts.Bot,
	}
}

func (tn *TelegramNotifier) getNewOfferAndSendToUser(userId int64, offerCooldownDuration time.Duration) error {
	offerCooldown := time.Now().Add(-offerCooldownDuration)
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

func (tn *TelegramNotifier) AddUser(userId int64, frequency int, offerCooldownDuration time.Duration, frequencyUnit units.FrequencyUnit) error {
	tn.mu.Lock()
	defer tn.mu.Unlock()

	if _, exists := tn.userFrequency[userId]; exists {
		return fmt.Errorf("пользователь с ID %d уже существует в списке рассылок", userId)
	}

	if frequency <= 0 {
		return fmt.Errorf("частота рассылки не может быть отрицательной для пользователя с ID %d", userId)
	}

	jobId, err := tn.scheduleUserJob(userId, frequency, offerCooldownDuration, frequencyUnit)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении задачи для пользователя с ID %d: %v", userId, err)
	}

	tn.userJobs[userId] = jobId

	tn.userFrequency[userId] = frequency

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

func (tn *TelegramNotifier) SetUserFrequency(userId int64, frequency int, offerCooldownDuration time.Duration, frequencyUnit units.FrequencyUnit) error {
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

	jobId, err := tn.scheduleUserJob(userId, frequency, offerCooldownDuration, frequencyUnit)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении задачи для пользователя с ID %d: %v", userId, err)
	}

	tn.userJobs[userId] = jobId
	tn.userFrequency[userId] = frequency

	return nil
}

func (tn *TelegramNotifier) scheduleUserJob(userId int64, frequency int, offerCooldownDuration time.Duration, frequencyUnit units.FrequencyUnit) (int, error) {
	return tn.scheduler.AddJob("@every "+strconv.Itoa(frequency)+frequencyUnit.String(), func() {
		if err := tn.getNewOfferAndSendToUser(userId, offerCooldownDuration); err != nil {
			tn.logger.Errorf(err.Error())
		}
	})
}

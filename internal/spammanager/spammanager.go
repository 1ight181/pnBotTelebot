package redis

import (
	"errors"
	"fmt"
	banifaces "pnBot/internal/banmanager/interfaces"
	cacheerrors "pnBot/internal/cache/errors"
	cacheifaces "pnBot/internal/cache/interfaces"
	"time"
)

type SpamManager struct {
	cacheProvider cacheifaces.CacheProvider
	messageLimit  int
	interval      time.Duration
	warnLimit     int
	banManager    banifaces.BanManager
	banDuration   time.Duration
	banReasonText string
	banAuthor     string
}

type SpamManagerOptions struct {
	CacheProvider cacheifaces.CacheProvider
	MessageLimit  int
	Interval      time.Duration
	WarnLimit     int
	BanManager    banifaces.BanManager
	BanDuration   time.Duration
	BanReasonText string
	BanAuthor     string
}

func NewSpamManager(opts SpamManagerOptions) *SpamManager {
	return &SpamManager{
		cacheProvider: opts.CacheProvider,
		messageLimit:  opts.MessageLimit,
		interval:      opts.Interval,
		warnLimit:     opts.WarnLimit,
		banManager:    opts.BanManager,
		banDuration:   opts.BanDuration,
		banReasonText: opts.BanReasonText,
		banAuthor:     opts.BanAuthor,
	}
}

func (sm *SpamManager) IsAllowed(userId int64) (banned bool, warned bool, remaining int, err error) {
	isBannedKey := fmt.Sprintf("user:ban:%d", userId)
	warnCountKey := fmt.Sprintf("user:warncount:%d", userId)
	messageCountKey := fmt.Sprintf("user:msgcount:%d", userId)

	val, err := sm.cacheProvider.Get(isBannedKey)
	if err != nil && !errors.Is(err, cacheerrors.ErrNilVal) {
		return false, false, 0, err
	}
	if val == "true" {
		return true, false, 0, nil
	}

	messageCount, err := sm.cacheProvider.Incr(messageCountKey)
	if err != nil {
		return false, false, 0, err
	}
	if messageCount == 1 {
		_ = sm.cacheProvider.Expire(messageCountKey, sm.interval)
	}

	if int(messageCount) <= sm.messageLimit {
		remaining = sm.warnLimit
		return false, false, remaining, nil
	}

	warnCount, err := sm.cacheProvider.Incr(warnCountKey)
	if err != nil {
		return false, false, 0, err
	}
	if warnCount == 1 {
		_ = sm.cacheProvider.Expire(warnCountKey, sm.interval)
	}

	remaining = sm.warnLimit - int(warnCount)
	warned = true

	if int(warnCount) >= sm.warnLimit {
		banned, err := sm.banManager.IsBanned(userId)
		if err != nil {
			return false, false, 0, err
		}
		if !banned {
			if err := sm.banManager.Ban(userId, sm.banReasonText, sm.banDuration, sm.banAuthor); err != nil {
				return false, false, 0, err
			}
		}
		return true, true, 0, nil
	}

	return false, warned, remaining, nil
}

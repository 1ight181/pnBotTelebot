package loaders

import (
	conf "pnBot/internal/config/models"
	"strconv"
	"time"
)

func LoadSpamManager(spamManager conf.SpamManager) (int, time.Duration, int, time.Duration, string, string) {
	messageLimit, _ := strconv.Atoi(spamManager.MessageLimit)
	interval, _ := strconv.Atoi(spamManager.Interval)
	intervalSecond := time.Duration(interval) * time.Second
	warnLimit, _ := strconv.Atoi(spamManager.WarnLimit)
	banDuration, _ := strconv.Atoi(spamManager.BanDuration)
	banDurationHours := time.Duration(banDuration) * time.Hour
	banReasonText := spamManager.BanReasonText
	banAuthor := spamManager.BanAuthor

	return messageLimit, intervalSecond, warnLimit, banDurationHours, banReasonText, banAuthor
}

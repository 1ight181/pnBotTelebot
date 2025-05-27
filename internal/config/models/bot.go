package models

import (
	"errors"
)

type Bot struct {
	Token      string `mapstructure:"token"`
	IsDebug    bool   `mapstructure:"is_debug"`
	Port       string `mapstructure:"port"`
	Host       string `mapstructure:"host"`
	WebhookUrl string `mapstructure:"webhook_url"`
}

func (b *Bot) Validate() error {
	if b.Token == "" {
		return errors.New("требуется указание токена бота")
	}
	if b.Port == "" {
		return errors.New("требуется указание порта")
	}
	if b.Host == "" {
		return errors.New("требуется указание хоста")
	}
	if b.WebhookUrl == "" {
		return errors.New("требуется указание Webhook_URL")
	}
	return nil
}

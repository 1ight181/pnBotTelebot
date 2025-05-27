package loaders

import (
	conf "pnBot/internal/config/models"
)

func LoadBotConfig(botConfig conf.Bot) (string, bool, string, string, string) {
	token := botConfig.Token
	is_debug := botConfig.IsDebug
	port := botConfig.Port
	host := botConfig.Host
	WebhookUrl := botConfig.WebhookUrl

	return token, is_debug, port, host, WebhookUrl
}

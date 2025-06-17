package loaders

import (
	conf "pnBot/internal/config/models"
)

func LoadSmtpConfig(smtpConfig conf.Smtp) (string, string, string, string, string) {
	host := smtpConfig.Host
	port := smtpConfig.Port
	from := smtpConfig.From
	password := smtpConfig.Password
	to := smtpConfig.To

	return host, port, from, password, to
}

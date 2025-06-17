package email

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

type SmtpEmailSender struct {
	from           string
	password       string
	host           string
	port           string
	adminEmail     string
	authentication smtp.Auth
}

type SmtpEmailSenderOptions struct {
	From       string
	Password   string
	Host       string
	Port       string
	AdminEmail string
}

func NewSmptEmailSender(opts SmtpEmailSenderOptions) *SmtpEmailSender {
	smptEmailSender := &SmtpEmailSender{
		from:       opts.From,
		password:   opts.Password,
		host:       opts.Host,
		port:       opts.Port,
		adminEmail: opts.AdminEmail,
	}
	smptEmailSender.auth()
	return smptEmailSender
}

func (s *SmtpEmailSender) auth() {
	s.authentication = smtp.PlainAuth("", s.from, s.password, s.host)
}

func (s *SmtpEmailSender) Send(to []string, subject string, body string) error {
	toHeader := strings.Join(to, ", ")
	encodedSubject := fmt.Sprintf("=?UTF-8?B?%s?=", base64.StdEncoding.EncodeToString([]byte(subject)))

	message := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"UTF-8\"\r\n\r\n%s",
		s.from, toHeader, encodedSubject, body,
	))
	address := fmt.Sprintf("%s:%s", s.host, s.port)

	return smtp.SendMail(address, s.authentication, s.from, to, message)
}
func (s *SmtpEmailSender) SendToAdmin(subject string, body string) error {
	encodedSubject := fmt.Sprintf("=?UTF-8?B?%s?=", base64.StdEncoding.EncodeToString([]byte(subject)))

	message := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"UTF-8\"\r\n\r\n%s",
		s.from, s.adminEmail, encodedSubject, body,
	))
	address := fmt.Sprintf("%s:%s", s.host, s.port)

	to := []string{s.adminEmail}
	return smtp.SendMail(address, s.authentication, s.from, to, message)
}

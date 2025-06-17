package interfaces

type EmailSender interface {
	Send(to []string, subject string, body string) error
	SendToAdmin(subject string, body string) error
}

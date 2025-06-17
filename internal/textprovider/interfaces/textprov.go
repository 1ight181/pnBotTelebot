package interfaces

type TextProvider interface {
	GetButtonText(string) string
	GetText(string) string
	GetInlineQueryTitle(string) string
	GetInlineQueryDescription(string) string
	GetCallbackText(string) string
	GetEmailSubject(key string) string
}

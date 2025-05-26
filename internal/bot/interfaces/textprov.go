package interfaces

type TextProvider interface {
	GetButtonText(string) string
	GetText(string) string
	GetInlineQueryTitle(string) string
	GetInlineQueryText(string) string
	GetCallbackText(string) string
}

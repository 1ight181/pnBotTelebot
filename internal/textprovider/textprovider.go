package textprovider

type TextProvider struct {
	texts                   map[string]string
	buttonTexts             map[string]string
	inlineQueryTitles       map[string]string
	inlineQueryDescriptions map[string]string
	callbackTexts           map[string]string
	emailSubject            map[string]string
}

type TextProviderOptions struct {
	Texts                   map[string]string
	ButtonTexts             map[string]string
	InlineQueryTitles       map[string]string
	InlineQueryDescriptions map[string]string
	CallbackTexts           map[string]string
	EmailSubject            map[string]string
}

func NewTextProvider(opts TextProviderOptions) *TextProvider {
	return &TextProvider{
		texts:                   opts.Texts,
		buttonTexts:             opts.ButtonTexts,
		inlineQueryTitles:       opts.InlineQueryTitles,
		inlineQueryDescriptions: opts.InlineQueryDescriptions,
		callbackTexts:           opts.CallbackTexts,
		emailSubject:            opts.EmailSubject,
	}
}

func (tp *TextProvider) GetText(key string) string {
	if text, exists := tp.texts[key]; exists {
		return text
	}
	return "Неизвестный текст"
}

func (tp *TextProvider) GetButtonText(key string) string {
	if text, exists := tp.buttonTexts[key]; exists {
		return text
	}
	return "Неизвестный текст"
}

func (tp *TextProvider) GetInlineQueryTitle(key string) string {
	if text, exists := tp.inlineQueryTitles[key]; exists {
		return text
	}
	return "Неизвестный текст"
}

func (tp *TextProvider) GetInlineQueryDescription(key string) string {
	if text, exists := tp.inlineQueryDescriptions[key]; exists {
		return text
	}
	return "Неизвестный текст"
}

func (tp *TextProvider) GetCallbackText(key string) string {
	if text, exists := tp.callbackTexts[key]; exists {
		return text
	}
	return "Неизвестный текст"
}

func (tp *TextProvider) GetEmailSubject(key string) string {
	if text, exists := tp.emailSubject[key]; exists {
		return text
	}
	return "Неизвестный текст"
}

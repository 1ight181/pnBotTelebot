package textprovider

type TextProvider struct {
	texts             map[string]string
	buttonTexts       map[string]string
	inlineQueryTitles map[string]string
	inlineQueryTexts  map[string]string
	callbackTexts     map[string]string
}

type TextProviderOptions struct {
	Texts             map[string]string
	ButtonTexts       map[string]string
	InlineQueryTitles map[string]string
	InlineQueryTexts  map[string]string
	CallbackTexts     map[string]string
}

func NewTextProvider(opts TextProviderOptions) *TextProvider {
	return &TextProvider{
		texts:             opts.Texts,
		buttonTexts:       opts.ButtonTexts,
		inlineQueryTitles: opts.InlineQueryTitles,
		inlineQueryTexts:  opts.InlineQueryTexts,
		callbackTexts:     opts.CallbackTexts,
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

func (tp *TextProvider) GetInlineQueryText(key string) string {
	if text, exists := tp.inlineQueryTexts[key]; exists {
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

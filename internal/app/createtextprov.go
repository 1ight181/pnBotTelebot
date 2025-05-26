package app

import (
	textproviface "pnBot/internal/bot/interfaces"
	textprov "pnBot/internal/bot/textprovider"
)

func CreateTextProvider() textproviface.TextProvider {
	texts := map[string]string{
		"start": "Привет\\! Я *PromoNotifyBot*\\. Я помогаю узнавать о новых акциях и скидках\\. Нажми 'Подписаться', чтобы подписаться на уведомления\\.",
		"help": "Доступные команды:\n\n" +
			"/start \\- Начать взаимодействие с ботом\n" +
			"/help \\- Получить помощь по командам\n" +
			"Если у вас есть вопросы или предложения, не стесняйтесь обращаться к нам через команду /feedback",
	}

	buttonTexts := map[string]string{
		"subscribe":   "Подписаться",
		"unsubscribe": "Отписаться",
	}

	inlineQueryTitles := map[string]string{
		"last":    "Самое свежее:",
		"empty":   "Пустой запрос\\!",
		"unknown": "Неизвестный запрос\\!",
	}

	inlineQueryTexts := map[string]string{
		"last":    "Последняя из акций: \n",
		"empty":   "Введите 'last' чтобы увидеть последние акции",
		"unknown": "Введите 'last' чтобы увидеть последние акции",
	}

	callbackTexts := map[string]string{
		"subscribe":   "Вы успешно подписались на уведомления!",
		"unsubscribe": "Вы успешно отписались от уведомлений!",
	}

	textProviderOpts := textprov.TextProviderOptions{
		Texts:             texts,
		ButtonTexts:       buttonTexts,
		InlineQueryTitles: inlineQueryTitles,
		InlineQueryTexts:  inlineQueryTexts,
		CallbackTexts:     callbackTexts,
	}

	return textprov.NewTextProvider(textProviderOpts)
}

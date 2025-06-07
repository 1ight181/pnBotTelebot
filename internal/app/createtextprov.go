package app

import (
	textproviface "pnBot/internal/bot/interfaces"
	textprov "pnBot/internal/bot/textprovider"
)

func CreateTextProvider() textproviface.TextProvider {
	texts := map[string]string{
		"first_start": "👋 Привет\\! Я *PromoNotifyBot* — твой верный помощник в мире *акций и скидок*💸✨\n\n" +
			"Хочешь всегда ловить самые горячие предложения и не пропускать лучшие скидки\\?🔥\n\n" +
			"Тогда жми *Подписаться* ✅ и начнем экономить вместе\\!",

		"help": "Доступные команды:\n\n" +
			"/start \\- начать работу с ботом\n" +
			"/help \\- показать это сообщение\n" +
			"/feedback \\- оставить отзыв или предложение\n\n" +
			"Мы всегда рады твоим сообщениям\\! 💬",

		"menu": "🏠 *Главное меню*\n\nВыбери действие ниже, чтобы продолжить 👇",

		"not_first_start_subscribed": "Ты уже подписан на свежие акции\\! 😉",

		"not_first_start_not_subscribed": "Привет снова\\! Я помню тебя, но вижу, что ты еще не подписан на уведомления\\.\n" +
			"Жми *Подписаться* ✅, чтобы не пропускать самые вкусные предложения\\!🎁",

		"not_allowed":     "Сначала тебе нужно подписаться\\!",
		"category_filter": "Доступны следующие категории:",
		"process":         "Обработка...",
	}

	buttonTexts := map[string]string{
		"subscribe":       "Подписаться ✅",
		"unsubscribe":     "Отписаться ❌",
		"filter_settings": "Настроить фильтры ⚙️",
		"last_promo":      "Последние акции 🔥",
		"apply_filter":    "Сохранить фильтры 💾",
	}

	inlineQueryTitles := map[string]string{
		"last":    "Самое свежее для тебя ✨",
		"empty":   "Пустой запрос\\! Введи что-нибудь 🔍",
		"unknown": "Неизвестный запрос\\! Попробуй еще раз\\.",
	}

	inlineQueryTexts := map[string]string{
		"last":    "Последняя акция: \n",
		"empty":   "Введите 'last', чтобы увидеть самые свежие акции",
		"unknown": "Введите 'last', чтобы увидеть самые свежие акции",
	}

	callbackTexts := map[string]string{
		"subscribe":          "Ты успешно подписался на уведомления! 🎉",
		"unsubscribe":        "Ты отписался от уведомлений. Надеемся, вернешься! 👋",
		"already_subscribed": "Ты уже подписан\\! 👍",
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

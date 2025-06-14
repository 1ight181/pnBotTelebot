package app

import (
	textproviface "pnBot/internal/bot/interfaces"
	textprov "pnBot/internal/bot/textprovider"
)

func CreateTextProvider() textproviface.TextProvider {
	texts := map[string]string{
		"first_start": "👋 Привет\\! Я *Промо* — твой верный помощник в мире *акций и скидок*💸✨\n\n" +
			"Хочешь всегда ловить самые горячие предложения и не пропускать лучшие скидки\\?🔥\n\n" +
			"Тогда жми *Подписаться* ✅ и начнем экономить вместе\\!",
		"not_first_start_subscribed": "Ты уже подписан на свежие акции\\! 😉",

		"not_first_start_not_subscribed": "Привет снова\\! Я помню тебя, но вижу, что ты еще не подписан на уведомления\\.\n" +
			"Жми *Подписаться* ✅, чтобы не пропускать самые вкусные предложения\\!🎁",

		"help": "Доступные команды:\n\n" +
			"/start \\- начать работу с ботом\n" +
			"/help \\- показать это сообщение\n" +
			"/feedback \\- оставить отзыв или предложение\n\n" +
			"Мы всегда рады твоим сообщениям\\! 💬",

		"menu": "🏠 *Главное меню*\n\nВыбери действие ниже, чтобы продолжить 👇",

		"not_allowed":        "Сначала тебе нужно подписаться\\!",
		"category_filter":    "Доступны следующие категории:",
		"process":            "Обработка...",
		"no_available_offer": "Пока ничего новенького нет\\! Поробуй позже 🕐 Или попробуй выбрать больше доступных категорий",
		"error":              "Произошла внутренняя ошибка\\. Пожалуйста, попробуйте позже\\.",
		"unknown":            "Неизвестнная команда\\! Список доступных команд можно узнать, написав /help\\.",
	}

	buttonTexts := map[string]string{
		"subscribe":       "Подписаться ✅",
		"unsubscribe":     "Отписаться ❌",
		"filter_settings": "Настроить фильтры ⚙️",
		"last_promo":      "Последние акции 🔥",
		"apply_filter":    "Сохранить фильтры 💾",
		"next_offer":      "Следующее ➡️",
		"menu":            "В меню 🏠",
	}

	inlineQueryTitles := map[string]string{
		"last":               "Самое свежее для тебя ✨",
		"empty":              "Пустой запрос!",
		"unknown":            "Неизвестный запроc!",
		"no_available_offer": "Пока тут пусто!😕",
	}

	inlineQueryDescription := map[string]string{
		"last":               "Последняя акция: ",
		"empty":              "Введите 'last', чтобы увидеть самые свежие акции🔍",
		"unknown":            "Введите 'last', чтобы увидеть самые свежие акции🔍",
		"no_available_offer": "Пока ничего новенького нет! Поробуй позже 🕐 Или попробуй выбрать больше доступных категорий",
	}

	callbackTexts := map[string]string{
		"subscribe":          "Ты успешно подписался на уведомления! 🎉",
		"unsubscribe":        "Ты отписался от уведомлений. Надеемся, вернешься! 👋",
		"already_subscribed": "Ты уже подписан\\! 👍",
	}

	textProviderOpts := textprov.TextProviderOptions{
		Texts:                   texts,
		ButtonTexts:             buttonTexts,
		InlineQueryTitles:       inlineQueryTitles,
		InlineQueryDescriptions: inlineQueryDescription,
		CallbackTexts:           callbackTexts,
	}

	return textprov.NewTextProvider(textProviderOpts)
}

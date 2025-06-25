package app

import (
	textprov "pnBot/internal/textprovider"
	textproviface "pnBot/internal/textprovider/interfaces"
)

func createTextProvider() textproviface.TextProvider {
	texts := map[string]string{
		"first_start": "👋 Привет\\! Я *Промо* — твой верный помощник в мире *акций и скидок*💸✨\n\n" +
			"Хочешь всегда ловить самые горячие предложения и не пропускать лучшие скидки\\?🔥\n\n" +
			"Тогда жми *Подписаться* ✅ и начнем экономить вместе\\!",
		"not_first_start_subscribed": "Ты уже подписан на свежие акции\\! 😉",

		"not_first_start_not_subscribed": "Привет\\! Я помню тебя, но вижу, что ты еще не подписан на уведомления\\.\n" +
			"Жми *Подписаться* ✅, чтобы не пропускать самые вкусные предложения\\!🎁",

		"help": "Доступные команды:\n\n" +
			"/start \\- начать работу с ботом\n" +
			"/help \\- показать это сообщение\n",

		"menu": "🏠 *Главное меню*\n\nВыбери действие ниже, чтобы продолжить 👇",

		"not_allowed":        "Сначала тебе нужно подписаться\\!",
		"category_filter":    "Доступны следующие категории:",
		"processing":         "Обработка\\.\\.\\.",
		"no_available_offer": "Пока ничего новенького нет\\! Поробуй позже 🕐 Или попробуй выбрать больше доступных категорий",
		"error":              "Произошла внутренняя ошибка\\. Пожалуйста, попробуйте позже\\.",
		"unknown":            "Неизвестнная команда\\! Список доступных команд можно узнать, написав /help\\.",
		"frequency_settings": "Настрой частоту уведомлений\\:",
		"frequency_setted":   "Частота уведомлений успешно изменена\\!",
		"not_subscribed":     "Ты не подписан на уведомления\\! Напиши /start\\!",
		"feedback_hint":      "Напиши свой отзыв или пожелания по боту\\. Мы все читаем\\!🌝",
		"bug_report_hint": "Напиши об ошибке, которую ты обнаружил\\. Постарайся ответить на следующие вопросы:\n" +
			"⭐️Как проявляется ошибка? \\(появилось соощение об ошибке, некорректное поведение и т\\.п\\.\\)\n" +
			"⭐️Какое именно действие вызвало проблему? \\(ты нажал кнопку, вызвал какую\\-то команду и т\\.д\\.\\)\n" +
			"⭐️Что ты делал до того, как появилась ошибка? \\(по возможности опиши свои последние действия до того, как возникла ошибка\\)\n" +
			"⭐️Повторялась ли ошибка?\n",
		"ban":  "🚫 Вы были забанены\\.",
		"warn": "⚠️ Пожалуйста, не спамьте\\. До блокировки осталось %d предупреждений\\.",
	}

	buttonTexts := map[string]string{
		"subscribe":          "Подписаться ✅",
		"unsubscribe":        "Отписаться ❌",
		"filter_settings":    "Настроить фильтры ⚙️",
		"last_promo":         "Последние акции 🔥",
		"apply_filter":       "Сохранить фильтры 💾",
		"next_offer":         "Следующее ➡️",
		"menu":               "В меню 🏠",
		"frequency_settings": "Настроить частоту уведомлений ⏰",
		"feedback":           "Оставить отзыв 💬",
		"every_x_hours":      "Каждые %d ч. ⏳",
		"everyday":           "Каждый день 🌅",
		"bug_report":         "Сообщить об ошибке 🐛",
	}

	inlineQueryTitles := map[string]string{
		"last":               "Самое свежее для тебя ✨",
		"empty":              "Пустой запрос!",
		"unknown":            "Неизвестный запроc!",
		"no_available_offer": "Пока тут пусто!😕",
		"error":              "Произошла ошибка при обработке запроса.",
	}

	inlineQueryDescription := map[string]string{
		"last":               "Последняя акция: ",
		"empty":              "Введите 'last', чтобы увидеть самые свежие акции🔍",
		"unknown":            "Введите 'last', чтобы увидеть самые свежие акции🔍",
		"no_available_offer": "Пока ничего новенького нет! Поробуй позже 🕐 Или попробуй выбрать больше доступных категорий",
		"error":              "Произошла ошибка при обработке запроса. Пожалуйста, попробуй позже",
	}

	callbackTexts := map[string]string{
		"subscribe":          "Ты успешно подписался на уведомления! 🎉",
		"unsubscribe":        "Ты отписался от уведомлений. Надеемся, вернешься! 👋",
		"already_subscribed": "Ты уже подписан! 👍",
	}

	emailSubject := map[string]string{
		"feedback":   "Новый отзыв из Promo",
		"bug_report": "Новый отчет об ошибке из Promo",
	}

	textProviderOpts := textprov.TextProviderOptions{
		Texts:                   texts,
		ButtonTexts:             buttonTexts,
		InlineQueryTitles:       inlineQueryTitles,
		InlineQueryDescriptions: inlineQueryDescription,
		CallbackTexts:           callbackTexts,
		EmailSubject:            emailSubject,
	}

	return textprov.NewTextProvider(textProviderOpts)
}

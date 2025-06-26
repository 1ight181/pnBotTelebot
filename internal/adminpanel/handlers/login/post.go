package login

import (
	"fmt"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	"time"

	cacheifaces "pnBot/internal/cache/interfaces"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
)

func LoginPost(expectedUsername, expectedPassword string,
	store adminifaces.SessionStore[*session.Session],
	cache cacheifaces.CacheProvider,
	maxAttempts int,
	blockDuration time.Duration,
) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		session, err := store.Get(context)
		if err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Ошибка при получении сессии</div>")
		}

		uid := context.Cookie("login_uid")
		if uid == "" {
			uid = uuid.New().String()
			maxAge := 86400
			context.SetCookie("login_uid", uid, maxAge)
		}

		blockedKey := "login_blocked:" + uid
		attemptsKey := "login_attempts:" + uid

		blocked, err := cache.Get(blockedKey)
		if err == nil && blocked == "1" {
			ttl, _ := cache.TTL(blockedKey)
			return context.Type("text/html").Status(200).
				SendString(fmt.Sprintf("<div class=error-box>Слишком много попыток. Попробуйте снова через %d секунд.</div>", int(ttl.Seconds())))
		}

		username := context.FormValue("username")
		password := context.FormValue("password")

		if username == expectedUsername && password == expectedPassword {
			err = cache.Del(attemptsKey)
			if err != nil {
				return context.Type("text/html").Status(200).
					SendString("<div class=error-box>Внутренняя ошибка сервера при попытке войти</div>")
			}

			session.Set("authenticated", true)
			session.Save()

			context.SetHeader("HX-Redirect", "/main")
			context.Status(200)
			return nil
		}

		attempts, err := cache.Incr(attemptsKey)
		if err != nil {
			return context.Type("text/html").Status(200).
				SendString("<div class=error-box>Внутренняя ошибка сервера при попытке войти</div>")
		}
		if attempts == 1 {
			cache.Expire(attemptsKey, blockDuration)
		}

		if attempts >= int64(maxAttempts) {
			err := cache.Set(blockedKey, "1", blockDuration)
			if err != nil {
				return context.Type("text/html").Status(200).
					SendString("<div class=error-box>Внутренняя ошибка сервера при попытке войти</div>")
			}

			err = cache.Set(blockedKey, "1", blockDuration)
			if err != nil {
				return context.Type("text/html").Status(200).
					SendString("<div class=error-box>Внутренняя ошибка сервера при попытке войти</div>")
			}

			err = cache.Del(attemptsKey)
			if err != nil {
				return context.Type("text/html").Status(200).
					SendString("<div class=error-box>Внутренняя ошибка сервера при попытке войти</div>")
			}

			return context.Type("text/html").Status(200).
				SendString(fmt.Sprintf("<div class=error-box>Превышено количество попыток. Заблокировано на %d секунд.</div>", int(blockDuration.Seconds())))
		}

		return context.Type("text/html").Status(200).
			SendString(fmt.Sprintf("<div class=error-box>Неверный логин или пароль. Осталось попыток: %d</div>", maxAttempts-int(attempts)))
	}
}

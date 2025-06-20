package login

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"

	"github.com/gofiber/fiber/v2/middleware/session"
)

func LoginPost(expectedUsername string, expectedPassword string, store adminifaces.SessionStore[*session.Session]) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		session, err := store.Get(context)
		if err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Ошибка при получении сессии из store</div>")
		}

		username := context.FormValue("username")
		password := context.FormValue("password")

		if username == expectedUsername && password == expectedPassword {
			session.Set("authenticated", true)
			session.Save()

			context.SetHeader("HX-Redirect", "/main")
			context.Status(200)
			return nil
		}

		return context.Type("text/html").Status(200).SendString("<div class=error-box>Неверный пароль или имя пользователя</div>")
	}
}

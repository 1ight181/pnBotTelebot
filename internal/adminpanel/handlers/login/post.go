package login

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"
)

func LoginPost(expectedUsername string, expectedPassword string, session adminifaces.Session) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		username := context.FormValue("username")
		password := context.FormValue("password")

		if username == expectedUsername && password == expectedPassword {
			session.Set("authenticated", true)
			session.Save()

			return context.Redirect("/create")
		}
		return context.SendString(401, "Неверный пароль или имя пользователя")
	}
}

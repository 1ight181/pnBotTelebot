package logout

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"

	"github.com/gofiber/fiber/v2/middleware/session"
)

func LogoutGet(store adminifaces.SessionStore[*session.Session]) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		session, err := store.Get(context)
		if err != nil {
			context.Status(500).SendString("Ошибка при получении сессии из store")
		}

		session.Destroy()
		return context.Redirect("/login")
	}
}

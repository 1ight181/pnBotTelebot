package middleware

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"

	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthMiddleware(store adminifaces.SessionStore[*session.Session]) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		sess, err := store.Get(context)
		if err != nil {
			return err
		}
		if auth, ok := sess.Get("authenticated").(bool); !ok || !auth {
			return context.Redirect("/login", 302)
		}
		return context.Next()
	}
}

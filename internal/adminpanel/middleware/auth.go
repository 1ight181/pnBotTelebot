package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthMiddleware(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}
		if auth, ok := sess.Get("authenticated").(bool); !ok || !auth {
			return c.Redirect("/login")
		}
		return c.Next()
	}
}

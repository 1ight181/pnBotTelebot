package logout

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"
)

func LogoutGet(session adminifaces.Session) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		session.Destroy()
		return context.Redirect("/login")
	}
}

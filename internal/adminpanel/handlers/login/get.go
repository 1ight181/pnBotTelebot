package login

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"
)

func LoginGet(c adminifaces.Context) error {
	return c.Render(200, "login", map[string]any{})
}

package create

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"
)

func CreateGet(context adminifaces.Context) error {
	return context.Render(200, "createform", map[string]any{})
}

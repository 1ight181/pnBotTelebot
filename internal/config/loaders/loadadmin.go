package loaders

import (
	conf "pnBot/internal/config/models"
)

func LoadAdminPanelConfig(adminPanelConfig conf.AdminPanel) (string, string, string, string, string) {
	username := adminPanelConfig.Username
	password := adminPanelConfig.Password
	templatesExtension := adminPanelConfig.TemplatesExtension
	host := adminPanelConfig.Host
	port := adminPanelConfig.Port

	return username, password, templatesExtension, host, port
}

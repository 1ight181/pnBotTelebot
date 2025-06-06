package loaders

import (
	conf "pnBot/internal/config/models"
)

func LoadAdminPanelConfig(adminPanelConfig conf.AdminPanel) (string, string, string, string, string, string) {
	username := adminPanelConfig.Username
	password := adminPanelConfig.Password
	templatesPath := adminPanelConfig.TemplatesPath
	templatesExtension := adminPanelConfig.TemplatesExtension
	host := adminPanelConfig.Host
	port := adminPanelConfig.Port

	return username, password, templatesPath, templatesExtension, host, port
}

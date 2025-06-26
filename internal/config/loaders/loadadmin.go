package loaders

import (
	conf "pnBot/internal/config/models"
	"strconv"
	"time"
)

func LoadAdminPanelConfig(adminPanelConfig conf.AdminPanel) (string, string, string, string, string, string, string, int, time.Duration) {
	username := adminPanelConfig.Username
	password := adminPanelConfig.Password
	templatesExtension := adminPanelConfig.TemplatesExtension
	host := adminPanelConfig.Host
	port := adminPanelConfig.Port
	staticRoot := adminPanelConfig.StaticRoot
	staticUrl := adminPanelConfig.StaticUrl
	maxLogginAttempts, _ := strconv.Atoi(adminPanelConfig.MaxLogginAttempts)
	logginBlockDuration, _ := strconv.Atoi(adminPanelConfig.LogginBlockDuration)
	logginBlockDurationSecond := time.Duration(logginBlockDuration) * time.Second

	return username, password, templatesExtension, host, port, staticRoot, staticUrl, maxLogginAttempts, logginBlockDurationSecond
}

package interfaces

import "pnBot/internal/config/models"

type ConfigProvider interface {
	Load(path, filename, configType string) (*models.Config, error)
}

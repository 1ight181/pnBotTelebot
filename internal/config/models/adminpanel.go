package models

import (
	"errors"
)

type AdminPanel struct {
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	TemplatesExtension string `mapstructure:"templates_extension"`
	Port               string `mapstructure:"port"`
	Host               string `mapstructure:"host"`
}

func (ap *AdminPanel) Validate() error {
	if ap.Username == "" {
		return errors.New("требуется указание имени админа")
	}
	if ap.Password == "" {
		return errors.New("требуется указание пароля админа")
	}
	if ap.TemplatesExtension == "" {
		return errors.New("требуется указание расширения для шаблонов")
	}
	if ap.Port == "" {
		return errors.New("требуется указание порта")
	}
	if ap.Host == "" {
		return errors.New("требуется указание хоста")
	}
	return nil
}

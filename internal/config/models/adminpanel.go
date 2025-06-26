package models

import (
	"errors"
	"strconv"
)

type AdminPanel struct {
	Username            string `mapstructure:"username"`
	Password            string `mapstructure:"password"`
	TemplatesExtension  string `mapstructure:"templates_extension"`
	Port                string `mapstructure:"port"`
	Host                string `mapstructure:"host"`
	StaticRoot          string `mapstructure:"static_root"`
	StaticUrl           string `mapstructure:"static_url"`
	MaxLogginAttempts   string `mapstructure:"max_loggin_attempts"`
	LogginBlockDuration string `mapstructure:"loggin_block_duration"`
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
	if ap.StaticRoot == "" {
		return errors.New("требуется указание пути к статическим данным")
	}
	if ap.StaticUrl == "" {
		return errors.New("требуется указание url к статическим данным")
	}
	if ap.MaxLogginAttempts == "" {
		if _, err := strconv.Atoi(ap.MaxLogginAttempts); err != nil {
			return errors.New("недопустимое значение для максимального количества попыток для входа")
		}
		return errors.New("требуется указание максимального количества попыток для входа")
	}
	if ap.LogginBlockDuration == "" {
		if _, err := strconv.Atoi(ap.MaxLogginAttempts); err != nil {
			return errors.New("недопустимое значение для продолжительности блокировки в секундах")
		}
		return errors.New("требуется указание продолжительности блокировки в секундах")
	}
	return nil
}

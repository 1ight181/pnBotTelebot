package models

import (
	"errors"
)

type Cache struct {
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func (c *Cache) Validate() error {
	if c.Port == "" {
		return errors.New("требуется указание порта для провайдера кэша")
	}
	if c.Host == "" {
		return errors.New("требуется указание хоста для провайдера кэша")
	}
	if c.Username == "" {
		return errors.New("требуется указание имени пользователя для провайдера кэша")
	}
	if c.Password == "" {
		return errors.New("требуется указание пароля для провайдера кэша")
	}

	return nil
}

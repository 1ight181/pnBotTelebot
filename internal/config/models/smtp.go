package models

import "errors"

type Smtp struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	From     string `mapstructure:"from"`
	Password string `mapstructure:"password"`
	To       string `mapstructure:"to"`
}

func (s *Smtp) Validate() error {
	if s.Host == "" {
		return errors.New("требуется указание почтового хоста")
	}
	if s.Port == "" {
		return errors.New("требуется указание почтового порта")
	}
	if s.From == "" {
		return errors.New("требуется указание требуется указание почтового адреса отправителя")
	}
	if s.Password == "" {
		return errors.New("требуется указание пароля почтового адреса отправителя")
	}
	if s.To == "" {
		return errors.New("требуется указание почтового адреса получателя")
	}
	return nil
}

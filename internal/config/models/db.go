package models

import (
	"errors"
)

type DataBase struct {
	Dsn            string `mapstructure:"dsn"`
	MigrationsPath string `mapstructure:"migrationsPath"`
}

func (b *DataBase) Validate() error {
	if b.Dsn == "" {
		return errors.New("требуется указание строки подключения")
	}
	if b.MigrationsPath == "" {
		return errors.New("требуется указание места хранения миграций")
	}

	return nil
}

package loaders

import (
	conf "pnBot/internal/config/models"
)

func LoadDbConfig(dbConfig conf.DataBase) (string, string) {
	dsn := dbConfig.Dsn
	migrationsPath := dbConfig.MigrationsPath

	return dsn, migrationsPath
}

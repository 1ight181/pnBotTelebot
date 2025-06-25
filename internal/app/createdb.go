package app

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	loaders "pnBot/internal/config/loaders"
	models "pnBot/internal/config/models"
	dbdao "pnBot/internal/db/dao"
	dbifaces "pnBot/internal/db/interfaces"
	migrationmanager "pnBot/internal/db/migrations/managers"
	gormprov "pnBot/internal/db/providers/gorm"
	gormcallbacks "pnBot/internal/db/providers/gorm/callbacks"
	loggerifaces "pnBot/internal/logger/interfaces"
)

func createDataBase(
	context context.Context,
	dbConfig models.DataBase,
	logger loggerifaces.Logger,
) (
	dbifaces.DataBaseProvider, dbifaces.OfferDao,
) {
	dsn, migrationsPath := loaders.LoadDbConfig(dbConfig)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		logger.Fatalf("Ошибка получения sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		logger.Fatalf("Не удалось проверить подключение: %v", err)
	}

	callbacks := []dbifaces.CallbackRegistrar[*gormprov.GormDataBaseProvider]{
		gormcallbacks.New(logger),
	}

	dbProviderOptions := gormprov.GormDataBaseProviderOptions{
		DataBase:  gormDB,
		Callbacks: callbacks,
	}

	dbProvider := gormprov.New(dbProviderOptions)

	offerDao := dbdao.NewOfferDao(gormDB)

	migrationManager, err := migrationmanager.New(
		migrationsPath,
		dsn,
	)
	if err != nil {
		logger.Fatalf("Ошибка при инициализации менеджера миграций: %v", err)
	}

	if err := migrationManager.Up(); err != nil {
		logger.Fatalf("Ошибка при миграции: %v", err)
	}

	maskedDsn := maskPasswordInDSN(dsn)

	logger.Infof("БД подключена по DSN: %s", maskedDsn)

	go func() {
		<-context.Done()
		if err := dbProvider.CloseConnection(); err != nil {
			logger.Fatalf("Ошибка при закрытии подключения: %v", err)
		} else {
			logger.Info("Пул подключений к БД успешно закрыт")
		}
	}()

	return dbProvider, offerDao
}

package migrations

import (
	ifaces "pnBot/internal/db/interfaces"

	"github.com/golang-migrate/migrate/v4"
)

type MigrateManager struct {
	migrate *migrate.Migrate
}

func New(migrationsPath, databaseURL string) (ifaces.MigrationManager, error) {
	migrate, err := migrate.New(migrationsPath, databaseURL)
	if err != nil {
		return nil, err
	}
	return &MigrateManager{
		migrate: migrate,
	}, nil
}

func (mm *MigrateManager) Up() error {
	err := mm.migrate.Up()
	if err == migrate.ErrNoChange {
		return nil
	}
	return err
}

func (mm *MigrateManager) Down() error {
	return mm.migrate.Down()
}

func (mm *MigrateManager) Steps(n int) error {
	return mm.migrate.Steps(n)
}

func (mm *MigrateManager) Force(version int) error {
	return mm.migrate.Force(version)
}

func (mm *MigrateManager) Version() (uint, bool, error) {
	return mm.migrate.Version()
}

func (mm *MigrateManager) Close() error {
	sourceErr, dbErr := mm.migrate.Close()
	if sourceErr != nil {
		return sourceErr
	}
	return dbErr
}

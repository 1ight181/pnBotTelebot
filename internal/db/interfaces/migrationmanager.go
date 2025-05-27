package interfaces

type MigrationManager interface {
	Up() error
	Down() error
	Steps(n int) error
	Force(version int) error
	Version() (uint, bool, error)
	Close() error
}

package interfaces

import (
	"context"

	"gorm.io/gorm"
)

type DataBaseProvider interface {
	Find(ctx context.Context, out any, where ...any) error
	First(ctx context.Context, out any, where ...any) error
	Create(ctx context.Context, value any) error
	Save(ctx context.Context, value any) error
	Delete(ctx context.Context, value any, where ...any) error
	Exec(ctx context.Context, sql string, values ...any) error
	WithTransaction(tx *gorm.DB) DataBaseProvider
	RunInTransaction(ctx context.Context, fn func(tx DataBaseProvider) error) error
}

package interfaces

import (
	ctx "context"

	"gorm.io/gorm"
)

type DataBaseProvider interface {
	Find(
		context ctx.Context,
		out any,
		where ...any,
	) error

	Count(
		context ctx.Context,
		source any,
		out *int64,
	) error

	First(
		context ctx.Context,
		out any,
		where ...any,
	) error

	FirstOrCreate(
		context ctx.Context,
		out any,
		where any,
		defaults any,
	) (bool, error)

	Create(
		context ctx.Context,
		value any,
	) error

	AddAssociation(
		context ctx.Context,
		source any,
		associationName string,
		values ...any,
	) error

	ReplaceAssociation(
		context ctx.Context,
		source any,
		associationName string,
		values ...any,
	) error

	GetAssociation(
		context ctx.Context,
		source any,
		associationName string,
		out any,
	) error

	Update(
		context ctx.Context,
		where any,
		column string,
		value any,
	) error

	Updates(
		context ctx.Context,
		where any,
		values any,
	) error

	Save(
		context ctx.Context,
		value any,
	) error

	Delete(
		context ctx.Context,
		value any,
		where ...any,
	) error

	Exec(
		context ctx.Context,
		sql string,
		values ...any,
	) error

	WithTransaction(
		tx *gorm.DB,
	) DataBaseProvider

	RunInTransaction(
		context ctx.Context,
		fn func(tx DataBaseProvider) error,
	) error

	CloseConnection() error
}

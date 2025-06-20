package providers

import (
	ctx "context"
	"fmt"

	dbifaces "pnBot/internal/db/interfaces"

	"gorm.io/gorm"
)

type GormDataBaseProvider struct {
	dataBase  *gorm.DB
	callbacks []dbifaces.CallbackRegistrar[*GormDataBaseProvider]
}

type GormDataBaseProviderOptions struct {
	DataBase  *gorm.DB
	Callbacks []dbifaces.CallbackRegistrar[*GormDataBaseProvider]
}

func New(
	opts GormDataBaseProviderOptions,
) dbifaces.DataBaseProvider {
	dbProvider := &GormDataBaseProvider{
		dataBase:  opts.DataBase,
		callbacks: opts.Callbacks,
	}
	dbProvider.attachCallbacks()
	return dbProvider
}

func (gdbp *GormDataBaseProvider) Find(
	context ctx.Context,
	out any,
	where ...any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Find(out, where...).
		Error
}

func (gdbp *GormDataBaseProvider) First(
	context ctx.Context,
	out any,
	where ...any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		First(out, where...).
		Error
}

func (gdbp *GormDataBaseProvider) FirstOrCreate(
	context ctx.Context,
	out any,
	where any,
	defaults any,
) (bool, error) {
	result := gdbp.dataBase.
		WithContext(context).
		Where(where).
		FirstOrCreate(out, defaults)

	return result.RowsAffected == 1,
		result.Error
}

func (gdbp *GormDataBaseProvider) Create(
	context ctx.Context,
	value any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Create(value).
		Error
}

func (gdbp *GormDataBaseProvider) AddAssociation(
	context ctx.Context,
	source any,
	associationName string,
	values ...any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Model(source).
		Association(associationName).
		Append(values...)
}

func (gdbp *GormDataBaseProvider) ReplaceAssociation(
	context ctx.Context,
	source any,
	associationName string,
	values ...any,
) error {
	return gdbp.dataBase.WithContext(context).
		Model(source).
		Association(associationName).
		Replace(values...)
}

func (gdbp *GormDataBaseProvider) GetAssociation(
	context ctx.Context,
	source any,
	associationName string,
	out any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Model(source).
		Association(associationName).
		Find(out)
}

func (gdbp *GormDataBaseProvider) Update(
	context ctx.Context,
	where any,
	column string,
	value any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Model(where).
		Where(where).
		Update(column, value).
		Error
}

func (gdbp *GormDataBaseProvider) Count(
	context ctx.Context,
	source any,
	out *int64,
) error {
	return gdbp.dataBase.Model(source).Count(out).Error
}

func (gdbp *GormDataBaseProvider) Updates(
	context ctx.Context,
	where any,
	values any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Where(where).
		Updates(values).
		Error
}

func (gdbp *GormDataBaseProvider) Save(
	context ctx.Context,
	value any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Save(value).
		Error
}

func (gdbp *GormDataBaseProvider) Delete(
	context ctx.Context,
	value any,
	where ...any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Delete(value, where...).
		Error
}

func (gdbp *GormDataBaseProvider) Exec(
	context ctx.Context,
	sql string,
	values ...any,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Exec(sql, values...).
		Error
}

func (gdbp *GormDataBaseProvider) WithTransaction(
	transaction *gorm.DB,
) dbifaces.DataBaseProvider {
	return &GormDataBaseProvider{
		dataBase: transaction,
	}
}

func (gdbp *GormDataBaseProvider) RunInTransaction(
	context ctx.Context,
	transactionFunction func(transaction dbifaces.DataBaseProvider) error,
) error {
	return gdbp.dataBase.
		WithContext(context).
		Transaction(func(transaction *gorm.DB) error {
			txProvider := &GormDataBaseProvider{
				dataBase: transaction,
			}
			return transactionFunction(txProvider)
		})
}

func (gdbp *GormDataBaseProvider) attachCallbacks() {
	for _, callback := range gdbp.callbacks {
		callback.RegisterCallback(gdbp)
	}
}

func (gdbp *GormDataBaseProvider) GetRawDb() *gorm.DB {
	return gdbp.dataBase
}

func (gdbp *GormDataBaseProvider) CloseConnection() error {
	db, err := gdbp.dataBase.DB()
	if err != nil {
		return fmt.Errorf(
			"не удалось закрыть пул соединений с БД: %w",
			err,
		)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf(
			"не удалось закрыть пул соединений с БД: %w",
			err,
		)
	}
	return nil
}

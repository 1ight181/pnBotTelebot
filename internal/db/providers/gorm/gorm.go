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

func New(opts GormDataBaseProviderOptions) dbifaces.DataBaseProvider {
	dbProvider := &GormDataBaseProvider{
		dataBase:  opts.DataBase,
		callbacks: opts.Callbacks,
	}
	dbProvider.attachCallbacks()
	return dbProvider
}

func (g *GormDataBaseProvider) Find(context ctx.Context, out any, where ...any) error {
	return g.dataBase.WithContext(context).Find(out, where...).Error
}

func (g *GormDataBaseProvider) First(context ctx.Context, out any, where ...any) error {
	return g.dataBase.WithContext(context).First(out, where...).Error
}

func (g *GormDataBaseProvider) FirstOrCreate(context ctx.Context, out any, where any, defaults any) (bool, error) {
	result := g.dataBase.WithContext(context).Where(where).FirstOrCreate(out, defaults)
	return result.RowsAffected == 1, result.Error
}

func (g *GormDataBaseProvider) Create(context ctx.Context, value any) error {
	return g.dataBase.WithContext(context).Create(value).Error
}

func (g *GormDataBaseProvider) Update(context ctx.Context, where any, column string, value any) error {
	return g.dataBase.WithContext(context).Model(where).Where(where).Update(column, value).Error
}

func (g *GormDataBaseProvider) Updates(context ctx.Context, where any, values any) error {
	return g.dataBase.WithContext(context).Where(where).Updates(values).Error
}

func (g *GormDataBaseProvider) Save(context ctx.Context, value any) error {
	return g.dataBase.WithContext(context).Save(value).Error
}

func (g *GormDataBaseProvider) Delete(context ctx.Context, value any, where ...any) error {
	return g.dataBase.WithContext(context).Delete(value, where...).Error
}

func (g *GormDataBaseProvider) Exec(context ctx.Context, sql string, values ...any) error {
	return g.dataBase.WithContext(context).Exec(sql, values...).Error
}

func (g *GormDataBaseProvider) WithTransaction(transaction *gorm.DB) dbifaces.DataBaseProvider {
	return &GormDataBaseProvider{dataBase: transaction}
}

func (g *GormDataBaseProvider) RunInTransaction(context ctx.Context, transactionFunction func(transaction dbifaces.DataBaseProvider) error) error {
	return g.dataBase.WithContext(context).Transaction(func(transaction *gorm.DB) error {
		txProvider := &GormDataBaseProvider{
			dataBase: transaction,
		}
		return transactionFunction(txProvider)
	})
}

func (g *GormDataBaseProvider) attachCallbacks() {
	for _, callback := range g.callbacks {
		callback.RegisterCallback(g)
	}
}

func (g *GormDataBaseProvider) GetRawDb() *gorm.DB {
	return g.dataBase
}

func (g *GormDataBaseProvider) CloseConnection() error {
	db, err := g.dataBase.DB()
	if err != nil {
		return fmt.Errorf("не удалось закрыть пул соединений с БД: %w", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("не удалось закрыть пул соединений с БД: %w", err)
	}
	return nil
}

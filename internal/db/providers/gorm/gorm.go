package providers

import (
	"context"

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

func (g *GormDataBaseProvider) Find(ctx context.Context, out any, where ...any) error {
	return g.dataBase.WithContext(ctx).Find(out, where...).Error
}

func (g *GormDataBaseProvider) First(ctx context.Context, out any, where ...any) error {
	return g.dataBase.WithContext(ctx).First(out, where...).Error
}

func (g *GormDataBaseProvider) Create(ctx context.Context, value any) error {
	return g.dataBase.WithContext(ctx).Create(value).Error
}

func (g *GormDataBaseProvider) Save(ctx context.Context, value any) error {
	return g.dataBase.WithContext(ctx).Save(value).Error
}

func (g *GormDataBaseProvider) Delete(ctx context.Context, value any, where ...any) error {
	return g.dataBase.WithContext(ctx).Delete(value, where...).Error
}

func (g *GormDataBaseProvider) Exec(ctx context.Context, sql string, values ...any) error {
	return g.dataBase.WithContext(ctx).Exec(sql, values...).Error
}

func (g *GormDataBaseProvider) WithTransaction(transaction *gorm.DB) dbifaces.DataBaseProvider {
	return &GormDataBaseProvider{dataBase: transaction}
}

func (g *GormDataBaseProvider) RunInTransaction(ctx context.Context, transactionFunction func(transaction dbifaces.DataBaseProvider) error) error {
	return g.dataBase.WithContext(ctx).Transaction(func(transaction *gorm.DB) error {
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

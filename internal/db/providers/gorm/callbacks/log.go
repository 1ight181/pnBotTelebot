package callbacks

import (
	dbifaces "pnBot/internal/db/interfaces"
	gormprov "pnBot/internal/db/providers/gorm"
	loggeriface "pnBot/internal/logger/interfaces"

	"gorm.io/gorm"
)

type GormLogCallbackRegistrar struct {
	logger loggeriface.Logger
}

func New(logger loggeriface.Logger) dbifaces.CallbackRegistrar[*gormprov.GormDataBaseProvider] {
	return &GormLogCallbackRegistrar{
		logger: logger,
	}
}

func (glcr *GormLogCallbackRegistrar) RegisterCallback(gormProvider *gormprov.GormDataBaseProvider) {
	rawDb := gormProvider.GetRawDb()

	rawDb.Callback().Create().After("gorm:create").Register("log:After_create", func(tx *gorm.DB) {
		glcr.logger.Infof("[After Create] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Create().After("gorm:create").Register("log:after_create", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Create] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Query().After("gorm:query").Register("log:After_query", func(tx *gorm.DB) {
		glcr.logger.Infof("[After Select] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Query().After("gorm:query").Register("log:after_query", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Select] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Update().After("gorm:update").Register("log:After_update", func(tx *gorm.DB) {
		glcr.logger.Infof("[After Update] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Update().After("gorm:update").Register("log:after_update", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Update] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Delete().After("gorm:delete").Register("log:After_delete", func(tx *gorm.DB) {
		glcr.logger.Infof("[After Delete] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Delete().After("gorm:delete").Register("log:after_delete", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Delete] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Row().After("gorm:row").Register("log:After_row", func(tx *gorm.DB) {
		glcr.logger.Infof("[After Row] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Row().After("gorm:row").Register("log:after_row", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Row] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Raw().After("gorm:raw").Register("log:After_raw", func(tx *gorm.DB) {
		glcr.logger.Infof("[After Raw Exec] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Raw().After("gorm:raw").Register("log:after_raw", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Raw Exec] error: %v", tx.Error)
		}
	})
}

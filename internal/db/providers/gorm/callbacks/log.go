package callbacks

import (
	gormprov "pnBot/internal/db/providers/gorm"
	loggeriface "pnBot/internal/logger/interfaces"

	"gorm.io/gorm"
)

type GormLogCallbackRegistrar struct {
	logger loggeriface.Logger
}

func (glcr *GormLogCallbackRegistrar) RegisterCallback(gormProvider gormprov.GormDataBaseProvider) {
	rawDb := gormProvider.GetRawDb()

	rawDb.Callback().Create().Before("gorm:create").Register("log:before_create", func(tx *gorm.DB) {
		glcr.logger.Infof("[Before Create] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Create().After("gorm:create").Register("log:after_create", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Create] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Query().Before("gorm:query").Register("log:before_query", func(tx *gorm.DB) {
		glcr.logger.Infof("[Before Select] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Query().After("gorm:query").Register("log:after_query", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Select] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Update().Before("gorm:update").Register("log:before_update", func(tx *gorm.DB) {
		glcr.logger.Infof("[Before Update] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Update().After("gorm:update").Register("log:after_update", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Update] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Delete().Before("gorm:delete").Register("log:before_delete", func(tx *gorm.DB) {
		glcr.logger.Infof("[Before Delete] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Delete().After("gorm:delete").Register("log:after_delete", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Delete] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Row().Before("gorm:row").Register("log:before_row", func(tx *gorm.DB) {
		glcr.logger.Infof("[Before Row] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Row().After("gorm:row").Register("log:after_row", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Row] error: %v", tx.Error)
		}
	})

	rawDb.Callback().Raw().Before("gorm:raw").Register("log:before_raw", func(tx *gorm.DB) {
		glcr.logger.Infof("[Before Raw Exec] SQL: %s, Vars: %v", tx.Statement.SQL.String(), tx.Statement.Vars)
	})
	rawDb.Callback().Raw().After("gorm:raw").Register("log:after_raw", func(tx *gorm.DB) {
		if tx.Error != nil {
			glcr.logger.Errorf("[After Raw Exec] error: %v", tx.Error)
		}
	})
}

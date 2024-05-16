package initialize

import (
	"os"

	"github.com/lonevle/gin-server/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GS_CONFIG.System.DbType {
	case "sqlite":
		return GormSqlite()
	// case "mysql":
	// 	return GormMysql()
	// case "pgsql":
	// 	return GormPgSql()
	// case "oracle":
	// 	return GormOracle()
	// case "mssql":
	// 	return GormMssql()
	default:
		return GormSqlite()
	}
}

func RegisterTables() {
	db := global.GS_DB
	err := db.AutoMigrate()
	if err != nil {
		zap.L().Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	zap.L().Info("register table success")
}

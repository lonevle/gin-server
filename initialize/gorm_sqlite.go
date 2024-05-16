package initialize

import (
	"github.com/lonevle/gin-server/global"
	"github.com/lonevle/gin-server/initialize/internal"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GormSqlite 初始化Sqlite数据库
func GormSqlite() *gorm.DB {
	s := global.GS_CONFIG.Sqlite

	// 如果没配置数据库名称，不初始化
	if s.Dbname == "" {
		zap.L().Warn("未配置sqlite数据库")
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular, s.GetLogMode())); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}

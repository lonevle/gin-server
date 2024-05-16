package internal

import (
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Config(prefix string, singular bool, logMode string) *gorm.Config {
	config := &gorm.Config{
		// 覆盖gorm默认配置
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,   // table name prefix, table for `User` would be `t_users`
			SingularTable: singular, // use singular table name, table for `User` would be `user` with this option enabled
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
	}
	// 自定义日志
	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond, // 慢查询阈值
		LogLevel:      logger.Warn,            // 日志级别
		Colorful:      true,                   // 禁用彩色打印
	})

	logMode = strings.ToLower(logMode)
	switch logMode {
	case "silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		// 默认使用info级别
		config.Logger = _default.LogMode(logger.Info)
	}

	return config
}

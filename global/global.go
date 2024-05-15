package global

import (
	"github.com/lonevle/gin-server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GS_ROOT   string        // 项目根目录
	GS_DB     *gorm.DB      // gorm数据库连接
	GS_VP     *viper.Viper  // 配置管理器
	GS_LOG    *zap.Logger   // 日志
	GS_CONFIG config.Server // 程序配置
)

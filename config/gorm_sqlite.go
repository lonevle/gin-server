package config

import (
	"github.com/lonevle/gin-server/utils"
)

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"` // yaml:",inline" 引入通用配置
}

// 数据源地址
func (s *Sqlite) Dsn() string {
	return utils.JoinPath(s.Path, s.Dbname+".db")
}

// 获取日志模式
func (s *Sqlite) GetLogMode() string {
	return s.LogMode
}

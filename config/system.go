package config

type System struct {
	Addr   int    `mapstructure:"addr" json:"addr" yaml:"addr"`          // 端口值
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"` // 数据库类型:sqlite(默认)|mysql|sqlserver|postgresql
}

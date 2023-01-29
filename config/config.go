package config

type Server struct {
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	// gorm 数据库
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql  Mssql  `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql  Pgsql  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle Oracle `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}

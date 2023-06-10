package config

type System struct {
	RunMode       string `mapstructure:"run-mode" json:"run-mode" yaml:"run-mode"`                   // 运行模式
	Host          string `mapstructure:"host" json:"host" yaml:"host"`                               // 监听服务器ip
	HttpPort      int    `mapstructure:"http-port" json:"http-port" yaml:"http-port"`                // 端口值
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 使用redis
	IsCors        bool   `mapstructure:"is-cors" json:"is-cors" yaml:"is-cors"`                      // 跨域
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql|mssql
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	GormLogMode   string `mapstructure:"gorm-log-mode" json:"gorm-log-mode" yaml:"gorm-log-mode"` // 是否开启Gorm全局日志
	GormLogZap    bool   `mapstructure:"gorm-log-zap" json:"gorm-log-zap" yaml:"gorm-log-zap"`    // 是否通过zap写入日志文件
}

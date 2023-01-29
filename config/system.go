package config

type System struct {
	RunMode       string `mapstructure:"run-mode" json:"run-mode" yaml:"run-mode"`                      // 运行模式
	HttpPort      int    `mapstructure:"http-port" json:"http-port" yaml:"http-port"`                   // 端口值
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                  // 使用redis
}
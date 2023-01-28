package config

type System struct {
	RunMode       string `mapstructure:"RunMode" json:"RunMode" yaml:"RunMode"`                      // 运行模式
	HttpPort      int    `mapstructure:"HttpPort" json:"HttpPort" yaml:"HttpPort"`                   // 端口值
	UseRedis      bool   `mapstructure:"UseRedis" json:"UseRedis" yaml:"UseRedis"`                  // 使用redis
}
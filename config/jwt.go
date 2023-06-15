package config

type JWT struct {
	SecretKey   string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`       // jwt签名
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
}

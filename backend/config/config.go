package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	// gorm 数据库
	Databases []GeneralDB `mapstructure:"databases" json:"databases" yaml:"databases"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
	// HTTPS 支持
	Ssl SSL `mapstructure:"ssl" json:"ssl" yaml:"ssl"`
}

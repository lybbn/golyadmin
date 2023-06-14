package config

type Captcha struct {
	StoreType      string `mapstructure:"store-type" json:"store-type" yaml:"store-type"`                // 存储位置
	CaptchaType    string `mapstructure:"captcha-type" json:"captcha-type" yaml:"captcha-type"`          // 验证码类型
	KeyLength      int    `mapstructure:"key-length" json:"key-length" yaml:"key-length"`                // 验证码长度
	ImgWidth       int    `mapstructure:"img-width" json:"img-width" yaml:"img-width"`                   // 验证码宽度
	ImgHeight      int    `mapstructure:"img-height" json:"img-height" yaml:"img-height"`                // 验证码高度
	CaptchaTimeout int    `mapstructure:"captcha-timeout" json:"captcha-timeout" yaml:"captcha-timeout"` // 验证码过期时间，单位：s(秒)
}

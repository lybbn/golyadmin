package config

type SSL struct {
	Enable   bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Domain   string `mapstructure:"domain" json:"domain" yaml:"domain"`
	KeyFile  string `mapstructure:"key-file" json:"key-file" yaml:"key-file"`
	CertFile string `mapstructure:"cert-file" json:"cert-file" yaml:"cert-file"`
}

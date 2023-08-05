package config

type Config struct {
	// jwt
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`

	// zap
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`

	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// system
	System System `mapstructure:"system" json:"system" yaml:"system"`

	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}

package config

type Config struct {
	// jwt
	JWT JWT `mapStructure:"jwt" json:"jwt" yaml:"jwt"`

	// zap
	Zap Zap `mapStructure:"zap" json:"zap" yaml:"zap"`

	// auto
	AutoCode Autocode `mapStructure:"autocode" json:"autocode" yaml:"autocode"`

	// gorm
	Mysql Mysql `mapStructure:"mysql" json:"mysql" yaml:"mysql"`

	// system
	System System `mapStructure:"system" json:"system" yaml:"system"`

	Captcha Captcha `mapStructure:"captcha" json:"captcha" yaml:"captcha"`
}

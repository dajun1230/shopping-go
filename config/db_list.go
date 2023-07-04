package config

type GeneralDB struct {
	Path         string `mapStructure:"path" json:"path" yaml:"path"`                               // 服务器地址:端口
	Port         string `mapStructure:"port" json:"port" yaml:"port"`                               //:端口
	Config       string `mapStructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapStructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapStructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password     string `mapStructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Prefix       string `mapStructure:"prefix" json:"prefix" yaml:"prefix"`                         //全局表前缀，单独定义TableName则不生效
	Singular     bool   `mapStructure:"singular" json:"singular" yaml:"singular"`                   //是否开启全局禁用复数，true表示开启
	Engine       string `mapStructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        //数据库引擎，默认InnoDB
	MaxIdleConns int    `mapStructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapStructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapStructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapStructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

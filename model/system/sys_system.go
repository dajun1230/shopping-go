package system

import "shopping-go/config"

// 配置文件结构体
type System struct {
	Config config.Config `json:"config"`
}

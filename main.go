package main

import (
	"encoding/json"
	"fmt"
	"shopping-go/core"
	"shopping-go/global"
	"shopping-go/initialize"

	"go.uber.org/zap"
)

func main() {
	// 初始化Viper
	global.SHOP_VP = core.Viper()

	// 查看配置文件
	data, _ := json.Marshal(global.SHOP_CONFIG)
	fmt.Println("配置文件Config:", string(data))

	initialize.OtherInit()

	// 初始化zap
	global.SHOP_LOG = core.Zap()
	zap.ReplaceGlobals(global.SHOP_LOG) // ReplaceGlobals 方法用于替换全局的 Logger 实例和 SugaredLogger 实例

	// gorm连接数据库

	global.SHOP_DB = initialize.Gorm()
	if global.SHOP_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库连接
		db, _ := global.SHOP_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}

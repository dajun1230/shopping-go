package main

import (
	"shopping-go/core"
	"shopping-go/global"
	"shopping-go/initialize"
)

func main() {
	// gorm连接数据库
	global.SHOP_DB = initialize.Gorm()
	if global.SHOP_DB != nil {
		// 程序结束前关闭数据库连接
		db, _ := global.SHOP_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}

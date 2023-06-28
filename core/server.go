package core

import (
	"shopping-go/global"
	"shopping-go/initialize"
)

func RunWindowsServer() {
	// 初始化路由
	Router := initialize.Routers()
	s := initServer(":8080", Router) // 启动HTTP服务，默认在0.0.0.0:8080启动服务，代替原有 Router.Run()
	global.SHOP_LOG.Error(s.ListenAndServe().Error())
}

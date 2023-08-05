package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-go/global"
	"shopping-go/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System

	// 公共路由，不做鉴权处理
	PublicGroup := Router.Group(global.SHOP_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由
	}

	// 私有路由，需要做鉴权处理
	PrivateGroup := Router.Group(global.SHOP_CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}

	//Router.GET("/hello", func(c *gin.Context) {
	//	// c.JSON：返回JSON格式的数据
	//	c.JSON(200, gin.H{
	//		"message": "Hello world!",
	//	})
	//})

	return Router
}

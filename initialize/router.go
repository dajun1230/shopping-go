package initialize

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	return Router
}

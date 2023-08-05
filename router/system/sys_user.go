package system

import (
	"github.com/gin-gonic/gin"
	"shopping-go/api"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	baseApi := api.ApiRouterGroup.System.BaseApi
	userRouter := Router.Group("user")
	{
		userRouter.POST("admin_register", baseApi.Register) // 管理员注册账号
	}
}

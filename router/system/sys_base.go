package system

import (
	"github.com/gin-gonic/gin"
	"shopping-go/api"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseApi := api.ApiRouterGroup.System.BaseApi
	baseRouter := Router.Group("api")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
}

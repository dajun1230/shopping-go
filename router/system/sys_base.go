package system

import (
	"shopping-go/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseApi := api.ApiRouterGroup.System.BaseApi
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
}

package system

import (
	"github.com/gin-gonic/gin"
	"shopping-go/api"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	api := api.ApiRouterGroup.System.BaseApi
	baseRouter := Router.Group("api")
	{
		baseRouter.POST("login", api.Login)
	}
	return baseRouter
}

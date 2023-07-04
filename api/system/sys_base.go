package system

import (
	"github.com/gin-gonic/gin"
	response "shopping-go/model/common"
)

type BaseApi struct {
}

// 用户登录
func (b *BaseApi) Login(c *gin.Context) {
	response.FailWithMsg("登录失败", c)
}

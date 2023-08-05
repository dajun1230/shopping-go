package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"shopping-go/global"
	response "shopping-go/model/common"
	"shopping-go/model/system"
	systemReq "shopping-go/model/system/request"
	systemRes "shopping-go/model/system/response"
	"shopping-go/utils"
	"time"
)

/*
*
Login
@Tags SysUser
@Summary 账号登录
@Produce application/json
@Params data body systemReq
@Success 200 {}
@Router /user/admin_register [post]
*/
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	// 判断验证码是否开启
	openCaptcha := global.SHOP_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.SHOP_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.SHOP_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMsg("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.SHOP_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMsg("用户被禁止登录", c)
			return
		}
		TokenNext(c, *user)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMsg("验证码错误", c)
}

func TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.SHOP_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		global.SHOP_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMsg("获取token失败", c)
		return
	}
	//if !global.SHOP_CONFIG.System.UseMultipoint {
	//	response.OkWithDetailed(systemRes.LoginResponse{
	//		User:      user,
	//		Token:     token,
	//		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	//	}, "登录成功", c)
	//	return
	//}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.SHOP_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMsg("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.SHOP_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMsg("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMsg("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMsg("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

/*
*
SysUser
@Tags SysUser
@Summary 用户注册账号
@Produce application/json
@Params data body systemReq.Register
@Success 200 {}
@Router /user/admin_register [post]
*/
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{
		Username:    r.Username,
		NickName:    r.NickName,
		Password:    r.Password,
		HeaderImg:   r.HeaderImg,
		AuthorityId: r.AuthorityId,
		Authorities: authorities,
		Enable:      r.Enable,
		Phone:       r.Phone,
		Email:       r.Email,
	}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.SHOP_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMsg("注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

package system

import (
	"context"
	"shopping-go/global"
	"shopping-go/model/system"
	"shopping-go/utils"
)

type JwtService struct {
}

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.SHOP_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (jwtService *JwtService) GetRedisJWT(username string) (redisJWT string, err error) {
	redisJWT, err = global.SHOP_REDIS.Get(context.Background(), username).Result()
	return redisJWT, err
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.SHOP_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.SHOP_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

package system

import (
	"shopping-go/global"
	"shopping-go/model/system"
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

package system

import "shopping-go/global"

type JwtBlacklist struct {
	global.SHOP_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

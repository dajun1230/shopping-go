package system

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"shopping-go/global"
	"shopping-go/model/system"
	"shopping-go/utils"
)

type UserService struct {
}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.SHOP_DB.Where("username =?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	// 否则  附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.SHOP_DB.Create(&u).Error
	return u, err
}

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.SHOP_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user system.SysUser
	err = global.SHOP_DB.Where("username=?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		// TODO: 用戶角色的默認路由
		//MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

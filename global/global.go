package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"shopping-go/config"
)

var (
	SHOP_DB                  *gorm.DB
	SHOP_LOG                 *zap.Logger
	SHOP_VP                  *viper.Viper
	SHOP_CONFIG              config.Config
	SHOP_Concurrency_Control = &singleflight.Group{}
)

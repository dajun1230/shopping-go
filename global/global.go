package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"shopping-go/config"
)

var (
	SHOP_DB                  *gorm.DB
	SHOP_REDIS               *redis.Client
	SHOP_VP                  *viper.Viper
	SHOP_CONFIG              config.Config
	SHOP_LOG                 *zap.Logger
	SHOP_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache
)

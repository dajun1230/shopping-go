package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SHOP_DB  *gorm.DB
	SHOP_LOG *zap.Logger
)

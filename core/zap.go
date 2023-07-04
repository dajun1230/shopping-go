package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"shopping-go/core/internal"
	"shopping-go/global"
	"shopping-go/utils"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.SHOP_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create directory: %v \n", global.SHOP_CONFIG.Zap.Director)
		_ = os.Mkdir(global.SHOP_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	return logger
}

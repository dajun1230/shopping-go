package internal

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"shopping-go/global"
	"time"
)

type fileRotateLogs struct{}

var FileRotateLogs = new(fileRotateLogs)

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotateLogs.New(
		path.Join(global.SHOP_CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
		rotateLogs.WithClock(rotateLogs.Local),
		rotateLogs.WithMaxAge(time.Duration(global.SHOP_CONFIG.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotateLogs.WithRotationTime(time.Hour*24),                                        // 24小时切割一次
	)
	if global.SHOP_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

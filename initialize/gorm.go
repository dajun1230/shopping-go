package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"shopping-go/global"
	system2 "shopping-go/model/system"
)

func Gorm() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       "root:zxcv1234@tcp(47.108.167.244:3306)/shopping_go?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 解决插入时表名为复数问题
		},
	})

	if err != nil {
		panic("连接数据库失败")
	}

	fmt.Print("数据库连接成功。")
	sqlDB, _ := db.DB()

	// 设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	return db
}

// 注册数据库表专用
func RegisterTables() {
	db := global.SHOP_DB
	err := db.AutoMigrate(
		// 系统模块表
		//system.SysUser{},
		//system.SysAuthority{},
		//system.SysBaseMenu{},
		system2.SysUser{},
		//system.SysBaseMenu{},
		system2.SysAuthority{},
		//system.SysBaseMenuParameter{},
		//system.SysBaseMenuBtn{},
	)
	if err != nil {
		global.SHOP_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.SHOP_LOG.Info("register table success")
}

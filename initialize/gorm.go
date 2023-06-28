package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       "root:zxcv1234@tcp(47.108.167.244:3306)/shopping?charset=utf8&parseTime=True&loc=Local&timeout=10s",
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

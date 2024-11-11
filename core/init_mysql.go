package core

import (
	"im_server/core/config"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func InitMysql() *gorm.DB {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/im_server_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	var mysqlLogger logger.Interface
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: mysqlLogger,
// 	})
// 	if err != nil {
// 		logx.Errorf("[%s] mysql 连接失败", dsn)
// 		return nil
// 	}
// 	sqlDB, _ := db.DB()
// 	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
// 	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
// 	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
// 	return db
// }

func InitGorm(MysqlDataSource string) *gorm.DB {
	// 加载配置
	config.InitConfig()
	// 初始化日志
	config.InitLogx()
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		logx.Info("Mysql 连接成功")
	}
	return db
}

package svc

import (
	"im_server/core"
	"im_server/im_auth/auth_api/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化mysql
	mysqlDB := core.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     mysqlDB,
	}
}

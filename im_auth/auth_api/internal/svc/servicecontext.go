package svc

import (
	"im_server/core"
	"im_server/im_auth/auth_api/internal/config"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	rdb := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
		RDB:    rdb,
	}
}

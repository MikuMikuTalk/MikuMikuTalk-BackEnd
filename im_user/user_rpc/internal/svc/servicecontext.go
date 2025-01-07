package svc

import (
	"im_server/core"
	"im_server/im_user/user_rpc/internal/config"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisDb := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
		Redis:  redisDb,
	}
}

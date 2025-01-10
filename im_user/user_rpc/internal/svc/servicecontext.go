package svc

import (
	"im_server/core"
	"im_server/im_user/user_rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisDb := redis.MustNewRedis(c.Redis.RedisConf)
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
		Redis:  redisDb,
	}
}

package svc

import (
	"im_server/core"
	"im_server/im_auth/auth_api/internal/config"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/im_user/user_rpc/users"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RDB     *redis.Client
	UserRpc user_rpc.UsersClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	rdb := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)
	return &ServiceContext{
		Config:  c,
		DB:      mysqlDb,
		RDB:     rdb,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}

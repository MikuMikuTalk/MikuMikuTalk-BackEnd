package svc

import (
	"im_server/core"
	"im_server/im_group/group_api/internal/config"
	"im_server/im_group/group_rpc/groups"
	"im_server/im_group/group_rpc/types/group_rpc"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/im_user/user_rpc/users"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	Redis    *redis.Redis
	UserRpc  user_rpc.UsersClient
	GroupRpc group_rpc.GroupsClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisDb := redis.MustNewRedis(c.Redis.RedisConf)
	return &ServiceContext{
		DB:       mysqlDb,
		Redis:    redisDb,
		UserRpc:  users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		GroupRpc: groups.NewGroups(zrpc.MustNewClient(c.GroupRpc)),
		Config:   c,
	}
}

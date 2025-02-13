package svc

import (
	"im_server/common/log_stash"
	"im_server/common/zrpc_interceptor"
	"im_server/core"
	"im_server/im_auth/auth_api/internal/config"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/im_user/user_rpc/users"

	"github.com/zeromicro/go-queue/kq"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	RDB            *redis.Redis
	UserRpc        user_rpc.UsersClient
	KqPusherClient *kq.Pusher
	ActionLogs     *log_stash.Pusher
	RuntimeLogs    *log_stash.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	rdb := redis.MustNewRedis(c.Redis.RedisConf)
	kqClient := kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic)
	return &ServiceContext{
		Config:         c,
		DB:             mysqlDb,
		RDB:            rdb,
		UserRpc:        users.NewUsers(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		KqPusherClient: kqClient,
		ActionLogs:     log_stash.NewActionPusher(kqClient, c.Name),
		RuntimeLogs:    log_stash.NewRuntimePusher(kqClient, c.Name),
	}
}

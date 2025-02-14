package svc

import (
	"im_server/common/log_stash"
	"im_server/core"
	"im_server/im_log/logs_api/internal/config"
	"im_server/im_log/logs_api/internal/middleware"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/im_user/user_rpc/users"
	"net/http"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	UserRpc         user_rpc.UsersClient
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
	KqPusherClient  *kq.Pusher
	ActionLogs      *log_stash.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	kqClient := kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic)
	return &ServiceContext{
		Config:          c,
		DB:              mysqlDb,
		UserRpc:         users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
		KqPusherClient:  kqClient,
		ActionLogs:      log_stash.NewActionPusher(kqClient, c.Name),
	}
}

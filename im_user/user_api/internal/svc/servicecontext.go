package svc

import (
	"im_server/core"
	"im_server/im_chat/chat_rpc/chat"
	"im_server/im_chat/chat_rpc/types/chat_rpc"
	"im_server/im_user/user_api/internal/config"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/im_user/user_rpc/users"

	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user_rpc.UsersClient
	ChatRpc chat_rpc.ChatClient
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		ChatRpc: chat.NewChat(zrpc.MustNewClient(c.ChatRpc)),
		DB:      mysqlDb,
	}
}

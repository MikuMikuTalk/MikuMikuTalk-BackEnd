package svc

import (
	"im_server/core"
	"im_server/im_chat/chat_api/internal/config"
	"im_server/im_file/file_rpc/files"
	"im_server/im_file/file_rpc/types/file_rpc"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/im_user/user_rpc/users"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	UserRpc user_rpc.UsersClient
	FileRpc file_rpc.FilesClient
	Redis   *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisClient := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config:  c,
		DB:      mysqlDb,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		FileRpc: files.NewFiles(zrpc.MustNewClient(c.FileRpc)),
		Redis:   redisClient,
	}
}

package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Etcd  string
	Mysql struct {
		DataSource string
	}
	Auth struct {
		AuthSecret string
		AuthExpire int64
	}
	Redis   redis.RedisKeyConf
	UserRpc zrpc.RpcClientConf
	FileRpc zrpc.RpcClientConf
}

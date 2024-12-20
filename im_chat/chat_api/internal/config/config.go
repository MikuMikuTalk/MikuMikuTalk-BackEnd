package config

import (
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
	UserRpc zrpc.RpcClientConf
}

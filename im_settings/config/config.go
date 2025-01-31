package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	Etcd  string
	Mysql struct {
		DataSource string
	}
	Auth struct {
		AuthSecret string
		AuthExpire int64
	}
	Redis redis.RedisKeyConf
	Log   logx.LogConf
}

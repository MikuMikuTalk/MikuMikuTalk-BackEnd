package core

import (
	"im_server/core/config"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
)

func InitRedis(addr, pwd string, db int) (client *redis.Client) {
	logx.MustSetup(config.GetConfig())
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	err := rdb.Ping().Err()
	if err != nil {
		logx.Error("redis connect failed ", "error:", err)
		panic(err)
	}
	logx.Info("Redis连接成功")
	return rdb
}

package core

import (
	"log/slog"

	"github.com/go-redis/redis"
)

func InitRedis(addr, pwd string, db int) (client *redis.Client) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	err := rdb.Ping().Err()
	if err != nil {
		slog.Error("redis connect failed ", "error:", err)
		panic(err)
	}
	slog.Info("Redis连接成功")
	return rdb
}

package core

import (
	"im_server/core/config"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcd(addr string) *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logx.MustSetup(config.GetConfig())
		logx.Error(err)
		panic(err)
	}
	return client
}

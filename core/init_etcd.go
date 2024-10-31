package core

import (
	"time"

	"im_server/utils/logs"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcd(addr string) *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	return client
}

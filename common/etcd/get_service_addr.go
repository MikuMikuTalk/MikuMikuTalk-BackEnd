package etcd

import (
	"context"

	"im_server/core"
)

func GetServiceAddr(etcdAddr string, serviceName string) (addr string) {
	client := core.InitEtcd(etcdAddr)
	res, err := client.Get(context.Background(), serviceName)
	if err == nil && len(res.Kvs) > 0 {
		return string(res.Kvs[0].Value)
	}
	return ""
}

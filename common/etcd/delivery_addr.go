package etcd

import (
	"context"
	"strings"

	"im_server/core"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
)

func DeliveryAddress(etcdAddr string, serviceName string, addr string) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		logx.Errorf("地址错误 :%s", addr)
		return
	}
	if list[0] == "0.0.0.0" {
		ip := netx.InternalIp()
		addr = strings.ReplaceAll(addr, "0.0.0.0", ip)
	}
	client := core.InitEtcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, addr)
	if err != nil {
		logx.Errorf("地址上送失败:%s", err)
		return
	}
	logx.Infof("地址上送成功 %s %s ", serviceName, addr)
}

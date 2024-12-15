package modules

import (
	"fmt"
	"im_server/common/etcd"
	"regexp"

	"github.com/zeromicro/go-zero/core/logx"
)

// 从路径中提取服务名称
func (g *GatewayService) extractService(path string) (string, error) {
	// 使用正则表达式提取 /api/{service}/ 中的 service 名称
	regex, _ := regexp.Compile(`/api/([^/]+)/`)
	matches := regex.FindStringSubmatch(path)
	if len(matches) != 2 {
		return "", fmt.Errorf("invalid service path")
	}
	return matches[1], nil
}

// 从 Etcd 获取服务地址
func (g *GatewayService) getServiceAddr(service string) string {
	addr := etcd.GetServiceAddr(g.config.Etcd, service+"_api")
	if addr == "" {
		logx.Error("不匹配的服务", service)
	}
	return addr
}

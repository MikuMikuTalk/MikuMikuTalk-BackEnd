package main

import (
	"flag"
	"fmt"
	"net/http"

	"im_server/common/etcd"
	"im_server/common/middleware"
	"im_server/im_auth/auth_api/internal/config"
	"im_server/im_auth/auth_api/internal/handler"
	"im_server/im_auth/auth_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		header.Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
	}, nil, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 使用中间件
	server.Use(middleware.LogActionMiddleware(ctx.ActionLogs))
	etcd.DeliveryAddress(c.Etcd, c.Name+"_api", fmt.Sprintf("%s:%d", c.Host, c.Port))
	logx.Infof("im_auth服务 正在监听 %s:%d...\n", c.Host, c.Port)
	server.Start()
}

package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"im_server/im_auth/auth_api/internal/config"
	"im_server/im_auth/auth_api/internal/handler"
	"im_server/im_auth/auth_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
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
	slog.Info(fmt.Sprintf("im_auth服务 正在监听 %s:%d...\n", c.Host, c.Port))
	// fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

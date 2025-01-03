package main

import (
	"flag"
	"fmt"
	"net/http"

	"im_server/im_gateway/modules"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

// 配置文件路径的命令行参数
var configFile = flag.String("f", "settings.yaml", "the config file")

func main() {
	flag.Parse()

	// 加载配置
	var config modules.Config
	conf.MustLoad(*configFile, &config)
	logx.MustSetup(config.Log)

	// 初始化网关服务
	gateway := modules.NewGatewayService(config)

	// 设置路由
	http.HandleFunc("/", gateway.HandleRequest)

	// 启动网关服务
	fmt.Printf("网关运行在 %s\n", config.Addr)
	http.ListenAndServe(config.Addr, nil)
}

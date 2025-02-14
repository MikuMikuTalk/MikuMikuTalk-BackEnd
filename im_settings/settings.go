package main

import (
	"encoding/json"
	"flag"

	"im_server/common/etcd"
	"im_server/im_settings/config"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/settings.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 配置日志
	logx.MustSetup(c.Log)
	etcdStorage := etcd.NewEtcdStorage(c.Etcd)
	// authExpireStr := strconv.FormatInt(c.Auth.AuthExpire, 10)
	result, err := json.Marshal(&c)
	if err != nil {
		logx.Error(err)
		return
	}

	etcdStorage.Put("config", string(result))

	logx.Info("配置文件输入成功")

}

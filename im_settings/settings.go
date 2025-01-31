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
	// // 存储jwt
	// etcdStorage.Put("AuthSecret", c.Auth.AuthSecret)
	// etcdStorage.Put("AuthExpire", authExpireStr)
	// // 存储数据库datasource
	// etcdStorage.Put("MySqlDataSource", c.Mysql.DataSource)
	// // 存储redis配置
	// etcdStorage.Put("RedisHost", c.Redis.Host)
	// etcdStorage.Put("RedisPass", c.Redis.Pass)
	// etcdStorage.Put("RedisType", c.Redis.Type)

}

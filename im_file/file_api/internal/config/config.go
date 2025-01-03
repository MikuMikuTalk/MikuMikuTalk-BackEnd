package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Etcd      string
	FileSize  float64
	ImageSize float64
	WhiteList []string
	Mysql     struct {
		DataSource string
	}
	UploadDir string // 上传文件保存目录
	UserRpc   zrpc.RpcClientConf
	Auth      struct {
		AuthSecret string
		AuthExpire int64
	}
}

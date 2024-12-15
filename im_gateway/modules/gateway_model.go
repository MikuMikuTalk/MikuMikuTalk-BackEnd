package modules

import "github.com/zeromicro/go-zero/core/logx"

// 配置结构体定义
type Config struct {
	Addr string // 网关服务监听地址
	Etcd string // Etcd 服务地址
	Log  logx.LogConf
}

// HTTP 响应的通用结构体
type Response struct {
	Code int    `json:"code"` // 状态码
	Data any    `json:"data"` // 数据
	Msg  string `json:"msg"`  // 消息
}

// 网关服务结构体
type GatewayService struct {
	config Config // 配置信息
}

// 创建新的网关服务实例
func NewGatewayService(config Config) *GatewayService {
	return &GatewayService{
		config: config,
	}
}

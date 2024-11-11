package config

import (
	"path/filepath"
	"runtime"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	Log logx.LogConf
}

var config Config

// 加载配置文件，只初始化一次
func InitConfig() {
	_, currentFile, _, _ := runtime.Caller(0)
	conf.MustLoad(filepath.Join(filepath.Dir(currentFile), "../etc/core.yaml"), &config)
}

// 初始化 logx 日志，只需调用一次
func InitLogx() {
	logx.MustSetup(config.Log)
}

// 获取日志配置
func GetConfig() logx.LogConf {
	return config.Log
}

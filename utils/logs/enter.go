package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

var MyLogger *logrus.Logger

func init() {
	MyLogger = logrus.New()
	MyLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05", // 自定义时间格式
	})
	MyLogger.SetOutput(os.Stdout)
	MyLogger.SetLevel(logrus.InfoLevel)
}

// 对外暴露日志记录函数
func Info(args ...interface{}) {
	MyLogger.Info(args...)
}

func Warn(args ...interface{}) {
	MyLogger.Warn(args...)
}

func Error(args ...interface{}) {
	MyLogger.Error(args...)
}

func Debug(args ...interface{}) {
	MyLogger.Debug(args...)
}

package pwd

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hash_pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		slog.Error("加密密码失败: ", "error", err) // 使用键值对形式记录日志
	}
	return string(hash_pwd)
}

func ComparePassword(hashpwd string, origin_pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(origin_pwd))
	if err != nil {
		slog.Error("比较密码失败：", "error", err) // 使用键值对形式记录日志
		return false
	}
	return true
}

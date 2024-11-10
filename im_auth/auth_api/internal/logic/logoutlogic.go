package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"im_server/im_auth/auth_api/internal/svc"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注销接口
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (string, error) {
	if token == "" {
		return "", errors.New("登陆后才可以注销哦")
	}

	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return "", err
	}
	//获取jti
	jti := jwts.ExtractJTI(claims)
	// 获取 Token 的过期时间
	now := time.Now()
	expiration := claims.ExpiresAt.Time.Sub(now)
	// 将 JTI 存入 Redis，设置为注销状态
	key := fmt.Sprintf("logout_%s", jti)
	// 设置redis中数据过期时间
	_, err = l.svcCtx.RDB.SetNX(key, "invalid", expiration).Result()
	if err != nil {
		l.Logger.Error("Redis 错误: ", err)
		return "", errors.New("注销失败，请稍后重试")
	}

	return "注销成功", nil
}

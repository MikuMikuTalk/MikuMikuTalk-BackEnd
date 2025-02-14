package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"im_server/common/contexts"
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

func (l *LogoutLogic) Logout() (string, error) {
	token, ok := l.ctx.Value(contexts.ContextKeyToken).(string)
	if !ok {
		return "", errors.New("登陆后才可以注销哦")
	}
	logx.Info(token)
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return "", err
	}
	// 获取jti
	jti := jwts.ExtractJTI(claims)
	// 获取 Token 的过期时间
	now := time.Now()
	expiration := claims.ExpiresAt.Time.Sub(now)
	// 将 JTI 存入 Redis，设置为注销状态
	key := fmt.Sprintf("logout_%s", jti)
	// 设置redis中数据过期时间
	_, err = l.svcCtx.RDB.SetnxEx(key, "invalid", int(expiration.Seconds()))
	if err != nil {
		l.Logger.Error("Redis 错误: ", err)
		return "", errors.New("注销失败，请稍后重试")
	}
	l.svcCtx.RuntimeLogs.Info("用户注销成功")
	l.svcCtx.RuntimeLogs.SetItem(claims.Nickname, "注销了")
	l.svcCtx.RuntimeLogs.Save(l.ctx)
	return "注销成功", nil
}

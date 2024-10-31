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

	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return "", err
	}
	now := time.Now()
	// 过期时间就是这个jwt的失效时间
	expiration := payload.ExpiresAt.Time.Sub(now)

	key := fmt.Sprintf("logout_%s", payload.Nickname)
	//设置redis中数据过期时间
	l.svcCtx.RDB.SetNX(key, "", expiration)
	return "注销成功", nil
}

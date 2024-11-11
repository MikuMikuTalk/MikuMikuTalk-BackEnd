package logic

import (
	"context"
	"errors"

	"im_server/im_auth/auth_api/internal/svc"
	"im_server/im_auth/auth_api/internal/types"
	"im_server/im_auth/whitelist"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 认证接口
func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp string, err error) {
	if whitelist.IsInList(req.ValidPath, l.svcCtx.Config.WhiteList) {
		return "ok", nil
	}
	// 检查是否提供了 Authorization
	if req.Authorization == "" {
		return "", errors.New("认证失败：缺少授权信息")
	}
	// 验证Token是否有效
	isValid, err := jwts.ValidateToken(req.Authorization, l.svcCtx.Config.Auth.AuthSecret, l.svcCtx.RDB)
	if err != nil || !isValid {
		err = errors.New("认证失败")
		return
	}
	return "ok", nil
}

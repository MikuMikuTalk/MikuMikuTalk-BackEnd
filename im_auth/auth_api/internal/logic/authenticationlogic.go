package logic

import (
	"context"
	"errors"

	"im_server/im_auth/auth_api/internal/svc"
	"im_server/im_auth/auth_api/internal/types"
	"im_server/utils/jwts"
	"im_server/utils/whitelist"

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
	// 常规方法
	if whitelist.IsInList(req.ValidPath, l.svcCtx.Config.WhiteList) {
		return "ok", nil
	}
	// 使用正则表达式

	if whitelist.IsInListByRegex(l.svcCtx.Config.WhiteList, req.ValidPath) {
		logx.Infof("%s 在白名单中", req.ValidPath)
		return "ok", nil
	}
	logx.Infof("%s 不在白名单中", req.ValidPath)
	// 检查是否提供了 Authorization
	if req.Authorization == "" {
		logx.Error("Authorization 请求头为空")
		return "", errors.New("认证失败：缺少授权信息")
	}
	// 验证Token是否有效
	isValid, err := jwts.ValidateToken(req.Authorization, l.svcCtx.Config.Auth.AuthSecret, l.svcCtx.RDB)
	if err != nil || !isValid {
		return "", errors.New("认证失败")
	}

	return "ok", nil
}

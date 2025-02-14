package logic

import (
	"context"
	"errors"
	"fmt"

	"im_server/im_auth/auth_api/internal/svc"
	"im_server/im_auth/auth_api/internal/types"
	"im_server/im_auth/auth_models"
	"im_server/utils/jwts"
	"im_server/utils/pwd"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nickname = ?", req.UserName).Error

	l.svcCtx.ActionLogs.Info("用户登录操作")
	l.svcCtx.ActionLogs.SetItem("nickname", req.UserName)
	l.svcCtx.ActionLogs.IsRequest()
	l.svcCtx.ActionLogs.IsResponse()
	l.svcCtx.ActionLogs.IsHeaders()
	defer l.svcCtx.ActionLogs.Save(l.ctx)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, fmt.Errorf("数据库错误❌:%w", err)
	}
	if !pwd.ComparePassword(user.Pwd, req.Password) {
		return nil, errors.New("用户密码不正确！")
	}
	token, err := jwts.GenerateJwtToken(jwts.JwtPayload{
		UserID:   user.ID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AuthSecret, int(l.svcCtx.Config.Auth.AuthExpire))
	if err != nil {
		logx.Error(err)
		l.svcCtx.ActionLogs.SetItem("error", err.Error())
		l.svcCtx.ActionLogs.Err("服务内部错误")
		err = fmt.Errorf("生成 JWT 失败: %w", err)
		return nil, err
	}

	loginResponse := types.LoginResponse{
		Token: token,
	}
	l.svcCtx.ActionLogs.Info("用户登录成功")
	l.svcCtx.ActionLogs.SetCtx(l.ctx)

	return &loginResponse, nil
}

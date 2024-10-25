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
	// todo: add your logic here and delete this line
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nickname = ?", req.UserName).Error
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
		err = fmt.Errorf("生成 JWT 失败: %w", err)
		return nil, err
	}
	return &types.LoginResponse{
		Token: token,
	}, nil
}

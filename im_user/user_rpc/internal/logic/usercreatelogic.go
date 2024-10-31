package logic

import (
	"context"
	"errors"

	"im_server/im_user/user_models"
	"im_server/im_user/user_rpc/internal/svc"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/logs"
	"im_server/utils/pwd"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {

	err := l.svcCtx.DB.Take("nickname = ?", in.NickName).Error
	if err == nil {
		return nil, errors.New("用户已经存在，请不要重复注册！")
	}
	var user user_models.UserModel = user_models.UserModel{
		Nickname:       in.NickName,
		Pwd:            pwd.HashPassword(in.Password),
		Role:           int8(in.Role),
		Avatar:         in.Avatar,
		RegisterSource: in.RegisterSource,
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		logs.Error("用户创建失败")
		return nil, errors.New("用户创建失败")
	}
	return &user_rpc.UserCreateResponse{
		UserName: user.Nickname,
	}, nil
}

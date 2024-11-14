package logic

import (
	"context"
	"errors"

	"im_server/im_user/user_models"
	"im_server/im_user/user_rpc/internal/svc"
	"im_server/im_user/user_rpc/types/user_rpc"
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
		logx.Error("用户创建失败")
		return nil, errors.New("用户创建失败")
	}
	var userconf user_models.UserConfModel = user_models.UserConfModel{
		UserID:        user.ID,
		RecallMessage: nil,   // 撤回消息的提示内容  撤回了一条消息
		FriendOnline:  false, // 关闭好友上线提醒
		Sound:         true,  // 开启声音
		SecureLink:    false, // 关闭安全链接
		SavePwd:       false, // 不保存密码
		SearchUser:    2,     // 可以通过用户id和昵称搜索
		Verification:  2,     // 需要验证消息
		Online:        true,
	}
	err = l.svcCtx.DB.Create(&userconf).Error
	if err != nil {
		logx.Error("用户配置信息创建失败")
		return nil, errors.New("用户配置信息创建失败")
	}

	return &user_rpc.UserCreateResponse{
		UserName: user.Nickname,
	}, nil
}

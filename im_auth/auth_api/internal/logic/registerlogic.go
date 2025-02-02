package logic

import (
	"context"

	"im_server/im_auth/auth_api/internal/svc"
	"im_server/im_auth/auth_api/internal/types"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/logs"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	logs.Info("注册服务调用")
	res, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user_rpc.UserCreateRequest{
		NickName:       req.UserName,
		Password:       req.Password,
		Role:           2,
		Avatar:         "",
		RegisterSource: "账户密码注册",
	})
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &types.RegisterResponse{
		UserName: res.GetUserName(),
	}, nil
}

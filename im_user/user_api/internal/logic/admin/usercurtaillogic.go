package admin

import (
	"context"

	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCurtailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCurtailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCurtailLogic {
	return &UserCurtailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCurtailLogic) UserCurtail(req *types.UserCurtailRequest) (resp *types.UserCurtailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

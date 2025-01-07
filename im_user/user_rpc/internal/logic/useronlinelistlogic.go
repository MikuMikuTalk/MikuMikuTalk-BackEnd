package logic

import (
	"context"
	"strconv"

	"im_server/im_user/user_rpc/internal/svc"
	"im_server/im_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOnlineListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOnlineListLogic {
	return &UserOnlineListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户在线id列表
func (l *UserOnlineListLogic) UserOnlineList(in *user_rpc.UserOnlineListRequest) (resp *user_rpc.UserOnlineListResponse, err error) {
	resp = new(user_rpc.UserOnlineListResponse)
	onlineMap := l.svcCtx.Redis.HGetAll("online").Val()
	for key, _ := range onlineMap {
		val, err := strconv.Atoi(key)
		if err != nil {
			logx.Error(err)
			continue
		}
		resp.UserIdList = append(resp.UserIdList, uint32(val))
	}

	return
}

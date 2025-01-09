package logic

import (
	"context"

	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_user/user_models"
	"im_server/im_user/user_rpc/internal/svc"
	"im_server/im_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *user_rpc.FriendListRequest) (*user_rpc.FriendListResponse, error) {
	friends, _, _ := list_query.ListQuery(l.svcCtx.DB, user_models.FriendModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Limit: -1, // 查看全部
		},
		Where:   l.svcCtx.DB.Where("send_user_id = ? or rev_user_id = ?", in.User, in.User),
		Preload: []string{"SendUserModel", "RevUserModel"},
	})
	var list []*user_rpc.FriendInfo
	for _, friend := range friends {
		info := user_rpc.FriendInfo{}
		if friend.SendUserID == uint(in.User) {
			// 我是发起方
			info.UserId = uint32(friend.RevUserID)
			info.NickName = friend.RevUserModel.Nickname
			info.Avatar = friend.RevUserModel.Avatar
		}

		if friend.RevUserID == uint(in.User) {
			// 如果我是接收方
			info.UserId = uint32(friend.SendUserID)
			info.NickName = friend.SendUserModel.Nickname
			info.Avatar = friend.SendUserModel.Avatar
		}
		list = append(list, &info)
	}

	return &user_rpc.FriendListResponse{
		FriendList: list,
	}, nil
}

package logic

import (
	"context"
	"strconv"

	"im_server/common/contexts"
	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友列表获取
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {
	user_id := l.ctx.Value(contexts.ContextKeyUserID).(uint)

	// 使用通用列表查询
	friend_list, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.FriendModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Where:   l.svcCtx.DB.Where("send_user_id = ? or rev_user_id = ?", user_id, user_id),
		Preload: []string{"SendUserModel", "RevUserModel"},
	})

	// 查询哪些好友在线
	onlineMap, err := l.svcCtx.Redis.Hgetall("online_user")
	if err != nil {
		return nil, err
	}
	onlineUserMap := map[uint]bool{}

	for key := range onlineMap {
		val, err1 := strconv.Atoi(key)
		if err1 != nil {
			logx.Error(err1)
			continue
		}
		onlineUserMap[uint(val)] = true
	}

	var friend_info_responses []types.FriendInfoResponse
	for _, friend := range friend_list {
		info := types.FriendInfoResponse{}
		if friend.SendUserID == user_id {
			// 我是发起方
			info = types.FriendInfoResponse{
				FriendID: friend.RevUserID,
				Nickname: friend.RevUserModel.Nickname,
				Abstract: friend.RevUserModel.Abstract,
				Avatar:   friend.RevUserModel.Avatar,
				Notice:   friend.SenUserNotice,
				IsOnline: onlineUserMap[friend.RevUserID],
			}
		}
		if friend.RevUserID == user_id {
			// 我是接收方
			info = types.FriendInfoResponse{
				FriendID: friend.SendUserID,
				Nickname: friend.SendUserModel.Nickname,
				Abstract: friend.SendUserModel.Abstract,
				Avatar:   friend.SendUserModel.Avatar,
				Notice:   friend.RevUserNotice,
			}
		}
		friend_info_responses = append(friend_info_responses, info)
	}

	return &types.FriendListResponse{
		List:  friend_info_responses,
		Count: count,
	}, nil
}

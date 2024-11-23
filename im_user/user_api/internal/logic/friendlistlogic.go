package logic

import (
	"context"

	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"
	"im_server/utils/jwts"

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

func (l *FriendListLogic) FriendList(req *types.FriendListRequest, token string) (resp *types.FriendListResponse, err error) {
	clamis, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	user_id := clamis.UserID
	var count int64
	l.svcCtx.DB.Model(user_models.FriendModel{}).Where("send_user_id = ? or rev_user_id = ?", user_id, user_id).Count(&count)
	var friend_list []user_models.FriendModel
	l.svcCtx.DB.Preload("SendUserModel").Preload("RevUserModel").Find(&friend_list, "send_user_id = ? or rev_user_id = ?", user_id, user_id)
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
				Notice:   friend.RevUserNotice,
			}
		}
		if friend.RevUserID == user_id {
			// 我是接收方
			info = types.FriendInfoResponse{
				FriendID: friend.SendUserID,
				Nickname: friend.SendUserModel.Nickname,
				Abstract: friend.SendUserModel.Abstract,
				Avatar:   friend.SendUserModel.Avatar,
				Notice:   friend.SenUserNotice,
			}
		}
		friend_info_responses = append(friend_info_responses, info)
	}

	return &types.FriendListResponse{
		List:  friend_info_responses,
		Count: count,
	}, nil
}

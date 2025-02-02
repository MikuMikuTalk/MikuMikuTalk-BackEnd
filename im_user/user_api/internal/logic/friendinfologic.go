package logic

import (
	"context"
	"encoding/json"

	"im_server/common/contexts"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"
	"im_server/im_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友信息接口
func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendInfoLogic) FriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	// 确定查的用户是自己的好友
	var user user_models.UserModel

	err = l.svcCtx.DB.Take(&user, "nickname = ?", req.FriendName).Error
	if err != nil {
		logx.Error("查找的好友不存在！")
		return
	}

	my_id := l.ctx.Value(contexts.ContextKeyUserID).(uint)

	var friend user_models.FriendModel

	err = l.svcCtx.DB.Take(&friend, "(send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)", my_id, user.ID, user.ID, my_id).Error
	if err != nil {
		logx.Errorf("用户%s不是您的好友", friend.RevUserModel.Nickname)
		return
	}
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user_rpc.UserInfoRequest{
		UserId: uint32(user.ID),
	})
	if err != nil {
		logx.Error(err)
		return
	}
	var friend_info user_models.UserModel
	err = json.Unmarshal(res.Data, &friend_info)
	if err != nil {
		logx.Error("解析用户信息失败:", err)
		return
	}
	resp = &types.FriendInfoResponse{
		FriendID: friend_info.ID,
		Nickname: friend_info.Nickname,
		Abstract: friend_info.Abstract,
		Avatar:   friend_info.Abstract,
		Notice:   friend.GetUserNotice(user.ID),
	}
	return resp, nil
}

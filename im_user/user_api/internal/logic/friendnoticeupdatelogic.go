package logic

import (
	"context"
	"errors"

	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendNoticeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友备注修改
func NewFriendNoticeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendNoticeUpdateLogic {
	return &FriendNoticeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendNoticeUpdateLogic) FriendNoticeUpdate(req *types.FriendNoticeUpdateRequest, token string) (resp *types.FriendNoticeUpdateResponse, err error) {

	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	user_id := claims.UserID
	var friend user_models.FriendModel

	if !friend.IsFriend(l.svcCtx.DB, user_id, req.FriendID) {
		err = errors.New("他还不是你的好友呢~")
		return
	}
	if friend.SendUserID == user_id {
		//我是发起方
		if friend.SenUserNotice == req.Notice {
			return
		}
		l.svcCtx.DB.Model(&friend).Update("sen_user_notice", req.Notice)
	}
	if friend.RevUserID == user_id {
		//我是接收方
		if friend.RevUserNotice == req.Notice {
			return
		}
		l.svcCtx.DB.Model(&friend).Update("rev_user_notice", req.Notice)
	}

	return
}

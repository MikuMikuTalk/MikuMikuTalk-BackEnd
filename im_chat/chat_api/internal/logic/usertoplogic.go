package logic

import (
	"context"
	"errors"

	"im_server/im_chat/chat_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"

	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友置顶
func NewUserTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTopLogic {
	return &UserTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserTopLogic) UserTop(req *types.UserTopRequest, token string) (resp *types.UserTopResponse, err error) {
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	my_id := claims.UserID
	if my_id != req.FriendID {
		res, err := l.svcCtx.UserRpc.IsFriend(l.ctx, &user_rpc.IsFriendRequest{
			User2: uint32(my_id),
			User1: uint32(req.FriendID),
		})
		if err != nil {
			return nil, err
		}
		if !res.IsFriend {
			return nil, errors.New("你们还不是好友呢")
		}
	}
	var top_user types.UserTopResponse
	err1 := l.svcCtx.DB.Take(&top_user, "user_id = ? and top_user_id = ?", my_id, req.FriendID).Error
	if err1 != nil {
		// 没有置顶
		l.svcCtx.DB.Create(&chat_models.TopUserModel{
			UserID:    my_id,
			TopUserID: req.FriendID,
		})
		return
	}
	// 已经有置顶了，就取消置顶
	l.svcCtx.DB.Delete(&top_user)
	return
}

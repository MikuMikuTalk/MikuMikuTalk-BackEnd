package logic

import (
	"context"
	"im_server/im_chat/chat_models"
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
	//已经有置顶了，就取消置顶
	l.svcCtx.DB.Model(chat_models.TopUserModel{}).Delete("user_id = ? and top_user_id = ?", my_id, req.FriendID)
	return
}

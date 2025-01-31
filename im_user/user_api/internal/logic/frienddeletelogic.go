package logic

import (
	"context"
	"errors"

	"im_server/common/contexts"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友删除接口
func NewFriendDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDeleteLogic {
	return &FriendDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendDeleteLogic) FriendDelete(req *types.FriendDeleteRequest) (resp *types.FriendDeleteResponse, err error) {
	my_id := l.ctx.Value(contexts.ContextKeyUserID).(uint)
	friend_name := req.FriendName
	var user user_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nickname = ?", friend_name).Error
	if err != nil {
		err = errors.New("用户不存在")
		return
	}
	friend_id := user.ID

	var friend_model user_models.FriendModel
	err = l.svcCtx.DB.Take(&friend_model, "(send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)", my_id, friend_id, friend_id, my_id).Error
	if err != nil {
		logx.Errorf("用户%s不是您的好友", friend_model.RevUserModel.Nickname)
		return
	}

	l.svcCtx.DB.Delete(&friend_model)
	return
}

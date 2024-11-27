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

type ValidStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友验证状态操作
func NewValidStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidStatusLogic {
	return &ValidStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidStatusLogic) ValidStatus(req *types.FriendValidStatusRequest, token string) (resp *types.FriendValidStatusResponse, err error) {
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return
	}
	my_id := claims.UserID

	// 别人给我发好友请求，在请求的验证表里找到了接受用户是我的验证
	var friendVerify user_models.FriendVerifyModel
	err = l.svcCtx.DB.Take(&friendVerify, "id = ? and rev_user_id = ?", req.VerifyID, my_id).Error
	if err != nil {
		err = errors.New("验证记录不存在")
		return
	}
	if friendVerify.RevStatus != 0 {
		err = errors.New("不可更改状态")
		return
	}
	switch req.Status {
	case 1: // 同意
		friendVerify.RevStatus = 1
		// 加入到好友表中
		l.svcCtx.DB.Create(&user_models.FriendModel{
			SendUserID: friendVerify.SendUserID,
			RevUserID:  friendVerify.RevUserID,
		})
	case 2: // 拒绝
		friendVerify.RevStatus = 2
	case 3: // 忽略
		friendVerify.RevStatus = 3
	case 4: // 删除
		// 删除验证记录
		l.svcCtx.DB.Delete(&friendVerify)
		err = nil
		return
	}
	l.svcCtx.DB.Save(&friendVerify)
	return
}

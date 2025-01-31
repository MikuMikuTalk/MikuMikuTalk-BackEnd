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

type UserValidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友验证接口
func NewUserValidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserValidLogic {
	return &UserValidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserValidLogic) UserValid(req *types.UserValidRequest) (resp *types.UserValidResponse, err error) {

	my_id := l.ctx.Value(contexts.ContextKeyUserID).(uint)
	friend_nickname := req.FriendName
	var user_conf user_models.UserConfModel
	var user user_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nickname = ?", friend_nickname).Error
	if err != nil {
		err = errors.New("查找的好友不存在")
		logx.Error("查找的好友不存在！")
		return
	}
	err = l.svcCtx.DB.Take(&user_conf, "user_id = ?", user.ID).Error
	if err != nil {
		err = errors.New("查找的好友不存在")
		logx.Error("查找的好友不存在！")
		return
	}
	// 查看找的是不是自己的好友
	var friend user_models.FriendModel
	err = l.svcCtx.DB.Take(&friend, "(send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)", my_id, user.ID, user.ID, my_id).Error
	// 如果查出来是自己的好友，那就不要重复添加
	if err == nil {
		err = errors.New("他已经是您的好友了")
		logx.Info("他已经是您的好友了")
		return
	}

	resp = new(types.UserValidResponse)
	resp.Verification = user_conf.Verification
	// 0 不允许任何人, 1 允许任何人, 2 验证消息, 3 回答问题, 4 正确回答问题
	switch user_conf.Verification {
	case 0: // 不允许任何人添加
	case 1:
		// 允许任何人添加
		// 直接成为好友
	case 2:
		// 需要先验证问题
	case 3, 4:
		// 需要正确回答问题
		if user_conf.VerificationQuestion != nil {
			resp.VerificationQuestion = types.VerificationQuestion{
				Problem1: user_conf.VerificationQuestion.Problem1,
				Problem2: user_conf.VerificationQuestion.Problem1,
				Problem3: user_conf.VerificationQuestion.Problem1,
			}
		}
	default:
	}

	return
}

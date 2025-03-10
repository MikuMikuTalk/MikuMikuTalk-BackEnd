package logic

import (
	"context"
	"errors"

	"im_server/common/contexts"
	"im_server/common/ctype"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友添加接口
func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddFriendRequest) (resp *types.AddFriendResponse, err error) {
	my_id := l.ctx.Value(contexts.ContextKeyUserID).(uint)
	friend_nickname := req.FriendName

	// 限制用户加好友
	var conf user_models.UserConfModel
	err = l.svcCtx.DB.Take(&conf, "user_id = ?", my_id).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if conf.CurtailAddUser {
		return nil, errors.New("当前用户限制加好友")
	}

	var user user_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nickname = ?", friend_nickname).Error
	if err != nil {
		err = errors.New("查找的好友不存在")
		logx.Error("查找的好友不存在！")
		return
	}
	var userConf user_models.UserConfModel
	err = l.svcCtx.DB.Take(&userConf, "user_id = ?", user.ID).Error
	if err != nil {
		return nil, errors.New("用户不存在")
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
	resp = new(types.AddFriendResponse)
	verifyModel := user_models.FriendVerifyModel{
		SendUserID:         my_id,
		RevUserID:          user.ID,
		AdditionalMessages: req.Verify,
	}

	switch userConf.Verification {
	case 0: // 不允许任何人添加
		return nil, errors.New("该用户不允许任何人添加")
	case 1: // 允许任何人添加
		// 直接成为好友
		// 先往验证表里面加一条记录，然后通过
		verifyModel.RevStatus = 1
		userFriend := user_models.FriendModel{
			SendUserID: my_id,
			RevUserID:  user.ID,
		}
		l.svcCtx.DB.Create(&userFriend)
	case 2: // 需要验证问题
	case 3: // 需要回答问题
		if req.VerificationQuestion != nil {
			verifyModel.VerificationQuestion = &ctype.VerificationQuestion{
				Problem1: req.VerificationQuestion.Problem1,
				Problem2: req.VerificationQuestion.Problem2,
				Problem3: req.VerificationQuestion.Problem3,
				Answer1:  req.VerificationQuestion.Answer1,
				Answer2:  req.VerificationQuestion.Answer2,
				Answer3:  req.VerificationQuestion.Answer3,
			}
		}
	case 4:
		// 判断问题是否回答正确
		if req.VerificationQuestion != nil && userConf.VerificationQuestion != nil {
			// 考虑到一个问题，两个问题，三个问题的情况
			var count int
			if userConf.VerificationQuestion.Answer1 != nil && req.VerificationQuestion.Answer1 != nil {
				if *userConf.VerificationQuestion.Answer1 == *req.VerificationQuestion.Answer1 {
					count += 1
				}
			}
			if userConf.VerificationQuestion.Answer2 != nil && req.VerificationQuestion.Answer2 != nil {
				if *userConf.VerificationQuestion.Answer2 == *req.VerificationQuestion.Answer2 {
					count += 1
				}
			}
			if userConf.VerificationQuestion.Answer3 != nil && req.VerificationQuestion.Answer3 != nil {
				if *userConf.VerificationQuestion.Answer3 == *req.VerificationQuestion.Answer3 {
					count += 1
				}
			}
			if count != userConf.ProblemCount() {
				return nil, errors.New("答案错误")
			}
			// 直接加好友
			verifyModel.RevStatus = 1
			verifyModel.VerificationQuestion = userConf.VerificationQuestion
			// 加好友
			userFriend := user_models.FriendModel{
				SendUserID: my_id,
				RevUserID:  user.ID,
			}
			l.svcCtx.DB.Create(&userFriend)
		}
	default:
		return nil, errors.New("不支持的验证参数")
	}
	err = l.svcCtx.DB.Create(&verifyModel).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("添加好友失败")
	}
	return
}

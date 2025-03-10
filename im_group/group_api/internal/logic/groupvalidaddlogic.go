package logic

import (
	"context"
	"encoding/json"
	"errors"

	"im_server/common/ctype"
	"im_server/im_group/group_models"
	"im_server/im_user/user_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupValidAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupValidAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupValidAddLogic {
	return &GroupValidAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupValidAddLogic) GroupValidAdd(req *types.AddGroupRequest) (resp *types.AddGroupResponse, err error) {
	// 解析 JWT 获取当前用户 ID
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, errors.New("无效的 Token")
	}
	myID := claims.UserID
	// 加群

	// 获取用户基本信息
	userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user_rpc.UserInfoRequest{
		UserId: uint32(myID),
	})
	if err != nil {
		logx.Error(err)
		return nil, errors.New("用户服务错误")
	}
	var userInfoModel user_models.UserModel
	// 如果被限制加群了，那就返回错误
	json.Unmarshal(userInfo.Data, &userInfoModel)
	if userInfoModel.UserConfModel.CurtailInGroupChat {
		return nil, errors.New("当前用户被限制加群")
	}

	// 如果自己已经在群里面了
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupID, myID).Error
	if err == nil {
		return nil, errors.New("请勿重复加群")
	}

	var group group_models.GroupModel
	err = l.svcCtx.DB.Take(&group, req.GroupID).Error
	if err != nil {
		return nil, errors.New("群不存在")
	}

	resp = new(types.AddGroupResponse)
	verifyModel := group_models.GroupVerifyModel{
		GroupID:            req.GroupID,
		UserID:             myID,
		Status:             0,
		AdditionalMessages: req.Verify,
		Type:               1, // 加群
	}
	switch group.Verification {
	case 0: // 不允许任何人添加
		return nil, errors.New("不允许任何人加群")
	case 1: // 允许任何人添加
		// 直接成为好友
		// 先往验证表里面加一条记录，然后通过
		verifyModel.Status = 1
		groupMember := group_models.GroupMemberModel{
			GroupID: req.GroupID,
			UserID:  myID,
			Role:    3,
		}
		l.svcCtx.DB.Create(&groupMember)
	case 2: // 需要验证问题

	case 3: // 需要回答问题
		if req.VerificationQuestion != nil {
			verifyModel.VerificationQuestion = &ctype.VerificationQuestion{
				Problem1: group.VerificationQuestion.Problem1,
				Problem2: group.VerificationQuestion.Problem2,
				Problem3: group.VerificationQuestion.Problem3,
				Answer1:  req.VerificationQuestion.Answer1,
				Answer2:  req.VerificationQuestion.Answer2,
				Answer3:  req.VerificationQuestion.Answer3,
			}
		}
	case 4: // 需要正确回答问题
		// 判断问题是否回答正确
		if req.VerificationQuestion != nil && group.VerificationQuestion != nil {
			// 考虑到一个问题，两个问题，三个问题的情况
			var count int
			if group.VerificationQuestion.Answer1 != nil && req.VerificationQuestion.Answer1 != nil {
				if *group.VerificationQuestion.Answer1 == *req.VerificationQuestion.Answer1 {
					count += 1
				}
			}
			if group.VerificationQuestion.Answer2 != nil && req.VerificationQuestion.Answer2 != nil {
				if *group.VerificationQuestion.Answer2 == *req.VerificationQuestion.Answer2 {
					count += 1
				}
			}
			if group.VerificationQuestion.Answer3 != nil && req.VerificationQuestion.Answer3 != nil {
				if *group.VerificationQuestion.Answer3 == *req.VerificationQuestion.Answer3 {
					count += 1
				}
			}
			if count != group.ProblemCount() {
				return nil, errors.New("答案错误")
			}
			// 直接加群
			verifyModel.Status = 1
			verifyModel.VerificationQuestion = group.VerificationQuestion
			// 把用户加到群里面
			groupMember := group_models.GroupMemberModel{
				GroupID: req.GroupID,
				UserID:  myID,
				Role:    3,
			}
			l.svcCtx.DB.Create(&groupMember)
		} else {
			return nil, errors.New("答案错误")
		}

	default:

	}
	err = l.svcCtx.DB.Create(&verifyModel).Error
	if err != nil {
		return
	}

	return
}

package logic

import (
	"context"
	"errors"

	"im_server/im_group/group_models"
	"im_server/utils/jwts"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberNicknameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberNicknameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberNicknameUpdateLogic {
	return &GroupMemberNicknameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberNicknameUpdateLogic) GroupMemberNicknameUpdate(req *types.GroupMemberNicknameUpdateRequest) (resp *types.GroupMemberNicknameUpdateResponse, err error) {
	// 解析 JWT 获取当前用户 ID
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, errors.New("无效的 Token")
	}
	myID := claims.UserID
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.ID, myID).Error
	if err != nil {
		return nil, errors.New("违规调用")
	}
	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.ID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("该用户不是群成员")
	}

	// 自己修改自己的
	if myID == req.MemberID {
		l.svcCtx.DB.Model(&member).Updates(map[string]any{
			"member_nickname": req.Nickname,
		})
		return
	}

	if !((member.Role == 1 && (member1.Role == 2 || member1.Role == 3)) || (member.Role == 2 && member1.Role == 3)) {
		return nil, errors.New("用户角色错误")
	}
	l.svcCtx.DB.Model(&member1).Updates(map[string]any{
		"member_nickname": req.Nickname,
	})
	return
}

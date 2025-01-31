package logic

import (
	"context"
	"errors"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberRoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberRoleUpdateLogic {
	return &GroupMemberRoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberRoleUpdateLogic) GroupMemberRoleUpdate(req *types.GroupMemberRoleUpdateRequest) (resp *types.GroupMemberRoleUpdateResponse, err error) {
	// 解析 JWT 获取当前用户 ID
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, errors.New("无效的 Token")
	}
	myID := claims.UserID
	// 看看我在不在这个群里
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.ID, myID).Error
	if err != nil {
		return nil, errors.New("违规调用")
	}
	// 如果我在这个群里，看看我是不是群主，群主才能让群成员变成管理员或者把管理员降级为群成员
	if member.Role != 1 {
		return nil, errors.New("权限错误")
	}
	// 看看要更新的人在不在群里
	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.ID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("用户还不是群成员呢")
	}
	// 一个群只能有一个群主
	if !(req.Role == 2 || req.Role == 3) {
		return nil, errors.New("用户角色设置错误")
	}
	if member1.Role == req.Role {
		return
	}
	l.svcCtx.DB.Model(&member1).Update("role", req.Role)

	return
}

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

type GroupMemberRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberRemoveLogic {
	return &GroupMemberRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberRemoveLogic) GroupMemberRemove(req *types.GroupMemberRemoveRequest) (resp *types.GroupMemberRemoveResponse, err error) {
	// 解析 JWT 获取当前用户 ID
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, errors.New("无效的 Token")
	}
	myID := claims.UserID

	// 验证当前用户是否为群成员
	var currentMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&currentMember, "group_id = ? and user_id = ?", req.ID, myID).Error
	if err != nil {
		return nil, errors.New("当前用户不是群成员")
	}

	// 验证当前用户是否有权限（群主或管理员）
	if currentMember.Role != 1 && currentMember.Role != 2 {
		return nil, errors.New("当前用户没有踢人权限")
	}

	// 验证被踢用户是否为群成员
	var targetMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&targetMember, "group_id = ? and user_id = ?", req.ID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("被踢用户不是群成员")
	}

	// 踢人逻辑
	if currentMember.Role == 1 {
		// 群主可以踢任何人
		err = l.svcCtx.DB.Delete(&targetMember).Error
		if err != nil {
			logx.Errorf("群主踢人失败: %v", err)
			return nil, errors.New("踢人失败")
		}
		logx.Infof("群主踢人成功: 群组ID=%d, 被踢用户ID=%d", req.ID, req.MemberID)
		return &types.GroupMemberRemoveResponse{}, nil
	} else if currentMember.Role == 2 {
		// 管理员只能踢普通用户
		if targetMember.Role == 3 {
			err = l.svcCtx.DB.Delete(&targetMember).Error
			if err != nil {
				logx.Errorf("管理员踢人失败: %v", err)
				return nil, errors.New("踢人失败")
			}
			logx.Infof("管理员踢人成功: 群组ID=%d, 被踢用户ID=%d", req.ID, req.MemberID)
			return &types.GroupMemberRemoveResponse{}, nil
		} else {
			return nil, errors.New("管理员只能踢普通用户")
		}
	}

	// 其他情况
	return nil, errors.New("没有权限")

}

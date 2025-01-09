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

type GroupRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupRemoveLogic {
	return &GroupRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupRemoveLogic) GroupRemove(req *types.GroupRemoveRequest) (resp *types.GroupRemoveResponse, err error) {
	token := req.Token
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	my_id := claims.UserID
	var groupMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&groupMember, "group_id = ? and user_id = ?", req.ID, my_id).Error
	if err != nil {
		return nil, errors.New("群不存在或用户不是群成员")
	}
	if groupMember.Role != 1 {
		return nil, errors.New("只有群主才能解散该群哦")
	}
	// 群关联的群消息要删除
	var msgList []group_models.GroupMsgModel
	l.svcCtx.DB.Find(&msgList, "group_id = ? ", req.ID).Delete(&msgList)
	// 关联的群成员也要删除
	var memberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&memberList, "group_id = ? ", req.ID).Delete(&memberList)
	// 群验证消息
	var verificationList []group_models.GroupVerifyModel
	l.svcCtx.DB.Find(&verificationList, "group_id = ? ", req.ID).Delete(&verificationList)
	// 群删除
	var group group_models.GroupModel
	l.svcCtx.DB.Find(&group, req.ID).Delete(&group)

	logx.Infof("删除群：%s", group.Title)
	logx.Infof("关联群成员数：%d", len(memberList))
	logx.Infof("关联群消息数：%d", len(msgList))
	logx.Infof("关联群验证消息数：%d", len(verificationList))

	return
}

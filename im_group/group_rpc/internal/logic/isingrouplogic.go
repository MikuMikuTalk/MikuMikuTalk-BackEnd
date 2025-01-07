package logic

import (
	"context"

	"im_server/im_group/group_models"
	"im_server/im_group/group_rpc/internal/svc"
	"im_server/im_group/group_rpc/types/group_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsInGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsInGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsInGroupLogic {
	return &IsInGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsInGroupLogic) IsInGroup(in *group_rpc.IsInGroupRequest) (resp *group_rpc.IsInGroupResponse, err error) {
	resp = new(group_rpc.IsInGroupResponse)
	var groupMember group_models.GroupMemberModel
	err1 := l.svcCtx.DB.Take(&groupMember, "group_id = ? and user_id = ?", in.GroupId, in.UserId).Error
	if err1 != nil {
		resp.IsInGroup = false
		return
	}
	resp.IsInGroup = true
	return
}

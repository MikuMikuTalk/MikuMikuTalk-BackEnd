package logic

import (
	"context"
	"errors"
	"fmt"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupProhibitionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupProhibitionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupProhibitionUpdateLogic {
	return &GroupProhibitionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 设置用户禁言时间
func (l *GroupProhibitionUpdateLogic) GroupProhibitionUpdate(req *types.GroupProhibitionUpdateRequest) (resp *types.GroupProhibitionUpdateResponse, err error) {
	token := req.Token
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	myID := claims.UserID
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupID, myID).Error
	if err != nil {
		return nil, errors.New("当前用户错误")
	}
	// 如果不是管理员或者群主，那么就没权限去禁言用户
	if !(member.Role == 1 || member.Role == 2) {
		return nil, errors.New("当前用户角色错误")
	}
	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.GroupID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("目标用户不是群成员")
	}
	if !((member.Role == 1 && member1.Role == 2 || member1.Role == 3) || (member.Role == 2 && member1.Role == 3)) {
		return nil, errors.New("角色错误")
	}
	l.svcCtx.DB.Model(&member1).Update("prohibition_time", req.ProhibitionTime)
	// 利用redis的过期时间去做这个禁言时间
	key := fmt.Sprintf("prohibition__%d", member1.ID)

	if req.ProhibitionTime != nil {
		// 给 Redis 设置一个 key，过期时间是 xxxx
		err := l.svcCtx.Redis.Setex(key, "1", *req.ProhibitionTime*60)
		if err != nil {
			return nil, err
		}
	} else {
		// 删除 Redis 中的 key
		_, err := l.svcCtx.Redis.Del(key)
		if err != nil {
			return nil, err
		}
	}
	return
}

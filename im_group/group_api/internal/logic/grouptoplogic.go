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

type GroupTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupTopLogic {
	return &GroupTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupTopLogic) GroupTop(req *types.GroupTopRequest) (resp *types.GroupTopResponse, err error) {
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	myID := claims.UserID
	// 谁能调这个接口 必须得是这个群的成员
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupID, myID).Error
	if err != nil {
		return nil, errors.New("你还不是群成员呢")
	}
	var userTop group_models.GroupUserTopModel
	err1 := l.svcCtx.DB.Take(&userTop, "group_id = ? and user_id = ?", req.GroupID, myID).Error
	if err1 != nil {
		// 查不到，还没有置顶
		if req.IsTop {
			// 我要置顶
			l.svcCtx.DB.Create(&group_models.GroupUserTopModel{
				GroupID: req.GroupID,
				UserID:  myID,
			})
		}
	} else {
		// 查得到
		if !req.IsTop {
			// 取消置顶
			l.svcCtx.DB.Delete(&userTop)
		}
	}
	return
}

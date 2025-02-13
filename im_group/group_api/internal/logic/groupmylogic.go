package logic

import (
	"context"

	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMyLogic {
	return &GroupMyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMyLogic) GroupMy(req *types.GroupMyRequest) (resp *types.GroupMyListResponse, err error) {
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	myID := claims.UserID
	// 查群id列表
	var groupIDList []uint
	query := l.svcCtx.DB.Model(&group_models.GroupMemberModel{}).Where("user_id = ?", myID)
	if req.Mode == 1 {
		// 我创建的群聊
		query.Where("role = ?", 1)
	}
	query.Select("group_id").Scan(&groupIDList)

	groupList, count, _ := list_query.ListQuery(l.svcCtx.DB, group_models.GroupModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preload: []string{"MemberList"},
		Where:   l.svcCtx.DB.Where("id in ?", groupIDList),
	})

	resp = new(types.GroupMyListResponse)
	for _, model := range groupList {

		var role int8
		for _, memberModel := range model.MemberList {
			if memberModel.UserID == myID {
				role = memberModel.Role
			}
		}

		resp.List = append(resp.List, types.GroupMyResponse{
			GroupID:          model.ID,
			GroupTitle:       model.Title,
			GroupAvatar:      model.Avatar,
			GroupMemberCount: len(model.MemberList),
			Role:             role,
			Mode:             req.Mode,
		})
	}
	resp.Count = int(count)

	return
}

package logic

import (
	"context"
	"errors"
	"fmt"

	"im_server/utils/set"

	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupSearchLogic {
	return &GroupSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupSearchLogic) GroupSearch(req *types.GroupSearchRequest) (resp *types.GroupSearchListResponse, err error) {
	// 解析 JWT 获取当前用户 ID
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, errors.New("无效的 Token")
	}
	myID := claims.UserID
	// 先找所有的用户
	// IsSearch 为false就表示不能被搜索  IsSearch = true就可以搜id加入群聊
	groups, count, err := list_query.ListQuery(l.svcCtx.DB, group_models.GroupModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preload: []string{"MemberList"},
		Where:   l.svcCtx.DB.Where("is_search = 1 and ( id = ? or title like ? )", req.Key, fmt.Sprintf("%%%s%%", req.Key)),
	})

	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(context.Background(), &user_rpc.UserOnlineListRequest{})
	var userOnlineIDList []uint
	if err == nil {
		// 称之为：服务降级，如果用户rpc方法挂了，只是页面上看到在线人数是0而已，不会影响这个群搜索功能
		for _, u := range userOnlineResponse.UserIdList {
			userOnlineIDList = append(userOnlineIDList, uint(u))
		}
	}

	resp = new(types.GroupSearchListResponse)
	for _, group := range groups {

		var groupMemberIdList []uint
		var isInGroup bool
		for _, model := range group.MemberList {
			groupMemberIdList = append(groupMemberIdList, model.UserID)
			if model.UserID == myID {
				isInGroup = true
			}
		}
		userOnlineCount := len(set.Intersect(groupMemberIdList, userOnlineIDList))
		resp.List = append(resp.List, types.GroupSearchResponse{
			GroupID:         group.ID,
			Title:           group.Title,
			Abstract:        group.Abstract,
			Avatar:          group.Avatar,
			UserCount:       len(group.MemberList),
			UserOnlineCount: userOnlineCount, // 这个群在线的用户总数
			IsInGroup:       isInGroup,       // 我是否在群里面
		})
	}
	resp.Count = int(count)
	return
}

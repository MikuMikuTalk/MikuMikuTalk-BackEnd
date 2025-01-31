package logic

import (
	"context"
	"fmt"

	"im_server/common/contexts"
	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友搜索接口
func NewFriendSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendSearchLogic {
	return &FriendSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendSearchLogic) FriendSearch(req *types.SearchRequest) (resp *types.SearchResponse, err error) {

	user_id := l.ctx.Value(contexts.ContextKeyUserID).(uint)
	// 所有用户
	user_list, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.UserConfModel{
		Online: req.Online,
	}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Joins:   "left join user_models um on um.id = user_conf_models.user_id",
		Preload: []string{"UserModel"},
		Where:   l.svcCtx.DB.Where("(user_conf_models.search_user <> 0 or user_conf_models.search_user is not null)   and (user_conf_models.search_user = 1 and um.id = ?)  or (user_conf_models.search_user = 2 and ( um.id = ? or um.nickname like ? ))", req.Key, req.Key, fmt.Sprintf("%%%s%%", req.Key)),
	})

	// 查自己这个用户的好友列表
	var my_friendlist []user_models.FriendModel
	l.svcCtx.DB.Find(&my_friendlist, "send_user_id = ? or rev_user_id = ?", user_id, user_id)
	userMap := map[uint]bool{}
	for _, friend := range my_friendlist {
		if friend.SendUserID == user_id {
			userMap[friend.RevUserID] = true
		} else {
			userMap[friend.SendUserID] = true
		}
	}

	search_info_list := make([]types.SearchInfo, 0)
	for _, uc := range user_list {
		search_info_list = append(search_info_list, types.SearchInfo{
			NickName: uc.UserModel.Nickname,
			Abstract: uc.UserModel.Abstract,
			Avatar:   uc.UserModel.Avatar,
			IsFriend: userMap[uc.UserID],
		})
	}
	return &types.SearchResponse{
		List:  search_info_list,
		Count: count,
	}, nil
}

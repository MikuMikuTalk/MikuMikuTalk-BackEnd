package admin

import (
	"context"

	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_group/group_rpc/types/group_rpc"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"
	"im_server/im_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListRequest) (resp *types.UserListResponse, err error) {
	list, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.UserModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Limit: req.Limit,
			Page:  req.Page,
			Key:   req.Key,
		},
		Likes: []string{"nickname", "ip"},
	})
	resp = new(types.UserListResponse)
	var userIDList []uint32
	for _, model := range list {
		userIDList = append(userIDList, uint32(model.ID))
	}
	// 去查用户在线状态
	var userOnlineMap = map[uint]bool{}
	userOnlineResponse, err1 := l.svcCtx.UserRpc.UserOnlineList(l.ctx, &user_rpc.UserOnlineListRequest{})
	if err1 == nil {
		for _, u := range userOnlineResponse.UserIdList {
			userOnlineMap[uint(u)] = true
		}
	} else {
		logx.Error(err1)
	}
	// 查用户创建的群聊个数
	groupResponse1, err2 := l.svcCtx.GroupRpc.UserGroupSearch(l.ctx, &group_rpc.UserGroupSearchRequest{
		UserIdList: userIDList,
		Mode:       1,
	})
	if err2 != nil {
		logx.Error(err2)
	}
	// 查用户加入的群聊个数
	groupResponse2, err3 := l.svcCtx.GroupRpc.UserGroupSearch(l.ctx, &group_rpc.UserGroupSearchRequest{
		UserIdList: userIDList,
		Mode:       2,
	})
	if err3 != nil {
		logx.Error(err3)
	}
	// 查用户发送的消息个数
	for _, model := range list {
		info := types.UserListInfoResponse{
			ID:              model.ID,
			CreatedAt:       model.CreatedAt.String(),
			Nickname:        model.Nickname,
			Avatar:          model.Avatar,
			IP:              model.IP,
			Addr:            model.Addr,
			IsOnline:        userOnlineMap[model.ID],
			GroupAdminCount: int(groupResponse1.Result[uint32(model.ID)]),
			GroupCount:      int(groupResponse2.Result[uint32(model.ID)]),
		}
		resp.List = append(resp.List, info)
	}

	resp.Count = count
	return
}

package logic

import (
	"context"
	"errors"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/set"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInfoLogic {
	return &GroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupInfoLogic) GroupInfo(req *types.GroupInfoRequest) (resp *types.GroupInfoResponse, err error) {
	// claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	var groupModel group_models.GroupModel
	err = l.svcCtx.DB.Take(&groupModel, "id = ?", req.ID).Error
	if err != nil {
		return nil, errors.New("群不存在")
	}
	// err = l.svcCtx.DB.Preload("MemberList").Take(&groupModel,my_id).Error
	// if err != nil {
	// 	return nil,errors.New("群不存在")
	// }

	resp = &types.GroupInfoResponse{
		GroupID:     req.ID,
		Title:       groupModel.Title,
		Abstract:    groupModel.Abstract,
		MemberCount: len(groupModel.MemberList),
		Avatar:      groupModel.Avatar,
	}
	//查询用户列表信息

	var userAllIDList []uint32
	for _, model := range groupModel.MemberList {
		//1 群主 2 管理员  3 普通成员
		userAllIDList = append(userAllIDList, uint32(model.UserID))
	}
	// 获取这些用户的基本信息 NickName Avatar
	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userAllIDList,
	})

	if err != nil {
		return
	}
	// 计算在线人数总数
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(context.Background(), &user_rpc.UserOnlineListRequest{})
	if err != nil {
		//所有的在线人数
		allOnlineUsersIDList := set.NewSet(userOnlineResponse.UserIdList)
		// 群里面所有的群成员id
		allUserIDList := set.NewSet(userAllIDList)
		//两个求交集，就能拿到群里面在线的人数了
		groupOnlineIDList := set.InterSet(allOnlineUsersIDList, allUserIDList)
		resp.MemberOnlineCount = groupOnlineIDList.Size()
	}

	/*
			type UserInfo struct {
			UserID   uint   `json:"userId"`
			Avatar   string `json:"avatar"`
			Nickname string `json:"nickname"`
		}
	*/
	var creator types.UserInfo
	var adminList []types.UserInfo
	for _, model := range groupModel.MemberList {
		userInfo := types.UserInfo{
			UserID:   model.UserID,
			Avatar:   userListResponse.UserInfo[uint32(model.UserID)].Avatar,
			Nickname: userListResponse.UserInfo[uint32(model.UserID)].NickName,
		}
		if model.Role == 1 {
			creator = userInfo
			continue
		}
		if model.Role == 2 {
			adminList = append(adminList, userInfo)
		}
	}
	resp.Creator = creator
	resp.AdminList = adminList

	return
}

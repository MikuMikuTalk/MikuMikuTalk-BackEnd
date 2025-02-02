package logic

import (
	"context"
	"errors"
	"fmt"

	"im_server/common/ctype"
	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_group/group_models"
	"im_server/im_user/user_rpc/types/user_rpc"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// GroupMemberLogic 结构体，用于处理群组成员相关的逻辑
type GroupMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGroupMemberLogic 创建一个新的 GroupMemberLogic 实例
func NewGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberLogic {
	return &GroupMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Data 结构体，用于映射数据库中的群组成员信息
type Data struct {
	GroupID        uint   `gorm:"column:group_id"`        // 群组ID
	UserID         uint   `gorm:"column:user_id"`         // 用户ID
	Role           int8   `gorm:"column:role"`            // 用户在群组中的角色
	CreatedAt      string `gorm:"column:created_at"`      // 用户加入群组的时间
	MemberNickname string `gorm:"column:member_nickname"` // 用户在群组中的昵称
	NewMsgDate     string `gorm:"column:new_msg_date"`    // 用户最新发言的时间
}

// GroupMember 处理获取群组成员列表的请求
func (l *GroupMemberLogic) GroupMember(req *types.GroupMemberRequest) (resp *types.GroupMemberResponse, err error) {
	// 检查排序参数是否合法
	switch req.Sort {
	case "new_msg_date desc", "new_msg_date asc": // 按照最新发言时间排序
	case "role asc": // 按照角色升序排序
	case "created_at desc", "created_at asc": // 按照加入群组时间排序
	default:
		return nil, errors.New("不支持的排序模式")
	}

	// 构造 SQL 查询语句，获取用户最新发言时间
	column := fmt.Sprintf(fmt.Sprintf("(select group_msg_models.created_at from group_msg_models  where group_member_models.group_id = %d  and group_msg_models.send_user_id = user_id) as new_msg_date", req.ID))

	// 使用 list_query.ListQuery 方法查询群组成员列表
	memberList, count, _ := list_query.ListQuery(l.svcCtx.DB, Data{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,  // 分页页码
			Limit: req.Limit, // 每页数量
			Sort:  req.Sort,  // 排序方式
		},
		Table: func() (string, any) {
			// 构造子查询，获取群组成员信息
			return "(?) as u", l.svcCtx.DB.Model(&group_models.GroupMemberModel{GroupID: req.ID}).
				Select("group_id",
					"user_id",
					"role",
					"created_at",
					"member_nickname",
					column)
		},
	})

	// 提取用户ID列表
	var userIDList []uint32
	for _, data := range memberList {
		userIDList = append(userIDList, uint32(data.UserID))
	}

	// 调用 UserRpc 服务获取用户信息
	userInfoMap := map[uint]ctype.UserInfo{}
	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{UserIdList: userIDList})
	if err == nil {
		// 将用户信息存入 map 中
		for u, info := range userListResponse.UserInfo {
			userInfoMap[uint(u)] = ctype.UserInfo{
				ID:       uint(u),
				NickName: info.NickName,
				Avatar:   info.Avatar,
			}
		}
	} else {
		logx.Error(err)
	}

	// 调用 UserRpc 服务获取在线用户列表
	userOnlineMap := map[uint]bool{}
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(l.ctx, &user_rpc.UserOnlineListRequest{})
	if err == nil {
		// 将在线用户信息存入 map 中
		for _, u := range userOnlineResponse.UserIdList {
			userOnlineMap[uint(u)] = true
		}
	} else {
		logx.Error(err)
	}

	// 构造响应数据
	resp = new(types.GroupMemberResponse)
	for _, data := range memberList {
		resp.List = append(resp.List, types.GroupMemberInfo{
			UserID:         data.UserID,                       // 用户ID
			UserNickname:   userInfoMap[data.UserID].NickName, // 用户昵称
			Avatar:         userInfoMap[data.UserID].Avatar,   // 用户头像
			IsOnline:       userOnlineMap[data.UserID],        // 用户是否在线
			Role:           data.Role,                         // 用户在群组中的角色
			MemberNickname: data.MemberNickname,               // 用户在群组中的昵称
			CreatedAt:      data.CreatedAt,                    // 用户加入群组的时间
			NewMsgDate:     data.NewMsgDate,                   // 用户最新发言的时间
		})
	}
	resp.Count = int(count) // 总成员数
	return
}

package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"
	"im_server/utils/set"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupCreateLogic) GroupCreate(req *types.GroupCreateRequest) (resp *types.GroupCreateResponse, err error) {
	token := req.Token
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return
	}
	my_id := claims.UserID
	var groupModel = group_models.GroupModel{
		Creator:      my_id,
		Abstract:     fmt.Sprintf("本群创建于%s:  群主很懒,什么都没有留下", time.Now().Format("2006-01-02")),
		IsSearch:     false,
		Verification: 2,
		Size:         50,
	}
	var groupUserList = []uint{my_id}
	switch req.Mode {
	case 1: //直接创建模式
		if req.Name == "" {
			return nil, errors.New("群名不可以为空")
		}
		if req.Size >= 1000 {
			return nil, errors.New("群人数不可以超过1000")
		}
		groupModel.Title = req.Name
		groupModel.Size = req.Size
		groupModel.IsSearch = req.IsSearch
	case 2: //选人创建模式
		if len(req.UserIDList) == 0 {
			return nil, errors.New("用户列表不可以为空")
		}
		var UserIDList = []uint32{uint32(my_id)} //把自己放进去
		for _, u := range req.UserIDList {
			UserIDList = append(UserIDList, uint32(u))
			groupUserList = append(groupUserList, u)
		}
		userFriendResponse, err := l.svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
			User: uint32(my_id),
		})
		var friendIDList []uint
		for _, i := range userFriendResponse.FriendList {
			friendIDList = append(friendIDList, uint(i.UserId))
		}
		// 判断它们两个是不是一致的,邀请人进群必须是自己的好友才行，不是自己的好友不能把人家拉入群
		slice := set.Difference(set.NewSet(req.UserIDList), set.NewSet(friendIDList))
		if slice.Size() != 0 {
			return nil, errors.New("选择的好友列表中有人不是你的好友")
		}
		userListResponse, err1 := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
			UserIdList: UserIDList,
		})
		if err1 != nil {
			logx.Error(err)
			return nil, errors.New("用户服务错误")
		}
		// 计算昵称长度,看到第几个人的时候长度会大于32
		var nameList []string
		for _, info := range userListResponse.UserInfo {
			if len(strings.Join(nameList, "、"))+len(info.NickName) >= 29 {
				break
			}
			nameList = append(nameList, info.NickName)
		}
		groupModel.Title = strings.Join(nameList, "、") + "的群聊"
	default:
		return nil, errors.New("不支持的模式")
	}
	//群头像
	// 1.默认头像 2.文字头像
	groupModel.Avatar = string([]rune(groupModel.Title)[0])
	err = l.svcCtx.DB.Create(&groupModel).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建群组失败")
	}
	var memebers []group_models.GroupMemberModel
	for i, u := range groupUserList {
		memberModel := group_models.GroupMemberModel{
			GroupID: groupModel.ID,
			UserID:  u,
			Role:    3,
		}
		if i == 0 {
			//设置为群主
			memberModel.Role = 1
		}
		memebers = append(memebers, memberModel)
	}
	err = l.svcCtx.DB.Create(&memebers).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("群成员添加失败")
	}
	return
}

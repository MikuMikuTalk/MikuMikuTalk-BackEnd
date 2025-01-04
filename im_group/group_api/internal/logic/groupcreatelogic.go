package logic

import (
	"context"
	"errors"
	"strings"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"

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
		IsSearch:     false,
		Verification: 2,
		Size:         50,
	}
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
	}
	//群头像
	// 1.默认头像 2.文字头像
	groupModel.Avatar = string([]rune(groupModel.Title)[0])
	err = l.svcCtx.DB.Create(&groupModel).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建群组失败")
	}
	return
}

package logic

import (
	"context"
	"errors"
	"fmt"

	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_api/internal/types"
	"im_server/im_chat/chat_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 最近聊天会话列表
func NewChatSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatSessionLogic {
	return &ChatSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Data struct {
	SU         uint   `gorm:"column:sU"`
	RU         uint   `gorm:"column:rU"`
	MaxDate    string `gorm:"column:maxDate"`
	MaxPreview string `gorm:"column:maxPreview"`
	IsTop      bool   `gorm:"column:isTop"`
}

func (l *ChatSessionLogic) ChatSession(req *types.ChatSessionRequest, token string) (resp *types.ChatSessionResponse, err error) {
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	my_id := claims.UserID
	column := fmt.Sprintf(" if((select 1 from top_user_models where user_id = %d and (top_user_id = sU or top_user_id = rU) limit 1), 1, 0)  as isTop", my_id)

	chatList, count, _ := list_query.ListQuery(l.svcCtx.DB, Data{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "isTop desc, maxDate desc",
		},
		Table: func() (string, any) {
			return "(?) as u", l.svcCtx.DB.Model(&chat_models.ChatModel{}).
				Select("least(send_user_id, rev_user_id) as sU",
					"greatest(send_user_id, rev_user_id) as rU",
					"max(created_at) as maxDate",
					fmt.Sprintf("(select msg_preview from chat_models  where ((send_user_id = sU and rev_user_id = rU) or (send_user_id = rU and rev_user_id = sU)) and id not in (select chat_id from user_chat_delete_models where user_id = %d) order by created_at desc  limit 1) as maxPreview", my_id),
					column).
				Where("(send_user_id = ? or rev_user_id = ?) and id not in (select chat_id from user_chat_delete_models where user_id = ?)", my_id, my_id, my_id).
				Group("least(send_user_id, rev_user_id)").
				Group("greatest(send_user_id, rev_user_id)")
		},
	})

	var userIDList []uint32
	// 获取所有用户ID
	for _, data := range chatList {
		// 排除自己
		if data.RU != my_id {
			userIDList = append(userIDList, uint32(data.RU))
		}
		// 排除自己
		if data.SU != my_id {
			userIDList = append(userIDList, uint32(data.SU))
		}
		if data.SU == my_id && my_id == data.RU {
			// 自己和自己聊
			userIDList = append(userIDList, uint32(my_id))
		}
	}

	response, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	if err != nil {
		logx.Error(err)
		return nil, errors.New("用户服务错误")
	}

	list := make([]types.ChatSession, 0)
	for _, data := range chatList {
		s := types.ChatSession{
			CreatedAt:  data.MaxDate,
			MsgPreview: data.MaxPreview,
			IsTop:      data.IsTop,
		}
		if data.RU != my_id {
			s.UserID = data.RU
			s.Avatar = response.UserInfo[uint32(s.UserID)].Avatar
			s.Nickname = response.UserInfo[uint32(s.UserID)].NickName
		}
		if data.SU != my_id {
			s.UserID = data.SU
			s.Avatar = response.UserInfo[uint32(s.UserID)].Avatar
			s.Nickname = response.UserInfo[uint32(s.UserID)].NickName
		}

		list = append(list, s)
	}

	return &types.ChatSessionResponse{List: list, Count: count}, nil
}

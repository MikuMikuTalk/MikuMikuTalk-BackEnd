package logic

import (
	"context"
	"errors"

	"im_server/common/ctype"
	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_api/internal/types"
	"im_server/im_chat/chat_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"
	"im_server/utils/list_util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 聊天记录接口
func NewChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryLogic {
	return &ChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type UserInfo struct {
	ID       uint   `json:"id"`
	NickName string `json:"username"`
	Avatar   string `json:"avatar"`
}

type ChatHistory struct {
	ID        uint             `json:"id"`
	SendUser  UserInfo         `json:"sendUser"`
	RevUser   UserInfo         `json:"revUser"`
	IsMe      bool             `json:"isMe"`       // 哪条消息是我发的
	CreatedAt string           `json:"created_at"` // 消息时间
	Msg       ctype.Msg        `json:"msg"`
	SystemMsg *ctype.SystemMsg `json:"systemMsg"`
}

type ChatHistoryResponse struct {
	List  []ChatHistory `json:"list"`
	Count int64         `json:"count"`
}

func (l *ChatHistoryLogic) ChatHistory(req *types.ChatHistoryRequest, token string) (resp *ChatHistoryResponse, err error) {
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		logx.Info(err)
		return
	}

	my_id := claims.UserID
	// 是否是好友
	res, err := l.svcCtx.UserRpc.IsFriend(context.Background(), &user_rpc.IsFriendRequest{
		User2: uint32(my_id),
		User1: uint32(req.FriendID),
	})
	if err != nil {
		return nil, err
	}
	if !res.IsFriend {
		return nil, errors.New("你们还不是好友呢")
	}
	chatList, count, _ := list_query.ListQuery(l.svcCtx.DB, chat_models.ChatModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
		//Debug: true,
		Where: l.svcCtx.DB.Where("((send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)) and id not in (select chat_id from user_chat_delete_models where user_id = ?)",
			my_id, req.FriendID, req.FriendID, my_id, my_id),
	})
	var userIDList []uint32
	for _, model := range chatList {
		userIDList = append(userIDList, uint32(model.SendUserID))
		userIDList = append(userIDList, uint32(model.RevUserID))
	}

	// 去重
	userIDList = list_util.DeduplicationList(userIDList)
	// 去调用户服务的rpc方法，获取用户信息 {用户id：{用户信息}}
	response, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	if err != nil {
		logx.Error(err)
		return nil, errors.New("用户服务错误")
	}

	list := make([]ChatHistory, 0)
	for _, model := range chatList {
		sendUser := UserInfo{
			ID:       model.SendUserID,
			NickName: response.UserInfo[uint32(model.SendUserID)].NickName,
			Avatar:   response.UserInfo[uint32(model.SendUserID)].Avatar,
		}
		revUser := UserInfo{
			ID:       model.RevUserID,
			NickName: response.UserInfo[uint32(model.RevUserID)].NickName,
			Avatar:   response.UserInfo[uint32(model.RevUserID)].Avatar,
		}
		info := ChatHistory{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.String(),
			SendUser:  sendUser,
			RevUser:   revUser,
			Msg:       model.Msg,
			SystemMsg: model.SystemMsg,
		}
		if info.SendUser.ID == my_id {
			info.IsMe = true
		}

		list = append(list, info)
	}

	resp = &ChatHistoryResponse{
		List:  list,
		Count: count,
	}
	return
}

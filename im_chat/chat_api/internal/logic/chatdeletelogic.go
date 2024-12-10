package logic

import (
	"context"
	"fmt"

	"im_server/im_chat/chat_models"
	"im_server/utils/jwts"

	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户聊天信息删除
func NewChatDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatDeleteLogic {
	return &ChatDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatDeleteLogic) ChatDelete(req *types.ChatDeleteRequest, token string) (resp *types.ChatDeleteResponse, err error) {
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	my_id := claims.UserID

	var chatList []chat_models.ChatModel
	l.svcCtx.DB.Find(&chatList, req.IdList)
	var userDeleteChatList []chat_models.UserChatDeleteModel
	l.svcCtx.DB.Find(&userDeleteChatList, req.IdList)
	chatDeleteMap := map[uint]struct{}{}
	for _, model := range userDeleteChatList {
		chatDeleteMap[model.ChatID] = struct{}{}
	}

	var deleteChatList []chat_models.UserChatDeleteModel

	if len(chatList) > 0 {
		for _, model := range chatList {
			if !(model.SendUserID == my_id || model.RevUserID == my_id) {
				// 不是自己的聊天记录
				fmt.Println("不是自己的聊天记录", model.ID)
				continue
			}
			// 已经删除过的记录
			_, ok := chatDeleteMap[model.ID]
			if ok {
				fmt.Println("已经删除过了", model.ID)
				continue
			}
			deleteChatList = append(deleteChatList, chat_models.UserChatDeleteModel{
				UserID: my_id,
				ChatID: model.ID,
			})
		}
	}
	if len(deleteChatList) > 0 {
		l.svcCtx.DB.Create(&deleteChatList)
	}
	logx.Infof("已删除聊天记录 %d 条", len(deleteChatList))
	return
}

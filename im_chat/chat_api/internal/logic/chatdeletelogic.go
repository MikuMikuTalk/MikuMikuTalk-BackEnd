package logic

import (
	"context"

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

func (l *ChatDeleteLogic) ChatDelete(req *types.ChatDeleteRequest) (resp *types.ChatDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

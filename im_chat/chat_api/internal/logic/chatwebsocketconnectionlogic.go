package logic

import (
	"context"

	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatWebsocketConnectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// websocket连接建立接口
func NewChatWebsocketConnectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatWebsocketConnectionLogic {
	return &ChatWebsocketConnectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatWebsocketConnectionLogic) ChatWebsocketConnection(req *types.ChatRequest) (resp *types.ChatResponse, err error) {
	return
}

// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"im_server/im_chat/chat_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 用户聊天信息删除
				Method:  http.MethodDelete,
				Path:    "/chat",
				Handler: chatDeleteHandler(serverCtx),
			},
			{
				// 聊天记录接口
				Method:  http.MethodGet,
				Path:    "/history",
				Handler: chatHistoryHandler(serverCtx),
			},
			{
				// 最近聊天会话列表
				Method:  http.MethodGet,
				Path:    "/session",
				Handler: chatSessionHandler(serverCtx),
			},
			{
				// 好友置顶
				Method:  http.MethodPost,
				Path:    "/user_top",
				Handler: userTopHandler(serverCtx),
			},
			{
				// websocket连接建立接口
				Method:  http.MethodGet,
				Path:    "/ws/chat",
				Handler: chatWebsocketConnectionHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/chat"),
	)
}

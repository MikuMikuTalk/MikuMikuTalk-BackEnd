// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"im_server/im_user/user_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 好友添加接口
				Method:  http.MethodPost,
				Path:    "/api/user/add",
				Handler: AddUserHandler(serverCtx),
			},
			{
				// 好友信息接口
				Method:  http.MethodGet,
				Path:    "/api/user/friend_info",
				Handler: FriendInfoHandler(serverCtx),
			},
			{
				// 好友列表获取
				Method:  http.MethodGet,
				Path:    "/api/user/friend_list",
				Handler: FriendListHandler(serverCtx),
			},
			{
				// 好友备注修改
				Method:  http.MethodPut,
				Path:    "/api/user/friends",
				Handler: FriendNoticeUpdateHandler(serverCtx),
			},
			{
				// 用户信息获取接口
				Method:  http.MethodGet,
				Path:    "/api/user/info",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				// 用户信息更新接口
				Method:  http.MethodPut,
				Path:    "/api/user/info",
				Handler: UserInfoUpdateHandler(serverCtx),
			},
			{
				// 好友搜索接口
				Method:  http.MethodGet,
				Path:    "/api/user/search",
				Handler: FriendSearchHandler(serverCtx),
			},
			{
				// 好友验证接口
				Method:  http.MethodPost,
				Path:    "/api/user/valid",
				Handler: UserValidHandler(serverCtx),
			},
		},
	)
}

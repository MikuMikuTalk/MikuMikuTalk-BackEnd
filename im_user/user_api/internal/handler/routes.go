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
		},
	)
}

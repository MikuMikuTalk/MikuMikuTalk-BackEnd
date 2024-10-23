package handler

import (
	"net/http"

	"im_server/common/response"
	"im_server/im_auth/auth_api/internal/logic"
	"im_server/im_auth/auth_api/internal/svc"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication()
		response.Response(r, w, resp, err)
		//❌就响应错误信息，使用common包中的response进行数据封装
	}
}

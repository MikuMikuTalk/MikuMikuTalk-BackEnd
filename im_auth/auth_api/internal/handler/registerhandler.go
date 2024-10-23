package handler

import (
	"net/http"

	"im_server/common/response"
	"im_server/im_auth/auth_api/internal/logic"
	"im_server/im_auth/auth_api/internal/svc"
	"im_server/im_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			//解析请求体失败
			//❌就响应错误信息，使用common包中的response进行数据封装
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Response(r, w, resp, err)
	}
}

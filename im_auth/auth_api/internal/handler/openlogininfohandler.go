package handler

import (
	"net/http"

	"im_server/common/response"
	"im_server/im_auth/auth_api/internal/logic"
	"im_server/im_auth/auth_api/internal/svc"
)

func open_login_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewOpen_login_infoLogic(r.Context(), svcCtx)
		resp, err := l.Open_login_info()
		// 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		response.Response(r, w, resp, err)
	}
}

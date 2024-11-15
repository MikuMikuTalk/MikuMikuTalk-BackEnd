package handler

import (
	"errors"
	"im_server/common/response"
	"im_server/im_user/user_api/internal/logic"
	"im_server/im_user/user_api/internal/types"
	"net/http"

	"im_server/im_user/user_api/internal/svc"
)

// 用户信息获取接口
func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		req.UserName = r.URL.Query().Get("username")
		if req.UserName == "" {
			response.Response(r, w, nil, errors.New("请在参数中传入username"))
			return
		}
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		// 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		response.Response(r, w, resp, err)
	}
}

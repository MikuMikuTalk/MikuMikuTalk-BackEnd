package handler

import (
	"im_server/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im_server/im_user/user_api/internal/logic"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
)

// 用户信息更新接口
func UserInfoUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoUpdateRequest
		//获取jwt token,提取userid
		token := r.Header.Get("Authorization")
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewUserInfoUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserInfoUpdate(token, &req)
		response.Response(r, w, resp, err)
	}
}

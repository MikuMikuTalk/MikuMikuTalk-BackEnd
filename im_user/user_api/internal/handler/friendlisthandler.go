package handler

import (
	"net/http"

	"im_server/common/response"
	"im_server/im_user/user_api/internal/logic"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 好友列表获取
func FriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		token := r.Header.Get("Authorization")
		l := logic.NewFriendListLogic(r.Context(), svcCtx)
		resp, err := l.FriendList(&req, token)
		response.Response(r, w, resp, err)
	}
}

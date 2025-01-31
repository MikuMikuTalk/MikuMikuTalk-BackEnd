package handler

import (
	"net/http"

	"im_server/common/response"
	"im_server/im_user/user_api/internal/logic"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 好友备注修改
func FriendNoticeUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendNoticeUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewFriendNoticeUpdateLogic(r.Context(), svcCtx)
		resp, err := l.FriendNoticeUpdate(&req)
		response.Response(r, w, resp, err)
	}
}

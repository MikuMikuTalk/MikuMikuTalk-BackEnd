package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im_server/im_user/user_api/internal/logic"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
)

// 好友备注修改
func FriendNoticeUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendNoticeUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFriendNoticeUpdateLogic(r.Context(), svcCtx)
		resp, err := l.FriendNoticeUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

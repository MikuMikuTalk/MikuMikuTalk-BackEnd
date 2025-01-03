package handler

import (
	"net/http"

	"im_server/common/response"
	"im_server/im_chat/chat_api/internal/logic"
	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func chatDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			//解析请求体失败
			//❌就响应错误信息，使用common包中的response进行数据封装
			response.Response(r, w, nil, err)
			return
		}
		token := r.Header.Get("Authorization")
		l := logic.NewChatDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ChatDelete(&req, token)
		// 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		response.Response(r, w, resp, err)
	}
}
package handler

import (
	"context"
	"fmt"
	"net/http"

	"im_server/common/response"
	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_api/internal/types"
	"im_server/im_user/user_rpc/types/user_rpc"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type UserInfo struct {
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	UserID   uint   `json:"userID"`
}

type UserWsInfo struct {
	UserInfo UserInfo        // 用户信息
	Conn     *websocket.Conn // 用户的ws连接对象
}

var UserWsMap = map[uint]UserWsInfo{}

func chatWebsocketConnectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		upGrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// 鉴权 true表示放行，false表示拦截
				return true
			},
		}

		conn, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		defer func() {
			conn.Close()
			delete(UserWsMap, req.UserID)
		}()
		// 调取用户服务，获取当前用户信息
		res, err := svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
			UserIdList: []uint32{uint32(req.UserID)},
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		var userWsInfo = UserWsInfo{
			UserInfo: UserInfo{
				UserID:   req.UserID,
				Avatar:   res.UserInfo[uint32(req.UserID)].Avatar,
				NickName: res.UserInfo[uint32(req.UserID)].NickName,
			},
			Conn: conn,
		}
		UserWsMap[req.UserID] = userWsInfo
		logx.Info("UserWsMap: ", UserWsMap)
		logx.Info("userWsInfo: ", userWsInfo)

		for {
			// 消息类型，消息，错误
			_, p, err := conn.ReadMessage()
			if err != nil {
				// 用户断开聊天
				fmt.Println(err)
				break
			}
			fmt.Println(string(p))
			// 发送消息
			conn.WriteMessage(websocket.TextMessage, []byte("xxx"))
		}
		//l := logic.NewChatWebsocketConnectionLogic(r.Context(), svcCtx)
		//resp, err := l.ChatWebsocketConnection(&req)
		//// 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		//response.Response(r, w, resp, err)
	}
}

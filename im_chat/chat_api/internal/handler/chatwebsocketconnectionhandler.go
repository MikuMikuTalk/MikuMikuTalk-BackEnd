package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"im_server/common/response"
	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_api/internal/types"
	"im_server/im_user/user_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type UserWsInfo struct {
	UserInfo user_models.UserModel // 用户信息
	Conn     *websocket.Conn       // 用户的ws连接对象
}

var UserWsMap = map[uint]UserWsInfo{}

func chatWebsocketConnectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		token := r.Header.Get("Authorization")
		claims, err := jwts.ParseToken(token, svcCtx.Config.Auth.AuthSecret)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		my_id := claims.UserID
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
		// 调取用户服务，获取好友信息
		res, err := svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
			UserId: uint32(req.UserID),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		var userInfo user_models.UserModel
		err = json.Unmarshal(res.Data, &userInfo)
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		var userWsInfo = UserWsInfo{
			UserInfo: userInfo,
			Conn:     conn,
		}
		UserWsMap[req.UserID] = userWsInfo
		// 把在线用户存入redis,如果在线，存入redis, key: online_user, field: 用户id, value: 用户id

		svcCtx.Redis.HSet("online_user", fmt.Sprintf("%d", req.UserID), req.UserID)
		logx.Info("UserWsMap: ", UserWsMap)
		logx.Info("userWsInfo: ", userWsInfo)
		/*
		   // 遍历在线的用户， 和当前这个人是好友的，就给他发送好友在线

		   		// 先把所有在线的用户id取出来，以及待确认的用户id，然后传到用户rpc服务中
		   		// [1,2,3]  3
		   		// 在rpc服务中，去判断哪些用户是好友关系

		   		//if userInfo.UserConfModel.FriendOnline {
		   		// 如果用户开启了好友上线提醒
		   		// 查一下自己的好友是不是上线了
		*/

		friendRes, err := svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
			User: uint32(my_id),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}

		for _, info := range friendRes.FriendList {
			logx.Info(info)
			friend, ok := UserWsMap[uint(info.UserId)]
			if ok {
				text := fmt.Sprintf("好友 %s 上线了", UserWsMap[req.UserID].UserInfo.Nickname)
				logx.Info(text)
				if friend.UserInfo.UserConfModel.FriendOnline {
					//判断好友是否开了好友上线提醒
					friend.Conn.WriteMessage(websocket.TextMessage, []byte("好友上线了"))
				}
			}
		}
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

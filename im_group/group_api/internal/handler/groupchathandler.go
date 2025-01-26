package handler

import (
	"context"
	"encoding/json"
	"im_server/common/ctype"
	"im_server/common/response"
	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type UserWsInfo struct {
	UserInfo    ctype.UserInfo             //用户信息
	WsClientMap map[string]*websocket.Conn // 这个用户管理的所有ws客户端
}

var UserOnlineWsMap = map[uint]*UserWsInfo{}

type ChatRequest struct {
	GroupID uint      `json:"groupID"` // 群id
	Msg     ctype.Msg `json:"msg"`     // 消息
}

type ChatResponse struct {
	UserID       uint      `json:"userID"`
	UserNickname string    `json:"userNickname"`
	UserAvatar   string    `json:"userAvatar"`
	Msg          ctype.Msg `json:"msg"`
	ID           uint      `json:"id"`
	MsgType      int8      `json:"msgType"`
	CreatedAt    time.Time `json:"createdAt"`
	IsMe         bool      `json:"isMe"`
}

func groupChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			//解析请求体失败
			//❌就响应错误信息，使用common包中的response进行数据封装
			response.Response(r, w, nil, err)
			return
		}

		claims, err := jwts.ParseToken(svcCtx.Config.Auth.AuthSecret, req.Token)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		userID := claims.UserID
		var upGrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		addr := conn.RemoteAddr().String()
		logx.Infof("用户建立ws连接%s", addr)
		defer func() {
			conn.Close()
			userWsInfo, ok := UserOnlineWsMap[userID]
			if ok {
				// 删除退出的ws信息
				delete(userWsInfo.WsClientMap, addr)
			}
			if userWsInfo != nil && len(userWsInfo.WsClientMap) == 0 {
				delete(UserOnlineWsMap, userID)
			}
		}()
		baseResponse, err := svcCtx.UserRpc.UserBaseInfo(context.Background(), &user_rpc.UserBaseInfoRequest{
			UserId: uint32(userID),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		userInfo := ctype.UserInfo{
			ID:       userID,
			NickName: baseResponse.NickName,
			Avatar:   baseResponse.Avatar,
		}
		userWsInfo, ok := UserOnlineWsMap[userID]
		if !ok {
			userWsInfo = &UserWsInfo{
				UserInfo: userInfo,
				WsClientMap: map[string]*websocket.Conn{
					addr: conn,
				},
			}
			UserOnlineWsMap[userID] = userWsInfo
		}
		_, ok1 := userWsInfo.WsClientMap[addr]
		if !ok1 {
			//用户第二次连接
			UserOnlineWsMap[userID].WsClientMap[addr] = conn
		}

		for {
			_, p, err1 := conn.ReadMessage()
			if err1 != nil {
				logx.Error(err)
				break
			}
			var request ChatRequest
			err = json.Unmarshal(p, &request)
			if err != nil {
				SendTipErrMsg(conn, "参数解析失败")
				continue
			}
			//判断自己是不是这个群的成员
			var member group_models.GroupMemberModel
			err = svcCtx.DB.Take(&member, "group_id = ? and user_id = ? ", request.GroupID, userID).Error
			if err != nil {
				//自己不是这个群的群成员
				SendTipErrMsg(conn, "你还不是这个群的群成员")
				continue
			}
			//查在线的用户列表
			userOnlineIDList := getOnlineUserIDList()

			//查这个群的成员 并且这个成员在线
			var groupMemberOnlineIDList []uint
			svcCtx.DB.Model(group_models.GroupMemberModel{}).Where("group_id = ? and user_id in ?", request.GroupID, userOnlineIDList).Select("user_id").Scan(&groupMemberOnlineIDList)
			// 遍历这个用户列表，找ws客户单
			for _, u := range groupMemberOnlineIDList {
				wsUserInfo, ok2 := UserOnlineWsMap[u]
				if !ok2 {
					continue
				}
				for _, w2 := range wsUserInfo.WsClientMap {
					w2.WriteMessage(websocket.TextMessage, []byte(""))
				}
			}
			logx.Info("message: ", string(p))
		}
		// l := logic.NewGroupChatLogic(r.Context(), svcCtx)
		// resp, err := l.GroupChat(&req)
		// // 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		// response.Response(r, w, resp, err)

	}
}

func getOnlineUserIDList() []uint {
	var userOnlineIDList []uint = make([]uint, 0)
	for u, _ := range UserOnlineWsMap {
		userOnlineIDList = append(userOnlineIDList, u)
	}
	return userOnlineIDList
}

func SendTipErrMsg(conn *websocket.Conn, msg string) {
	resp := ChatResponse{
		Msg: ctype.Msg{
			Type: ctype.TipMsgType,
			TipMsg: &ctype.TipMsg{
				Status:  "error",
				Content: msg,
			},
		},
		CreatedAt: time.Now(),
	}
	byteData, _ := json.Marshal(resp)
	conn.WriteMessage(websocket.TextMessage, byteData)
}

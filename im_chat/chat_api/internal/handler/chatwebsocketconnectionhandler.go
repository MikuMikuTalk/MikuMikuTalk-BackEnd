package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"im_server/common/ctype"
	"im_server/common/response"
	"im_server/im_chat/chat_api/internal/svc"
	"im_server/im_chat/chat_models"
	"im_server/im_file/file_rpc/types/file_rpc"
	"im_server/im_user/user_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// UserWsInfo表示已连接用户的WebSocket连接和用户信息。
type UserWsInfo struct {
	UserInfo user_models.UserModel
	Conn     *websocket.Conn
}

// ChatRequest表示一个用户发送给另一个用户的聊天消息。
type ChatRequest struct {
	RevUserID uint      `json:"revUserID"` // Receiver User ID
	Msg       ctype.Msg `json:"msg"`
}

// ChatResponse表示对聊天消息的响应。
type ChatResponse struct {
	RevUser   ctype.UserInfo `json:"revUser"`
	SendUser  ctype.UserInfo `json:"sendUser"`
	Msg       ctype.Msg      `json:"msg"`
	CreatedAt time.Time      `json:"createdAt"`
}

var UserWsMap = make(map[uint]UserWsInfo) // 全局映射以存储用户WebSocket连接。

func chatWebsocketConnectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析用户ID
		token := r.Header.Get("Authorization")
		claims, err := jwts.ParseToken(token, svcCtx.Config.Auth.AuthSecret)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		myID := claims.UserID

		// 升级HTTP连接为WebSocket
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		defer func() {
			// 关闭连接时删除用户的WebSocket连接
			conn.Close()
			// 删除用户的WebSocket连接
			delete(UserWsMap, myID)
			// 删除在线用户
			svcCtx.Redis.HDel("online_user", fmt.Sprintf("%d", myID))
		}()

		// 获取我的用户信息
		resMine, err := svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
			UserId: uint32(myID),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		var userInfoMine user_models.UserModel
		if err := json.Unmarshal(resMine.Data, &userInfoMine); err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}

		// 存储我的WebSocket连接
		UserWsMap[myID] = UserWsInfo{
			UserInfo: userInfoMine,
			Conn:     conn,
		}
		// 存储在线用户
		svcCtx.Redis.HSet("online_user", fmt.Sprintf("%d", myID), myID)

		// 通知我的好友我的在线状态
		friendRes, err := svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
			User: uint32(myID),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		for _, friend := range friendRes.FriendList {
			friendWs, ok := UserWsMap[uint(friend.UserId)]
			if ok && friendWs.UserInfo.UserConfModel.FriendOnline {
				text := fmt.Sprintf("好友 %s 上线了", userInfoMine.Nickname)
				friendWs.Conn.WriteMessage(websocket.TextMessage, []byte(text))
			}
		}

		// 处理WebSocket消息
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				logx.Error("WebSocket read error: ", err)
				break
			}

			var chatReq ChatRequest
			if err := json.Unmarshal(message, &chatReq); err != nil {
				errorMsg := fmt.Sprintf("消息格式错误: %s", err.Error())
				// 发送错误消息
				SendTipErrMsg(conn, errorMsg)
				logx.Error(errorMsg)
				conn.WriteMessage(websocket.TextMessage, []byte(errorMsg))
				continue
			}
			//检查接收者是否是好友
			if myID != chatReq.RevUserID {
				res, err := svcCtx.UserRpc.IsFriend(context.Background(), &user_rpc.IsFriendRequest{
					User2: uint32(myID),
					User1: uint32(chatReq.RevUserID),
				})
				if err != nil {
					// 用户乱发消息
					logx.Error("用户服务错误: ", err)
					conn.WriteMessage(websocket.TextMessage, []byte("用户服务错误"))
					continue
				}
				//如果不是好友，返回不是好友的消息
				if !res.GetIsFriend() {
					errorMsg := fmt.Sprintf("%v 和 %v 还不是好友呢", myID, chatReq.RevUserID)
					SendTipErrMsg(conn, errorMsg)
					logx.Error(errorMsg)
					conn.WriteMessage(websocket.TextMessage, []byte(errorMsg))
					continue
				}
			}
			//检查请求的类型，如果是文件类型，就调用文件rpc服务，获取文件相关信息
			switch chatReq.Msg.Type {
			case ctype.FileMsgType:
				//如果是文件类型，就要去请求文件rpc服务
				nameList := strings.Split(chatReq.Msg.FileMsg.Src, ".")
				if len(nameList) == 0 {
					SendTipErrMsg(conn, "请上传文件")
					continue
				}
				fileID := nameList[len(nameList)-1]
				fileResponse, err := svcCtx.FileRpc.FileInfo(context.Background(), &file_rpc.FileInfoRequest{
					FildId: fileID,
				})
				if err != nil {
					logx.Error(err)
					SendTipErrMsg(conn, err.Error())
					continue
				}
				chatReq.Msg.FileMsg.Title = fileResponse.FileName
				chatReq.Msg.FileMsg.Size = fileResponse.FileSize
				chatReq.Msg.FileMsg.Type = fileResponse.FileType
			}

			// 先入库
			InsertMsgByChat(svcCtx.DB, chatReq.RevUserID, myID, chatReq.Msg)
			SendMsgByUser(chatReq.RevUserID, myID, chatReq.Msg)

		}
	}
}

// 发送错误提示的消息
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

// InsertMsgByChat 消息入库
func InsertMsgByChat(db *gorm.DB, revUserID uint, sendUserID uint, msg ctype.Msg) {
	chatModel := chat_models.ChatModel{
		SendUserID: sendUserID,
		RevUserID:  revUserID,
		MsgType:    msg.Type,
		Msg:        msg,
	}
	chatModel.MsgPreview = chatModel.MsgPreviewMethod()
	err := db.Create(&chatModel).Error
	if err != nil {
		logx.Error(err)
		sendUser, ok := UserWsMap[sendUserID]
		if !ok {
			return
		}
		SendTipErrMsg(sendUser.Conn, "消息保存失败")
	}
}

// SendMsgByUser 发消息，给谁发，谁发的
func SendMsgByUser(revUserID uint, sendUserID uint, msg ctype.Msg) {
	revUser, ok := UserWsMap[revUserID]
	if !ok {
		return
	}
	sendUser, ok := UserWsMap[sendUserID]
	if !ok {
		return
	}
	resp := ChatResponse{
		RevUser: ctype.UserInfo{
			ID:       revUserID,
			NickName: revUser.UserInfo.Nickname,
			Avatar:   revUser.UserInfo.Avatar,
		},
		SendUser: ctype.UserInfo{
			ID:       sendUserID,
			NickName: sendUser.UserInfo.Nickname,
			Avatar:   sendUser.UserInfo.Avatar,
		},
		Msg:       msg,
		CreatedAt: time.Now(),
	}
	byteData, _ := json.Marshal(resp)
	revUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

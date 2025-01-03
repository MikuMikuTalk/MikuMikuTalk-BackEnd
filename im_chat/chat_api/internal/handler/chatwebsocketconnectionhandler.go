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

var UserOnlineMap = make(map[uint]UserWsInfo) // 全局映射以存储用户WebSocket连接。

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
			delete(UserOnlineMap, myID)
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
		UserOnlineMap[myID] = UserWsInfo{
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
			friendWs, ok := UserOnlineMap[uint(friend.UserId)]
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
			// 检查接收者是否是好友
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
				// 如果不是好友，返回不是好友的消息
				if !res.GetIsFriend() {
					errorMsg := fmt.Sprintf("%v 和 %v 还不是好友呢", myID, chatReq.RevUserID)
					SendTipErrMsg(conn, errorMsg)
					logx.Error(errorMsg)
					conn.WriteMessage(websocket.TextMessage, []byte(errorMsg))
					continue
				}
			}
			// 检查请求的类型，如果是文件类型，就调用文件rpc服务，获取文件相关信息
			switch chatReq.Msg.Type {
			case ctype.TextMsgType:
				if chatReq.Msg.TextMsg == nil {
					SendTipErrMsg(conn, "请输入内容")
					continue
				}
				if chatReq.Msg.TextMsg.Content == "" {
					SendTipErrMsg(conn, "请输入内容")
					logx.Error("请输入内容")
					continue
				}
			case ctype.FileMsgType:
				// 如果是文件类型，就要去请求文件rpc服务
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
			case ctype.WithdrawMsgType:
				//撤回消息id必填
				if chatReq.Msg.WithdrawMsg.MsgID == 0 {
					SendTipErrMsg(conn, "撤回消息id必填")
					continue
				}
				//自己只能撤回自己的消息
				var msgModel chat_models.ChatModel
				//查看消息是否存在
				if err = svcCtx.DB.Take(&msgModel, chatReq.Msg.WithdrawMsg.MsgID).Error; err != nil {
					SendTipErrMsg(conn, "消息不存在，无法撤回")
					continue
				}
				// 判断是不是自己发的
				if msgModel.SendUserID != myID {
					//如果不是，提醒用户只能撤回自己的消息
					SendTipErrMsg(conn, "只能撤回自己的消息")
					continue
				}
				//判断消息时间，如果超过三分钟，就提示不能撤回了
				now := time.Now()
				subTime := now.Sub(msgModel.CreatedAt)
				if subTime >= time.Minute*3 {
					SendTipErrMsg(conn, "超过三分钟的消息不能被撤回")
					continue
				}
				//撤回逻辑
				var content string = fmt.Sprintf("%s 撤回了一条消息", userInfoMine.Nickname)
				if userInfoMine.UserConfModel.RecallMessage != nil {
					content = *userInfoMine.UserConfModel.RecallMessage
				}
				svcCtx.DB.Model(&msgModel).Updates(chat_models.ChatModel{
					Msg: ctype.Msg{
						Type: ctype.WithdrawMsgType,
						WithdrawMsg: &ctype.WithdrawMsg{
							Content:   content,
							MsgID:     chatReq.Msg.WithdrawMsg.MsgID,
							OriginMsg: &msgModel.Msg,
						},
					},
				})

			}

			// 先入库
			InsertMsgByChat(svcCtx.DB, chatReq.RevUserID, myID, chatReq.Msg)
			// 发送消息给好友
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
	switch msg.Type {
	case ctype.WithdrawMsgType:
		logx.Info("撤回消息不入库")
		return
	}
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
		sendUser, ok := UserOnlineMap[sendUserID]
		if !ok {
			return
		}
		SendTipErrMsg(sendUser.Conn, "消息保存失败")
	}
}

// SendMsgByUser 发消息，给谁发，谁发的
func SendMsgByUser(revUserID uint, sendUserID uint, msg ctype.Msg) {
	revUser, ok1 := UserOnlineMap[revUserID]

	sendUser, ok2 := UserOnlineMap[sendUserID]
	resp := ChatResponse{
		Msg:       msg,
		CreatedAt: time.Now(),
	}
	if ok1 && ok2 && sendUserID == revUserID {
		// 自己给自己发消息
		resp.RevUser = ctype.UserInfo{
			ID:       revUserID,
			NickName: revUser.UserInfo.Nickname,
			Avatar:   revUser.UserInfo.Avatar,
		}
		resp.SendUser = ctype.UserInfo{
			ID:       sendUserID,
			NickName: sendUser.UserInfo.Nickname,
			Avatar:   sendUser.UserInfo.Avatar,
		}
		byteData, _ := json.Marshal(resp)
		revUser.Conn.WriteMessage(websocket.TextMessage, byteData)
		return
	}

	//在线情况下，可以拿到对方用户信息
	//对方不在线的情况下，只能通过调用用户信息的Rpc方法获取用户信息
	if ok1 {
		//接收者在线
		resp.RevUser = ctype.UserInfo{
			ID:       revUserID,
			NickName: revUser.UserInfo.Nickname,
			Avatar:   revUser.UserInfo.Avatar,
		}
		byteData, _ := json.Marshal(resp)
		revUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}

	if ok2 {
		//发送者也在线
		resp.SendUser = ctype.UserInfo{
			ID:       sendUserID,
			NickName: sendUser.UserInfo.Nickname,
			Avatar:   sendUser.UserInfo.Avatar,
		}
		byteData, _ := json.Marshal(resp)
		sendUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}

}

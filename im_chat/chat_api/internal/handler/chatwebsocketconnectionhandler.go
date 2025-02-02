package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"im_server/common/ctype"
	"im_server/common/response"
	"im_server/common/service/redis_cache"
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
	UserInfo    user_models.UserModel
	WsClientMap map[string]*websocket.Conn // 这个用户管理的所有ws客户端
	CurrentConn *websocket.Conn            // 当前的连接对象
}

// ChatRequest表示一个用户发送给另一个用户的聊天消息。
type ChatRequest struct {
	RevUserID uint      `json:"revUserID"` // Receiver User ID
	Msg       ctype.Msg `json:"msg"`
}

// ChatResponse表示对聊天消息的响应。
type ChatResponse struct {
	ID        uint           `json:"id"`
	IsMe      bool           `json:"isMe"`
	RevUser   ctype.UserInfo `json:"revUser"`
	SendUser  ctype.UserInfo `json:"sendUser"`
	Msg       ctype.Msg      `json:"msg"`
	CreatedAt time.Time      `json:"created_at"`
}

var UserOnlineWsMap = map[uint]*UserWsInfo{}

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
		addr := conn.RemoteAddr().String()
		defer func() {
			// 关闭连接时删除用户的WebSocket连接
			conn.Close()
			UserWsInfo, ok := UserOnlineWsMap[myID]
			if ok {
				// 删除退出的ws连接
				delete(UserWsInfo.WsClientMap, addr)
			}
			if UserWsInfo != nil && len(UserWsInfo.WsClientMap) == 0 {
				// 代表这个用户最后一个连接断开了
				delete(UserOnlineWsMap, myID)
				svcCtx.Redis.Hdel("online_user", fmt.Sprintf("%d", myID))
			}
		}()

		// 获取我的用户信息
		resMine, err := svcCtx.UserRpc.UserInfo(r.Context(), &user_rpc.UserInfoRequest{
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

		userWsInfo, ok := UserOnlineWsMap[myID]
		if !ok {
			userWsInfo = &UserWsInfo{
				UserInfo: userInfoMine,
				WsClientMap: map[string]*websocket.Conn{
					addr: conn, // 当前的连接对象
				},
				CurrentConn: conn,
			}
			// 代表这个用户第一次连接服务器
			UserOnlineWsMap[myID] = userWsInfo
			// 存储在线用户
			svcCtx.Redis.Hset("online_user", fmt.Sprintf("%d", myID), fmt.Sprintf("%d", myID))
		}
		// 检查是否已经存在相同的连接
		if userWsInfo.WsClientMap[addr] == nil {
			// 相当于在另一个ip登录,允许这个ip登入
			logx.Info(fmt.Sprintf("%v 再另一台设备登录了", addr))
			UserOnlineWsMap[myID].WsClientMap[addr] = conn
			UserOnlineWsMap[myID].CurrentConn = conn
		}

		// 通知我的好友我的在线状态
		friendRes, err := svcCtx.UserRpc.FriendList(r.Context(), &user_rpc.FriendListRequest{
			User: uint32(myID),
		})
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		for _, friend := range friendRes.FriendList {
			friendWs, ok := UserOnlineWsMap[uint(friend.UserId)]
			if ok && friendWs.UserInfo.UserConfModel.FriendOnline {
				text := fmt.Sprintf("好友 %s 上线了", userInfoMine.Nickname)
				sendWsMapMsg(friendWs.WsClientMap, []byte(text))
				// friendWs.CurrentConn.WriteMessage(websocket.TextMessage, []byte(text))
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
				sendWsMapMsg(userWsInfo.WsClientMap, []byte(errorMsg))
				// conn.WriteMessage(websocket.TextMessage, []byte(errorMsg))
				continue
			}
			// 检查接收者是否是好友
			if myID != chatReq.RevUserID {
				res, err := svcCtx.UserRpc.IsFriend(r.Context(), &user_rpc.IsFriendRequest{
					User2: uint32(myID),
					User1: uint32(chatReq.RevUserID),
				})
				if err != nil {
					// 用户乱发消息
					logx.Error("用户服务错误: ", err)
					sendWsMapMsg(userWsInfo.WsClientMap, []byte("用户服务错误"))
					// conn.WriteMessage(websocket.TextMessage, []byte("用户服务错误"))
					continue
				}
				// 如果不是好友，返回不是好友的消息
				if !res.GetIsFriend() {
					errorMsg := fmt.Sprintf("%v 和 %v 还不是好友呢", myID, chatReq.RevUserID)
					SendTipErrMsg(conn, errorMsg)
					logx.Error(errorMsg)
					sendWsMapMsg(userWsInfo.WsClientMap, []byte(errorMsg))
					// conn.WriteMessage(websocket.TextMessage, []byte(errorMsg))
					continue
				}
			}
			// 判断type  1 - 12 检查消息类型是否正确
			if !(chatReq.Msg.Type >= 1 && chatReq.Msg.Type <= 12) {
				SendTipErrMsg(conn, "消息类型错误")
				continue
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
				if chatReq.Msg.FileMsg == nil {
					SendTipErrMsg(conn, "请上传文件")
					return
				}
				// 如果是文件类型，就要去请求文件rpc服务
				nameList := strings.Split(chatReq.Msg.FileMsg.Src, ".")
				if len(nameList) == 0 {
					SendTipErrMsg(conn, "请上传文件")
					continue
				}
				fileID := nameList[len(nameList)-1]
				fileResponse, err := svcCtx.FileRpc.FileInfo(r.Context(), &file_rpc.FileInfoRequest{
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
				if chatReq.Msg.WithdrawMsg == nil {
					SendTipErrMsg(conn, "撤回消息id必填")
					continue
				}
				// 撤回消息id必填
				if chatReq.Msg.WithdrawMsg.MsgID == 0 {
					SendTipErrMsg(conn, "撤回消息id必填")
					continue
				}
				// 自己只能撤回自己的消息
				var msgModel chat_models.ChatModel
				// 查看消息是否存在
				if err = svcCtx.DB.Take(&msgModel, chatReq.Msg.WithdrawMsg.MsgID).Error; err != nil {
					SendTipErrMsg(conn, "消息不存在，无法撤回")
					continue
				}
				// 如果已经是撤回消息就不能再撤回了
				if msgModel.Msg.Type == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "不能撤回已经撤回的消息")
					continue
				}
				// 判断是不是自己发的
				if msgModel.SendUserID != myID {
					// 如果不是，提醒用户只能撤回自己的消息
					SendTipErrMsg(conn, "只能撤回自己的消息")
					continue
				}
				// 判断消息时间，如果超过三分钟，就提示不能撤回了
				now := time.Now()
				subTime := now.Sub(msgModel.CreatedAt)
				if subTime >= time.Minute*3 {
					SendTipErrMsg(conn, "超过三分钟的消息不能被撤回")
					continue
				}
				// 撤回逻辑
				var content string = fmt.Sprintf("%s 撤回了一条消息", userInfoMine.Nickname)
				if userInfoMine.UserConfModel.RecallMessage != nil {
					content = *userInfoMine.UserConfModel.RecallMessage
				}
				originMsg := msgModel.Msg
				originMsg.WithdrawMsg = nil // 这里可能会出现循环引用，所以拷贝了这个值，并且把撤回消息置空

				svcCtx.DB.Model(&msgModel).Updates(chat_models.ChatModel{
					MsgPreview: "[撤回消息] - " + content,
					MsgType:    ctype.WithdrawMsgType,
					Msg: ctype.Msg{
						Type: ctype.WithdrawMsgType,
						WithdrawMsg: &ctype.WithdrawMsg{
							Content:   content,
							MsgID:     chatReq.Msg.WithdrawMsg.MsgID,
							OriginMsg: &originMsg,
						},
					},
				})
			case ctype.ReplyMsgType:
				// 回复消息
				// 先校验
				if chatReq.Msg.ReplyMsg == nil || chatReq.Msg.ReplyMsg.MsgID == 0 {
					SendTipErrMsg(conn, "回复消息id必填")
					continue
				}
				// 找到回复的原消息
				var msgModel chat_models.ChatModel
				err = svcCtx.DB.Take(&msgModel, chatReq.Msg.ReplyMsg).Error
				if err != nil {
					SendTipErrMsg(conn, "消息不存在")
					continue
				}
				// 不能回复撤回消息
				if msgModel.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "该消息已撤回")
					continue
				}

				userBaseInfo, err := redis_cache.GetUserBaseInfo(svcCtx.Redis, svcCtx.UserRpc, msgModel.SendUserID)
				if err != nil {
					logx.Error(err)
					return
				}
				chatReq.Msg.ReplyMsg.Msg = &msgModel.Msg
				chatReq.Msg.ReplyMsg.UserID = msgModel.SendUserID
				chatReq.Msg.ReplyMsg.UserNickName = userBaseInfo.NickName
				chatReq.Msg.ReplyMsg.OriginMsgDate = msgModel.CreatedAt
			case ctype.QuoteMsgType:
				// 回复消息
				// 先校验
				if chatReq.Msg.QuoteMsg == nil || chatReq.Msg.QuoteMsg.MsgID == 0 {
					SendTipErrMsg(conn, "引用消息id必填")
					continue
				}
				// 找这个原消息
				var msgModel chat_models.ChatModel
				err = svcCtx.DB.Take(&msgModel, chatReq.Msg.QuoteMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "消息不存在")
					continue
				}

				// 不能回复撤回消息
				if msgModel.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "该消息已撤回")
					continue
				}
				userBaseInfo, err := redis_cache.GetUserBaseInfo(svcCtx.Redis, svcCtx.UserRpc, msgModel.SendUserID)
				if err != nil {
					logx.Error(err)
					return
				}
				chatReq.Msg.QuoteMsg.Msg = &msgModel.Msg
				chatReq.Msg.QuoteMsg.UserID = msgModel.SendUserID
				chatReq.Msg.QuoteMsg.UserNickName = userBaseInfo.NickName
				chatReq.Msg.QuoteMsg.OriginMsgDate = msgModel.CreatedAt
			}

			// 先入库
			msgID := InsertMsgByChat(svcCtx.DB, chatReq.RevUserID, myID, chatReq.Msg)
			// 发送消息给好友
			SendMsgByUser(svcCtx, chatReq.RevUserID, myID, chatReq.Msg, msgID)

		}
	}
}

// sendWsMapMsg 函数用于向给定的 WebSocket 连接映射中的所有连接发送消息。
// 参数:
//   - wsMap: 一个映射，键是字符串，值是指向 websocket.Conn 类型的指针，表示 WebSocket 连接。
//   - byteData: 要发送的消息数据，类型为 []byte。
func sendWsMapMsg(wsMap map[string]*websocket.Conn, byteData []byte) {
	// 遍历 wsMap 中的所有连接
	for _, conn := range wsMap {
		// 使用 conn.WriteMessage 方法向每个连接发送文本消息
		conn.WriteMessage(websocket.TextMessage, byteData)
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
func InsertMsgByChat(db *gorm.DB, revUserID uint, sendUserID uint, msg ctype.Msg) (msgID uint) {
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
		sendUser, ok := UserOnlineWsMap[sendUserID]
		if !ok {
			return
		}
		SendTipErrMsg(sendUser.CurrentConn, "消息保存失败")
	}

	return chatModel.ID
}

// SendMsgByUser 发消息，给谁发，谁发的
func SendMsgByUser(svcCtx *svc.ServiceContext, revUserID uint, sendUserID uint, msg ctype.Msg, msgID uint) {
	revUser, ok1 := UserOnlineWsMap[revUserID]
	sendUser, ok2 := UserOnlineWsMap[sendUserID]
	resp := ChatResponse{
		ID:        msgID,
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
		sendWsMapMsg(revUser.WsClientMap, byteData)
		// revUser.CurrentConn.WriteMessage(websocket.TextMessage, byteData)
		return
	}

	// 在线情况下，可以拿到对方用户信息
	// 对方不在线的情况下，只能通过调用用户信息的Rpc方法获取用户信息
	if ok1 {
		// 接收者在线
		resp.RevUser = ctype.UserInfo{
			ID:       revUserID,
			NickName: revUser.UserInfo.Nickname,
			Avatar:   revUser.UserInfo.Avatar,
		}
		resp.IsMe = true
		byteData, _ := json.Marshal(resp)
		sendWsMapMsg(revUser.WsClientMap, byteData)
		// revUser.CurrentConn.WriteMessage(websocket.TextMessage, byteData)
	} else {
		userBaseInfo, err := redis_cache.GetUserBaseInfo(svcCtx.Redis, svcCtx.UserRpc, revUserID)
		if err != nil {
			logx.Error(err)
			return
		}
		resp.RevUser = ctype.UserInfo{
			ID:       revUserID,
			NickName: userBaseInfo.NickName,
			Avatar:   userBaseInfo.Avatar,
		}
	}
	if ok2 {
		// 发送者也在线
		resp.SendUser = ctype.UserInfo{
			ID:       sendUserID,
			NickName: sendUser.UserInfo.Nickname,
			Avatar:   sendUser.UserInfo.Avatar,
		}
		resp.IsMe = false
		byteData, _ := json.Marshal(resp)
		sendWsMapMsg(sendUser.WsClientMap, byteData)
		// sendUser.CurrentConn.WriteMessage(websocket.TextMessage, byteData)
	}
}

package handler

import (
	"context"
	"encoding/json"
	"fmt"
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
	"gorm.io/gorm"
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
	UserID         uint          `json:"userID"`
	UserNickname   string        `json:"userNickname"`
	UserAvatar     string        `json:"userAvatar"`
	Msg            ctype.Msg     `json:"msg"`
	ID             uint          `json:"id"`
	MsgType        ctype.MsgType `json:"msgType"`
	CreatedAt      time.Time     `json:"createdAt"`
	IsMe           bool          `json:"isMe"`
	MemberNickname string        `json:"memberNickname"` // 群好友备注
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
			switch request.Msg.Type {
			case ctype.WithdrawMsgType: //撤回消息
				withdrawMsg := request.Msg.WithdrawMsg
				if withdrawMsg == nil {
					SendTipErrMsg(conn, "撤回消息的格式错误")
					continue
				}
				if withdrawMsg.MsgID == 0 {
					SendTipErrMsg(conn, "撤回消息id为空")
					continue
				}
				// 去找消息
				var groupMsg group_models.GroupMsgModel
				err = svcCtx.DB.Take(&groupMsg, withdrawMsg.MsgID).Error
				if err != nil {
					SendTipErrMsg(conn, "原消息不存在")
					continue
				}
				// 原消息不能是撤回消息
				if groupMsg.MsgType == ctype.WithdrawMsgType {
					SendTipErrMsg(conn, "该消息已撤回")
					continue
				}
				// 要去拿我在这个群的角色
				// 如果是自己撤自己的 并且自己是普通用户
				if member.Role == 3 {
					// 要判断时间是不是大于了2分钟
					now := time.Now()
					if now.Sub(groupMsg.CreatedAt) > 2*time.Minute {
						SendTipErrMsg(conn, "只能撤回两分钟以内的消息")
						continue
					}
				}
				// 查这个消息的用户在这个群的角色
				var msgUserRole int8 = 3
				err = svcCtx.DB.Model(group_models.GroupMemberModel{}).
					Where("group_id = ? and user_id = ?", request.GroupID, groupMsg.SendUserID).
					Select("role").
					Scan(&msgUserRole).Error
				// 这里有可能查不到  原因是这个消息的用户退群了，那么也是可以撤回的
				// 如果是管理员撤回  它能撤自己和用户的，没有时间限制
				if member.Role == 2 {
					// 不能撤群主和别的管理员
					if msgUserRole == 1 || (msgUserRole == 2 && groupMsg.SendUserID != userID) {
						SendTipErrMsg(conn, "管理员只能撤回自己或者普通用户的消息")
						continue
					}
				}

				// 如果是群主，那就能撤管理员和用户的
				var content = "撤回了一条消息"
				content = "你" + content
				// 前端可以判断，这个消息如果不是isMe，可以把你替换成对方的名称
				originMsg := groupMsg.Msg
				originMsg.WithdrawMsg = nil
				svcCtx.DB.Model(&groupMsg).Updates(group_models.GroupMsgModel{
					MsgPreview: "[撤回消息] - " + content,
					MsgType:    ctype.WithdrawMsgType,
					Msg: ctype.Msg{
						Type: ctype.WithdrawMsgType,
						WithdrawMsg: &ctype.WithdrawMsg{
							Content:   content,
							MsgID:     request.Msg.WithdrawMsg.MsgID,
							OriginMsg: &originMsg,
						},
					},
				})
			}
			// 群聊消息入库
			msgID := insertMsg(svcCtx.DB, conn, member, request.Msg)
			// 遍历这个用户列表，去找ws的客户端
			sendGroupOnlineUserMsg(svcCtx.DB, member, request.Msg, msgID)
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

// 给这个群的用户发消息
func sendGroupOnlineUserMsg(db *gorm.DB, member group_models.GroupMemberModel, msg ctype.Msg, msgID uint) {

	// 查在线的用户列表
	userOnlineIDList := getOnlineUserIDList()
	// 查这个群的成员 并且在线
	var groupMemberOnlineIDList []uint
	db.Model(group_models.GroupMemberModel{}).
		Where("group_id = ? and user_id in ?", member.GroupID, userOnlineIDList).
		Select("user_id").Scan(&groupMemberOnlineIDList)

	// 构造响应
	var chatResponse = ChatResponse{
		UserID:         member.UserID,
		Msg:            msg,
		ID:             msgID,
		MsgType:        msg.Type,
		CreatedAt:      time.Now(),
		MemberNickname: member.MemberNickname,
	}

	wsInfo, ok := UserOnlineWsMap[member.UserID]
	if ok {
		chatResponse.UserNickname = wsInfo.UserInfo.NickName
		chatResponse.UserAvatar = wsInfo.UserInfo.Avatar
	}

	for _, u := range groupMemberOnlineIDList {
		wsUserInfo, ok2 := UserOnlineWsMap[u]
		if !ok2 {
			continue
		}
		chatResponse.IsMe = false
		// 判断isMe
		if wsUserInfo.UserInfo.ID == member.UserID {
			chatResponse.IsMe = true
		}

		byteData, _ := json.Marshal(chatResponse)
		for _, w2 := range wsUserInfo.WsClientMap {
			w2.WriteMessage(websocket.TextMessage, byteData)
		}
	}
}

func insertMsg(db *gorm.DB, conn *websocket.Conn, member group_models.GroupMemberModel, msg ctype.Msg) uint {
	switch msg.Type {
	case ctype.WithdrawMsgType:
		fmt.Println("撤回消息自己是不入库的")
		return 0
	}
	groupMsg := group_models.GroupMsgModel{
		GroupID:       member.GroupID,
		SendUserID:    member.UserID,
		GroupMemberID: member.ID,
		MsgType:       msg.Type,
		Msg:           msg,
	}
	groupMsg.MsgPreview = groupMsg.MsgPreviewMethod()
	err := db.Create(&groupMsg).Error
	if err != nil {
		logx.Error(err)
		SendTipErrMsg(conn, "消息保存失败")
		return 0
	}
	return groupMsg.ID
}

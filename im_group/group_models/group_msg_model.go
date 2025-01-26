package group_models

import (
	"im_server/common/ctype"
	"im_server/common/models"
)

// GroupMsgModel 群消息表
type GroupMsgModel struct {
	models.Models
	GroupID          uint              `json:"groupID"`                           // 群id
	GroupModel       GroupModel        `gorm:"foreignKey:GroupID" json:"-"`       // 群
	SendUserID       uint              `json:"sendUserID"`                        // 发送者id
	GroupMemberID    uint              `json:"groupMemberID"`                     // 群成员id
	GroupMemberModel *GroupMemberModel `gorm:"foreignKey:GroupMemberID" json:"-"` // 对应的群成员
	MsgType          ctype.MsgType     `json:"msgType"`                           // 消息类型 1 文本类型  2 图片消息  3 视频消息 4 文件消息 5 语音消息  6 语言通话  7 视频通话  8 撤回消息 9回复消息 10 引用消息 11 at消息
	MsgPreview       string            `gorm:"size:64" json:"msgPreview"`         // 消息预览
	Msg              ctype.Msg         `json:"msg"`                               // 消息内容
	SystemMsg        *ctype.SystemMsg  `json:"systemMsg"`                         // 系统提示
}

func (chat GroupMsgModel) MsgPreviewMethod() string {
	if chat.SystemMsg != nil {
		switch chat.SystemMsg.Type {
		case 1:
			return "[系统消息]- 该消息涉黄，已被系统拦截"
		case 2:
			return "[系统消息]- 该消息涉恐，已被系统拦截"
		case 3:
			return "[系统消息]- 该消息涉政，已被系统拦截"
		case 4:
			return "[系统消息]- 该消息不正当言论，已被系统拦截"
		}
		return "[系统消息]"
	}
	return chat.Msg.MsgPreview()
}

package chat_models

import (
	"im_server/common/ctype"
	"im_server/common/models"
)

type ChatModel struct {
	models.Models
	SendUserID uint             `json:"sendUserID"`
	RevUserID  uint             `json:"revUserID"`
	MsgType    int8             `json:"msgType"`                   // 消息类型 1 文本类型  2 图片消息  3 视频消息 4 文件消息 5 语音消息  6 语言通话  7 视频通话  8 撤回消息 9回复消息 10 引用消息
	MsgPreview string           `gorm:"size:64" json:"msgPreview"` // 消息预览
	Msg        ctype.Msg        `json:"msg"`                       // 消息类容
	SystemMsg  *ctype.SystemMsg `json:"systemMsg"`                 // 系统提示
}

package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Msg struct {
	Type         int8          `json:"type"`         // 消息类型 和msgType一模一样
	Content      *string       `json:"content"`      // 为1的时候使用
	ImageMsg     *ImageMsg     `json:"imageMsg"`     // 图片消息
	VideoMsg     *VideoMsg     `json:"videoMsg"`     // 视频消息
	FileMsg      *FileMsg      `json:"fileMsg"`      // 文件消息
	VoiceMsg     *VoiceMsg     `json:"voiceMsg"`     // 语音消息
	VoiceCallMsg *VoiceCallMsg `json:"voiceCallMsg"` // 语言通话
	VideoCallMsg *VideoCallMsg `json:"videoCallMsg"` // 视频通话
	WithdrawMsg  *WithdrawMsg  `json:"withdrawMsg"`  // 撤回消息
	ReplyMsg     *ReplyMsg     `json:"replyMsg"`     // 回复消息
	QuoteMsg     *QuoteMsg     `json:"quoteMsg"`     // 引用消息
	AtMsg        *AtMsg        `json:"atMsg"`        // @用户的消息 群聊才有
}

// Scan 取出来的时候的数据
func (c *Msg) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 入库的数据
func (c Msg) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type ImageMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}
type VideoMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` // 时长 单位秒
}
type FileMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"` // 文件大小
	Type  string `json:"type"` // 文件类型 word
}
type VoiceMsg struct {
	Src  string `json:"src"`
	Time int    `json:"time"` // 时长 单位秒
}
type VoiceCallMsg struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
}
type VideoCallMsg struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断  2  网络原因挂断  3 未打通
}

// WithdrawMsg 撤回消息
type WithdrawMsg struct {
	Content   string `json:"content"` // 撤回的提示词
	OriginMsg *Msg   `json:"-"`       // 原消息
}
type ReplyMsg struct {
	MsgID   uint   `json:"msgID"`   // 消息id
	Content string `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Msg     *Msg   `json:"msg"`
}
type QuoteMsg struct {
	MsgID   uint   `json:"msgID"`   // 消息id
	Content string `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Msg     *Msg   `json:"msg"`
}

// AtMsg @消息
type AtMsg struct {
	UserID  uint   `json:"userID"`
	Content string `json:"content"` // 回复的文本消息
	Msg     *Msg   `json:"msg"`
}

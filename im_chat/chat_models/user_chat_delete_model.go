package chat_models

// UserChatDeleteModel 用户删除聊天记录表
type UserChatDeleteModel struct {
	UserID uint `json:"userID"`
	ChatID uint `json:"chatID"` // 聊天记录的id
}

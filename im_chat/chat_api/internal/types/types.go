// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type ChatDeleteRequest struct {
	IdList []uint `json:"idList"`
}

type ChatDeleteResponse struct {
}

type ChatHistoryRequest struct {
	Page     int  `form:"page,optional"`
	Limit    int  `form:"limit,optional"`
	FriendID uint `form:"friendId"` // 好友id
}

type ChatHistoryResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"userId"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	CreatedAt string `json:"created_at"` // 消息时间
}

type ChatSession struct {
	UserID     uint   `json:"userId"`
	Avatar     string `json:"avatar"`
	Nickname   string `json:"nickname"`
	CreatedAt  string `json:"created_at"` // 消息时间
	MsgPreview string `json:"msgPreview"` // 消息预览
	IsTop      bool   `json:"isTop"`      // 是否置顶
}

type ChatSessionRequest struct {
	Page  int `form:"page,optional"`
	Limit int `form:"limit,optional"`
	Key   int `form:"key,optional"`
}

type ChatSessionResponse struct {
	List  []ChatSession `json:"list"`
	Count int64         `json:"count"`
}

type UserTopRequest struct {
	FriendID uint `json:"friendId"` // 好友id
}

type UserTopResponse struct {
}

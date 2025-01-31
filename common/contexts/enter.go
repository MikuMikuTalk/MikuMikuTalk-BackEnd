package contexts

// 定义自己的类型
type contextKey string

const (
	ContextKeyClientIP contextKey = "clientIP"
	ContextKeyToken    contextKey = "token"
	ContextKeyUserID   contextKey = "userID"
)

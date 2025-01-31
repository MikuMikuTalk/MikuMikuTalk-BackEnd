package middleware

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 定义自己的类型
type contextKey string

const (
	clientIPKey contextKey = "clientIP"
	tokenKey    contextKey = "token"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := httpx.GetRemoteAddr(r)
		ctx := context.WithValue(r.Context(), clientIPKey, clientIP)
		token := r.Header.Get("Authorization")
		ctx = context.WithValue(ctx, tokenKey, token)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

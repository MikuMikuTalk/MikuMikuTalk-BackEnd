package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"im_server/common/contexts"
	"im_server/common/etcd"
	"im_server/common/log_stash"
	"im_server/im_settings/config"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Writer struct {
	http.ResponseWriter
	Body []byte
}

func (w *Writer) Write(b []byte) (int, error) {
	w.Body = append(w.Body, b...)
	return w.ResponseWriter.Write(b)
}

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := httpx.GetRemoteAddr(r)
		ctx := context.WithValue(r.Context(), contexts.ContextKeyClientIP, clientIP)
		// header中取出authorization 中的jwt字段值
		token := r.Header.Get("Authorization")
		etcdStorage := etcd.NewEtcdStorage("127.0.0.1:2379")
		appConfiguration, _ := etcdStorage.Get("config")

		var appConfig config.Config

		err := json.Unmarshal([]byte(appConfiguration), &appConfig)
		if err != nil {
			http.Error(w, "json解析失败", http.StatusBadRequest)
		}

		claims, err := jwts.ParseToken(token, appConfig.Auth.AuthSecret)
		if err != nil {
			http.Error(w, "jwt解析失败", http.StatusBadRequest)
		}

		ctx = context.WithValue(ctx, contexts.ContextKeyToken, token)
		if claims != nil {
			logx.Info("log_middleware parsed claims:", claims)
			userID := claims.UserID
			ctx = context.WithValue(ctx, contexts.ContextKeyUserID, userID)
		}
		next(w, r.WithContext(ctx))
	}
}

func LogActionMiddleware(pusher *log_stash.Pusher) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			clientIP := httpx.GetRemoteAddr(r)

			// 设置入参
			pusher.SetRequest(r)
			pusher.SetHeaders(r)

			ctx := context.WithValue(r.Context(), contexts.ContextKeyClientIP, clientIP)

			// // header中取出authorization 中的jwt字段值
			// token := r.Header.Get("Authorization")
			// etcdStorage := etcd.NewEtcdStorage("127.0.0.1:2379")
			// appConfiguration, _ := etcdStorage.Get("config")

			// var appConfig config.Config

			// err := json.Unmarshal([]byte(appConfiguration), &appConfig)
			// if err != nil {
			// 	http.Error(w, "json解析失败", http.StatusBadRequest)
			// }

			// claims, err := jwts.ParseToken(token, appConfig.Auth.AuthSecret)
			// if err != nil {
			// 	http.Error(w, "jwt解析失败", http.StatusBadRequest)
			// }

			// ctx = context.WithValue(ctx, contexts.ContextKeyToken, token)
			// if claims != nil {
			// 	logx.Info("log_middleware parsed claims:", claims)
			// 	userID := claims.UserID
			// 	ctx = context.WithValue(ctx, contexts.ContextKeyUserID, userID)
			// }

			var nw = Writer{
				ResponseWriter: w,
			}
			next(&nw, r.WithContext(ctx))
			if pusher.GetResponse() {
				pusher.SetResponse(string(nw.Body))
			}
		}
	}
}

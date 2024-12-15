package modules

import (
	"encoding/json"
	"net/http"
)

// 处理错误响应
func (g *GatewayService) handleError(res http.ResponseWriter, status int, message string) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Cache-Control", "no-store")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(Response{
		Code: 7,
		Msg:  message,
	})
}

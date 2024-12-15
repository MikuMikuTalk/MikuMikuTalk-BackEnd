package modules

import (
	"encoding/json"
	"fmt"
	"im_server/common/etcd"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// 对请求进行认证
func (g *GatewayService) authenticate(req *http.Request) error {
	remoteAddr := strings.Split(req.RemoteAddr, ":")[0]
	logx.Info("remoteAddr:", req.RemoteAddr)

	// 获取认证服务地址
	authAPIAddr := etcd.GetServiceAddr(g.config.Etcd, "auth_api")
	authURL := fmt.Sprintf("http://%s/api/auth/authentication", authAPIAddr)

	// 构建认证请求
	authReq, err := http.NewRequest("POST", authURL, nil)
	if err != nil {
		return err
	}
	authReq.Header.Set("Authorization", req.Header.Get("Authorization"))
	authReq.Header.Set("X-Forwarded-For", remoteAddr)
	authReq.Header.Set("ValidPath", req.URL.Path)

	// 执行认证请求
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		return err
	}
	defer authRes.Body.Close()

	// 解析认证响应
	var authResponse Response
	if err := json.NewDecoder(authRes.Body).Decode(&authResponse); err != nil {
		return err
	}

	if authResponse.Data != "ok" {
		return fmt.Errorf("authentication failed")
	}

	return nil
}

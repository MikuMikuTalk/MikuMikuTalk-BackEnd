package modules

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// 判断是否为 WebSocket 请求
func IsWebSocketRequest(req *http.Request) bool {
	upgrade := strings.ToLower(req.Header.Get("Connection"))
	websocket := strings.ToLower(req.Header.Get("Upgrade"))
	return strings.Contains(upgrade, "upgrade") && websocket == "websocket"
}

// singleJoiningSlash 确保路径拼接时只有一个 "/"
func SingleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// 处理 WebSocket 请求

func (g *GatewayService) handleWebSocket(addr string, res http.ResponseWriter, req *http.Request) error {
	targetURL := fmt.Sprintf("http://%s", addr)
	logx.Info("targetUrl: ", targetURL)
	logx.Info("WebSocket Proxy to: ", targetURL)

	// parsedUrl

	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return fmt.Errorf("invalid target URL: %v", err)
	}
	logx.Info("parsed Url:", parsedURL)
	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	// 自定义修改请求方法
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = parsedURL.Scheme
		req.URL.Host = parsedURL.Host
		req.URL.Path = SingleJoiningSlash(parsedURL.Path, req.URL.Path)
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Origin-Host", parsedURL.Host)
	}
	// 启动 WebSocket 反向代理
	proxy.ServeHTTP(res, req)
	return nil
}

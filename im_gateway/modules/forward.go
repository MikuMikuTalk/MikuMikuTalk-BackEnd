package modules

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// 转发 HTTP 请求
func (g *GatewayService) forwardRequest(addr string, originalReq *http.Request, res http.ResponseWriter) error {
	targetURL := fmt.Sprintf("http://%s%s", addr, originalReq.URL.RequestURI())
	logx.Info("Proxy Request URL: ", targetURL)

	// 读取并重置请求体
	body, err := io.ReadAll(originalReq.Body)
	if err != nil {
		return err
	}
	originalReq.Body = io.NopCloser(bytes.NewBuffer(body))

	// 创建代理请求
	proxyReq, err := http.NewRequest(originalReq.Method, targetURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// 复制原始请求头
	for header, values := range originalReq.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	// 设置 X-Forwarded-For
	remoteAddr := strings.Split(originalReq.RemoteAddr, ":")[0]
	xff := originalReq.Header.Get("X-Forwarded-For")
	if xff != "" {
		proxyReq.Header.Set("X-Forwarded-For", fmt.Sprintf("%s, %s", xff, remoteAddr))
	} else {
		proxyReq.Header.Set("X-Forwarded-For", remoteAddr)
	}

	// 发送代理请求
	response, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 复制响应头
	for key, values := range response.Header {
		for _, value := range values {
			res.Header().Add(key, value)
		}
	}

	// 复制响应体
	if _, err := io.Copy(res, response.Body); err != nil {
		return err
	}

	logx.Info("Proxy Response: ", response.Status)
	return nil
}

// 处理网关请求
func (g *GatewayService) HandleRequest(res http.ResponseWriter, req *http.Request) {
	service, err := g.extractService(req.URL.Path)
	if err != nil {
		g.handleError(res, http.StatusBadRequest, "服务错误")
		return
	}

	// 获取服务地址
	addr := g.getServiceAddr(service)
	if addr == "" {
		g.handleError(res, http.StatusBadGateway, "服务错误")
		return
	}

	// 认证
	if err := g.authenticate(req); err != nil {
		g.handleError(res, http.StatusUnauthorized, err.Error())
		return
	}

	// 判断是否为 WebSocket 请求
	if IsWebSocketRequest(req) {

		if err := g.handleWebSocket(addr, res, req); err != nil {
			g.handleError(res, http.StatusInternalServerError, "WebSocket 代理失败")
		}
		return
	}

	// 转发 HTTP 请求
	if err := g.forwardRequest(addr, req, res); err != nil {
		g.handleError(res, http.StatusBadGateway, "服务错误")
		return
	}
}

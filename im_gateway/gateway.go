package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"im_server/common/etcd"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

// 配置结构体定义
type Config struct {
	Addr string // 网关服务监听地址
	Etcd string // Etcd 服务地址
	Log  logx.LogConf
}

// HTTP 响应的通用结构体
type Response struct {
	Code int    `json:"code"` // 状态码
	Data any    `json:"data"` // 数据
	Msg  string `json:"msg"`  // 消息
}

// 配置文件路径的命令行参数
var configFile = flag.String("f", "etc/settings.yaml", "the config file")

// 网关服务结构体
type GatewayService struct {
	config Config // 配置信息
}

// 创建新的网关服务实例
func NewGatewayService(config Config) *GatewayService {
	return &GatewayService{
		config: config,
	}
}

// 从路径中提取服务名称
func (g *GatewayService) extractService(path string) (string, error) {
	// 使用正则表达式提取 /api/{service}/ 中的 service 名称
	regex, _ := regexp.Compile(`/api/([^/]+)/`)
	matches := regex.FindStringSubmatch(path)
	if len(matches) != 2 {
		return "", fmt.Errorf("invalid service path")
	}
	return matches[1], nil
}

// 从 Etcd 获取服务地址
func (g *GatewayService) getServiceAddr(service string) string {
	addr := etcd.GetServiceAddr(g.config.Etcd, service+"_api")
	if addr == "" {
		logx.Error("不匹配的服务", service)
	}
	return addr
}

// 创建认证请求
func (g *GatewayService) createAuthRequest(authURL string, originalReq *http.Request, remoteAddr string) (*http.Request, error) {
	authReq, err := http.NewRequest("POST", authURL, nil)
	if err != nil {
		return nil, err
	}

	authReq.Header.Set("Authorization", originalReq.Header.Get("Authorization"))
	authReq.Header.Set("X-Forwarded-For", remoteAddr)
	authReq.Header.Set("ValidPath", originalReq.URL.Path)

	return authReq, nil
}

// 对请求进行认证
func (g *GatewayService) authenticate(req *http.Request) error {
	remoteAddr := strings.Split(req.RemoteAddr, ":")[0]
	logx.Info("remoteAddr:", req.RemoteAddr)

	// 获取认证服务地址
	authAPIAddr := etcd.GetServiceAddr(g.config.Etcd, "auth_api")
	auth_url := fmt.Sprintf("http://%s/api/auth/authentication", authAPIAddr)

	// 构建认证请求
	authReq, err := g.createAuthRequest(auth_url, req, remoteAddr)
	if err != nil {
		return err
	}

	// 执行认证请求
	return g.performAuthentication(authReq)
}

// 执行认证请求
func (g *GatewayService) performAuthentication(authReq *http.Request) error {
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

// 转发请求
func (g *GatewayService) forwardRequest(addr string, originalReq *http.Request, res http.ResponseWriter) error {
	// 构建目标 URL
	targetURL := fmt.Sprintf("http://%s%s", addr, originalReq.URL.RequestURI())
	logx.Info("Proxy Request URL: ", targetURL)

	// 读取并重置请求体
	body, err := g.readAndResetRequestBody(originalReq)
	if err != nil {
		return err
	}

	// 创建代理请求
	proxyReq, err := g.createProxyRequest(targetURL, originalReq, body)
	if err != nil {
		return err
	}

	// 发送代理请求
	return g.sendProxyRequest(proxyReq, res)
}

// 读取并重置请求体
func (g *GatewayService) readAndResetRequestBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body)) // 重置请求体
	return body, nil
}

// 创建代理请求
func (g *GatewayService) createProxyRequest(url string, originalReq *http.Request, body []byte) (*http.Request, error) {
	proxyReq, err := http.NewRequest(originalReq.Method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 复制原始请求头
	g.copyHeaders(originalReq, proxyReq)

	// 设置 X-Forwarded-For
	g.setForwardedFor(originalReq, proxyReq)

	return proxyReq, nil
}

// 复制请求头
func (g *GatewayService) copyHeaders(src *http.Request, dst *http.Request) {
	for header, values := range src.Header {
		for _, value := range values {
			dst.Header.Add(header, value)
		}
	}
}

// 设置 X-Forwarded-For 头部
func (g *GatewayService) setForwardedFor(originalReq *http.Request, proxyReq *http.Request) {
	remoteAddr := strings.Split(originalReq.RemoteAddr, ":")[0]
	xff := originalReq.Header.Get("X-Forwarded-For")
	if xff != "" {
		proxyReq.Header.Set("X-Forwarded-For", fmt.Sprintf("%s, %s", xff, remoteAddr))
	} else {
		proxyReq.Header.Set("X-Forwarded-For", remoteAddr)
	}
}

// 发送代理请求
func (g *GatewayService) sendProxyRequest(proxyReq *http.Request, res http.ResponseWriter) error {
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

// 处理错误响应
func (g *GatewayService) handleError(res http.ResponseWriter, status int, message string) {
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(Response{
		Code: 7,
		Msg:  message,
	})
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
		g.handleError(res, http.StatusUnauthorized, "认证失败")
		return
	}

	// 转发请求
	if err := g.forwardRequest(addr, req, res); err != nil {
		g.handleError(res, http.StatusBadGateway, "服务错误")
		return
	}
}

func main() {
	flag.Parse()

	// 加载配置
	var config Config
	conf.MustLoad(*configFile, &config)
	logx.MustSetup(config.Log)

	// 初始化网关服务
	gateway := NewGatewayService(config)

	// 设置路由
	http.HandleFunc("/", gateway.HandleRequest)

	// 启动网关服务
	fmt.Printf("网关运行在 %s\n", config.Addr)
	http.ListenAndServe(config.Addr, nil)
}

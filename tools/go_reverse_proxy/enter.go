package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Proxy struct {
	targetURL *url.URL
	proxy     *httputil.ReverseProxy
}

// NewProxy 创建一个新的反向代理实例
func NewProxy(target string) (*Proxy, error) {
	parsedURL, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("invalid target URL: %v", err)
	}
	return &Proxy{
		targetURL: parsedURL,
		proxy:     httputil.NewSingleHostReverseProxy(parsedURL),
	}, nil
}

// ServeHTTP 处理反向代理的请求
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 日志记录请求信息
	log.Printf("Proxying request from %s: %s %s -> %s", r.RemoteAddr, r.Method, r.URL.Path, p.targetURL.String())

	// 自定义 Director 以调整请求
	p.proxy.Director = func(req *http.Request) {
		req.URL.Scheme = p.targetURL.Scheme
		req.URL.Host = p.targetURL.Host
		req.URL.Path = singleJoiningSlash(p.targetURL.Path, req.URL.Path)
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Origin-Host", p.targetURL.Host)
	}
	p.proxy.ServeHTTP(w, r)
}

// singleJoiningSlash 确保路径拼接时只有一个 "/"
func singleJoiningSlash(a, b string) string {
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

func main() {
	target := "http://127.0.0.1:20023"
	addr := "127.0.0.1:8081"

	proxy, err := NewProxy(target)
	if err != nil {
		log.Fatalf("Error creating proxy: %v", err)
	}

	fmt.Printf("Reverse proxy server listening on: %s, targeting: %s\n", addr, target)
	if err := http.ListenAndServe(addr, proxy); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

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
	"im_server/utils/logs"

	"github.com/zeromicro/go-zero/core/conf"
)

type Data struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func toJson(data Data) []byte {
	byteData, _ := json.MarshalIndent(data, "", "  ")
	return byteData
}
func gateway(res http.ResponseWriter, req *http.Request) {
	// 更精确地匹配请求前缀 /api/user/xx
	regex, _ := regexp.Compile(`/api/([^/]+)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}

	// 从etcd中查找对应的服务
	service := addrList[1]
	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		logs.MyLogger.Errorln("不匹配的服务", service)
		res.WriteHeader(http.StatusBadGateway)
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}

	// 获取客户端 IP 地址
	remoteAddr := strings.Split(req.RemoteAddr, ":")[0]
	logs.MyLogger.Infoln("remoteAddr:", req.RemoteAddr)

	// 生成目标请求 URL
	url := fmt.Sprintf("http://%s%s", addr, req.URL.RequestURI())
	logs.MyLogger.Infoln("Proxy Request URL: ", url)

	// 读取原始请求体
	body, err := io.ReadAll(req.Body)
	if err != nil {
		logs.MyLogger.Errorln("读取请求体失败:", err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	req.Body = io.NopCloser(bytes.NewBuffer(body)) // 重新设置请求体

	// 创建代理请求
	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewBuffer(body))
	if err != nil {
		logs.MyLogger.Errorln("创建代理请求失败:", err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}

	// 复制原始请求的头信息到代理请求
	for header, values := range req.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	// 设置请求头，保留原有 X-Forwarded-For 头信息
	xff := req.Header.Get("X-Forwarded-For")
	if xff != "" {
		proxyReq.Header.Set("X-Forwarded-For", fmt.Sprintf("%s, %s", xff, remoteAddr))
	} else {
		proxyReq.Header.Set("X-Forwarded-For", remoteAddr)
	}

	// 执行请求并处理响应
	response, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		logs.MyLogger.Errorln("请求目标服务失败:", err)
		res.WriteHeader(http.StatusBadGateway)
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	defer response.Body.Close()

	// 将目标服务的响应头传递给客户端
	for key, values := range response.Header {
		for _, value := range values {
			res.Header().Add(key, value)
		}
	}

	// 将响应体内容复制给客户端
	_, err = io.Copy(res, response.Body)
	if err != nil {
		logs.MyLogger.Errorln("响应写入失败:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	logs.MyLogger.Infoln("Proxy Response: ", response.Status)
}

var configFile = flag.String("f", "settings.yaml", "the config file")

type Config struct {
	Addr string
	Etcd string
}

var config Config

func main() {
	flag.Parse()
	conf.MustLoad(*configFile, &config)

	// 回调函数
	http.HandleFunc("/", gateway)
	fmt.Printf("网关运行在 %s\n", config.Addr)
	// 绑定服务
	http.ListenAndServe(config.Addr, nil)
}

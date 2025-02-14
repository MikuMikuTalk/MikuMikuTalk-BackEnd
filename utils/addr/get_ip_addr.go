package addr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/netip"

	"github.com/zeromicro/go-zero/core/logx"
)

// BaiduAPIResponse 是百度 IP 归属地 API 的返回结构
type BaiduAPIResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Location string `json:"location"`
	} `json:"data"`
}

// GetAddr 使用百度 API 查询 IP 归属地
func GetAddr(ip string) string {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		logx.Info("错误IP")
		return "错误IP"
	}
	if IsInternalIP(addr) {
		return "内网IP"
	}

	// 调用百度 API
	url := fmt.Sprintf("https://opendata.baidu.com/api.php?query=%s&co=&resource_id=6006&oe=utf8", ip)
	resp, err := http.Get(url)
	if err != nil {
		logx.Errorf("查询 IP 归属地失败: %v", err)
		return "查询失败"
	}
	defer resp.Body.Close()

	var result BaiduAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logx.Errorf("解析 API 响应失败: %v", err)
		return "解析失败"
	}

	if result.Status != "0" || len(result.Data) == 0 {
		logx.Errorf("API 返回错误: status=%s", result.Status)
		return "API 错误"
	}

	// 返回归属地信息
	return result.Data[0].Location
}

// IsInternalIP 判断是否为内网 IP
func IsInternalIP(addr netip.Addr) bool {
	if addr.IsLoopback() {
		return true
	}

	// 处理 IPv6 地址
	if !addr.Is4() {
		return false
	}

	ip4 := addr.As4()
	return (ip4[0] == 192 && ip4[1] == 168) ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 32) ||
		(ip4[0] == 10) ||
		(ip4[0] == 169 && ip4[1] == 254)
}

package addr

import (
	"fmt"
	"net/netip"

	"github.com/oschwald/maxminddb-golang/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

var db *maxminddb.Reader

func init() {
	var err error
	db, err = maxminddb.Open("GeoLite2-City.mmdb")
	if err != nil {
		panic(err)
	}
}

// Close 关闭数据库连接
func Close() {
	if db != nil {
		db.Close()
	}
}

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

func GetAddr(ip string) string {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		logx.Info("错误IP")
		return "错误IP"
	}
	if IsInternalIP(addr) {
		return "内网IP"
	}

	var record struct {
		City struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"city"`
		Subdivisions []struct {
			Names   map[string]string `maxminddb:"names"`
			IsoCode string            `maxminddb:"iso_code"`
		} `maxminddb:"subdivisions"`
		Country struct {
			Names   map[string]string `maxminddb:"names"`
			IsoCode string            `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}

	// 使用 db.Lookup 获取 Result
	lookup := db.Lookup(addr)
	if err := lookup.Decode(&record); err != nil {
		return "错误的地址"
	}

	// 获取国家名称作为后备选项
	country := record.Country.Names["zh-CN"]
	if country == "" {
		return "未知地址"
	}

	// 获取省份名称
	var province string
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}

	// 获取城市名称
	city := record.City.Names["zh-CN"]

	// 根据可用信息构建地址字符串
	if province != "" && city != "" {
		return fmt.Sprintf("%s-%s", province, city)
	} else if province != "" {
		return fmt.Sprintf("%s-%s", country, province)
	} else if city != "" {
		return fmt.Sprintf("%s-%s", country, city)
	}

	return country
}

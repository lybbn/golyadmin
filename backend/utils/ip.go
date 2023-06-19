package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取客户端真实IP
func GetRealClientIP(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = c.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	ClientIP := c.ClientIP()
	if ClientIP != "127.0.0.1" {
		ip = ClientIP
	}
	return ip
}

// 获取外网ip地址详情
func GetIpLocation(ip, key string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	url := "https://restapi.amap.com/v5/ip?ip=" + ip + "&type=4&key=" + key
	fmt.Println("url", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("restapi.amap.com failed:", err)
		return "未知位置"
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}

	return m["country"] + "-" + m["province"] + "-" + m["city"] + "-" + m["district"] + "-" + m["isp"]
}

// 获取局域网ip地址(此获取ip方式存在多网卡时，ip不准确问题)
func GetLocaHost() string {
	netifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netifaces); i++ {
		if (netifaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netifaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}

// GetLocalIpAddr 利用udp 网络连接 获取本地IP地址(出口流量的IP地址)
func GetLocalIpAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return GetLocaHost()
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// 192.168.0.101:61085
	ip := strings.Split(localAddr.String(), ":")[0]

	return ip

}

package utils

import (
	"strconv"
	"strings"
	"time"

	"gitee.com/lybbn/go-vue-lyadmin/global"
)

// 获取当前日期时间格式为 2006-01-02 15:04:05
func GetNowTimeFormatStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 转换int类型为字符串
func FormatInt2String(n int) string {
	return strconv.Itoa(n)
}

// 转换字符串类型为int
func FormatString2Int(e string) (int, error) {
	return strconv.Atoi(e)
}

// 获取配置文件的端口字符串形式 如   :9000
func GetServerPort() string {
	if global.GVLA_CONFIG.System.HttpPort == 80 {
		return ""
	}
	return ":" + FormatInt2String(global.GVLA_CONFIG.System.HttpPort)
}

// 解析1d 2h 30m格式时间为time.Duration
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

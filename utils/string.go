package utils

import (
	"strconv"
	"time"

	"gitee.com/lybbn/go-vue-lyadmin/global"
)

func GetNowTimeFormatStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatInt2String(n int) string {
	return strconv.Itoa(n)
}

func GetServerPort() string {
	if global.GVLA_CONFIG.System.HttpPort == 80 {
		return ""
	}
	return ":" + FormatInt2String(global.GVLA_CONFIG.System.HttpPort)
}

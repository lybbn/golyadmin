package utils

import (
	"crypto/rand"
	"strconv"
	"strings"
	"time"

	"gitee.com/lybbn/golyadmin/global"
)

const (
	symbol = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"
	letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
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

func generateRandString(length int, s string) string {
	var chars = []byte(s)
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// GenerateRandomKey6 生成6为随机字符串
func GenerateRandomKey6() string {
	return generateRandString(6, letter)
}

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

// 把数据库日期时间格式为 2006-01-02 15:04:05
func DateTimeFormat(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04:05")
}

// 转换int类型为字符串
func FormatInt2String(n int) string {
	return strconv.Itoa(n)
}

// 转换uint类型为字符串
func FormatUint2String(n uint64) string {
	return strconv.FormatUint(n, 10)
}

// 转换字符串类型为int
func FormatString2Int(e string) (int, error) {
	return strconv.Atoi(e)
}

// 转换字符串类型为int
func FormatString2Bool(e string) (bool, error) {
	return strconv.ParseBool(e)
}

// 获取配置文件的端口字符串形式 如   :9000
func GetServerPort() string {
	if global.GL_CONFIG.System.HttpPort == 80 || global.GL_CONFIG.System.HttpPort == 443 {
		return ""
	}
	return ":" + FormatInt2String(global.GL_CONFIG.System.HttpPort)
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

// 字符型数组去重
func RemoveDuplicatesArrStr(arr []string) []string {
	uniqueMap := make(map[string]bool)

	for _, v := range arr {
		if _, ok := uniqueMap[v]; !ok {
			uniqueMap[v] = true
		}
	}

	var uniqueArr []string
	for k := range uniqueMap {
		uniqueArr = append(uniqueArr, k)
	}
	return uniqueArr
}

// 数字型数组去重
func RemoveDuplicatesArrInt(arr []int) []int {
	result := []int{}
	tempMap := map[int]byte{}
	for _, e := range arr {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

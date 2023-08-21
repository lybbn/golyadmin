package utils

import (
	"crypto/rand"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"

	"gitee.com/lybbn/golyadmin/global"
)

const (
	symbol         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"
	letter         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	formatDateTime = "2006-01-02 15:04:05"
)

// 获取时间戳
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// GetNowDayStr 获取年月日20060102
func GetNowDayStr() string {
	t := "20060102"
	return time.Now().Format(t)
}

// 获取年月日2006-01-02
func GetNowDayStr2() string {
	t := "2006-01-02"
	return time.Now().Format(t)
}

// 获取当前日期时间格式为 2006-01-02 15:04:05
func GetNowTimeFormatStr() string {
	return time.Now().Format(formatDateTime)
}

// 把数据库日期时间格式为 2006-01-02 15:04:05
func DateTimeFormat(datetime time.Time) string {
	return datetime.Format(formatDateTime)
}

// 转换int类型为字符串
func FormatInt2String(n int) string {
	return strconv.Itoa(n)
}

// 转换int64类型为字符串
func FormatInt642String(n int64) string {
	return strconv.FormatInt(n, 10)
}

// 转换uint类型为字符串
func FormatUint2String(n uint64) string {
	return strconv.FormatUint(n, 10)
}

// 转换float类型为字符串 s为要保留的位数
func FormatFloat2String(n float64, s int) string {
	return strconv.FormatFloat(n, 'f', s, 64)
}

// 保留float64类型小数点位数 s为要保留的位数
func FormatFloat64ToFloat64(n float64, s int) float64 {
	value, _ := strconv.ParseFloat(strconv.FormatFloat(n, 'f', s, 64), 64)
	return value
}

// 转换字符串类型为int
func FormatString2Int(v string) (int, error) {
	return strconv.Atoi(v)
}

// 转换字符串类型为int64
func FormatString2Int64(v string) int64 {
	value, _ := strconv.ParseInt(v, 10, 64)
	return value
}

// 转换字符串类型为uint64
func FormatString2Uint64(v string) uint64 {
	value, _ := strconv.ParseUint(v, 0, 64)
	return value
}

// 转换字符串类型为float64
func FormatString2Float64(v string) float64 {
	value, _ := strconv.ParseFloat(v, 64)
	return value
}

// 转换字符串类型为bool
func FormatString2Bool(v string) (bool, error) {
	if strings.ToLower(v) == "on" || strings.ToLower(v) == "1" || strings.ToLower(v) == "yes" {
		return true, nil
	}
	if strings.ToLower(v) == "off" || strings.ToLower(v) == "0" || strings.ToLower(v) == "no" {
		return false, nil
	}
	return strconv.ParseBool(v)
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

// GenerateRandomKey6 生成6位随机字符串
func GenerateRandomKey6() string {
	return generateRandString(6, letter)
}

func GenerateRandomNumsInt(min, max int) int {
	mrand.Seed(time.Now().Unix()) //Seed生成的随机数
	return mrand.Intn(max-min) + min
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

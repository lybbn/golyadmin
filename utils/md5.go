package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//@function: MD5
//@description: md5加密
//@param: str []byte
//@return: md5 string

func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

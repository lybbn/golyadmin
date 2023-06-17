package utils

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
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

// MakePassowrd 使用 bcrypt 对密码进行加密
func MakePassowrd(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// CheckPassword 对比明文密码和数据库的哈希值
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MakePasswordSalt 根据明文密码和加盐值生成加密密码
func MakePasswordSalt(password string, salt string) string {
	var rb []byte
	rb, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(rb)
}

// CheckPasswordSalt 对比明文密码和数据库的哈希值和加盐值
func CheckPasswordSalt(password, hash string, salt string) bool {
	passwordHash := MakePasswordSalt(password, salt)
	return hash == passwordHash
}

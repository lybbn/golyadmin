package utils

import (
	"strings"

	"github.com/google/uuid"
)

// @function: MakeUUID
// @description: 随机生成UUID，替换-为空，转换为32位字符串
// @param:
// @return: uuid string
func MakeUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
